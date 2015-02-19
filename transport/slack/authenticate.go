package slack

import (
	"fmt"
)

func authenticate(api SlackAPICaller) (wsUrl string, userId string, users []User) {
	m := api.Call("rtm.start").(map[string]interface{})

	wsUrl = m["url"].(string)

	usersData := m["users"].([]interface{})

	userId = m["self"].(map[string]interface{})["id"].(string)

	for _, u := range usersData {
		m = u.(map[string]interface{})
		users = append(users, User{m["id"].(string), m["name"].(string)})

		fmt.Println("id:", m["id"].(string))
		fmt.Println("name:", m["name"].(string))
	}

	// store this for use by the slack transport
	knownUsers = users

	return
}
