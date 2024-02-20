package main

import (
	"fmt"

	"imagine_everything_exercise/extractor"
)

func main() {
	records, err := extractor.ExtractAuditJSON()
	if err != nil {
		fmt.Println("Invalid JSON")

		return
	}

	riskiestUser, err := records.GetLargestRiskLevelAndFrequency()
	if err != nil {
		fmt.Println("Error finding risk level")

		return
	}

	activeUsers := records.GetActiveUsersInAuditRecords()

	fmt.Println("Largest risk level:", riskiestUser.RiskLevel)
	fmt.Println("Largest risk level user:", riskiestUser.Email)
	fmt.Println("Frequency of largest risk:", riskiestUser.Frequency)

	fmt.Println("Active Users:", activeUsers)
	fmt.Println("Length of records:", len(records.AuditRecords))
}
