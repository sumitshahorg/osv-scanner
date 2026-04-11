package helper

import "fmt"

const AllowUnsafeRustCallAnalysisFlag = "allow-unsafe-rust-call-analysis"

// ValidateRustCallAnalysisAcknowledgement enforces an explicit acknowledgement
// before enabling Rust call analysis, because it executes cargo build scripts.
func ValidateRustCallAnalysisAcknowledgement(enabledCallAnalysis, disabledCallAnalysis []string, allowUnsafeRustCallAnalysis bool) error {
	callAnalysisStates := CreateCallAnalysisStates(enabledCallAnalysis, disabledCallAnalysis)

	if callAnalysisStates["rust"] && !allowUnsafeRustCallAnalysis {
		return fmt.Errorf("rust call analysis can execute arbitrary code through Cargo build scripts (build.rs); rerun with --%s only for trusted projects or in a sandbox", AllowUnsafeRustCallAnalysisFlag)
	}

	return nil
}