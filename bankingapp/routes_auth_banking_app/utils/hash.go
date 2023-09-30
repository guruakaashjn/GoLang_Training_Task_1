package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(payload string) ([]byte, error) {
	payloadBytes := []byte(payload)
	hashedPassword, err := bcrypt.GenerateFromPassword(
		payloadBytes,
		16,
	)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

func IsHashSame(hashedPassword string, payload string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(payload),
	)
	if err != nil {
		fmt.Println(err)
	}
	return err == nil

}
