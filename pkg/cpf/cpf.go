package cpf

import (
	"regexp"
	"strconv"
)

func Validate(cpf string) bool {
	cpf = removeNonDigits(cpf)

	if len(cpf) != 11 {
		return false
	}

	if checkRepeatedDigits(cpf) {
		return false
	}

	if !validateCPFCheckDigits(cpf) {
		return false
	}

	return true
}

func removeNonDigits(cpf string) string {
	reg := regexp.MustCompile(`\D`)
	return reg.ReplaceAllString(cpf, "")
}

func checkRepeatedDigits(cpf string) bool {
	for i := 0; i < 10; i++ {
		if cpf == strconv.Itoa(i*11) {
			return true
		}
	}
	return false
}

func validateCPFCheckDigits(cpf string) bool {
	ninthDigit, _ := strconv.Atoi(string(cpf[9]))
	tenthDigit, _ := strconv.Atoi(string(cpf[10]))

	sum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (10 - i)
	}

	remainder := sum % 11
	ninthDigitExpected := 0
	if remainder >= 2 {
		ninthDigitExpected = 11 - remainder
	}

	if ninthDigit != ninthDigitExpected {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (11 - i)
	}

	remainder = sum % 11
	tenthDigitExpected := 0
	if remainder >= 2 {
		tenthDigitExpected = 11 - remainder
	}

	return tenthDigit == tenthDigitExpected
}
