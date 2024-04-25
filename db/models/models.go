package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	Username     string
	PasswordHash string
	Role         string
}

type TokenClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
	Csrf string `json:"csrf"`
}

const RefreshTokenValidTime = time.Hour * 72
const AuthTokenValidTime = time.Minute * 15
