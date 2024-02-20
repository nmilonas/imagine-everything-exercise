package model

import (
	"strconv"
)

type AuditRecords struct {
	AuditRecords []AuditRecord `json:"records"`
}

type AuditRecord struct {
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
	Frequency int
}

// GetLargestRiskLevelAndFrequency finds the user with the largest risk level along with the frequency.
func (r *AuditRecords) GetLargestRiskLevelAndFrequency() (UserRisk, error) {
	riskiest := UserRisk{}

	for _, record := range r.AuditRecords {
		riskLevel, err := strconv.Atoi(record.RiskLevel)
		if err != nil {
			return UserRisk{}, err
		}

		if riskLevel > riskiest.RiskLevel {
			riskiest.RiskLevel = riskLevel
			riskiest.Email = record.Email
			riskiest.Frequency = 0

			continue
		}

		riskiest.Frequency++
	}

	return riskiest, nil
}

// GetActiveUsersInAuditRecords finds the active users among the audit records.
func (r *AuditRecords) GetActiveUsersInAuditRecords() []string {
	usersMap := map[string]bool{}
	activeEmails := []string{}

	for _, record := range r.AuditRecords {
		if record.Active == "t" {
			if _, ok := usersMap[record.Email]; !ok {
				usersMap[record.Email] = true
			}
		}
	}

	for email := range usersMap {
		activeEmails = append(activeEmails, email)
	}

	return activeEmails
}
