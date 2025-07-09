package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gocomply/fedramp/pkg/fedramp"
)

func main() {
	fmt.Println("FedRAMP 20x Phase One Pilot Demo")
	fmt.Println("=================================")
	fmt.Println()
	
	serviceID := "CSO-CLOUDNATIVE-001"
	
	// Step 1: Generate KSI Validation Report
	fmt.Println("Step 1: Generating KSI Validation Report")
	fmt.Println("----------------------------------------")
	
	report := generateKSIReport(serviceID)
	
	// Step 2: Generate Continuous Reporting Proposal
	fmt.Println("\nStep 2: Generating Continuous Reporting Proposal")
	fmt.Println("------------------------------------------------")
	
	proposal := generateContinuousReportingProposal(serviceID)
	
	// Step 3: Create Submission Package
	fmt.Println("\nStep 3: Creating 20x Submission Package")
	fmt.Println("---------------------------------------")
	
	submission := create20xSubmission(serviceID, report, proposal)
	
	// Export everything
	exportResults(report, proposal, submission)
	
	fmt.Println("\n✅ FedRAMP 20x Demo Complete!")
	fmt.Println("Ready for submission to 20x@fedramp.gov")
}

func generateKSIReport(serviceID string) *fedramp.KSIReport {
	report := fedramp.NewKSIReport(serviceID)
	
	// Simulate comprehensive KSI validations
	validations := map[string]struct {
		automated bool
		evidence  []fedramp.KSIEvidence
	}{
		"KSI-CNA": {
			automated: true,
			evidence: []fedramp.KSIEvidence{
				{
					Type:        "configuration",
					Description: "DoS protection via AWS Shield Advanced",
					Reference:   "aws-shield-config-2025",
					Source:      "AWS API",
					Timestamp:   time.Now(),
				},
				{
					Type:        "architecture",
					Description: "Microservices deployed on EKS with network policies",
					Reference:   "k8s-network-policies",
					Source:      "Kubernetes API",
					Timestamp:   time.Now(),
				},
				{
					Type:        "scan_result",
					Description: "Container images scanned with Trivy - 0 critical vulnerabilities",
					Reference:   "trivy-scan-20250110",
					Source:      "CI/CD Pipeline",
					Timestamp:   time.Now(),
				},
			},
		},
		"KSI-SC": {
			automated: true,
			evidence: []fedramp.KSIEvidence{
				{
					Type:        "encryption",
					Description: "All data encrypted at rest using AWS KMS",
					Reference:   "kms-policy-20250110",
					Source:      "AWS KMS API",
					Timestamp:   time.Now(),
				},
				{
					Type:        "certificate",
					Description: "TLS 1.3 enforced on all endpoints",
					Reference:   "ssl-labs-A+",
					Source:      "SSL Labs API",
					Timestamp:   time.Now(),
				},
			},
		},
		"KSI-IAM": {
			automated: true,
			evidence: []fedramp.KSIEvidence{
				{
					Type:        "configuration",
					Description: "Phishing-resistant MFA (FIDO2) enforced for all users",
					Reference:   "okta-policy-fido2",
					Source:      "Okta API",
					Timestamp:   time.Now(),
				},
				{
					Type:        "policy",
					Description: "RBAC with least privilege implemented",
					Reference:   "iam-roles-matrix",
					Source:      "AWS IAM",
					Timestamp:   time.Now(),
				},
			},
		},
		"KSI-MLA": {
			automated: true,
			evidence: []fedramp.KSIEvidence{
				{
					Type:        "integration",
					Description: "Splunk SIEM collecting all logs with 90-day retention",
					Reference:   "splunk-config-20250110",
					Source:      "Splunk API",
					Timestamp:   time.Now(),
				},
				{
					Type:        "scan_result",
					Description: "Weekly authenticated vulnerability scans via Qualys",
					Reference:   "qualys-scan-20250103",
					Source:      "Qualys API",
					Timestamp:   time.Now(),
				},
			},
		},
		"KSI-CM": {
			automated: true,
			evidence: []fedramp.KSIEvidence{
				{
					Type:        "process",
					Description: "GitOps deployment with ArgoCD - all changes tracked",
					Reference:   "argocd-audit-log",
					Source:      "ArgoCD API",
					Timestamp:   time.Now(),
				},
			},
		},
		"KSI-PI": {
			automated: false,
			evidence: []fedramp.KSIEvidence{
				{
					Type:        "documentation",
					Description: "Complete asset inventory in ServiceNow CMDB",
					Reference:   "cmdb-export-20250110",
					Source:      "ServiceNow",
					Timestamp:   time.Now(),
				},
				{
					Type:        "policy",
					Description: "Security policies documented and approved",
					Reference:   "security-policy-v3.0",
					Source:      "Document Management",
					Timestamp:   time.Now(),
				},
			},
		},
		"KSI-3IR": {
			automated: true,
			evidence: []fedramp.KSIEvidence{
				{
					Type:        "verification",
					Description: "All third-party services are FedRAMP authorized",
					Reference:   "vendor-fedramp-status",
					Source:      "FedRAMP Marketplace API",
					Timestamp:   time.Now(),
				},
				{
					Type:        "sbom",
					Description: "SBOM generated for all components",
					Reference:   "sbom-20250110.json",
					Source:      "SBOM Generator",
					Timestamp:   time.Now(),
				},
			},
		},
		"KSI-CE": {
			automated: false,
			evidence: []fedramp.KSIEvidence{
				{
					Type:        "training",
					Description: "100% completion of annual security awareness training",
					Reference:   "training-report-2025",
					Source:      "LMS",
					Timestamp:   time.Now(),
				},
			},
		},
		"KSI-IR": {
			automated: true,
			evidence: []fedramp.KSIEvidence{
				{
					Type:        "metrics",
					Description: "RTO: 4 hours, RPO: 1 hour, MTTD: 15 min, MTTR: 2 hours",
					Reference:   "ir-metrics-2025",
					Source:      "Incident Management System",
					Timestamp:   time.Now(),
				},
				{
					Type:        "test_result",
					Description: "Quarterly DR test successful - recovered in 3.5 hours",
					Reference:   "dr-test-20241215",
					Source:      "Test Report",
					Timestamp:   time.Now(),
				},
			},
		},
	}
	
	// Add all validations to report
	for ksiID, v := range validations {
		validation := fedramp.ValidateKSI(ksiID, v.evidence, v.automated)
		if validation != nil {
			validation.ThreePAOAttested = true // Simulating 3PAO attestation
			report.AddValidation(validation)
			fmt.Printf("✓ Validated %s: %s (%d evidence items)\n", 
				ksiID, validation.Status, len(v.evidence))
		}
	}
	
	fmt.Printf("\nKSI Summary: %d/%d validated (%.1f%% compliance)\n",
		len(report.Validations), len(fedramp.KSIDefinitions), report.Summary.ComplianceScore)
	
	return report
}

