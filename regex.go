package main

import (
	"github.com/Pallinder/go-randomdata" 
	"github.com/brianvoe/gofakeit"
	"regexp"
	"math/rand"
	"strconv"
	"time"
	"reflect"
)

func Replace(str, rStr string) string {
	var re = regexp.MustCompile(`XXXX XXXX XXXX XXXX|XXXX XXXX XXXX|XXXX XXXX|XXXX|XX/XX/XXXX|XX/XX/`)
	return re.ReplaceAllString(str, rStr)
}

func RandomStr(typeStr string) string {
	rand.Seed(time.Now().UnixNano())

	product := []string{
		"Master Card",
		"Visa",
		"RuPay Card",
    	"American Express",
		"Maestro Card",
	    "Master Card",
		"Visa",
		"RuPay Card",
    	"American Express",
		"Maestro Card",
	}
//10 Random Inputs each.	
	}

	company := []string{
        "ANDHRA BANK",
        "VIJAYA BANK",
        "CANARA BANK",
		"BANK OF BARODA",
        "INDIAN OVERSEAS BANK",
        "THE ORIENTAL BANK OF COMMERCE",
        "PUNJAB NATIONAL BANK",
        "STATE BANK OF INDIA",
        "SYNDICATE BANK",
   		"BANK OF BARODA",
	}
	str := "DEFAULT"

	switch typeStr {
	
	case "/bio/age":
		str = strconv.Itoa(rand.Intn(100))
	case "/profession/job_title":
		str = randomdata.Title(randomdata.RandomGender)
	case "/finance/credit_score":
		str = strconv.Itoa(rand.Intn(200))
	case "/location":
		str = randomdata.Address()
	case "/location/address":
		str = randomdata.Address()
	case "/location/city":
		str = randomdata.City()
	case "/location/zipcode":
		str = gofakeit.Zip()
	case "/id/account_number":
		str = strconv.Itoa(rand.Intn(200))
	case "/id/us_social_security_number":
		str = strconv.Itoa(rand.Intn(500000))
	case "/id/credit_card_number":
		str = strconv.Itoa(gofakeit.CreditCardNumber())
	case "/contact/name":
		str = randomdata.FullName(randomdata.RandomGender)
	case "/contact/email_address":
		str = randomdata.Email()
	case "/contact/phonenumber":
		str = randomdata.PhoneNumber()
	case "/other/number":
		str = strconv.Itoa(rand.Intn(100))
	case "/other/datetime":
		str = randomdata.FullDate()
	case "/other/datetime/date":
		str = strconv.Itoa(rand.Intn(28))
	case "/other/datetime/month":
		str = randomdata.Month()
	case "/other/datetime/year":
		str = strconv.Itoa(gofakeit.Year())
	case "/other/organization":
		str = company[rand.Intn(10)]
	case "/other/organization/company":
		str = company[rand.Intn(10)]
	case "/interest/product":
		str = product[rand.Intn(10)]
	
	
	}

	return str
}

func InArray(val interface{}, array interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				exists = true
				return
			}
		}
	}

	return
}
