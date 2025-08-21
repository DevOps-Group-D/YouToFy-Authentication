package utils

import (
	"testing"
)

func TestHash(t *testing.T) {
	password := "testpassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if hashedPassword == "" {
		t.Error("expected hashed password to be non-empty")
	}
	err = CheckHashedPassword(hashedPassword, password)
	if err != nil {
		t.Errorf("expected no error when checking hashed password, got %v", err)
	}
}

func TestHashWithEmptyPassword(t *testing.T) {
	password := ""
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if hashedPassword == "" {
		t.Error("expected hashed password to be non-empty")
	}
	err = CheckHashedPassword(hashedPassword, password)
	if err != nil {
		t.Errorf("expected no error when checking hashed password, got %v", err)
	}
}

func TestCheckHashedPasswordWithWrongPassword(t *testing.T) {
	password := "testpassword"
	wrongPassword := "wrongpassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	err = CheckHashedPassword(hashedPassword, wrongPassword)
	if err == nil {
		t.Error("expected error when checking wrong password, got nil")
	}
}

func TestCheckHashedPasswordWithEmptyHash(t *testing.T) {
	password := "testpassword"
	emptyHash := ""
	err := CheckHashedPassword(emptyHash, password)
	if err == nil {
		t.Error("expected error when checking empty hash, got nil")
	}
}

func TestCheckHashedPasswordWithEmptyPassword(t *testing.T) {
	emptyPassword := ""
	hashedPassword, err := HashPassword(emptyPassword)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	err = CheckHashedPassword(hashedPassword, emptyPassword)
	if err != nil {
		t.Errorf("expected no error when checking empty password, got %v", err)
	}
}

func TestToken(t *testing.T) {
	token, err := GenerateToken(32)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(token) != 32 {
		t.Errorf("expected token length to be 32, got %d", len(token))
	}
	if token == "" {
		t.Error("expected token to be non-empty")
	}
	otherToken, err := GenerateToken(32)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if token == otherToken {
		t.Error("expected tokens to be unique")
	}
}

func TestTokenWithZeroLength(t *testing.T) {
	token, err := GenerateToken(0)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if token != "" {
		t.Error("expected token to be empty for zero length")
	}
}

func TestTokenWithNegativeLength(t *testing.T) {
	_, err := GenerateToken(-1)
	if err == nil {
		t.Error("expected error when checking negative sized token, got nil")
	}

}