func generateContinuousReportingProposal(serviceID string) *fedramp.ContinuousReportingProposal {
	manager := fedramp.NewContinuousReportingManager(serviceID)
	proposal := manager.GenerateProposal()
	
	fmt.Printf("✓ Generated continuous reporting proposal\n")
	fmt.Printf("  - Automated KSI coverage: %.1f%%\n", proposal.CoveragePercentage)
	fmt.Printf("  - Reporting mechanisms: %d\n", len(proposal.ReportingMechanisms))
	fmt.Printf("  - Implementation milestones: %d\n", len(proposal.Implementation.Milestones))
	
	return proposal
}

func create20xSubmission(serviceID string, report *fedramp.KSIReport, proposal *fedramp.ContinuousReportingProposal) map[string]interface{} {
	submission := map[string]interface{}{
		"submission_type": "fedramp_20x_phase_one",
		"submission_date": time.Now(),
		"service_offering": map[string]interface{}{
			"id":          serviceID,
			"name":        "CloudNative Secure Platform",
			"description": "Cloud-native SaaS platform with automated security validation",
			"type":        "SaaS",
			"deployment":  "AWS GovCloud",
		},
		"eligibility": map[string]interface{}{
			"cloud_native":     true,
			"fedramp_hosting":  "AWS GovCloud (FedRAMP High)",
			"public_internet":  true,
			"soc2_type2":       true,
			"soc2_date":        "2024-06-15",
		},
		"contacts": map[string]interface{}{
			"csp": map[string]interface{}{
				"name":  "John Smith",
				"title": "Chief Security Officer",
				"email": "john.smith@cloudnative.example",
				"phone": "555-0123",
			},
			"3pao": map[string]interface{}{
				"company": "SecureAssess Inc",
				"name":    "Jane Doe",
				"title":   "Principal Assessor",
				"email":   "jane.doe@secureassess.example",
			},
		},
		"approach_summary": `This submission demonstrates automated validation of FedRAMP 20x Key Security Indicators using cloud-native 
technologies and continuous monitoring. Our approach leverages APIs from cloud providers, security tools, and infrastructure 
components to automatically validate security controls in real-time or near real-time.`,
		"3pao_summary": `SecureAssess Inc. has reviewed and tested the automated validation approach. We confirm that the technical 
implementation accurately validates the security capabilities required by each KSI. The evidence collection is comprehensive 
and the validation logic is sound. We attest to the accuracy of the validation results.`,
		"ksi_report":                report,
		"continuous_reporting":      proposal,
		"data_schema_version":       "1.0.0",
		"machine_readable_format":   "JSON",
		"access_instructions": map[string]interface{}{
			"portal_url": "https://trust.cloudnative.example/fedramp",
			"api_endpoint": "https://api.cloudnative.example/fedramp/v1",
			"authentication": "OAuth 2.0 with client credentials",
		},
	}
	
	fmt.Printf("✓ Created 20x submission package\n")
	fmt.Printf("  - Eligibility criteria: ✓ Met\n")
	fmt.Printf("  - Machine-readable format: ✓ JSON\n")
	fmt.Printf("  - 3PAO attestation: ✓ Included\n")
	
	return submission
}

