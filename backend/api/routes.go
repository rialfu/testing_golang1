package api

import (
	"rema/kredit/kredit"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "DELETE", "PUT"},
		AllowHeaders: []string{"*"},
	}))
	
	s.Router.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
						"message": "sukes",
					})
	})
	kreditRepo := kredit.NewRepository(s.DB)
	kreditService := kredit.NewService(kreditRepo)
	kreditHandler := kredit.NewHandler(kreditService)
	s.Router.GET("/checklist-pencairan", kreditHandler.HandleGetChecklistPencairan)
	s.Router.POST("/checklist-pencairan", kreditHandler.HandleApprovePencairan)
	s.Router.GET("/report-pencairan", kreditHandler.HandleGetDataReport)
	s.Router.GET("/get-search", kreditHandler.HandleGetSearch)
	// s.Router.GET('')
	// authRepo := auth.NewRepository(s.DB)
	// authService := auth.NewService(authRepo)
	// authHandler := auth.NewHandler(authService)
	// s.Router.POST("/create-account", authHandler.CreateUser)
	// s.Router.POST("/login", authHandler.Login)
	// s.Router.POST("/send-otp", authHandler.SendOTP)

}
