package utils

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

func CalculateIdealCost() int {
	cost := bcrypt.DefaultCost

	startTime := time.Now()
	HashPassword("microbenchmark") // possible infinite loop
	durationMS := time.Since(startTime).Milliseconds()

	for durationMS < 250 {
		cost++
		durationMS *= 2
	}

	return cost
}

func HashPassword(password string) string {
	cost := CalculateIdealCost()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(hashedPassword)
}

func ComparePasswords(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
