package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocomply/fedramp/pkg/fedramp"
)

func main() {
	fmt.Println("FedRAMP R5 Balance Initiatives Demo")
	fmt.Println("===================================")
	
	// Demo 1: Significant Change Notification (R5.SCN)
	fmt.Println("\n1. R5.SCN - Significant Change Notification Demo")
	demoSCN()
	
	// Demo 2: Continuous Reporting Standard (R5.CRS)
	fmt.Println("\n2. R5.CRS - Continuous Reporting Standard Demo")
	demoCRS()
	
	fmt.Println("\n✅ R5 Balance Initiatives Demo Complete!")
	fmt.Println("These tools support FedRAMP's modernization from compliance-based to security-based assessment.")
}

func demoSCN() {
	// Create a new SCN for an adaptive change
	scn := fedramp.NewSCN(
		"CSO-12345", 
		"security-patch", 
		"Apply critical security patches to web servers", 
		"Address CVE-2024-1234 vulnerability")
	
	// Add details
	scn.AddAffectedControl("SI-2")
	scn.AddAffectedControl("RA-5")
	scn.AddAffectedComponent("web-server-cluster")
	scn.ApproverName = "Jane Smith"
	scn.ApproverTitle = "CISO"
	
	// Set change date for adaptive change
	changeDate := time.Now().AddDate(0, 0, 1) // Tomorrow
	scn.DateOfChange = &changeDate
	scn.VerificationSteps = "Automated vulnerability scan post-patching"
	
	// Classify the SCN type
	if err := scn.ClassifySCNType(); err != nil {
		log.Printf("Error classifying SCN: %v", err)
		return
	}
	
	// Validate the SCN
	if err := scn.ValidateForSubmission(); err != nil {
		log.Printf("SCN validation failed: %v", err)
		return
	}
	
	// Export to JSON (machine-readable format requirement)
	jsonData, err := scn.ToJSON()
	if err != nil {
		log.Printf("Error exporting SCN: %v", err)
		return
	}
	
	fmt.Printf("✅ SCN Created: Type=%s, Service=%s\n", scn.SCNType, scn.ServiceOfferingID)
	fmt.Printf("   Affected Controls: %v\n", scn.ControlsAffected)
	fmt.Printf("   JSON Size: %d bytes\n", len(jsonData))
	
	// Demo SCN Manager
	manager := fedramp.NewSCNManager()
	manager.AddNotification("scn-001", scn)
	
	report := manager.GenerateSCNReport()
	fmt.Printf("   SCN Report: %d total notifications\n", report["total_notifications"])
}

func demoCRS() {
	// Create CRS Manager
	crsManager := fedramp.NewCRSManager()
	
	// Create a reporting period (monthly)
	period := fedramp.ReportingPeriod{
		StartDate: time.Now().AddDate(0, -1, 0), // Last month
		EndDate:   time.Now(),
		Type:      "monthly",
	}
	
	// Create a continuous monitoring report
	report := crsManager.CreateReport("CSO-12345", period)
	
	// Generate standard FedRAMP metrics
	standardMetrics := crsManager.GenerateStandardMetrics()
	
	// Add metrics to the report
	for _, metric := range standardMetrics {
		if err := crsManager.AddMetric(report.ReportID, metric); err != nil {
			log.Printf("Error adding metric: %v", err)
			continue
		}
	}
	
	// Add attestation
	report.Attestation = fedramp.ReportAttestation{
		AttestorName:    "John Doe",
		AttestorTitle:   "System Owner",
		AttestorEmail:   "john.doe@example.com",
		AttestationDate: time.Now(),
		Statement:       "I attest that the information in this report is accurate and complete.",
	}
	
	// Validate the report
	if err := crsManager.ValidateReport(report.ReportID); err != nil {
		log.Printf("Report validation failed: %v", err)
		return
	}
	
	// Export the report
	reportJSON, err := crsManager.ExportReport(report.ReportID)
	if err != nil {
		log.Printf("Error exporting report: %v", err)
		return
	}
	
	fmt.Printf("✅ Continuous Report Created: ID=%s\n", report.ReportID)
	fmt.Printf("   Metrics: %d total\n", report.Summary.TotalMetrics)
	fmt.Printf("   Status Breakdown: Green=%d, Yellow=%d, Red=%d\n", 
		report.Summary.MetricsByStatus[fedramp.MetricStatusGreen],
		report.Summary.MetricsByStatus[fedramp.MetricStatusYellow],
		report.Summary.MetricsByStatus[fedramp.MetricStatusRed])
	fmt.Printf("   JSON Size: %d bytes\n", len(reportJSON))
	
	// Generate dashboard data
	dashboard := crsManager.GenerateDashboardData("CSO-12345")
	fmt.Printf("   Dashboard Status: %s\n", dashboard["overall_status"])
} 