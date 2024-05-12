package domain

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
}

type Report struct {
	ID     string     `bson:"_id" json:"_id"`
	Report ReportData `bson:"report" json:"report"`
}
