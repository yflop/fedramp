package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocomply/fedramp/pkg/fedramp/frmr"
)

func main() {
	fmt.Println("FedRAMP Machine Readable (FRMR) Demo")
	fmt.Println("=====================================")

	// Fetch and process KSI document
	ksiDoc, err := fetchFRMRDocument("https://raw.githubusercontent.com/FedRAMP/docs/main/FRMR.KSI.key-security-indicators.json")
	if err != nil {
		log.Fatalf("Failed to fetch KSI document: %v", err)
	}

	fmt.Printf("\nDocument: %s\n", ksiDoc.Info.Name)
	fmt.Printf("Current Release: %s\n", ksiDoc.Info.CurrentRelease)
	fmt.Printf("Number of KSIs: %d\n\n", len(ksiDoc.KSI))

	// Create sample evidence for validation
	evidence := generateSampleEvidence(ksiDoc)

	// Validate each KSI
	fmt.Println("KSI Validation Results:")
	fmt.Println("=======================")

	totalKSIs := 0
	metKSIs := 0
	totalRequirements := 0
	metRequirements := 0

	for _, ksi := range ksiDoc.KSI {
		result := ksi.ValidateKSI(evidence)
		totalKSIs++
		totalRequirements += result.TotalCount
		metRequirements += result.MetCount
		
		if result.FullyMet {
			metKSIs++
		}

		status := "❌"
		if result.FullyMet {
			status = "✅"
		}

		fmt.Printf("\n%s %s: %s\n", status, result.KSIID, result.KSIName)
		fmt.Printf("   Requirements: %d/%d met (%.1f%%)\n", 
			result.MetCount, result.TotalCount, 
			float64(result.MetCount)/float64(result.TotalCount)*100)

		// Show first few unmet requirements
		unmetShown := 0
		for _, req := range result.Requirements {
			if !req.Met && unmetShown < 3 {
				fmt.Printf("   ❌ %s: %s\n", req.ID, truncate(req.Statement, 60))
				unmetShown++
			}
		}
		if result.UnmetCount > 3 {
			fmt.Printf("   ... and %d more unmet requirements\n", result.UnmetCount-3)
		}
	}

	fmt.Printf("\n\nOverall Summary:\n")
	fmt.Printf("================\n")
	fmt.Printf("KSIs Fully Met: %d/%d (%.1f%%)\n", 
		metKSIs, totalKSIs, float64(metKSIs)/float64(totalKSIs)*100)
	fmt.Printf("Total Requirements Met: %d/%d (%.1f%%)\n", 
		metRequirements, totalRequirements, 
		float64(metRequirements)/float64(totalRequirements)*100)

	// Export validation report
	exportValidationReport(ksiDoc, evidence)

	// Demonstrate SCN document parsing
	fmt.Println("\n\nFetching SCN Document...")
	scnDoc, err := fetchFRMRDocument("https://raw.githubusercontent.com/FedRAMP/docs/main/FRMR.SCN.significant-change-notifications.json")
	if err != nil {
		log.Printf("Failed to fetch SCN document: %v", err)
	} else {
		fmt.Printf("SCN Document: %s\n", scnDoc.Info.Name)
		fmt.Printf("Current Release: %s\n", scnDoc.Info.CurrentRelease)
		
		// Show SCN definitions
		if scnDefs, ok := scnDoc.FRD["SCN"]; ok {
			fmt.Printf("\nSCN Definitions:\n")
			for _, def := range scnDefs {
				fmt.Printf("- %s (%s): %s\n", def.Term, def.ID, truncate(def.Definition, 60))
			}
		}
	}
}

func fetchFRMRDocument(url string) (*frmr.FRMRDocument, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	return frmr.ParseFRMR(resp.Body)
}

