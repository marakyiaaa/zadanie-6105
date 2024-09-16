package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/internal/config"
	handlers2 "main/internal/handlers"
	repository2 "main/internal/repository"
	service2 "main/internal/service"
)

func main() {
	cfg := config.MustConfig()

	db, err := repository2.NewPostgresRepository(cfg.PostgresConn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	userRepo := repository2.NewUserRepository(db.DB)
	orgRepo := repository2.NewOrganizationRepository(db.DB)
	tenderRepo := repository2.NewTenderRepository(db.DB)
	proposalRepo := repository2.NewProposalRepository(db.DB)

	userSvc := service2.NewUserService(userRepo)
	orgSvc := service2.NewOrganizationService(orgRepo)
	tenderSvc := service2.NewTenderService(tenderRepo)
	proposalSvc := service2.NewProposalService(proposalRepo)

	userHandler := handlers2.NewUserHandler(userSvc)
	orgHandler := handlers2.NewOrganizationHandler(orgSvc)
	tenderHandler := handlers2.NewTenderHandler(tenderSvc)
	proposalHandler := handlers2.NewProposalHandler(proposalSvc)

	r := gin.Default()

	r.GET("/api/ping", handlers2.PingHandler)

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

	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
