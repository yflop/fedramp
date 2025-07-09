package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocomply/fedramp/pkg/fedramp"
)

func main() {
	fmt.Println("FedRAMP 25.05C Key Security Indicators Demo")
	fmt.Println("==========================================")
	fmt.Println()

	// Create KSI report
	serviceID := "CS-DEMO-2025"
	report := fedramp.NewKSIReport(serviceID)

	// Demonstrate validation for each of the 11 KSIs
	
	// 1. KSI-CED: Cybersecurity Education
	fmt.Println("Validating KSI-CED: Cybersecurity Education")
	cedEvidence := []fedramp.KSIEvidence{
		{
			Type:        "KSI-CED-01",
			Description: "All employees completed annual security awareness training",
			Reference:   "training-records-2025",
			Source:      "HR Training System",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-CED-02",
			Description: "Role-specific training completed for all privileged users",
			Reference:   "privileged-training-2025",
			Source:      "Security Training Platform",
			Timestamp:   time.Now(),
		},
	}
	report.AddValidation(fedramp.ValidateKSI("KSI-CED", cedEvidence, true))

	// 2. KSI-CMT: Change Management
	fmt.Println("Validating KSI-CMT: Change Management")
	cmtEvidence := []fedramp.KSIEvidence{
		{
			Type:        "KSI-CMT-01",
			Description: "All changes logged in JIRA with automated monitoring",
			Reference:   "jira-change-log",
			Source:      "JIRA API",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-CMT-02",
			Description: "GitOps deployment with immutable containers",
			Reference:   "k8s-deployments",
			Source:      "Kubernetes API",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-CMT-03",
			Description: "CI/CD pipeline with automated testing",
			Reference:   "github-actions",
			Source:      "GitHub Actions",
			Timestamp:   time.Now(),
		},
	}
	report.AddValidation(fedramp.ValidateKSI("KSI-CMT", cmtEvidence, true))

	// 3. KSI-CNA: Cloud Native Architecture
	fmt.Println("Validating KSI-CNA: Cloud Native Architecture")
	cnaEvidence := []fedramp.KSIEvidence{
		{
			Type:        "KSI-CNA-01",
			Description: "Network policies configured for all pods",
			Reference:   "k8s-netpol",
			Source:      "Kubernetes NetworkPolicy",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-CNA-04",
			Description: "All deployments use immutable containers",
			Reference:   "container-policy",
			Source:      "OPA Policy Engine",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-CNA-05",
			Description: "AWS Shield and WAF configured",
			Reference:   "aws-shield-config",
			Source:      "AWS API",
			Timestamp:   time.Now(),
		},
	}
	report.AddValidation(fedramp.ValidateKSI("KSI-CNA", cnaEvidence, true))

	// 4. KSI-IAM: Identity and Access Management
	fmt.Println("Validating KSI-IAM: Identity and Access Management")
	iamEvidence := []fedramp.KSIEvidence{
		{
			Type:        "KSI-IAM-01",
			Description: "FIDO2 WebAuthn enforced for all users",
			Reference:   "auth0-config",
			Source:      "Auth0 Management API",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-IAM-04",
			Description: "RBAC with least privilege implemented",
			Reference:   "iam-policies",
			Source:      "AWS IAM",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-IAM-05",
			Description: "Zero Trust architecture with BeyondCorp",
			Reference:   "zt-architecture",
			Source:      "Architecture Docs",
			Timestamp:   time.Now(),
		},
	}
	report.AddValidation(fedramp.ValidateKSI("KSI-IAM", iamEvidence, true))

	// 5. KSI-INR: Incident Reporting
	fmt.Println("Validating KSI-INR: Incident Reporting")
	inrEvidence := []fedramp.KSIEvidence{
		{
			Type:        "KSI-INR-01",
			Description: "Automated incident reporting to FedRAMP",
			Reference:   "incident-api",
			Source:      "SOAR Platform",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-INR-02",
			Description: "Incident database with pattern analysis",
			Reference:   "incident-db",
			Source:      "ServiceNow",
			Timestamp:   time.Now(),
		},
	}
	report.AddValidation(fedramp.ValidateKSI("KSI-INR", inrEvidence, true))

	// 6. KSI-MLA: Monitoring, Logging, and Auditing
	fmt.Println("Validating KSI-MLA: Monitoring, Logging, and Auditing")
	mlaEvidence := []fedramp.KSIEvidence{
		{
			Type:        "KSI-MLA-01",
			Description: "Splunk SIEM with tamper-resistant logging",
			Reference:   "splunk-config",
			Source:      "Splunk API",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-MLA-04",
			Description: "Weekly authenticated vulnerability scans",
			Reference:   "qualys-scans",
			Source:      "Qualys API",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-MLA-05",
			Description: "Terraform security scanning in CI/CD",
			Reference:   "tfsec-reports",
			Source:      "GitHub Actions",
			Timestamp:   time.Now(),
		},
	}
	report.AddValidation(fedramp.ValidateKSI("KSI-MLA", mlaEvidence, true))

	// 7. KSI-PIY: Policy and Inventory
	fmt.Println("Validating KSI-PIY: Policy and Inventory")
	piyEvidence := []fedramp.KSIEvidence{
		{
			Type:        "KSI-PIY-01",
			Description: "Terraform state as complete inventory",
			Reference:   "terraform-state",
			Source:      "Terraform Cloud",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-PIY-04",
			Description: "SSDLC with security gates implemented",
			Reference:   "sdlc-policy",
			Source:      "Policy Docs",
			Timestamp:   time.Now(),
		},
	}
	report.AddValidation(fedramp.ValidateKSI("KSI-PIY", piyEvidence, false))

	// 8. KSI-RPL: Recovery Planning
	fmt.Println("Validating KSI-RPL: Recovery Planning")
	rplEvidence := []fedramp.KSIEvidence{
		{
			Type:        "KSI-RPL-01",
			Description: "RTO: 4 hours, RPO: 1 hour defined",
			Reference:   "dr-plan",
			Source:      "DR Documentation",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-RPL-03",
			Description: "Automated hourly backups to S3",
			Reference:   "backup-config",
			Source:      "AWS Backup",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-RPL-04",
			Description: "Quarterly DR tests performed",
			Reference:   "dr-test-results",
			Source:      "Test Reports",
			Timestamp:   time.Now(),
		},
	}
	report.AddValidation(fedramp.ValidateKSI("KSI-RPL", rplEvidence, true))

	// 9. KSI-SVC: Service Configuration
	fmt.Println("Validating KSI-SVC: Service Configuration")
	svcEvidence := []fedramp.KSIEvidence{
		{
			Type:        "KSI-SVC-02",
			Description: "TLS 1.3 enforced for all traffic",
			Reference:   "tls-config",
			Source:      "Load Balancer Config",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-SVC-03",
			Description: "AES-256 encryption at rest",
			Reference:   "kms-config",
			Source:      "AWS KMS",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-SVC-06",
			Description: "Automated key rotation every 90 days",
			Reference:   "key-rotation-policy",
			Source:      "AWS KMS",
			Timestamp:   time.Now(),
		},
	}
	report.AddValidation(fedramp.ValidateKSI("KSI-SVC", svcEvidence, true))

	// 10. KSI-TPR: Third-Party Information Resources
	fmt.Println("Validating KSI-TPR: Third-Party Information Resources")
	tprEvidence := []fedramp.KSIEvidence{
		{
			Type:        "KSI-TPR-01",
			Description: "Complete SBOM maintained",
			Reference:   "sbom-inventory",
			Source:      "Dependency Track",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-TPR-02",
			Description: "All cloud services are FedRAMP authorized",
			Reference:   "fedramp-auth-check",
			Source:      "Compliance Dashboard",
			Timestamp:   time.Now(),
		},
		{
			Type:        "KSI-TPR-04",
			Description: "Snyk monitoring for vulnerabilities",
			Reference:   "snyk-integration",
			Source:      "Snyk API",
			Timestamp:   time.Now(),
		},
	}
	report.AddValidation(fedramp.ValidateKSI("KSI-TPR", tprEvidence, true))

	// Export report
	data, err := report.ToJSON()
	if err != nil {
		log.Fatalf("Error generating report: %v", err)
	}

	filename := fmt.Sprintf("ksi-v25-report-%s.json", time.Now().Format("20060102"))
	if err := os.WriteFile(filename, data, 0644); err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	// Display summary
	fmt.Println("\nKSI Validation Summary (FedRAMP 25.05C)")
	fmt.Println("========================================")
	fmt.Printf("Service ID: %s\n", report.ServiceOfferingID)
	fmt.Printf("Report ID: %s\n", report.ReportID)
	fmt.Printf("Generated: %s\n", report.GeneratedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Schema: %s\n", report.DataSchema)
	fmt.Println("\nValidation Results:")
	fmt.Printf("  Total KSIs: %d\n", report.Summary.TotalKSIs)
	fmt.Printf("  Compliance Score: %.1f%%\n", report.Summary.ComplianceScore)
	fmt.Printf("  Automated Validations: %d\n", report.Summary.AutomatedCount)
	fmt.Printf("  Manual Validations: %d\n", report.Summary.ManualCount)
	
	fmt.Println("\nStatus by KSI:")
	for _, v := range report.Validations {
		fmt.Printf("  %s (%s): %s\n", v.ID, v.Name, v.Status)
	}

	fmt.Printf("\nReport saved to: %s\n", filename)
	
	// Show compliance recommendations
	fmt.Println("\nCompliance Recommendations:")
	fmt.Println("===========================")
	for _, v := range report.Validations {
		if v.Status != fedramp.KSIStatusTrue {
			fmt.Printf("\n%s - %s (Status: %s)\n", v.ID, v.Name, v.Status)
			def := fedramp.KSIDefinitions[v.ID]
			fmt.Println("Missing validation points:")
			
			// Check which points are missing
			evidenceTypes := make(map[string]bool)
			for _, e := range v.Evidence {
				evidenceTypes[e.Type] = true
			}
			
			for _, point := range def.ValidationPoints {
				pointID := point[:10]
				if !evidenceTypes[pointID] {
					fmt.Printf("  ❌ %s\n", point)
				}
			}
		}
	}
} 