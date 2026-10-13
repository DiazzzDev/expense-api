package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const UserIDKey = "userID"

type Authenticator struct {
	keyfunc  keyfunc.Keyfunc
	issuer   string
	clientID string
}

func NewAuthenticator(region, userPoolID, clientID string) (*Authenticator, error) {
	issuer := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s", region, userPoolID)

	k, err := keyfunc.NewDefault([]string{issuer + "/.well-known/jwks.json"})
	if err != nil {
		return nil, fmt.Errorf("load jwks: %w", err)
	}
	return &Authenticator{keyfunc: k, issuer: issuer, clientID: clientID}, nil
}

func (a *Authenticator) Require() gin.HandlerFunc {
	return func(c *gin.Context) {
		raw, ok := strings.CutPrefix(c.GetHeader("Authorization"), "Bearer ")
		if !ok || raw == "" {
			unauthorized(c)
			return
		}

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(raw, claims, a.keyfunc.Keyfunc,
			jwt.WithValidMethods([]string{"RS256"}),
			jwt.WithIssuer(a.issuer),
			jwt.WithExpirationRequired(),
		)
		if err != nil {
			unauthorized(c)
			return
		}

		if claims["token_use"] != "access" || claims["client_id"] != a.clientID {
			unauthorized(c)
			return
		}

		sub, _ := claims["sub"].(string)
		if sub == "" {
			unauthorized(c)
			return
		}

		c.Set(UserIDKey, sub)
		c.Next()
	}
}

type UserEnsurer interface {
	Ensure(ctx context.Context, id string) error
}

func EnsureUser(u UserEnsurer) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := u.Ensure(c.Request.Context(), c.GetString(UserIDKey)); err != nil {
			log.Printf("ensure user: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}
		c.Next()
	}
}

func unauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
}