func generateSampleEvidence(doc *frmr.FRMRDocument) map[string]bool {
	evidence := make(map[string]bool)
	
	// Simulate partial compliance - some requirements met, some not
	// This represents a realistic scenario where a CSP has implemented
	// some but not all security requirements
	
	metPatterns := []string{
		"KSI-CNA-01", "KSI-CNA-02", "KSI-CNA-03", // Some CNA requirements
		"KSI-SVC-01", "KSI-SVC-02", // Some service config
		"KSI-IAM-01", "KSI-IAM-02", "KSI-IAM-03", // Most IAM
		"KSI-MLA-01", "KSI-MLA-02", // Some monitoring
		"KSI-CMT-01", "KSI-CMT-04", // Some change management
		"KSI-PIY-01", "KSI-PIY-02", // Basic policies
		"KSI-CED-01", // Basic education
	}

	// Mark specific requirements as met
	for _, pattern := range metPatterns {
		evidence[pattern] = true
	}

	// Add some additional random requirements to simulate ~40% compliance
	for _, ksi := range doc.KSI {
		for i, req := range ksi.Requirements {
			// Meet every 3rd requirement not already set
			if i%3 == 0 && !evidence[req.ID] {
				evidence[req.ID] = true
			}
		}
	}

	return evidence
}

func truncate(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length] + "..."
}

func exportValidationReport(doc *frmr.FRMRDocument, evidence map[string]bool) {
	report := struct {
		DocumentInfo struct {
			Name           string `json:"name"`
			CurrentRelease string `json:"current_release"`
			ValidatedAt    string `json:"validated_at"`
		} `json:"document_info"`
		ValidationResults []struct {
			KSIID        string `json:"ksi_id"`
			KSIName      string `json:"ksi_name"`
			FullyMet     bool   `json:"fully_met"`
			MetCount     int    `json:"met_count"`
			TotalCount   int    `json:"total_count"`
			Percentage   float64 `json:"percentage"`
			UnmetReqs    []string `json:"unmet_requirements"`
		} `json:"validation_results"`
		Summary struct {
			TotalKSIs           int     `json:"total_ksis"`
			FullyMetKSIs        int     `json:"fully_met_ksis"`
			TotalRequirements   int     `json:"total_requirements"`
			MetRequirements     int     `json:"met_requirements"`
			OverallPercentage   float64 `json:"overall_percentage"`
		} `json:"summary"`
	}{}

	report.DocumentInfo.Name = doc.Info.Name
	report.DocumentInfo.CurrentRelease = doc.Info.CurrentRelease
	report.DocumentInfo.ValidatedAt = "2025-01-17T12:00:00Z"

	totalReqs := 0
	metReqs := 0
	fullyMetKSIs := 0

	for _, ksi := range doc.KSI {
		result := ksi.ValidateKSI(evidence)
		
		valResult := struct {
			KSIID        string `json:"ksi_id"`
			KSIName      string `json:"ksi_name"`
			FullyMet     bool   `json:"fully_met"`
			MetCount     int    `json:"met_count"`
			TotalCount   int    `json:"total_count"`
			Percentage   float64 `json:"percentage"`
			UnmetReqs    []string `json:"unmet_requirements"`
		}{
			KSIID:      result.KSIID,
			KSIName:    result.KSIName,
			FullyMet:   result.FullyMet,
			MetCount:   result.MetCount,
			TotalCount: result.TotalCount,
			Percentage: float64(result.MetCount) / float64(result.TotalCount) * 100,
			UnmetReqs:  []string{},
		}

		for _, req := range result.Requirements {
			if !req.Met {
				valResult.UnmetReqs = append(valResult.UnmetReqs, req.ID)
			}
		}

		report.ValidationResults = append(report.ValidationResults, valResult)
		
		totalReqs += result.TotalCount
		metReqs += result.MetCount
		if result.FullyMet {
			fullyMetKSIs++
		}
	}

	report.Summary.TotalKSIs = len(doc.KSI)
	report.Summary.FullyMetKSIs = fullyMetKSIs
	report.Summary.TotalRequirements = totalReqs
	report.Summary.MetRequirements = metReqs
	report.Summary.OverallPercentage = float64(metReqs) / float64(totalReqs) * 100

	// Write report
	file, err := os.Create("frmr-validation-report.json")
	if err != nil {
		log.Printf("Failed to create report file: %v", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(report); err != nil {
		log.Printf("Failed to write report: %v", err)
		return
	}

	fmt.Println("\nValidation report saved to: frmr-validation-report.json")
} 