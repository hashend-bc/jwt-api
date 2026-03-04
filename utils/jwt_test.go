package utils

import (
	"testing"
	"user-auth-api/config"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateToken(t *testing.T) {
	config.JwtSecret = []byte("test-secret")
	
	token, err := GenerateToken("testuser")
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}
	if token == "" {
		t.Fatal("Token is empty")
	}
}

func TestValidateToken(t *testing.T) {
	config.JwtSecret = []byte("test-secret")
	
	token, _ := GenerateToken("testuser")
	
	parsed, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("ValidateToken failed: %v", err)
	}
	
	claims := parsed.Claims.(jwt.MapClaims)
	if claims["username"] != "testuser" {
		t.Errorf("Expected username 'testuser', got %v", claims["username"])
	}
}

func TestValidateToken_Invalid(t *testing.T) {
	config.JwtSecret = []byte("test-secret")
	
	_, err := ValidateToken("invalid-token")
	if err == nil {
		t.Fatal("Expected error for invalid token")
	}
}
