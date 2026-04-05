package bird

import "testing"

func TestResolveLabel(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Eagle exact", "Eagle", "American Bald Eagle"},
		{"sparrow", "sparrow", "Chickadee"},
		{"empty", "", "Chickadee"},
		{"whitespace only", "   ", "Chickadee"},
		{"lowercase eagle", "eagle", "Chickadee"},
		{"uppercase EAGLE", "EAGLE", "Chickadee"},
		{"Eagle with space", "Eagle ", "Chickadee"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResolveLabel(tt.input); got != tt.want {
				t.Errorf("ResolveLabel(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}
