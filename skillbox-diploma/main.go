package main

import (
    d "diploma/domain"
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "io/ioutil"
    "math/rand"
    "net/http"
    "strconv"
    "strings"
    "time"
)

var firstSMSRowForCorrupt int
var secondSMSRowForCorrupt int

var firstVoiceRowForCorrupt int
var secondVoiceRowForCorrupt int

var firstEmailRowForCorrupt int
var secondEmailRowForCorrupt int

var MMSCollection []MMSItem
var SupportCollection []SupportItem
var AccendentCollection []AccendentItem

type MMSItem struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type SupportItem struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

type AccendentItem struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

const accendentStatusActive = "active"
const accendentStatusClosed = "closed"

var AccendentTopics = []string{
	"SMS delivery in EU",
	"MMS connection stability",
	"Voice call connection purity",
	"Checkout page is down",
	"Support overload",
	"Buy phone number not working in US",
	"API Slow latency",
}

func init() {
	rand.Seed(time.Now().UnixNano())

	firstSMSRowForCorrupt = rand.Intn(70)
	fmt.Printf("First SMS row for currupt %d\n", firstSMSRowForCorrupt+1)

	secondSMSRowForCorrupt = rand.Intn(90)
	fmt.Printf("Second SMS row for currupt %d\n", secondSMSRowForCorrupt+1)

	firstVoiceRowForCorrupt = rand.Intn(70)
	fmt.Printf("First Voice row for currupt %d\n", firstVoiceRowForCorrupt+1)

	secondVoiceRowForCorrupt = rand.Intn(90)
	fmt.Printf("Second Voice row for currupt %d\n", secondVoiceRowForCorrupt+1)

	firstEmailRowForCorrupt = rand.Intn(70)
	fmt.Printf("First Email row for currupt %d\n", firstEmailRowForCorrupt+1)

	secondEmailRowForCorrupt = rand.Intn(90)
	fmt.Printf("Second Email row for currupt %d\n", secondEmailRowForCorrupt+1)
}

func main() {
	shuffleSmsData()

	MMSCollection = shuffleMMSData()

	shuffleVoiceData()
	shuffleEmailData()
	shuffleBillingData()

	SupportCollection = shuffleSupportData()
	AccendentCollection = shuffleAccendentData()

	listenAndServeHTTP()
}

func shuffleSmsData() {
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

	err := ioutil.WriteFile(getFilapathByFilename(d.SmsFilename), []byte(data), 0644)
	if err != nil {
		fmt.Printf("Error in write sms data: %s", err.Error())
	}
}

func shuffleMMSData() []MMSItem {
	data := make([]MMSItem, 0)
	for _, country := range getCountriesList() {
		data = append(
			data,
			MMSItem{
				Country:      country,
				Provider:     getMMSProviderByCountry(country),
				Bandwidth:    getRandomBandwidthInString(),
				ResponseTime: getRandomResponseTimeInString(),
			},
		)
	}

	return data
}

func shuffleVoiceData() {
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

	err := ioutil.WriteFile(getFilapathByFilename(d.VoiceFilename), []byte(data), 0644)
	if err != nil {
		fmt.Printf("Error in write sms data: %s", err.Error())
	}
}

func shuffleEmailData() {
	var data string
	providersList := getEmailProvidersList()
	i := 0
	for _, country := range getCountriesList() {
		for _, provider := range providersList {
			row := strings.Join([]string{
				country,
				provider,
				getRandomEmailDeliveryTime(),
			}, ";") + "\n"

			if i == firstEmailRowForCorrupt || i == secondEmailRowForCorrupt {
				row = strings.Replace(row, ";", "", rand.Intn(4))
				row = strings.Replace(row, "A", "", rand.Intn(3))
				row = strings.Replace(row, "a", "", rand.Intn(3))
				row = strings.Replace(row, "O", "", rand.Intn(3))
				row = strings.Replace(row, "o", "", rand.Intn(3))
				row = strings.Replace(row, "M", "", rand.Intn(3))
				row = strings.Replace(row, "m", "", rand.Intn(3))
				row = strings.Replace(row, "P", "", rand.Intn(3))
				row = strings.Replace(row, "p", "", rand.Intn(3))

				fmt.Println("Email row corrupted")
			}

			data += row
			i++
		}
	}

	err := ioutil.WriteFile(getFilapathByFilename(d.EmailFilename), []byte(data), 0644)
	if err != nil {
		fmt.Printf("Error in write email data: %s", err.Error())
	}
}

func shuffleBillingData() {
	data := ""
	for i := 0; i < 6; i++ {
		value := getRandomIntBetweenValues(0, 150)
		if value > 50 {
			value = 1
		} else {
			value = 0
		}

		data = data + fmt.Sprintf("%d", value)
		// create customer
		// purchase
		// payout
		// recurring
		// fraud control
		// checkout page
	}

	err := ioutil.WriteFile(getFilapathByFilename(d.BillingFilename), []byte(data), 0644)
	if err != nil {
		fmt.Printf("Error in write sms data: %s", err.Error())
	}
}

func shuffleSupportData() []SupportItem {
	data := make([]SupportItem, 0)
	for _, topic := range getSupportTopicsList() {
		data = append(data, SupportItem{Topic: topic, ActiveTickets: getRandomSupportTickets()})
	}

	return data
}

func shuffleAccendentData() []AccendentItem {
	collection := make([]AccendentItem, 0)
	status := ""
	for _, topic := range AccendentTopics {
		if getRandomIntBetweenValues(0, 1) == 1 {
			status = accendentStatusActive
		} else {
			status = accendentStatusClosed
		}

		collection = append(collection, AccendentItem{Topic: topic, Status: status})
	}

	return collection
}

func getCountriesList() []string {
	return []string{"RU", "US", "GB", "FR", "BL", "AT", "BG", "DK", "CA", "ES", "CH", "TR", "PE", "NZ", "MC"}
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

func getEmailProvidersList() []string {
	return []string{
		"Gmail",
		"Yahoo",
		"Hotmail",
		"MSN",
		"Orange",
		"Comcast",
		"AOL",
		"Live",
		"RediffMail",
		"GMX",
		"Protonmail",
		"Yandex",
		"Mail.ru",
	}
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

func getRandomSupportTickets() int {
	return getRandomIntBetweenValues(0, 8)
}

func getFilapathByFilename(filename string) string {
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

func listenAndServeHTTP() {
	router := mux.NewRouter()

	router.HandleFunc("/mms", handleMMS)
	router.HandleFunc("/support", handleSupport)
	router.HandleFunc("/accendent", handleAccendent)
	router.HandleFunc("/test", handleTest).Methods("GET", "OPTIONS")

	http.ListenAndServe("localhost:8383", router)
}

func handleMMS(w http.ResponseWriter, r *http.Request) {
	response(w, r, MMSCollection)
}

func handleSupport(w http.ResponseWriter, r *http.Request) {
	response(w, r, SupportCollection)
}

func handleAccendent(w http.ResponseWriter, r *http.Request) {
	response(w, r, AccendentCollection)
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write([]byte(d.TestByteString))
}

func response(w http.ResponseWriter, r *http.Request, responseStruct interface{}) {
	response, _ := json.Marshal(responseStruct)

	w.Write(response)
}
