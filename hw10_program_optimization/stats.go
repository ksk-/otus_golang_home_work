package hw10programoptimization

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/mailru/easyjson"
)

var (
	ErrInvalidUser  = errors.New("invalid user")
	ErrInvalidEmail = errors.New("invalid email")
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)
	suffix := "." + strings.ToLower(domain)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var user User
		if err := easyjson.Unmarshal(scanner.Bytes(), &user); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInvalidUser, err)
		}

		email := strings.ToLower(user.Email)
		if matched := strings.HasSuffix(email, suffix); matched {
			d, err := extractDomain(email)
			if err != nil {
				return nil, fmt.Errorf("extract domain: %w", err)
			}
			result[d]++
		}
	}
	return result, nil
}

func extractDomain(email string) (string, error) {
	if parts := strings.SplitN(email, "@", 2); len(parts) == 2 {
		return parts[1], nil
	}
	return "", ErrInvalidEmail
}
