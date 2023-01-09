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
		// Buatlah 1 Service Automatic Generate Skala Angsuran Debitur
		// Scheduler Service Automatic Run For Checking data setiap 15 menit sekali.
		// Data-data yang di-generate otomatis angsuran per bulan masing-masing debitur 
// adalah data-data yang “approval_status” = 0
// Data Angsuran tiap bulan masing-masing debitur yang di-generate tersebut diinsert 
// ke Tabel “Skala_Rental_Tab”
// Debitur yang skala angsurannya selesai degenerate secara otomatis, update “approval_status” = 1.
	})
	cron.Every(30).Minute().Do(func(){
		automatic.GenerateValidatePengajuanKredit()
		// ervice Run secara schedular (Tiap 30 menit) 
		// untuk membaca dan mem-proses validasi data dan insert data 
		// pengajuan kredit pada table “Staging_Customer” berdasarkan:
		// Create_date = Current_date
		// SC_Flag = 0
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