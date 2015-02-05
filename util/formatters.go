package util

import (
	"fmt"
	"encoding/json"
)

func DumpJSON(o *interface{}) {
	pretty, _ := json.MarshalIndent(&o, "", "  ")
	fmt.Printf("%v\n", string(pretty))
}
