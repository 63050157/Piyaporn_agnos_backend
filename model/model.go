package model

import (
	"database/sql"
	"regexp"
	"unicode"
)

func CalculateNumStep(passwordInput string) int {
	var count int

	if len(passwordInput) < 6 || len(passwordInput) >= 20 {
		count++
	}

	var hasLower, hasUpper, hasDigit bool
	for _, char := range passwordInput {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	if !hasLower || !hasUpper || !hasDigit {
		count++
	}

	for i := 0; i < len(passwordInput)-2; i++ {
		if passwordInput[i] == passwordInput[i+1] && passwordInput[i] == passwordInput[i+2] {
			count++
			break
		}
	}

	return count
}


func AddData(db *sql.DB, passwordInput string) error {
	numStep := CalculateNumStep(passwordInput)

	query := "INSERT INTO strongpassword (password_input, num_of_steps) VALUES ($1, $2)"
	_, err := db.Exec(query, passwordInput, numStep)
	if err != nil {
		return err
	}
	return nil
}

func IsValidInput(input string) bool {
	if len(input) < 1 || len(input) > 40 {
		return false
	}

	validChars := regexp.MustCompile("^[a-zA-Z0-9.!]+$")
	return validChars.MatchString(input)
}