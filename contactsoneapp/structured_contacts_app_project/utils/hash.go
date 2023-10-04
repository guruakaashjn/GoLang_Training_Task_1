package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(payload string) ([]byte, error) {
	payloadBytes := []byte(payload)
	hash, err := bcrypt.GenerateFromPassword(payloadBytes, 10)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func IsHashMatch(hash []byte, payload string) bool {
	// fmt.Println("Required user password: ")
	// fmt.Println(hash)
	// fmt.Println("Claims password: ")
	// fmt.Println([]byte(payload))
	err := bcrypt.CompareHashAndPassword(hash, []byte(payload))
	fmt.Println(err)
	return err == nil
}
