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
	fmt.Println("FedRAMP R5.MAS - Minimum Assessment Standard Demo")
	fmt.Println("=================================================")

	// Create initial assessment
	serviceID := "CS-DEMO-2024"
	assessment := fedramp.NewMASAssessment(serviceID, fedramp.MASInitial)

	// Configure 3PAO
	assessment.ThreePAO = fedramp.AssessmentOrganization{
		Name:              "SecureAssess Partners LLC",
		A2LAAccreditation: "R311-2024-DEMO",
		LeadAssessor:      "Sarah Johnson, CISSP",
		TeamMembers:       []string{"Mike Chen, CISA", "Emily Davis, PMP", "Robert Wilson, CEH"},
		ContactEmail:      "assessments@secureassess.example",
	}

	// Define scope
	assessment.Scope = fedramp.AssessmentScope{
		FullAssessment:     true,
		Locations:          []string{"AWS us-east-1", "AWS us-west-2", "On-premise DC1"},
		DataTypes:          []string{"CUI", "PII", "PHI", "Financial"},
		UserPopulation:     25000,
		ExcludedComponents: []string{"Legacy System A (decommissioning)"},
	}

	// Add assessment methods
	fmt.Println("\nAdding assessment methods...")

	// Document examination
	assessment.AddMethod(fedramp.AssessmentMethod{
		MethodType:      "examine",
		Description:     "Review of security policies, procedures, and system documentation",
		Automated:       false,
		StartTime:       time.Now().AddDate(0, 0, -30),
		EndTime:         time.Now().AddDate(0, 0, -28),
		ControlsCovered: []string{"AC-1", "AC-2", "AU-1", "CA-1", "CM-1"},
	})

	// Interviews
	assessment.AddMethod(fedramp.AssessmentMethod{
		MethodType:      "interview",
		Description:     "Interviews with system administrators and security personnel",
		Automated:       false,
		StartTime:       time.Now().AddDate(0, 0, -25),
		EndTime:         time.Now().AddDate(0, 0, -23),
		ControlsCovered: []string{"AC-2", "AC-3", "AU-2", "IR-1", "PE-1"},
	})

	// Automated vulnerability scanning
	assessment.AddMethod(fedramp.AssessmentMethod{
		MethodType:      "test",
		Description:     "Automated vulnerability scanning of all system components",
		Automated:       true,
		ToolsUsed:       []string{"Nessus Professional", "Qualys VMDR"},
		StartTime:       time.Now().AddDate(0, 0, -20),
		EndTime:         time.Now().AddDate(0, 0, -19),
		ControlsCovered: []string{"RA-5", "SI-2", "CM-6", "SC-7"},
	})

	// Penetration testing
	assessment.AddMethod(fedramp.AssessmentMethod{
		MethodType:      "test",
		Description:     "Application and infrastructure penetration testing",
		Automated:       false,
		ToolsUsed:       []string{"Burp Suite Pro", "Metasploit", "Custom scripts"},
		StartTime:       time.Now().AddDate(0, 0, -15),
		EndTime:         time.Now().AddDate(0, 0, -10),
		ControlsCovered: []string{"CA-8", "SC-7", "SI-3", "AC-4"},
	})

	// Add evidence
	fmt.Println("Collecting evidence...")

	evidence := []fedramp.AssessmentEvidence{
		{
			EvidenceID:      "EVD-001",
			Type:            "screenshots",
			Description:     "Access control configuration screenshots",
			CollectionDate:  time.Now().AddDate(0, 0, -25),
			CollectedBy:     "Mike Chen",
			Location:        "evidence/access-control/",
			ControlsCovered: []string{"AC-2", "AC-3", "AC-6"},
			Automated:       false,
		},
		{
			EvidenceID:      "EVD-002",
			Type:            "scan_reports",
			Description:     "Vulnerability scan reports for all subnets",
			CollectionDate:  time.Now().AddDate(0, 0, -19),
			CollectedBy:     "Automated Scanner",
			Location:        "evidence/scans/",
			ControlsCovered: []string{"RA-5", "SI-2"},
			Automated:       true,
		},
		{
			EvidenceID:      "EVD-003",
			Type:            "test_results",
			Description:     "Penetration test results and remediation evidence",
			CollectionDate:  time.Now().AddDate(0, 0, -10),
			CollectedBy:     "Robert Wilson",
			Location:        "evidence/pentest/",
			ControlsCovered: []string{"CA-8"},
			Automated:       false,
		},
		{
			EvidenceID:      "EVD-004",
			Type:            "policies",
			Description:     "Security policies and procedures documentation",
			CollectionDate:  time.Now().AddDate(0, 0, -28),
			CollectedBy:     "Emily Davis",
			Location:        "evidence/policies/",
			ControlsCovered: []string{"AC-1", "AU-1", "CA-1", "CM-1"},
			Automated:       false,
		},
		{
			EvidenceID:      "EVD-005",
			Type:            "configurations",
			Description:     "System configuration baselines and hardening evidence",
			CollectionDate:  time.Now().AddDate(0, 0, -22),
			CollectedBy:     "Mike Chen",
			Location:        "evidence/configs/",
			ControlsCovered: []string{"CM-2", "CM-6", "SC-7"},
			Automated:       false,
		},
	}

	for _, e := range evidence {
		assessment.AddEvidence(e)
	}

	// Add findings
	fmt.Println("Recording findings...")

	findings := []fedramp.AssessmentFinding{
		{
			FindingID:      "FIND-001",
			ControlID:      "AC-2",
			Severity:       "moderate",
			Status:         "remediated",
			Description:    "Service accounts were not reviewed within required timeframe",
			Recommendation: "Implement automated quarterly service account review process",
			Evidence:       []string{"EVD-001"},
			DateIdentified: time.Now().AddDate(0, 0, -25),
			DateRemediated: &[]time.Time{time.Now().AddDate(0, 0, -5)}[0],
		},
		{
			FindingID:      "FIND-002",
			ControlID:      "SI-2",
			Severity:       "high",
			Status:         "open",
			Description:    "Critical patches not applied within required timeframe on 3 servers",
			Recommendation: "Apply patches immediately and implement automated patch management",
			Evidence:       []string{"EVD-002"},
			DateIdentified: time.Now().AddDate(0, 0, -19),
		},
		{
			FindingID:      "FIND-003",
			ControlID:      "AU-4",
			Severity:       "low",
			Status:         "risk_accepted",
			Description:    "Log storage capacity warning threshold set at 85% instead of 75%",
			Recommendation: "Adjust monitoring threshold to 75% per policy",
			Evidence:       []string{"EVD-001", "EVD-005"},
			DateIdentified: time.Now().AddDate(0, 0, -22),
		},
	}

	for _, f := range findings {
		assessment.AddFinding(f)
	}

	// Complete assessment
	fmt.Println("Completing assessment...")

	assessment.Attestation = &fedramp.AssessmentAttestation{
		AttestorName:  "Sarah Johnson",
		AttestorTitle: "Lead Assessor, CISSP",
		Organization:  assessment.ThreePAO.Name,
		Date:          time.Now(),
		Statement:     "I attest that this assessment was conducted in accordance with FedRAMP Minimum Assessment Standards and all findings have been accurately documented.",
		Signature:     fmt.Sprintf("SJ-%s", time.Now().Format("20060102")),
	}

	assessment.Complete()

	// Validate completeness
	if err := assessment.ValidateCompleteness(); err != nil {
		log.Printf("Warning: %v\n", err)
	}

	// Generate summary
	summary := assessment.GenerateSummary()
	fmt.Println("\nAssessment Summary:")
	fmt.Println("==================")
	fmt.Printf("Assessment ID: %s\n", summary["assessment_id"])
	fmt.Printf("Type: %s\n", summary["type"])
	fmt.Printf("Status: %s\n", summary["status"])
	fmt.Printf("Duration: %s\n", summary["duration"])
	fmt.Printf("3PAO: %s\n", summary["3pao"])
	fmt.Printf("Total Findings: %d\n", summary["total_findings"])
	fmt.Printf("Evidence Collected: %d\n", summary["evidence_collected"])

	// Export assessment
	data, err := assessment.ToJSON()
	if err != nil {
		log.Fatalf("Error generating JSON: %v", err)
	}

	filename := fmt.Sprintf("mas-assessment-%s.json", time.Now().Format("20060102"))
	if err := os.WriteFile(filename, data, 0644); err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	fmt.Printf("\nAssessment exported to: %s\n", filename)

	// Show MAS compliance
	fmt.Println("\nMAS Compliance Check:")
	fmt.Println("====================")
	requirements := assessment.GetRequirements()
	for _, req := range requirements {
		fmt.Printf("\n%s:\n", req.Name)
		fmt.Printf("  Description: %s\n", req.Description)
		fmt.Printf("  Required Methods: %v ✓\n", req.RequiredMethods)
		fmt.Printf("  Minimum Evidence: %v ✓\n", req.MinimumEvidence)
		fmt.Printf("  Automation Level: %s\n", req.AutomationLevel)
	}

	// Pretty print the assessment
	fmt.Println("\nFull Assessment Details:")
	fmt.Println("=======================")
	prettyJSON, _ := json.MarshalIndent(assessment, "", "  ")
	fmt.Println(string(prettyJSON))
} 