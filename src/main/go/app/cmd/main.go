package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/internal/config"
	"main/internal/handlers"
	"main/internal/repository"
	"main/internal/service"
)

func main() {
	cfg := config.MustConfig()

	db, err := repository.NewPostgresRepository(cfg.PostgresConn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	userRepo := repository.NewUserRepository(db.DB)
	orgRepo := repository.NewOrganizationRepository(db.DB)
	tenderRepo := repository.NewTenderRepository(db.DB)
	proposalRepo := repository.NewProposalRepository(db.DB)

	userSvc := service.NewUserService(userRepo)
	orgSvc := service.NewOrganizationService(orgRepo)
	tenderSvc := service.NewTenderService(tenderRepo)
	proposalSvc := service.NewProposalService(proposalRepo)

	userHandler := handlers.NewUserHandler(userSvc)
	orgHandler := handlers.NewOrganizationHandler(orgSvc)
	tenderHandler := handlers.NewTenderHandler(tenderSvc)
	proposalHandler := handlers.NewProposalHandler(proposalSvc)

	r := gin.Default()

	r.GET("/api/ping", handlers.PingHandler)

	r.POST("/api/users", userHandler.CreateUser)
	r.GET("/api/users/:id", userHandler.GetUserByID)
	r.PUT("/api/users/:id", userHandler.UpdateUser)
	r.DELETE("/api/users/:id", userHandler.DeleteUser)

	r.POST("/api/organizations", orgHandler.CreateOrganization)
	r.GET("/api/organizations/:id", orgHandler.GetOrganizationByID)
	r.PUT("/api/organizations/:id", orgHandler.UpdateOrganization)
	r.DELETE("/api/organizations/:id", orgHandler.DeleteOrganization)

	r.POST("/api/tenders", tenderHandler.CreateTender)
	r.PUT("/api/tenders/:id/publish", tenderHandler.PublishTender)
	r.PUT("/api/tenders/:id/close", tenderHandler.CloseTender)
	r.PUT("/api/tenders/:id", tenderHandler.UpdateTender)
	r.GET("/api/tenders/:id", tenderHandler.GetTenderByID)

	r.POST("/api/proposals", proposalHandler.CreateProposal)
	r.GET("/api/proposals/:id", proposalHandler.GetProposalByID)
	r.PUT("/api/proposals/:id", proposalHandler.UpdateProposal)
	r.DELETE("/api/proposals/:id", proposalHandler.DeleteProposal)
	r.GET("/api/proposals/tender/:tenderID", proposalHandler.ListProposalsByTenderID)
}
