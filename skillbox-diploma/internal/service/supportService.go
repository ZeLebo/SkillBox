package service

import "diploma/internal/entity"

var SupportCollection []entity.SupportItem

func init() {
	SupportCollection = ShuffleSupportData()
}

func ShuffleSupportData() []entity.SupportItem {
	data := make([]entity.SupportItem, 0)
	for _, topic := range getSupportTopicsList() {
		data = append(data, entity.SupportItem{Topic: topic, ActiveTickets: getRandomSupportTickets()})
	}

	return data
}

func getSupportTopicsList() []string {
	return []string{
		"SMS",
		"MMS",
		"Email",
		"Billing",
		"Create account",
		"API",
		"Marketing",
		"Privacy",
		"GDPR",
		"Other",
	}
}
