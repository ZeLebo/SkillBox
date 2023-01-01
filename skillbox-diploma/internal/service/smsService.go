package service

import (
	d "diploma/domain"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

var firstSMSRowForCorrupt int
var secondSMSRowForCorrupt int

func init() {
	rand.Seed(time.Now().UnixNano())

	firstSMSRowForCorrupt = rand.Intn(70)
	fmt.Printf("First SMS row for currupt %d\n", firstSMSRowForCorrupt+1)

	secondSMSRowForCorrupt = rand.Intn(90)
	fmt.Printf("Second SMS row for currupt %d\n", secondSMSRowForCorrupt+1)
}

func ShuffleSmsData() {
	var data string
	for i, country := range getCountriesList() {
		row := strings.Join([]string{
			country,
			getRandomBandwidthInString(),
			getRandomResponseTimeInString(),
			getSmsProviderByCountry(country),
		}, ";") + "\n"

		if i == firstSMSRowForCorrupt || i == secondSMSRowForCorrupt {
			row = strings.Replace(row, ";", "", rand.Intn(4))
			row = strings.Replace(row, "R", "", rand.Intn(3))
			row = strings.Replace(row, "C", "", rand.Intn(3))

			fmt.Println("SMS row corrupted")
		}

		data += row
	}

	err := ioutil.WriteFile(getFilepathByFilename(d.SmsFilename), []byte(data), 0644)
	if err != nil {
		fmt.Printf("Error in write sms data: %s", err.Error())
	}
}

func getSmsProviderByCountry(country string) string {
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
