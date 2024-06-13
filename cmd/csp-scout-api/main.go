package main

import (
	"github.com/AchimGrolimund/CSP-Scout-API/pkg/api"
	"github.com/AchimGrolimund/CSP-Scout-API/pkg/repository"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/datasource/mongo"
	"log"
	//_ "net/http/pprof"
)

var (
	Build   = "raw"
	Version = "raw"
)

func main() {
	//
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()

	// Print the version and build number
	log.Printf("Version: %s\n", Version)
	log.Printf("Build: %s\n", Build)

	app := gofr.New()

	// using the mongo driver from `vipul-rawat/gofr-mongo`
	db := mongo.New(mongo.Config{URI: "mongodb://root:toor@10.90.0.10:27017/", Database: "csp-report"})

	// inject the mongo into gofr to use mongoDB across the application
	// using gofr context
	app.AddMongo(db)

	reportRepo := repository.NewReportRepository(db)
	handler := api.NewHandler(reportRepo)

	app.GET("/report", handler.FindOne)
	app.GET("/reports", handler.FindAll)
	app.GET("/reportbyid", handler.FindByID)
	app.GET("/reportsbytimelt", handler.FindByTimeLT)
	app.GET("/reportsbytimegt", handler.FindByTimeGT)
	app.GET("/reportsbyuseragent", handler.FindByUserAgent)

	app.Run()
}
