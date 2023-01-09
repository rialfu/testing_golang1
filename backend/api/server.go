package api

import (
	"os"
	"rema/kredit/auto"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
)

type server struct {
	Router *gin.Engine
	DB *gorm.DB
}

func MakeServer(db *gorm.DB) *server {
	s := &server{
		Router: gin.Default(),
		DB: db,
	}
	// s := auto.NewRepository(db)
	automatic := auto.NewRepository(db)
	cron := gocron.NewScheduler(time.UTC)
	cron.Every(15).Minute().Do(func(){
		automatic.GenerateSkalaAngsuran()
		
	})
	cron.Every(30).Minute().Do(func(){
		automatic.GenerateValidatePengajuanKredit()
		
	})
	cron.StartAsync()
	// s.GenerateSkalaAngsuran()
	// s.GenerateValidatePengajuanKredit()
	return s
}

func (s *server) RunServer() {
	s.SetupRouter()
	port := os.Getenv("port")
	if err := s.Router.Run(":" + port); err != nil{
		panic(err)
	}
}