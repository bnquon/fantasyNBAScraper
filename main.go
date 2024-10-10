package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
)

func main() {

	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)

	defer cancel()	

	var tableData string

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://fantasy.espn.com/basketball/leaders?leagueId=264010336&statSplit=lastSeason&scoringPeriodId=0&view=stats"),
		chromedp.WaitVisible(`body`),
		chromedp.OuterHTML("div.jsx-1811044066.player-column-table2.justify-start.pa0.relative.flex.items-center.player-info", &tableData),
	)
	if err != nil {
		fmt.Println("Error scraping data")
		return
	}

	fmt.Println(tableData)
}