package handlers

import "main.go/graph/model"

var member = []*model.Member{
	{
		ID:       "1",
		Name:     "member001",
		Position: model.AllPositionName[0],
		Worker: []*model.List{
			{
				Name:  "view",
				Point: 10,
			},
		},
	},
	{
		ID:       "2",
		Name:     "member002",
		Position: model.AllPositionName[1],
		Worker: []*model.List{
			{
				Name:  "edit",
				Point: 5,
			},
		},
	},
}
