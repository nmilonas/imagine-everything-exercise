package main

import (
	"fmt"

	"imagine_everything_exercise/extractor"
)

func main() {
	records, err := extractor.ExtractJSON()
	if err != nil {
		fmt.Println("Invalid JSON")

		return
	}
	riskiestUser, err := records.GetLargestRiskLevel()
	if err != nil {
		fmt.Println("Error finding risk level")

		return
	}

	fmt.Println("Largest risk level", riskiestUser.RiskLevel)
	fmt.Println("Largest risk level user", riskiestUser.Email)

	fmt.Println("Length of records:", len(records.Records))
}
