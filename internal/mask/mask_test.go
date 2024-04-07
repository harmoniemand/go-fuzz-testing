package mask_test

import (
	"errors"
	"testing"

	"github.com/harmoniemand/go-fuzz-testing/internal/mask"
)

func TestIfMaskEmailWorks(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		email    string
		expected string
		err      error
	}{
		{
			name:     "mask email",
			email:    "max.mustermann@example.com",
			expected: "max***********@*******.com",
			err:      nil,
		},
		{
			name:     "mask empty string",
			email:    "",
			expected: "",
			err:      errors.New("email is empty"),
		}, {
			name:     "mask faulty email",
			email:    "max.mustermannexample.com",
			expected: "",
			err:      errors.New("email faulty, no @ found"),
		}, {
			name:     "mask faulty domain",
			email:    "max.mustermann@com",
			expected: "",
			err:      errors.New("email faulty, domain faulty"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			masked, err := mask.MaskEmail(tt.email)

			if err != nil {
				if tt.err == nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if tt.err.Error() != err.Error() {
					t.Fatalf("expected error %v, got %v", tt.err, err)
				}
			}

			if masked != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, masked)
			}
		})
	}
}

func FuzzMask(f *testing.F) {
	f.Add("max.mustermann@example.com")

	f.Fuzz(func(t *testing.T, s string) {
		mask.MaskEmail(s)
	})
}
