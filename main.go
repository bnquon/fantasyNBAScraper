package main

import (
	"encoding/csv"
	"log"
	"os"
	"time"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type Player struct {
	name string
	fantasyPoints string
}

func main() {

	var topPlayers []Player

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
	
	playerNamesRows, err := players[0].FindElements(selenium.ByTagName, "tr")
	if err != nil {
		log.Fatal("Error:", err)
	}
	for _, row := range playerNamesRows {
		text, err := row.Text()
		if err != nil {
			log.Fatal("Error:", err)
		}
		var temp string;
		for _, char := range text {
			if char == '\n' {
				break
			}
			temp += string(char)
		}
		player := Player{}
		player.name = temp
		topPlayers = append(topPlayers, player)
	}
	
	fantasyPointsRows, err := players[2].FindElements(selenium.ByTagName, "tr")
	if err != nil {
		log.Fatal("Error:", err)
	}
	for index, row := range fantasyPointsRows {
		text, err := row.Text()
		if err != nil {
			log.Fatal("Error:", err)
		}
		topPlayers[index].fantasyPoints = text
	}
	
	file, err := os.Create("players.csv")
	if err != nil {
		log.Fatal("Error:", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	headers := []string{
		"name",
		"fantasyPoints",
	}

	writer.Write(headers)

	for _, player := range topPlayers {
		record := []string{
			player.name,
			player.fantasyPoints,
		}
		writer.Write(record)
	}

	defer writer.Flush()
}