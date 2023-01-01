package service

import (
	d "diploma/domain"
	"fmt"
	"math/rand"
	"strconv"
)

func getCountriesList() []string {
	return []string{"RU", "US", "GB", "FR", "BL", "AT", "BG", "DK", "CA", "ES", "CH", "TR", "PE", "NZ", "MC"}
}

func getRandomSupportTickets() int {
	return getRandomIntBetweenValues(0, 8)
}

func getFilepathByFilename(filename string) string {
	return "" + filename
}

func getRandomBandwidthInString() string {
	return strconv.Itoa(getRandomIntBetweenValues(d.MinBandwidth, d.MaxBandwidth))
}

func getRandomResponseTimeInString() string {
	return strconv.Itoa(getRandomIntBetweenValues(d.MinResponseTime, d.MaxResponseTime))
}

func getRandomConnectionStability() string {
	stability := getRandomIntBetweenValues(d.MinConnectionStability, d.MaxConnectionStability)

	return fmt.Sprintf("%.2f", float32(stability)/1000)
}

func getRandomTTFB() string {
	return strconv.Itoa(getRandomIntBetweenValues(d.MinTTFB, d.MaxTTFB))
}

func getRandomVoicePurity() string {
	return strconv.Itoa(getRandomIntBetweenValues(d.MinVoicePurity, d.MaxVoicePurity))
}

func getRandomMedianOfCallsTime() string {
	return strconv.Itoa(getRandomIntBetweenValues(d.MinVoiceCallMedian, d.MaxVoiceCallMedian))
}

func getRandomEmailDeliveryTime() string {
	return strconv.Itoa(getRandomIntBetweenValues(d.MinEmailDeliveryTime, d.MaxEmailDeliveryTime))
}

func getRandomIntBetweenValues(min int, max int) int {
	return rand.Intn(max-min) + min
}