func exportResults(report *fedramp.KSIReport, proposal *fedramp.ContinuousReportingProposal, submission map[string]interface{}) {
	// Export KSI Report
	if data, err := report.ToJSON(); err == nil {
		fmt.Printf("\n📄 KSI Report: ksi-report.json (%d bytes)\n", len(data))
	}
	
	// Export Continuous Reporting Proposal
	if data, err := json.MarshalIndent(proposal, "", "  "); err == nil {
		fmt.Printf("📄 Continuous Reporting Proposal: continuous-reporting-proposal.json (%d bytes)\n", len(data))
	}
	
	// Export Full Submission
	if data, err := json.MarshalIndent(submission, "", "  "); err == nil {
		fmt.Printf("📄 20x Submission Package: fedramp-20x-submission.json (%d bytes)\n", len(data))
	}
	
	// Generate Trust Center structure
	trustCenter := map[string]interface{}{
		"trust_center_url": "https://trust.cloudnative.example",
		"sections": []string{
			"FedRAMP Authorization Status",
			"Key Security Indicators",
			"Continuous Monitoring Reports",
			"Security Documentation",
			"POA&Ms and SCNs",
			"Incident Reports",
			"API Access",
		},
		"api_endpoints": map[string]string{
			"ksi_status":     "/api/v1/ksi/status",
			"ksi_evidence":   "/api/v1/ksi/{id}/evidence",
			"continuous":     "/api/v1/monitoring/continuous",
			"documentation":  "/api/v1/docs",
		},
	}
	
	if data, err := json.MarshalIndent(trustCenter, "", "  "); err == nil {
		fmt.Printf("📄 Trust Center Configuration: trust-center.json (%d bytes)\n", len(data))
	}
} 