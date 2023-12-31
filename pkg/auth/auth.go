package auth

import (
	"fmt"
	"math/rand"
)

//implement proper authentication later

var passCode string

func CodeGen() string {
	randInt := 0
	//should be minimum 10000 to generate a 4 digit password
	for randInt < 1000 {
		randInt = rand.Int() % 10_000
	}
	passCode = fmt.Sprintf("%d", randInt)
	return passCode
}

func MatchAndVerify(code string) bool {
	return code == passCode
}
