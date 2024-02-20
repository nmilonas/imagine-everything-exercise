package model

import (
	"strconv"
)

type Records struct {
	Records []Record `json:"records"`
}

type Record struct {
	ID        string `json:"id"`
	Created   string `json:"created"`
	Email     string `json:"email"`
	Risk      string `json:"risk"`
	RiskLevel string `json:"risk_level"`
	Meta      string `json:"meta"`
	Active    string `json:"active"`
}

// A nice future additon would be to add a parser to extract the contents of Meta in order to perform further analysis of the data.
type Meta struct {
	Content     string
	UserAgent   string
	IPInternal  string
	IPExternal  string
	BrowserUUID int64
}

type UserRisk struct {
	Email     string
	RiskLevel int
}

func (r *Records) GetLargestRiskLevel() (UserRisk, error) {
	riskiest := UserRisk{
		Email:     "",
		RiskLevel: 0,
	}

	for _, record := range r.Records {
		riskLevel, err := strconv.Atoi(record.RiskLevel)
		if err != nil {
			return UserRisk{}, err
		}

		if riskLevel > riskiest.RiskLevel {
			riskiest.RiskLevel = riskLevel
			riskiest.Email = record.Email
		}
	}

	return riskiest, nil
}
