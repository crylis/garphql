package handlers

import (
	"fmt"
	"strconv"

	"main.go/graph/model"
)

func EditMember(id int, newName string, newPosition model.PositionName, listName string, point int) error {
	strID := strconv.Itoa(id)
	for i, m := range member {
		if m.ID == strID {
			member[i].Name = newName
			member[i].Position = newPosition

			for j, worker := range member[i].Worker {
				if worker.Name == listName {
					member[i].Worker[j].Point = point
					break
				}
			}

			return nil
		}
	}
	return fmt.Errorf("member with ID %d not found", id)
}
