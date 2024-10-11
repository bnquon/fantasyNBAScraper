package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

// type Player struct {
// 	name string
// 	fantasyPoints float32
// }

func main() {

	// var topPlayers []Player
	// fmt.Println(topPlayers)

	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string {
		"--headless-new",
	}})
	
	driver, err := selenium.NewRemote(caps, "")

	if err != nil {
		log.Fatal("Error:", err)
	}

	err = driver.MaximizeWindow("")
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = driver.Get("https://fantasy.espn.com/basketball/leaders")
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = driver.WaitWithTimeout(func(driver selenium.WebDriver) (bool, error) {
		lastPlayer, _ := driver.FindElement(selenium.ByCSSSelector, "tbody.Table__TBODY:last-child")
		if lastPlayer != nil {
			return lastPlayer.IsDisplayed()
		}
		return false, nil
	}, 10*time.Second)

	if err != nil {
		log.Fatal("Error:", err)
	}

	players, err := driver.FindElements(selenium.ByCSSSelector, "tbody.Table__TBODY")
	if err != nil {
	 log.Fatal("Error:", err)
	}

	fmt.Println("Found the third tbody:", players)

	// for _, productElement  := range productElements {
	// 	nameElement, err := productElement.FindElement(selenium.ByCSSSelector, "h4")
	// 	priceElement, err := productElement.FindElement(selenium.ByCSSSelector, "h5")
	// 	name, err := nameElement.Text()
	// 	price, err := priceElement.Text()
	// 	if err != nil {
	// 		log.Fatal("Error:", err)
	// 	}
	// 	product := Product{}
	// 	product.name = name
	// 	product.price = price
	// 	products = append(products, product)
	// }
	
	// file, err := os.Create("products.csv")
	// if err != nil {
	// 	log.Fatal("Error:", err)
	// }

	// defer file.Close()

	// writer := csv.NewWriter(file)

	// headers := []string{
	// 	"name",
	// 	"price",
	// }

	// writer.Write(headers)

	// for _, product := range products {
	// 	record := []string{
	// 		product.name,
	// 		product.price,
	// 	}
	// 	writer.Write(record)
	// }

	// defer writer.Flush()
}