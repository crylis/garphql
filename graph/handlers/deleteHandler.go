package handlers

import (
	"fmt"
	"strconv"
)

func DeleteMember(id int) error {
	strID := strconv.Itoa(id)
	for i, m := range member {
		if m.ID == strID {
			member = append(member[:i], member[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("member with ID %d not found", id)
}
