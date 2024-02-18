package main

import (
	"fmt"
	"imagine_everything_exercise/json"
)

func main() {
	recordsJSON := `{
   "RECORDS":[
      {
         "id":"f895c9dd-94c2-4f97-8d25-c27f695d051a",
         "created":"25/11/2019 19:58:40.183337+00",
         "email":"test@imagineeverything.com",
         "risk":"",
         "risk_level":"0",
         "meta":"{\"content\": \"https://www.abcya.com/\", \"user_agent\": \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36\", \"ip_external\": \"172.10.10.14\", \"ip_internal\": [\"10.10.10.10\"], \"browser_uuid\": \"116540447491807533196\"}",
         "active":"t"
      },
      {
         "id":"a0691552-3268-4297-b6df-f9bd7e349272",
         "created":"25/11/2019 19:53:36.600218+00",
         "email":"test@imagineeverything.com",
         "risk":"",
         "risk_level":"0",
         "meta":"{\"content\": \"https://www.pubg.com/\", \"user_agent\": \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36\", \"ip_external\": \"172.10.10.14\", \"ip_internal\": [\"10.10.10.10\"], \"browser_uuid\": \"105974343207421462009\"}",
         "active":"t"
      }
   ]
}`

	records := json.ExtractJSON(recordsJSON)

	fmt.Println("ID 1", records.Records[0].ID)
	fmt.Println("Email 1", records.Records[0].Email)

	fmt.Println("ID 2", records.Records[1].ID)
	fmt.Println("Email 2", records.Records[1].Email)
}
