package controllers

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/Ellie-Yen/go_scraping_house/api"
	"github.com/gin-gonic/gin"
)

func HouseList(c *gin.Context) {

	priceMin := 15000
	priceMax := 25000

	items, err := api.QueryHouseList(priceMin, priceMax, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	filterItems := filterDisplayItem(*items)
	c.HTML(http.StatusOK, "house_list.html", gin.H{
		"title":        fmt.Sprintf("%d - %d", priceMin, priceMax),
		"filter_count": len(filterItems),
		"total_count":  len(items.HousePreviews),
		"price_min":    priceMin,
		"price_max":    priceMax,
		"items":        filterItems,
	})
}

func filterDisplayItem(resp api.HouseListResponse) []api.HousePreview {
	excludeWords := []string{
		"萬華",
		"內湖",
		"天母",
		"明德",
		"大橋頭",
		"北投",
		"樓中樓",
		"挑高",
		"3米6",
		"奇岩",
		"和益金銀大樓",
		"復興SOGO",
		"寄居蟹",
		"頂樓",
		"頂加",
		"頂樓加蓋",
		"共生",
	}

	items := make([]api.HousePreview, 0)
	uniqueItems := make(map[string]bool)
	for _, p := range resp.HousePreviews {
		if len(p.ImageURLs) == 0 {
			continue
		}
		isExclude := false
		for _, word := range excludeWords {
			if strings.Contains(p.Title, word) {
				isExclude = true
				break
			}
			if strings.Contains(p.Address, word) {
				isExclude = true
				break
			}
		}
		if isExclude {
			continue
		}
		if p.Type == "雅房" {
			continue
		}
		if _, existed := uniqueItems[p.Title]; existed {
			continue
		}
		uniqueItems[p.Title] = true
		items = append(items, p)
	}
	// sort by create time desc
	sort.Slice(items, func(i, j int) bool {
		return items[i].CreateTime > items[j].CreateTime
	})
	return items
}
