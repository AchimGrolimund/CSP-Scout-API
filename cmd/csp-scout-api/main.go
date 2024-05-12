package main

import (
	"github.com/AchimGrolimund/CSP-Scout-API/pkg/api"
	"github.com/AchimGrolimund/CSP-Scout-API/pkg/repository"
	"github.com/vipul-rawat/gofr-mongo"
	"gofr.dev/pkg/gofr"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var (
	Build   = "raw"
	Version = "raw"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Print the version and build number
	log.Printf("Version: %s\n", Version)
	log.Printf("Build: %s\n", Build)

	app := gofr.New()

	// using the mongo driver from `vipul-rawat/gofr-mongo`
	db := mongo.New(app.Config, app.Logger(), app.Metrics())

	// inject the mongo into gofr to use mongoDB across the application
	// using gofr context
	app.UseMongo(db)

	reportRepo := repository.NewReportRepository(db)
	handler := api.NewHandler(reportRepo)

	app.GET("/report", handler.FindOne)
	app.GET("/reports", handler.FindAll)

	app.Run()
}
