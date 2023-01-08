package service

import (
	"diploma/internal/entity"
	"fmt"
	"math/rand"
	"time"
)

var MMSCollection []entity.MMSItem

func init() {
	rand.Seed(time.Now().UnixNano())
	firstVoiceRowForCorrupt = rand.Intn(70)
	fmt.Printf("First Voice row for currupt %d\n", firstVoiceRowForCorrupt+1)

	secondVoiceRowForCorrupt = rand.Intn(90)
	fmt.Printf("Second Voice row for currupt %d\n", secondVoiceRowForCorrupt+1)

	MMSCollection = ShuffleMMSData()
}

func ShuffleMMSData() []entity.MMSItem {
	data := make([]entity.MMSItem, 0)
	for _, country := range getCountriesList() {
		data = append(
			data,
			entity.MMSItem{
				Country:      country,
				Provider:     getMMSProviderByCountry(country),
				Bandwidth:    getRandomBandwidthInString(),
				ResponseTime: getRandomResponseTimeInString(),
			},
		)
	}

	return data
}

func getMMSProviderByCountry(country string) string {
	smsProviderMap := map[string]string{
		"RU": "Topolo",
		"US": "Rond",
		"GB": "Topolo",
		"FR": "Topolo",
		"BL": "Kildy",
		"AT": "Topolo",
		"BG": "Rond",
		"DK": "Topolo",
		"CA": "Rond",
		"ES": "Topolo",
		"CH": "Topolo",
		"TR": "Rond",
		"PE": "Topolo",
		"NZ": "Kildy",
		"MC": "Kildy",
	}

	return smsProviderMap[country]
}
