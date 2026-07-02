package brdocs

import "testing"

func TestIsCPF(t *testing.T) {
	tests := []struct {
		name string
		cpf  string
		want bool
	}{
		{"987.654.321-00", "987.654.321-00", true},
		{"123.456.789-09", "123.456.789-09", true},
		{"529.982.247-25", "529.982.247-25", true},
		{"52998224725", "52998224725", true},
		{"529.982.247-26", "529.982.247-26", false},
		{"111.111.111-11", "111.111.111-11", false},
		{"123", "123", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsCPF(test.cpf)

			if got != test.want {
				t.Errorf("IsCPF(%s) = %v; wnat %v", test.cpf, got, test.want)
			}
		})
	}
}

func TestIsCNPJ(t *testing.T) {
	tests := []struct {
		name string
		cnpj string
		want bool
	}{
		{"04.252.011/0001-10", "04.252.011/0001-10", true},
		{"40.688.134/0001-61", "40.688.134/0001-61", true},
		{"11.222.333/0001-81", "11.222.333/0001-81", true},
		{"04252011000110", "04252011000110", true},
		{"04.252.011/0001-11", "04.252.011/0001-11", false},
		{"11.111.111/1111-11", "11.111.111/1111-11", false},
		{"123", "123", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsCNPJ(test.cnpj)

			if got != test.want {
				t.Errorf("IsCNPJ(%s) = %v; want %v", test.cnpj, got, test.want)
			}
		})
	}
}

func BenchmarkIsCPF(b *testing.B) {
	for range b.N {
		IsCPF("987.654.321-00")
	}
}

func BenchmarkIsCNPJ(b *testing.B) {
	for range b.N {
		IsCNPJ("04.252.011/0001-10")
	}
}
