package jwt

import (
	"auth-service/internal/config"
	"encoding/hex"
	"errors"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
	"golang.org/x/exp/rand"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token is expired")
)

type Manager struct {
	secret    []byte
	issuer    string
	audience  string
	accessTTL time.Duration
}

type Claims struct {
	UserID string `json:"uid"`
	Email  string `json:"email"`

	jwtlib.RegisteredClaims
}

func NewManagerFromConfig(cfg config.TokensConfiguration) *Manager {
	return &Manager{
		secret:    []byte(cfg.JwtSecret),
		issuer:    cfg.Issuer,
		audience:  cfg.Audience,
		accessTTL: cfg.AccessTokenTTL,
	}
}

// GenerateAccessToken generate access-токен with TTL from config.
func (m *Manager) GenerateAccessToken(userID, email string) (string, error) {
	now := time.Now().UTC()
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwtlib.RegisteredClaims{
			Issuer:    m.issuer,
			Audience:  []string{m.audience},
			Subject:   userID,
			IssuedAt:  jwtlib.NewNumericDate(now),
			ExpiresAt: jwtlib.NewNumericDate(now.Add(m.accessTTL)),
		},
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	return token.SignedString(m.secret)
}

// ParseAndValidate parse token, check signature, exp, issuer, audience.
func (m *Manager) ParseAndValidate(tokenStr string) (*Claims, error) {
	token, err := jwtlib.ParseWithClaims(tokenStr, &Claims{}, func(t *jwtlib.Token) (interface{}, error) {
		// defence from algorithm substitution
		if _, ok := t.Method.(*jwtlib.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return m.secret, nil
	})
	if err != nil {
		if errors.Is(err, jwtlib.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	// issuer
	if claims.Issuer != "" && claims.Issuer != m.issuer {
		return nil, ErrInvalidToken
	}

	// audience
	if len(claims.Audience) > 0 {
		match := false
		for _, aud := range claims.Audience {
			if aud == m.audience {
				match = true
				break
			}
		}
		if !match {
			return nil, ErrInvalidToken
		}
	}

	return claims, nil
}

func (m *Manager) GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 64)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
