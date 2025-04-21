package sdk

import (
	"fmt"

	"github.com/Ellie-Yen/go_scraping_house/api"
)

const PAGE_SIZE = 30

func QueryHouseList(priceMin int, priceMax int, saveRaw bool) (*api.HouseListResponse, error) {
	firstPage, err := api.HouseList(priceMin, priceMax, 0, saveRaw)
	if err != nil {
		return nil, err
	}
	totalCount := firstPage.TotalCount
	pageCount := totalCount / PAGE_SIZE

	// use goroutine to query house list
	ch := make(chan *api.HouseListResponse, pageCount)
	for i := range pageCount {
		go func(i int) {
			resp, err := api.HouseList(priceMin, priceMax, i, saveRaw)
			if err != nil {
				fmt.Println("error: ", err)
			}
			ch <- resp
		}(i)
	}

	// collect all house list
	houseList := firstPage.HousePreviews
	for _ = range pageCount {
		resp := <-ch
		houseList = append(houseList, resp.HousePreviews...)
	}
	firstPage.HousePreviews = houseList

	return firstPage, nil
}
