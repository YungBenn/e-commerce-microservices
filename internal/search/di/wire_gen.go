// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/YungBenn/tech-shop-microservices/internal/search/pb"
	"github.com/YungBenn/tech-shop-microservices/internal/search/repository"
	"github.com/YungBenn/tech-shop-microservices/internal/search/usecase"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/sirupsen/logrus"
)

// Injectors from wire.go:

func InitSearchService(es *elasticsearch.Client, log *logrus.Logger, Topic string) pb.SearchServiceServer {
	esProduct := repository.NewSearchProduct(es)
	searchServiceServer := usecase.NewSearchServiceServer(log, esProduct)
	return searchServiceServer
}