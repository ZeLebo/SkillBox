package service

import (
	d "diploma/domain"
	"fmt"
	"io/ioutil"
)

func ShuffleBillingData() {
	data := ""
	for i := 0; i < 6; i++ {
		value := 0
		if getRandomIntBetweenValues(0, 150) > 50 {
			value = 1
		}
		data += fmt.Sprintf("%d", value)
		// create customer
		// purchase
		// payout
		// recurring
		// fraud control
		// checkout page
	}

	err := ioutil.WriteFile(getFilepathByFilename(d.BillingFilename), []byte(data), 0644)
	if err != nil {
		fmt.Printf("Error in write sms data: %s", err.Error())
	}
}
