// +build !skippackr

package bundled

import "github.com/markbates/pkger"

func init() {
	// Include template files
	pkger.Include("/bundled/templates/FedRAMP-SSP-Low-Baseline-Template.docx")
	pkger.Include("/bundled/templates/FedRAMP-SSP-Moderate-Baseline-Template.docx")
	pkger.Include("/bundled/templates/FedRAMP-SSP-High-Baseline-Template.docx")
	pkger.Include("/bundled/templates/FedRAMP-SSP-OSCAL-Template.xml")
	
	// Include catalog files
	pkger.Include("/bundled/catalogs/FedRAMP_LOW-baseline-resolved-profile_catalog.xml")
	pkger.Include("/bundled/catalogs/FedRAMP_MODERATE-baseline-resolved-profile_catalog.xml")
	pkger.Include("/bundled/catalogs/FedRAMP_HIGH-baseline-resolved-profile_catalog.xml")
} 