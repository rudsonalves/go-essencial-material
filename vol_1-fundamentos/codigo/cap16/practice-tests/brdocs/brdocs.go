package brdocs

import (
	"practice-tests/textutil"
	"strings"
)

func IsCPF(cpf string) bool {
	cpfNumbers := textutil.OnlyDigits(cpf)
	if len(cpfNumbers) != 11 {
		return false
	}

	for _, b := range "0123456789" {
		if cpfNumbers == strings.Repeat(string(b), 11) {
			return false
		}
	}

	base := cpfNumbers[:9]
	d1 := calcCpfDigit(base)
	base = base + string(intToByte(d1))
	d2 := calcCpfDigit(base)

	if d1 != byteToInt(cpfNumbers[9]) || d2 != byteToInt(cpfNumbers[10]) {
		return false
	}

	return true
}

func IsCNPJ(cnpj string) bool {
	cnpjNumbers := textutil.OnlyDigits(cnpj)
	if len(cnpjNumbers) != 14 {
		return false
	}

	for _, b := range "0123456789" {
		if cnpjNumbers == strings.Repeat(string(b), 14) {
			return false
		}
	}

	base := cnpjNumbers[:12]

	d1 := calcCnpjDigit(base)
	base += string(intToByte(d1))

	d2 := calcCnpjDigit(base)

	return d1 == byteToInt(cnpjNumbers[12]) &&
		d2 == byteToInt(cnpjNumbers[13])
}

func calcCnpjDigit(digits string) int {
	weight := len(digits) - 7
	sum := 0

	for i := range len(digits) {
		sum += byteToInt(digits[i]) * weight

		weight--
		if weight < 2 {
			weight = 9
		}
	}

	rest := sum % 11

	if rest < 2 {
		return 0
	}

	return 11 - rest
}

func calcCpfDigit(digits string) int {
	sum := 0
	length := len(digits)
	weight := length + 1

	for i := range length {
		sum += (weight - i) * int(digits[i]-'0')
	}

	rest := sum % 11

	if rest < 2 {
		return 0
	}
	return 11 - rest
}

func intToByte(i int) byte {
	return byte('0' + i)
}

func byteToInt(b byte) int {
	return int(b - '0')
}
