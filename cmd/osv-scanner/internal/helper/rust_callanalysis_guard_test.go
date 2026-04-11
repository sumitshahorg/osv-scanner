package helper

import "testing"

func TestValidateRustCallAnalysisAcknowledgement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name               string
		enabled            []string
		disabled           []string
		allowUnsafeRust    bool
		expectErr          bool
	}{
		{
			name:            "default settings do not require acknowledgement",
			enabled:         []string{},
			disabled:        []string{},
			allowUnsafeRust: false,
			expectErr:       false,
		},
		{
			name:            "explicit rust enable without acknowledgement fails",
			enabled:         []string{"rust"},
			disabled:        []string{},
			allowUnsafeRust: false,
			expectErr:       true,
		},
		{
			name:            "explicit rust enable with acknowledgement succeeds",
			enabled:         []string{"rust"},
			disabled:        []string{},
			allowUnsafeRust: true,
			expectErr:       false,
		},
		{
			name:            "enable all without acknowledgement fails because rust is included",
			enabled:         []string{"all"},
			disabled:        []string{},
			allowUnsafeRust: false,
			expectErr:       true,
		},
		{
			name:            "enable all then disable rust does not require acknowledgement",
			enabled:         []string{"all"},
			disabled:        []string{"rust"},
			allowUnsafeRust: false,
			expectErr:       false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := ValidateRustCallAnalysisAcknowledgement(tc.enabled, tc.disabled, tc.allowUnsafeRust)
			if tc.expectErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tc.expectErr && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}