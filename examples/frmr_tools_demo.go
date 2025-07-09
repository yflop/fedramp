package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocomply/fedramp/pkg/fedramp/frmr"
)

func main() {
	fmt.Println("FedRAMP Machine Readable (FRMR) Tools Demo")
	fmt.Println("===========================================\n")

	// Fetch documents
	fmt.Println("1. Fetching FRMR Documents from GitHub...")
	fmt.Println("------------------------------------------")
	
	ksiDoc := fetchDocument("KSI", "https://raw.githubusercontent.com/FedRAMP/docs/main/FRMR.KSI.key-security-indicators.json")
	masDoc := fetchDocument("MAS", "https://raw.githubusercontent.com/FedRAMP/docs/main/FRMR.MAS.minimum-assessment-standard.json")
	scnDoc := fetchDocument("SCN", "https://raw.githubusercontent.com/FedRAMP/docs/main/FRMR.SCN.significant-change-notifications.json")

	// Schema validation
	fmt.Println("\n2. Schema Validation")
	fmt.Println("--------------------")
	validateDocument("KSI", ksiDoc)
	validateDocument("MAS", masDoc)
	validateDocument("SCN", scnDoc)

	// Filtering
	fmt.Println("\n3. Filtering Documents")
	fmt.Println("----------------------")
	
	// Filter KSI for Low impact only
	lowKSI := frmr.FilterDocument(ksiDoc, frmr.FilterOptions{
		ImpactLevels: []string{"Low"},
	})
	fmt.Printf("Original KSI requirements: %d\n", countKSIRequirements(ksiDoc))
	fmt.Printf("Filtered KSI requirements (Low only): %d\n", countKSIRequirements(lowKSI))

	// Filter for specific KSIs
	iamKSI := frmr.FilterDocument(ksiDoc, frmr.FilterOptions{
		KSIIDs: []string{"KSI-IAM", "KSI-MLA"},
	})
	fmt.Printf("IAM and MLA KSIs only: %d KSIs\n", len(iamKSI.KSI))

	// Combining documents
	fmt.Println("\n4. Combining Documents")
	fmt.Println("----------------------")
	
	// Note: Can't combine KSI with MAS/SCN due to different structures
	// But we can demonstrate the concept
	fmt.Println("Combining filtered documents...")
	combined, err := frmr.CombineFRMRDocuments(lowKSI, iamKSI)
	if err != nil {
		log.Printf("Error combining: %v", err)
	} else {
		fmt.Printf("Combined document has %d KSIs\n", len(combined.KSI))
	}

	// Generate evidence template
	fmt.Println("\n5. Evidence Template Generation")
	fmt.Println("-------------------------------")
	
	evidenceTemplate := frmr.GenerateEvidenceTemplate(ksiDoc)
	fmt.Printf("Generated evidence template with %d requirements\n", len(evidenceTemplate)-1) // -1 for metadata
	
	// Save template
	saveEvidenceTemplate(evidenceTemplate, "ksi-evidence-template.json")

	// Markdown export
	fmt.Println("\n6. Markdown Export")
	fmt.Println("------------------")
	
	exporter := frmr.NewMarkdownExporter()
	var buf bytes.Buffer
	if err := exporter.Export(lowKSI, &buf); err != nil {
		log.Printf("Error exporting to markdown: %v", err)
	} else {
		// Save first 1000 chars as preview
		preview := buf.String()
		if len(preview) > 1000 {
			preview = preview[:1000] + "..."
		}
		fmt.Printf("Markdown preview:\n%s\n", preview)
		
		// Save full markdown
		saveMarkdown(buf.String(), "frmr-low-ksi.md")
	}

	// Demonstrate validation with evidence
	fmt.Println("\n7. KSI Validation with Evidence")
	fmt.Println("--------------------------------")
	
	// Create sample evidence
	evidence := createSampleEvidence(ksiDoc)
	
	// Validate each KSI
	totalMet := 0
	totalReqs := 0
	for _, ksi := range ksiDoc.KSI {
		result := ksi.ValidateKSI(evidence)
		totalReqs += result.TotalCount
		totalMet += result.MetCount
		
		if result.FullyMet {
			fmt.Printf("✅ %s: Fully compliant\n", ksi.ID)
		} else {
			fmt.Printf("⚠️  %s: %d/%d requirements met\n", ksi.ID, result.MetCount, result.TotalCount)
		}
	}
	
	fmt.Printf("\nOverall compliance: %.1f%% (%d/%d requirements met)\n", 
		float64(totalMet)/float64(totalReqs)*100, totalMet, totalReqs)

	// Advanced filtering example
	fmt.Println("\n8. Advanced Filtering")
	fmt.Println("--------------------")
	
	// Get only FRR and KSI types
	rulesAndKSI := frmr.FilterDocument(ksiDoc, frmr.FilterOptions{
		Types: []string{"FRR", "KSI"},
	})
	fmt.Printf("Document filtered to FRR and KSI only\n")
	fmt.Printf("- FRR categories: %d\n", len(rulesAndKSI.FRR))
	fmt.Printf("- KSI indicators: %d\n", len(rulesAndKSI.KSI))
	fmt.Printf("- FRD definitions: %d (filtered out)\n", len(rulesAndKSI.FRD))

	fmt.Println("\n✅ Demo complete! Check generated files:")
	fmt.Println("- ksi-evidence-template.json")
	fmt.Println("- frmr-low-ksi.md")
}

