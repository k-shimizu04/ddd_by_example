// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package infrastructure

import (
	"github.com/k-shimizu04/ddd_by_example/config"
	companies2 "github.com/k-shimizu04/ddd_by_example/interfaces/controllers/companies"
	"github.com/k-shimizu04/ddd_by_example/interfaces/controllers/middleware"
	"github.com/k-shimizu04/ddd_by_example/interfaces/query_services"
	"github.com/k-shimizu04/ddd_by_example/interfaces/repositories"
	"github.com/k-shimizu04/ddd_by_example/usecase/companies"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitRouting(configAPIServer *config.APIServer, d *gorm.DB) *Routing {
	corsMiddleware := middleware.NewCorsMiddleware(configAPIServer)
	companyRepository := repositories.NewCompanyRepository(d)
	companyUsecase := companies.NewCompanyUsecase(companyRepository)
	companyController := companies2.NewCompanyController(companyUsecase)
	companyIdQueryService := query_services.NewCompanyIdQueryService(d)
	companyIDUsecase := companies.NewCompanyIDUsecase(companyIdQueryService)
	companyIDController := companies2.NewCompanyIDController(companyIDUsecase)
	routing := NewRouting(configAPIServer, corsMiddleware, companyController, companyIDController)
	return routing
}
