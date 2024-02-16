// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/YungBenn/tech-shop-microservices/internal/auth/handler"
	"github.com/YungBenn/tech-shop-microservices/internal/auth/pb"
	"github.com/YungBenn/tech-shop-microservices/internal/auth/repository"
	"github.com/YungBenn/tech-shop-microservices/internal/auth/token"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitAuthService(db *gorm.DB, log *logrus.Logger, rdb *redis.Client) pb.AuthServiceServer {
	tokenRepository := token.NewTokenRepository(rdb)
	authRepository := repository.NewAuthRepository(db, log)
	authServiceServer := handler.NewAuthServiceServer(log, tokenRepository, authRepository)
	return authServiceServer
}
