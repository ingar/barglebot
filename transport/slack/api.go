package slack

import (
	"encoding/json"
	"fmt"
	"github.com/ingar/barglebot/util"
	"io/ioutil"
	"net/http"
)

type SlackAPICaller interface {
	Call(apiMethod string) interface{}
}

type RestAPICaller struct {
	apiToken string
}

func (self RestAPICaller) Call(apiMethod string) (result interface{}) {
	resp, _ := http.Get(fmt.Sprintf("https://slack.com/api/%s?token=%s", apiMethod, self.apiToken))
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &result) // ignore error!

	fmt.Println("Authenticated:")
	util.DumpJSON(&result)
	return
}
