package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReportData struct {
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
	ClientIP           string `bson:"clientip" json:"clientip"`
	UserAgent          string `bson:"useragent" json:"useragent"`
	ReportTime         int    `bson:"reporttime" json:"reporttime"`
}

type Report struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id"`
	Report ReportData         `bson:"report" json:"report"`
}
