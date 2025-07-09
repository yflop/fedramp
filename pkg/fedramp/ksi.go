package fedramp

import (
	"encoding/json"
	"fmt"
	"time"
)

// KSIValidationStatus represents the validation status of a KSI
type KSIValidationStatus string

const (
	KSIStatusTrue    KSIValidationStatus = "True"
	KSIStatusFalse   KSIValidationStatus = "False"
	KSIStatusPartial KSIValidationStatus = "Partial"
)

// KSIEvidence represents supporting evidence for a KSI validation
type KSIEvidence struct {
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Reference   string    `json:"reference"`
	Timestamp   time.Time `json:"timestamp"`
	Source      string    `json:"source"`
}

// KSIValidation represents a single KSI validation result
type KSIValidation struct {
	ID               string              `json:"id"`
	Name             string              `json:"name"`
	Category         string              `json:"category"`
	Status           KSIValidationStatus `json:"status"`
	Evidence         []KSIEvidence       `json:"evidence"`
	AutomatedCheck   bool                `json:"automated_check"`
	LastValidated    time.Time           `json:"last_validated"`
	ValidationMethod string              `json:"validation_method"`
	RelatedControls  []string            `json:"related_controls"`
	ThreePAOAttested bool                `json:"3pao_attested"`
	Notes            string              `json:"notes,omitempty"`
}

// KSIReport represents a complete KSI validation report
type KSIReport struct {
	ServiceOfferingID string                    `json:"service_offering_id"`
	ReportID          string                    `json:"report_id"`
	GeneratedAt       time.Time                 `json:"generated_at"`
	Validations       map[string]*KSIValidation `json:"validations"`
	Summary           KSISummary                `json:"summary"`
	DataSchema        string                    `json:"data_schema"`
	Version           string                    `json:"version"`
}

// KSISummary provides summary statistics for KSI validations
type KSISummary struct {
	TotalKSIs          int                            `json:"total_ksis"`
	ValidationsByStatus map[KSIValidationStatus]int    `json:"validations_by_status"`
	AutomatedCount     int                            `json:"automated_count"`
	ManualCount        int                            `json:"manual_count"`
	ComplianceScore    float64                        `json:"compliance_score"`
}

// KSIDefinitions for FedRAMP 20x Phase One
var KSIDefinitions = map[string]KSIDefinition{
	"KSI-CNA": {
		ID:       "KSI-CNA",
		Name:     "Cloud Native Architecture",
		Category: "Architecture",
		Description: "A secure cloud service offering will use cloud native architecture and design principles",
		ValidationPoints: []string{
			"DoS protection",
			"Firewall/proxy configuration",
			"Immutable containers/serverless",
			"Micro-services design",
			"Virtual network controls",
			"Continuous scanning",
			"High availability design",
		},
		RelatedControls: []string{"SC-5", "SC-7", "SC-12", "SC-39", "SR-12"},
	},
	"KSI-SC": {
		ID:       "KSI-SC",
		Name:     "Service Configuration",
		Category: "Configuration",
		Description: "Enforce approved cryptography, verify integrity, restrict external services",
		ValidationPoints: []string{
			"Hardened configurations",
			"Encrypted network traffic",
			"Encrypted data at rest",
			"Central configuration management",
			"Cryptographic integrity",
			"Key rotation",
			"Security patching",
		},
		RelatedControls: []string{"CM-2", "CM-4", "CM-8", "IA-7", "RA-7", "SC-8", "SC-8(1)", "SC-13", "SC-28", "SC-28(1)", "SI-3", "SI-4"},
	},
	"KSI-IAM": {
		ID:       "KSI-IAM",
		Name:     "Identity and Access Management",
		Category: "Access Control",
		Description: "Protect user data, control access, implement zero trust",
		ValidationPoints: []string{
			"Phishing-resistant MFA",
			"Strong passwords",
			"Secure API authentication",
			"Least-privileged access",
		},
		RelatedControls: []string{"AC-2", "AC-3", "AU-9", "AC-14", "IA-2", "IA-2(1)", "IA-2(2)", "IA-2(8)", "IA-2(12)", "IA-4", "IA-5", "IA-5(1)", "IA-6", "IA-8", "IA-8(1)", "IA-8(2)", "IA-8(4)", "IA-11", "PS-2", "PS-3", "PS-4", "PS-5", "PS-7", "PS-9"},
	},
	"KSI-MLA": {
		ID:       "KSI-MLA",
		Name:     "Monitoring, Logging, and Auditing",
		Category: "Monitoring",
		Description: "Monitor, log, and audit all important events",
		ValidationPoints: []string{
			"SIEM system",
			"Log review and audit",
			"Vulnerability detection",
			"Authenticated scanning",
			"IaC scanning",
			"Vulnerability tracking",
		},
		RelatedControls: []string{"AC-7", "AU-2", "AU-3", "AU-4", "AU-8", "AU-11", "AU-12", "RA-5", "SI-2"},
	},
	"KSI-CM": {
		ID:       "KSI-CM",
		Name:     "Change Management",
		Category: "Change Control",
		Description: "Ensure all system changes are properly documented",
		ValidationPoints: []string{
			"Change logging",
			"Immutable deployments",
			"Automated testing",
			"Change procedures",
			"Risk evaluation",
		},
		RelatedControls: []string{"CM-6", "CM-7", "CM-10", "CM-11"},
	},
	"KSI-PI": {
		ID:       "KSI-PI",
		Name:     "Policy and Inventory",
		Category: "Governance",
		Description: "Organized guidance for securing all assets",
		ValidationPoints: []string{
			"Asset inventory",
			"Security policies",
			"Vulnerability disclosure",
			"Secure SDLC",
			"Automated evaluation",
			"Security budget",
		},
		RelatedControls: []string{"AC-1", "AU-1", "CA-1", "CM-1", "CM-8", "CP-1", "IA-1", "IR-1", "PL-1", "PL-2", "PS-1", "RA-1", "SA-1", "SA-2", "SA-3", "SA-5", "SA-8", "SC-1", "SI-1", "SR-1"},
	},
	"KSI-3IR": {
		ID:       "KSI-3IR",
		Name:     "Third Party Information Resources",
		Category: "Supply Chain",
		Description: "Understand, monitor, and manage supply chain risks",
		ValidationPoints: []string{
			"FedRAMP authorization verification",
			"Supply chain risk identification",
			"SBOM collection",
			"CISA attestation",
			"Zero trust implementation",
		},
		RelatedControls: []string{"AC-2", "AC-20", "AC-23", "CA-3", "CA-9", "RA-3(1)", "SA-4", "SA-9", "SA-22", "SI-5", "SR-2", "SR-2(1)", "SR-3", "SR-5", "SR-8", "SR-10", "SR-11", "SR-11(2)"},
	},
	"KSI-CE": {
		ID:       "KSI-CE",
		Name:     "Cybersecurity Education",
		Category: "Training",
		Description: "Continuously educate employees on cybersecurity",
		ValidationPoints: []string{
			"Security awareness training",
			"Role-specific training",
		},
		RelatedControls: []string{"AT-2", "AT-3", "AT-6"},
	},
	"KSI-IR": {
		ID:       "KSI-IR",
		Name:     "Incident Response",
		Category: "Incident Management",
		Description: "Maintain, test, and execute effective Incident Response Plans",
		ValidationPoints: []string{
			"RTO/RPO definition",
			"System backups",
			"Recovery testing",
			"Incident reporting",
			"Incident logging",
			"MTTD/MTTR metrics",
		},
		RelatedControls: []string{"CP-2", "CP-4", "CP-9", "CP-10", "IR-4", "IR-5", "IR-6", "IR-7", "IR-8", "PS-8", "RA-3", "RA-5(2)", "RA-5(11)"},
	},
}

