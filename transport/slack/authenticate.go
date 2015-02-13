package slack

func authenticate(api SlackAPICaller) (wsUrl string, userId string) {
	m := api.Call("rtm.start").(map[string]interface{})
	wsUrl = m["url"].(string)
	userId = m["self"].(map[string]interface{})["id"].(string)
	return
}
