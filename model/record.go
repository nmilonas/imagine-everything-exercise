package model

type Records struct {
	Records []Record
}

type Record struct {
	ID        string
	Created   string
	Email     string
	Risk      string
	RiskLevel int64
	Meta      string
	Active    string
}

type Meta struct {
	Content     string
	UserAgent   string
	IPInternal  string
	IPExternal  string
	BrowserUUID int64
}
