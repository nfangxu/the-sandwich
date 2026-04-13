package auth

import "testing"

func TestHashAndCheckPassword(t *testing.T) {
	password := "secret123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}
	if !CheckPasswordHash(password, hash) {
		t.Errorf("Expected password check to pass")
	}
	if CheckPasswordHash("wrong", hash) {
		t.Errorf("Expected password check to fail on wrong password")
	}
}

func TestGenerateAndValidateJWT(t *testing.T) {
    secret := "test_secret"
    userID := uint(1)
    
    token, err := GenerateJWT(userID, secret)
    if err != nil {
        t.Fatalf("Failed to generate JWT: %v", err)
    }
    
    id, err := ValidateJWT(token, secret)
    if err != nil {
        t.Fatalf("Failed to validate JWT: %v", err)
    }
    
    if id != userID {
        t.Errorf("Expected userID %d, got %d", userID, id)
    }
}