// KSIDefinition defines the structure of a KSI
type KSIDefinition struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Category         string   `json:"category"`
	Description      string   `json:"description"`
	ValidationPoints []string `json:"validation_points"`
	RelatedControls  []string `json:"related_controls"`
}

// NewKSIReport creates a new KSI validation report
func NewKSIReport(serviceID string) *KSIReport {
	return &KSIReport{
		ServiceOfferingID: serviceID,
		ReportID:          fmt.Sprintf("KSI-%s-%s", serviceID, time.Now().Format("20060102-150405")),
		GeneratedAt:       time.Now(),
		Validations:       make(map[string]*KSIValidation),
		Version:           "1.0.0",
		DataSchema:        "https://github.com/FedRAMP/docs/blob/main/FRMR.KSI.key-security-indicators.json",
	}
}

// AddValidation adds a KSI validation to the report
func (r *KSIReport) AddValidation(validation *KSIValidation) {
	r.Validations[validation.ID] = validation
	r.updateSummary()
}

// updateSummary recalculates the summary statistics
func (r *KSIReport) updateSummary() {
	r.Summary.TotalKSIs = len(r.Validations)
	r.Summary.ValidationsByStatus = make(map[KSIValidationStatus]int)
	r.Summary.AutomatedCount = 0
	r.Summary.ManualCount = 0
	
	trueCount := 0
	
	for _, v := range r.Validations {
		r.Summary.ValidationsByStatus[v.Status]++
		if v.AutomatedCheck {
			r.Summary.AutomatedCount++
		} else {
			r.Summary.ManualCount++
		}
		if v.Status == KSIStatusTrue {
			trueCount++
		}
	}
	
	if r.Summary.TotalKSIs > 0 {
		r.Summary.ComplianceScore = float64(trueCount) / float64(r.Summary.TotalKSIs) * 100
	}
}

// ToJSON exports the KSI report as JSON
func (r *KSIReport) ToJSON() ([]byte, error) {
	return json.MarshalIndent(r, "", "  ")
}

// ValidateKSI performs validation for a specific KSI
func ValidateKSI(ksiID string, evidence []KSIEvidence, automated bool) *KSIValidation {
	def, exists := KSIDefinitions[ksiID]
	if !exists {
		return nil
	}
	
	validation := &KSIValidation{
		ID:               ksiID,
		Name:             def.Name,
		Category:         def.Category,
		Evidence:         evidence,
		AutomatedCheck:   automated,
		LastValidated:    time.Now(),
		RelatedControls:  def.RelatedControls,
		ThreePAOAttested: false,
	}
	
	// Simple validation logic - in practice this would be more sophisticated
	if len(evidence) >= len(def.ValidationPoints)/2 {
		validation.Status = KSIStatusTrue
	} else if len(evidence) > 0 {
		validation.Status = KSIStatusPartial
	} else {
		validation.Status = KSIStatusFalse
	}
	
	if automated {
		validation.ValidationMethod = "automated"
	} else {
		validation.ValidationMethod = "manual"
	}
	
	return validation
} 