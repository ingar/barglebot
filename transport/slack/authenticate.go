package slack

func authenticate(api SlackAPICaller) (wsUrl string) {
	m := api.Call("rtm.start").(map[string]interface{})
	wsUrl = m["url"].(string)
	return
}
