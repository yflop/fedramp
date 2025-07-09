# FedRAMP R5 Balance Implementations

This repository now includes initial implementations for **FedRAMP's R5 Balance Improvement Releases (BIRs)** as outlined in the [FedRAMP roadmap](https://github.com/FedRAMP/roadmap/issues?q=type:%22R5%20Balance%22).

## 🚀 What We've Built

### 1. R5.SCN - Significant Change Notification Support
**Status**: ✅ Initial Implementation Complete  
**GitHub Issue**: [#16](https://github.com/FedRAMP/roadmap/issues/16)

**Features:**
- Complete SCN data structures following RFC-0007 specifications
- Automatic classification of change types (Adaptive, Transformative, Impact Change)
- JSON export for machine-readable format requirements
- Validation engine for submission readiness
- CLI tools for SCN creation, validation, and export

**Key Files:**
- `pkg/fedramp/scn.go` - Core SCN data structures and logic
- `cli/cmd/scn.go` - CLI commands for SCN operations

**Usage Example:**
```bash
# Create a new SCN
gocomply_fedramp scn create CSO-12345 "security-patch" "Apply critical patches" "CVE remediation" \
  --affected-controls SI-2,RA-5 \
  --approver-name "Jane Smith" \
  --approver-title "CISO"

# Validate an SCN
gocomply_fedramp scn validate scn.json

# Export summary
gocomply_fedramp scn export summary scn.json summary.txt
```

### 2. R5.CRS - Continuous Reporting Standard Support  
**Status**: ✅ Initial Implementation Complete  
**GitHub Issue**: [#23](https://github.com/FedRAMP/roadmap/issues/23)

**Features:**
- Key Security Metrics (KSM) data structures
- Standard FedRAMP metrics (vulnerability scanning, patch compliance, MFA coverage, etc.)
- Continuous monitoring report generation
- Dashboard data aggregation
- Automated trend analysis framework

**Key Files:**
- `pkg/fedramp/crs.go` - Continuous Reporting Standard implementation

**Standard Metrics Included:**
- Vulnerability Scan Coverage
- Security Patch Compliance  
- Failed Login Attempts
- Backup Success Rate
- Data Encryption Coverage
- Multi-Factor Authentication Coverage

## 🛠️ Technical Implementation

### Architecture Principles
These implementations follow FedRAMP's modernization goals:

1. **Machine-Readable First**: All outputs are JSON-based for automation
2. **Security-Based Assessment**: Focus on actual security metrics vs. compliance documentation
3. **Continuous Monitoring**: Real-time data streams replace periodic reports
4. **Automation-Friendly**: Designed for integration with CI/CD and monitoring tools

### Integration Points
- **OSCAL Compatibility**: Works with existing OSCAL SSP processing
- **FedRAMP Baselines**: Integrates with Low/Moderate/High baseline catalogs
- **CLI Integration**: Extends existing gocomply_fedramp command structure

## 🧪 Demo & Testing

Run the demonstration:
```bash
go run examples/r5_balance_demo.go
```

This demo shows:
- SCN creation and validation workflow
- Continuous reporting with standard metrics
- JSON export capabilities
- Integration between SCN and CRS systems

## 📈 R5 Balance Timeline Support

Based on the [roadmap issue #16](https://github.com/FedRAMP/roadmap/issues/16):

- ✅ **Preparation Phase** (June 2-27): Core implementations ready
- 🔄 **Closed Beta** (July 7 - August 29): Ready for beta testing
- 🔄 **Open Beta** (September 1 - October 31): Community feedback integration
- 🔄 **Wide Release** (November 3): Production readiness

## 🤝 Contributing to R5 Balance

### Immediate Contribution Opportunities:

1. **Beta Testing**: Test the SCN and CRS implementations with real-world scenarios
2. **Integration Examples**: Build connectors to popular monitoring tools
3. **Metric Expansion**: Add additional Key Security Metrics
4. **Validation Enhancement**: Improve SCN classification logic
5. **Dashboard Development**: Create visualization tools for CRS data

### Next Steps:

1. **R5.MAS** (Minimum Assessment Standard) - Issue [#19](https://github.com/FedRAMP/roadmap/issues/19)
2. **R5.SSAD** (Storing and Sharing Authorization Data) - Issue [#28](https://github.com/FedRAMP/roadmap/issues/28)

## 🔗 Related Resources

- [RFC-0007 Significant Change Notification Standard](https://www.fedramp.gov/rfcs/0007/)
- [RFC-0008 Continuous Reporting Standard](https://www.fedramp.gov/rfcs/0008/)
- [FedRAMP 20x Phase One Key Security Indicators](https://www.fedramp.gov/rfcs/0006/)
- [FedRAMP Roadmap Repository](https://github.com/FedRAMP/roadmap)

## 💡 Why This Matters

These R5 Balance implementations represent FedRAMP's shift from:
- **Manual → Automated** assessment processes
- **Compliance → Security** focused evaluation  
- **Point-in-time → Continuous** monitoring
- **Narrative → Machine-readable** documentation

By contributing to these implementations, you're helping modernize federal cloud security assessment for the entire government.

---

**Ready to contribute?** Join the [Rev5 Community Working Group](https://github.com/FedRAMP/roadmap/issues/16) and help shape the future of FedRAMP automation! 