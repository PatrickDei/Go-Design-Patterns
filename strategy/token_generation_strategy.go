package strategy

import (
	"encoding/base64"
	"fmt"
)

type User struct {
	username string
	password string
}

type TokenGenerationStrategy interface {
	generateToken(User) string
}

type BasicTokenGenerationStrategy struct{}

func (basic *BasicTokenGenerationStrategy) generateToken(u User) string {
	return base64.StdEncoding.EncodeToString([]byte(u.username + ":" + u.password))
}

type JWTTokenGenerationStrategy struct{}

func (basic *JWTTokenGenerationStrategy) generateToken(u User) string {
	fmt.Println(fmt.Errorf("jwt token implementation doesn't exist"))
	return ""
}
