package api

import (
	"fmt"
)

const PAGE_SIZE = 30

// FIXME: current the api has some issue, some data is not fetched correctly
func QueryHouseList(priceMin int, priceMax int, saveRaw bool) (*HouseListResponse, error) {
	firstPage, err := houseList(priceMin, priceMax, 0, saveRaw)
	if err != nil {
		return nil, err
	}
	totalCount := firstPage.TotalCount
	pageCount := totalCount / PAGE_SIZE

	// use goroutine to query house list
	ch := make(chan *HouseListResponse, pageCount)
	for i := range pageCount {
		go func(i int) {
			resp, err := houseList(priceMin, priceMax, i, saveRaw)
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
