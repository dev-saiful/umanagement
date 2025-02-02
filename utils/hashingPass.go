package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a plain text password and returns the hashed password or an error.
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

// CheckPassword compares a hashed password with a plain text password and returns an error if they don't match.
func CheckPassword(hashedPassword, password string) error {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
        return err
    }
    return nil
}