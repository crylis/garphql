package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"main.go/graph/model"
)

func FindMemberByID(id int) (*model.Member, error) {
	strID := strconv.Itoa(id)
	for _, m := range member {
		if m.ID == strID {

			memberJSON, err := json.Marshal(m)
			if err != nil {
				log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
			}

			msg := string(memberJSON)

			SendDiscordMessage(msg)

			return m, nil
		}
	}

	return nil, fmt.Errorf("member with ID %d not found", id)
}
