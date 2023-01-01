package service

import (
	d "diploma/domain"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var firstEmailRowForCorrupt int
var secondEmailRowForCorrupt int

func init() {
	rand.Seed(time.Now().UnixNano())

	firstEmailRowForCorrupt = rand.Intn(70)
	fmt.Printf("First Email row for currupt %d\n", firstEmailRowForCorrupt+1)

	secondEmailRowForCorrupt = rand.Intn(90)
	fmt.Printf("Second Email row for currupt %d\n", secondEmailRowForCorrupt+1)
}

func ShuffleEmailData() {
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
				for j := range []string{"A", "a", "O", "o", "M", "m", "P", "p"} {
					row = strings.Replace(row, strconv.Itoa(j), "", rand.Intn(3))
				}

				fmt.Println("Email row corrupted")
			}

			data += row
			i++
		}
	}

	err := ioutil.WriteFile(getFilepathByFilename(d.EmailFilename), []byte(data), 0644)
	if err != nil {
		fmt.Printf("Error in write email data: %s", err.Error())
	}
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
