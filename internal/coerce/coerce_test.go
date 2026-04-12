package coerce

import (
	"testing"
)

func TestInt(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    int
		wantErr bool
	}{
		{"float64", float64(5), 5, false},
		{"int", 3, 3, false},
		{"string", "42", 42, false},
		{"nil", nil, 0, false},
		{"invalid string", "abc", 0, true},
		{"unsupported type", true, 0, true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Int(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("Int(%v) error = %v, wantErr %v", tc.input, err, tc.wantErr)
			}
			if !tc.wantErr && got != tc.want {
				t.Errorf("Int(%v) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    float64
		wantErr bool
	}{
		{"float64", float64(3.14), 3.14, false},
		{"int", 7, 7.0, false},
		{"string", "2.5", 2.5, false},
		{"nil", nil, 0, false},
		{"invalid string", "xyz", 0, true},
		{"unsupported type", true, 0, true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Float64(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("Float64(%v) error = %v, wantErr %v", tc.input, err, tc.wantErr)
			}
			if !tc.wantErr && got != tc.want {
				t.Errorf("Float64(%v) = %f, want %f", tc.input, got, tc.want)
			}
		})
	}
}
