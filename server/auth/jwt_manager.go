package auth

import (
	"context"
	"log"
	"time"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type JWTManager interface {
	Validate(ctx context.Context) (context.Context, error)
}

type jwtManager struct {
	// TODO
	secretKey string
	duration  time.Duration
}

// NewJWTManager returns a new jwt manager to parse and validate jwt token
func NewJWTManager(secretKey string, duration time.Duration) JWTManager {
	return &jwtManager{
		secretKey: secretKey,
		duration:  duration,
	}
}

func (jm jwtManager) Validate(ctx context.Context) (context.Context, error) {
	log.Printf("Validate")
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil || token == "" {
		return nil, status.Errorf(codes.Unauthenticated, "%v", err)
	}

	//TODO

	return ctx, nil
}
