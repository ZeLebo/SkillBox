package service

import (
	d "diploma/domain"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

var firstVoiceRowForCorrupt int
var secondVoiceRowForCorrupt int

func ShuffleVoiceData() {
	var data string
	for i, country := range getCountriesList() {
		row := strings.Join([]string{
			country,
			getRandomBandwidthInString(),
			getRandomResponseTimeInString(),
			getVoiceCallProviderByCountry(country),
			getRandomConnectionStability(),
			getRandomTTFB(),
			getRandomVoicePurity(),
			getRandomMedianOfCallsTime(),
		}, ";") + "\n"

		if i == firstVoiceRowForCorrupt || i == secondVoiceRowForCorrupt {
			row = strings.Replace(row, ";", "", rand.Intn(4))
			row = strings.Replace(row, "R", "", rand.Intn(3))
			row = strings.Replace(row, "C", "", rand.Intn(3))

			fmt.Println("Voice row corrupted")
		}

		data += row
	}

	err := ioutil.WriteFile(getFilepathByFilename(d.VoiceFilename), []byte(data), 0644)
	if err != nil {
		fmt.Printf("Error in write sms data: %s", err.Error())
	}
}

func getVoiceCallProviderByCountry(country string) string {
	voiceProviderMap := map[string]string{
		"RU": "TransparentCalls",
		"US": "E-Voice",
		"GB": "TransparentCalls",
		"FR": "TransparentCalls",
		"BL": "E-Voice",
		"AT": "TransparentCalls",
		"BG": "E-Voice",
		"DK": "JustPhone",
		"CA": "JustPhone",
		"ES": "E-Voice",
		"CH": "JustPhone",
		"TR": "TransparentCalls",
		"PE": "JustPhone",
		"NZ": "JustPhone",
		"MC": "E-Voice",
	}

	return voiceProviderMap[country]
}
