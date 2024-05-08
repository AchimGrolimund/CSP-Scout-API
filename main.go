package main

import (
	mongo "github.com/vipul-rawat/gofr-mongo"
	"go.mongodb.org/mongo-driver/bson"

	"gofr.dev/pkg/gofr"
)

type Report struct {
	DocumentUri        string `bson:"documenturi" json:"documenturi"`
	Referrer           string `bson:"referrer" json:"referrer"`
	ViolatedDirective  string `bson:"violateddirective" json:"violateddirective"`
	EffectiveDirective string `bson:"effectivedirective" json:"effectivedirective"`
	OriginalPolicy     string `bson:"originalpolicy" json:"originalpolicy"`
	Disposition        string `bson:"disposition" json:"disposition"`
	BlockedUri         string `bson:"blockeduri" json:"blockeduri"`
	LineNumber         int    `bson:"linenumber" json:"linenumber"`
	SourceFile         string `bson:"sourcefile" json:"sourcefile"`
	StatusCode         int    `bson:"statuscode" json:"statuscode"`
	ScriptSample       string `bson:"scriptsample" json:"scriptsample"`
}

func main() {
	app := gofr.New()

	// using the mongo driver from `vipul-rawat/gofr-mongo`
	db := mongo.New(app.Config, app.Logger(), app.Metrics())

	// inject the mongo into gofr to use mongoDB across the application
	// using gofr context
	app.UseMongo(db)

	//app.POST("/mongo", Insert)
	app.GET("/report", Get)

	app.Run()
}

func Get(ctx *gofr.Context) (interface{}, error) {
	var result Report

	err := ctx.Mongo.Find(ctx, "reports", bson.D{{}} /* valid filter */, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
