package main

import (
	"rema/kredit/api"
	"time"

	_ "github.com/joho/godotenv/autoload"
)


func main() {
	db, err := api.SetupDb()
	if err != nil{
		panic(err)
	}
	// s := auto.NewRepository(db)
	// s.GenerateSkalaAngsuran()
	// s.GenerateValidatePengajuanKredit()
	server := api.MakeServer(db)
	server.RunServer()

	// s := gocron.NewScheduler(time.UTC)
	// s.Every(1).Minute().Do(func(){
	// 	s := auto.NewRepository(db)
	// 	s.GenerateSkalaAngsuran()
	// 	// fmt.Println("every minute", time.Now().String())
	// })
	// s.StartAsync()
	for{
		time.Sleep(10* time.Second)
	}
 
	// s.Every(2).Minute().Do(func(){fmt.Print("every 2")})


}
