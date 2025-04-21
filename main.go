package main

import (
	"github.com/Ellie-Yen/go_scraping_house/sdk"
	"github.com/Ellie-Yen/go_scraping_house/view"
)

func main() {
	items, err := sdk.QueryHouseList(15000, 25000, false)
	if err != nil {
		return
	}
	view.GenView(*items)
}
