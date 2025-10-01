package usecases

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
)

type LoginUC struct {
	config      *entity.Config
	redisClient *redis.Client
}

func NewLoginUC(config *entity.Config, redisClient *redis.Client) *LoginUC {
	return &LoginUC{
		config:      config,
		redisClient: redisClient,
	}
}

func (u *LoginUC) Execute(ctx context.Context, req entity.LoginRequest, ip string) (string, error) {
	attemptsKey := fmt.Sprintf("login_attempts:%s", ip)
	blockedKey := fmt.Sprintf("login_blocked:%s", ip)

	_, err := u.redisClient.Get(ctx, blockedKey).Result()
	if err == nil {
		return "", errors.New("too many failed login attempts. Please try again later")
	}

	loginUser := u.config.Login.User
	loginPass := u.config.Login.Password

	if req.Username != loginUser || req.Password != loginPass {
		attempts, err := u.redisClient.Incr(ctx, attemptsKey).Result()
		if err != nil {
			return "", errors.New("failed to update login attempts")
		}

		if attempts >= 5 {
			u.redisClient.Set(ctx, blockedKey, "1", 10*time.Minute)
			u.redisClient.Del(ctx, attemptsKey)
			return "", errors.New("too many failed login attempts. Please try again later")
		} else {
			u.redisClient.Expire(ctx, attemptsKey, 10*time.Minute)
		}

		return "", errors.New("invalid credentials")
	}

	u.redisClient.Del(ctx, attemptsKey)

	jwtSecret := u.config.JWT.Secret
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &entity.Claims{
		Username: req.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}