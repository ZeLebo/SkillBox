package service

import (
	d "diploma/domain"
	"diploma/internal/entity"
)

var AccendentCollection []entity.AccendentItem

func init() {
	AccendentCollection = ShuffleAccendentData()
}

var AccendentTopics = []string{
	"SMS delivery in EU",
	"MMS connection stability",
	"Voice call connection purity",
	"Checkout page is down",
	"Support overload",
	"Buy phone number not working in US",
	"API Slow latency",
}

func ShuffleAccendentData() []entity.AccendentItem {
	collection := make([]entity.AccendentItem, 0)
	status := ""
	for _, topic := range AccendentTopics {
		if getRandomIntBetweenValues(0, 1) == 1 {
			status = d.AccendentStatusActive
		} else {
			status = d.AccendentStatusClosed
		}

		collection = append(collection, entity.AccendentItem{Topic: topic, Status: status})
	}

	return collection
}
