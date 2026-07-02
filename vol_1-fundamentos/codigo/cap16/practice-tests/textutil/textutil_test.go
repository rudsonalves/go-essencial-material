package textutil

import "testing"

func TestOnlyDigits(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"texto sem números", "abs", ""},
		{"texto com números", "tenho 28 anos", "28"},
		{"123", "123", "123"},
		{"123.456.789-00", "123.456.789-00", "12345678900"},
		{"(11) 91234-5678", "(11) 91234-5678", "11912345678"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := OnlyDigits(test.value)

			if got != test.want {
				t.Errorf("OnlyDigits('%s')= '%s'; want '%s'", test.value, got, test.want)
			}
		})
	}
}
