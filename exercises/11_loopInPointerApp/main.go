package main

import "fmt"

type SendgridItem struct {
	SendgridId int
}

func main() {
	templates := []SendgridItem{
		{
			SendgridId: 11,
		},
		{
			SendgridId: 12,
		},
		{
			SendgridId: 13,
		},
	}
	sendgridTemplates := make(map[int]*SendgridItem)
	for i := range templates {
		t := &templates[i]
		sendgridTemplates[t.SendgridId] = t
	}
	fmt.Println(sendgridTemplates)
}
