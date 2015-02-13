package util

import (
	"encoding/json"
	"fmt"
)

func FmtJSON(o interface{}) string {
	pretty, _ := json.MarshalIndent(o, "", "  ")
	return fmt.Sprintf("%v", string(pretty))
}
