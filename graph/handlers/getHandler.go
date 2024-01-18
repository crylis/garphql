package handlers

import (
	"encoding/json"
	"log"

	"main.go/graph/model"
)

func GetMember() []*model.Member {

	//add webhook
	memberJSON, err := json.Marshal(member)
	if err != nil {
		log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
	}

	msg := string(memberJSON)

	SendDiscordMessage(msg)

	return member
}
