package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
    Sub   string `json:"sub"`   // User ID
    Email string `json:"email"`
    jwt.RegisteredClaims
}