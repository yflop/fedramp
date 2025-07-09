package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocomply/fedramp/pkg/fedramp"
)

func main() {
	fmt.Println("FedRAMP R5.SSAD - Storing and Sharing Authorization Data Demo")
	fmt.Println("============================================================")

	// Create SSAD package metadata
	serviceID := "CS-DEMO-2024"
	metadata := fedramp.SSADMetadata{
		Title:             "CloudNative Platform - FedRAMP Authorization Package",
		Description:       "Complete FedRAMP Moderate authorization documentation for CloudNative Platform v2.0",
		AuthorizationType: "Agency",
		ImpactLevel:       "Moderate",
		AuthorizationDate: time.Now().AddDate(0, -1, 0), // Authorized 1 month ago
		ExpirationDate:    time.Now().AddDate(3, -1, 0), // Expires in 3 years
		CSPName:           "CloudNative Inc",
		PackageFormat:     "OSCAL",
		Tags:              []string{"saas", "cloud", "moderate", "agency-ato"},
		Keywords:          []string{"cloudnative", "platform", "kubernetes", "containers"},
		CustomMetadata: map[string]string{
			"agency":         "Department of Examples",
			"system_version": "2.0.3",
			"oscal_version":  "1.1.2",
		},
	}

	// Create package
	pkg := fedramp.NewSSADPackage(serviceID, metadata)
	pkg.AccessControl.Owner = "CloudNative Inc"
	pkg.AccessControl.DataClassification = "Controlled Unclassified Information (CUI)"
	pkg.AccessControl.AccessRestrictions = []string{
		"Must have valid NDA",
		"US persons only",
		"No export without approval",
	}

	// Add SSP document
	fmt.Println("\nAdding authorization documents...")
	
	sspDoc := fedramp.SSADDocument{
		DocumentID:   "DOC-SSP-001",
		Title:        "CloudNative Platform System Security Plan",
		Type:         "ssp",
		Format:       "JSON",
		Version:      "2.0.3",
		CreatedDate:  time.Now().AddDate(0, -6, 0),
		LastModified: time.Now().AddDate(0, -1, -15),
		Author:       "CloudNative Security Team",
		Size:         2457600, // 2.4MB
		Location:     "packages/cloudnative/ssp-v2.0.3.json",
		AccessLevel:  "restricted",
		Metadata: map[string]string{
			"controls_implemented": "325",
			"oscal_version":        "1.1.2",
			"validation_status":    "passed",
		},
	}
	pkg.AddDocument("ssp", sspDoc)

	// Add SAP document
	sapDoc := fedramp.SSADDocument{
		DocumentID:   "DOC-SAP-001",
		Title:        "Security Assessment Plan - CloudNative Platform",
		Type:         "sap",
		Format:       "JSON",
		Version:      "1.2",
		CreatedDate:  time.Now().AddDate(0, -4, 0),
		LastModified: time.Now().AddDate(0, -3, -20),
		Author:       "SecureAssess Partners LLC",
		Size:         876543,
		Location:     "packages/cloudnative/sap-v1.2.json",
		AccessLevel:  "restricted",
		Metadata: map[string]string{
			"3pao":              "SecureAssess Partners LLC",
			"assessment_scope":  "full",
		},
	}
	pkg.AddDocument("sap", sapDoc)

	// Add SAR document
	sarDoc := fedramp.SSADDocument{
		DocumentID:   "DOC-SAR-001",
		Title:        "Security Assessment Report - CloudNative Platform",
		Type:         "sar",
		Format:       "JSON",
		Version:      "1.0",
		CreatedDate:  time.Now().AddDate(0, -2, 0),
		LastModified: time.Now().AddDate(0, -1, -5),
		Author:       "SecureAssess Partners LLC",
		Size:         3456789,
		Location:     "packages/cloudnative/sar-v1.0.json",
		AccessLevel:  "restricted",
		Metadata: map[string]string{
			"findings_total":      "12",
			"findings_high":       "2",
			"findings_moderate":   "7",
			"findings_low":        "3",
		},
	}
	pkg.AddDocument("sar", sarDoc)

	// Add POA&M document
	poamDoc := fedramp.SSADDocument{
		DocumentID:   "DOC-POAM-001",
		Title:        "Plan of Action and Milestones",
		Type:         "poam",
		Format:       "JSON",
		Version:      "1.3",
		CreatedDate:  time.Now().AddDate(0, -1, -5),
		LastModified: time.Now().AddDate(0, 0, -2),
		Author:       "CloudNative Security Team",
		Size:         234567,
		Location:     "packages/cloudnative/poam-v1.3.json",
		AccessLevel:  "restricted",
		Metadata: map[string]string{
			"open_items":        "8",
			"overdue_items":     "0",
			"next_review_date":  time.Now().AddDate(0, 1, 0).Format("2006-01-02"),
		},
	}
	pkg.AddDocument("poam", poamDoc)

	// Add continuous monitoring reports
	for i := 0; i < 3; i++ {
		conmonDoc := fedramp.SSADDocument{
			DocumentID:   fmt.Sprintf("DOC-CONMON-%03d", i+1),
			Title:        fmt.Sprintf("Monthly ConMon Report - %s", time.Now().AddDate(0, -i, 0).Format("January 2006")),
			Type:         "conmon",
			Format:       "JSON",
			Version:      "1.0",
			CreatedDate:  time.Now().AddDate(0, -i, -5),
			LastModified: time.Now().AddDate(0, -i, -3),
			Author:       "CloudNative Security Operations",
			Size:         567890 + int64(i*10000),
			Location:     fmt.Sprintf("packages/cloudnative/conmon/%s.json", time.Now().AddDate(0, -i, 0).Format("2006-01")),
			AccessLevel:  "restricted",
			Metadata: map[string]string{
				"vulnerabilities_found": fmt.Sprintf("%d", 15-i*3),
				"patches_applied":       fmt.Sprintf("%d", 42+i*5),
				"incidents":             "0",
			},
		}
		pkg.AddDocument("conmon", conmonDoc)
	}

	// Add KSI reports (for 20x)
	ksiDoc := fedramp.SSADDocument{
		DocumentID:   "DOC-KSI-001",
		Title:        "Key Security Indicators Report - FedRAMP 20x",
		Type:         "ksi",
		Format:       "JSON",
		Version:      "1.0",
		CreatedDate:  time.Now().AddDate(0, 0, -7),
		LastModified: time.Now().AddDate(0, 0, -1),
		Author:       "CloudNative Compliance Team",
		Size:         123456,
		Location:     "packages/cloudnative/ksi-report.json",
		AccessLevel:  "restricted",
		Metadata: map[string]string{
			"ksi_validated":     "9",
			"compliance_score":  "88.9%",
			"automation_level":  "high",
		},
	}
	pkg.AddDocument("ksi", ksiDoc)

	// Share with entities
	fmt.Println("Setting up access controls...")

	// Share with FedRAMP PMO
	pkg.ShareWith(fedramp.SSADSharee{
		EntityID:       "ENT-FEDRAMP-PMO",
		EntityType:     "agency",
		Name:           "FedRAMP PMO",
		Email:          "info@fedramp.gov",
		AccessLevel:    "read",
		AccessConditions: []string{
			"Official use only",
			"No redistribution without approval",
		},
	})

	// Share with authorizing agency
	expirationDate := time.Now().AddDate(1, 0, 0)
	pkg.ShareWith(fedramp.SSADSharee{
		EntityID:       "ENT-DOE-001",
		EntityType:     "agency",
		Name:           "Department of Examples - CISO Office",
		Email:          "ciso@examples.gov",
		AccessLevel:    "admin",
		ExpirationDate: &expirationDate,
		AccessConditions: []string{
			"Authorizing agency full access",
		},
	})

	// Share with 3PAO for annual assessment
	assessmentExpiry := time.Now().AddDate(0, 3, 0)
	pkg.ShareWith(fedramp.SSADSharee{
		EntityID:       "ENT-3PAO-SA",
		EntityType:     "3pao",
		Name:           "SecureAssess Partners LLC",
		Email:          "assessments@secureassess.example",
		AccessLevel:    "read",
		ExpirationDate: &assessmentExpiry,
		AccessConditions: []string{
			"For assessment purposes only",
			"Must maintain confidentiality",
		},
	})

	// Log distribution events
	fmt.Println("Logging distribution events...")

	pkg.LogDistribution("FedRAMP PMO", "Initial authorization package submission", "api")
	pkg.LogDistribution("Department of Examples", "Agency review and authorization", "portal")
	pkg.LogDistribution("SecureAssess Partners", "Annual assessment preparation", "email")

	// Finalize package
	fmt.Println("Finalizing package...")
	
	if err := pkg.Finalize("John Smith, Authorizing Official"); err != nil {
		log.Fatalf("Error finalizing package: %v", err)
	}

	// Create repository and add package
	repo := fedramp.NewSSADRepository("fedramp-central", "central")
	if err := repo.AddPackage(pkg); err != nil {
		log.Fatalf("Error adding package to repository: %v", err)
	}

	// Export package
	data, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		log.Fatalf("Error generating JSON: %v", err)
	}

	filename := fmt.Sprintf("ssad-package-%s.json", time.Now().Format("20060102"))
	if err := os.WriteFile(filename, data, 0644); err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	// Display summary
	fmt.Println("\nSSAD Package Summary:")
	fmt.Println("====================")
	fmt.Printf("Package ID: %s\n", pkg.PackageID)
	fmt.Printf("Service: %s\n", pkg.ServiceOfferingID)
	fmt.Printf("CSP: %s\n", pkg.Metadata.CSPName)
	fmt.Printf("Impact Level: %s\n", pkg.Metadata.ImpactLevel)
	fmt.Printf("Authorization Type: %s\n", pkg.Metadata.AuthorizationType)
	fmt.Printf("Status: %s\n", pkg.Status)
	fmt.Printf("Package Hash: %s\n", pkg.IntegrityCheck.PackageHash)

	fmt.Println("\nDocuments included:")
	if pkg.Components.SSP != nil {
		fmt.Printf("  - SSP: %s (v%s)\n", pkg.Components.SSP.Title, pkg.Components.SSP.Version)
	}
	if pkg.Components.SAP != nil {
		fmt.Printf("  - SAP: %s (v%s)\n", pkg.Components.SAP.Title, pkg.Components.SAP.Version)
	}
	if pkg.Components.SAR != nil {
		fmt.Printf("  - SAR: %s (v%s)\n", pkg.Components.SAR.Title, pkg.Components.SAR.Version)
	}
	if pkg.Components.POAM != nil {
		fmt.Printf("  - POA&M: %s (v%s)\n", pkg.Components.POAM.Title, pkg.Components.POAM.Version)
	}
	fmt.Printf("  - ConMon Reports: %d\n", len(pkg.Components.ConMon))
	fmt.Printf("  - KSI Reports: %d\n", len(pkg.Components.KSIReports))

	fmt.Println("\nAccess Control:")
	fmt.Printf("  Owner: %s\n", pkg.AccessControl.Owner)
	fmt.Printf("  Classification: %s\n", pkg.AccessControl.DataClassification)
	fmt.Printf("  Shared with: %d entities\n", len(pkg.AccessControl.SharedWith))
	for _, sharee := range pkg.AccessControl.SharedWith {
		fmt.Printf("    - %s (%s): %s access\n", sharee.Name, sharee.EntityType, sharee.AccessLevel)
	}

	fmt.Println("\nDistribution Log:")
	for _, dist := range pkg.DistributionLog {
		fmt.Printf("  - %s: %s via %s\n", dist.RecipientName, dist.Purpose, dist.Method)
	}

	fmt.Printf("\nPackage exported to: %s\n", filename)

	// Demonstrate repository search
	fmt.Println("\nRepository Search Demo:")
	fmt.Println("======================")
	
	searchCriteria := map[string]string{
		"impact_level": "Moderate",
		"csp":          "CloudNative Inc",
	}
	
	results := repo.Search(searchCriteria)
	fmt.Printf("Found %d packages matching criteria\n", len(results))
	for _, result := range results {
		fmt.Printf("  - %s: %s (%s)\n", result.PackageID, result.Metadata.Title, result.Metadata.ImpactLevel)
	}
} 