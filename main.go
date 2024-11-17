package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	"strings"
)

const (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}<>?,./"
)

func generatePassword(length int, useDigits, useSpecialChars, useUpper bool) (string, error) {
	charSet := lowerLetters
	if useUpper {
		charSet += upperLetters
	}

	if useDigits {
		charSet += digits
	}

	if useSpecialChars {
		charSet += specialChars
	}

	if len(charSet) == 0 {
		return "", fmt.Errorf("No character set is selected for password generation")
	}

	var password strings.Builder
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			return "", fmt.Errorf("Random number generation error: %v", err)
		}
		password.WriteByte(charSet[randomIndex.Int64()])
	}

	return password.String(), nil
}

func main() {
	length := flag.Int("length", 12, "Password length")
	useDigits := flag.Bool("digits", true, "Use numbers in the password")
	useSpecialChars := flag.Bool("special", true, "Use special characters in the password")
	useUpper := flag.Bool("upper", true, "Use capital letters in a password")

	flag.Parse()

	password, err := generatePassword(*length, *useDigits, *useSpecialChars, *useUpper)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	fmt.Printf("Сгенерированный пароль: %s\n", password)
}
