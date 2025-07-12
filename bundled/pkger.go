package bundled

import (
	"github.com/markbates/pkger"
)

// Include directives for pkger
// These comments tell pkger which files to include in the build
func init() {
	// Include template files
	_ = pkger.Include("/bundled/templates/FedRAMP-SSP-Low-Baseline-Template.docx")
	_ = pkger.Include("/bundled/templates/FedRAMP-SSP-Moderate-Baseline-Template.docx")
	_ = pkger.Include("/bundled/templates/FedRAMP-SSP-High-Baseline-Template.docx")
	_ = pkger.Include("/bundled/templates/FedRAMP-SSP-OSCAL-Template.xml")
	
	// Include catalog files
	_ = pkger.Include("/bundled/catalogs/FedRAMP_LOW-baseline-resolved-profile_catalog.xml")
	_ = pkger.Include("/bundled/catalogs/FedRAMP_MODERATE-baseline-resolved-profile_catalog.xml")
	_ = pkger.Include("/bundled/catalogs/FedRAMP_HIGH-baseline-resolved-profile_catalog.xml")
} 