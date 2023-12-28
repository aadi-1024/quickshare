package auth

import (
	"fmt"
	"math/rand"
)

//implement proper authentication later

var passCode string

func CodeGen() string {
	passCode = fmt.Sprintf("%d", rand.Int()%10_000)
	return passCode
}

func MatchAndVerify(code string) bool {
	return code == passCode
}
