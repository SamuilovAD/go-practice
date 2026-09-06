package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Event struct {
	UserID string
	Amount int
	Status string
}
type UserTotal struct {
	UserID string
	Total  int
}

func main() {
	events := []Event{
		{UserID: "alice", Amount: 100, Status: "success"},
		{UserID: "bob", Amount: 200, Status: "success"},
		{UserID: "alice", Amount: 150, Status: "success"},
		{UserID: "charlie", Amount: 500, Status: "failed"},
		{UserID: "david", Amount: 250, Status: "success"},
		{UserID: "bob", Amount: 50, Status: "success"},
	}
	limit := 2
	fmt.Print(TopUsers(events, limit))
}
func TopUsers(events []Event, limit int) []UserTotal {
	userTotal := []UserTotal{}
	if limit <= 0 {
		return userTotal
	}
	eventsMap := make(map[string]int)
	for _, event := range events {
		if event.Status != "success" {
			continue
		}
		eventsMap[event.UserID] += event.Amount
	}
	for userIdMap, userTotalAmountMap := range eventsMap {
		userTotal = append(userTotal, UserTotal{
			UserID: userIdMap,
			Total:  userTotalAmountMap,
		})
	}
	slices.SortFunc(userTotal, func(a, b UserTotal) int {
		if n := cmp.Compare(b.Total, a.Total); n != 0 {
			return n
		}

		return cmp.Compare(a.UserID, b.UserID)
	})
	if limit > len(userTotal) {
		limit = len(userTotal)
	}

	return userTotal[0:limit]
}