func fetchDocument(name, url string) *frmr.FRMRDocument {
	fmt.Printf("Fetching %s...", name)
	
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch %s: %v", name, err)
	}
	defer resp.Body.Close()

	doc, err := frmr.ParseFRMR(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse %s: %v", name, err)
	}

	fmt.Printf(" ✓ (%s v%s)\n", doc.Info.ShortName, doc.Info.CurrentRelease)
	return doc
}

func validateDocument(name string, doc *frmr.FRMRDocument) {
	errors := frmr.ValidateSchema(doc)
	if len(errors) == 0 {
		fmt.Printf("✅ %s: Valid\n", name)
	} else {
		fmt.Printf("❌ %s: %d errors\n", name, len(errors))
		for _, err := range errors[:3] { // Show first 3 errors
			fmt.Printf("   - %s\n", err)
		}
		if len(errors) > 3 {
			fmt.Printf("   ... and %d more\n", len(errors)-3)
		}
	}
}

func countKSIRequirements(doc *frmr.FRMRDocument) int {
	count := 0
	for _, ksi := range doc.KSI {
		count += len(ksi.Requirements)
	}
	return count
}

func saveEvidenceTemplate(template map[string]interface{}, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Failed to create evidence template: %v", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(template); err != nil {
		log.Printf("Failed to write evidence template: %v", err)
		return
	}

	fmt.Printf("Evidence template saved to %s\n", filename)
}

func saveMarkdown(content, filename string) {
	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		log.Printf("Failed to save markdown: %v", err)
		return
	}
	fmt.Printf("Markdown saved to %s\n", filename)
}

func createSampleEvidence(doc *frmr.FRMRDocument) map[string]bool {
	evidence := make(map[string]bool)
	
	// Simulate partial compliance
	compliantPatterns := []string{
		"KSI-IAM-01", "KSI-IAM-02", "KSI-IAM-03", // IAM mostly compliant
		"KSI-MLA-01", "KSI-MLA-02", // Some monitoring
		"KSI-CED-01", // Basic education
		"KSI-PIY-01", "KSI-PIY-02", // Basic policies
	}
	
	// Set specific requirements as met
	for _, pattern := range compliantPatterns {
		evidence[pattern] = true
	}
	
	// Add some random compliance (every 3rd requirement)
	i := 0
	for _, ksi := range doc.KSI {
		for _, req := range ksi.Requirements {
			if i%3 == 0 && !evidence[req.ID] {
				evidence[req.ID] = true
			}
			i++
		}
	}
	
	return evidence
} 