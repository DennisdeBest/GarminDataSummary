package output

import (
	"activitesSummary/pkg/args"
	"activitesSummary/pkg/data"
	"encoding/json"
	"fmt"
	"os"
)

func PrintJson(data *data.Data, args args.Args) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	fmt.Println(string(jsonData))

	os.Exit(0)
}
