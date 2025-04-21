package view

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Ellie-Yen/go_scraping_house/api"
	"github.com/Ellie-Yen/go_scraping_house/utils"
)

func GenView(resp api.HouseListResponse) {
	items := filterDisplayItem(resp)
	html := toHtml(resp.TotalCount, items)
	utils.SaveFile("test.html", html)
}

func filterDisplayItem(resp api.HouseListResponse) []api.HousePreview {
	excludeWords := []string{
		"萬華",
		"內湖",
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

func toHtml(totalCount int, items []api.HousePreview) string {
	html := "<html>\n"
	html += fmt.Sprintf("<div>TotalCount: %v / (%v)</div>\n", len(items), totalCount)
	for i, p := range items {
		html += ("<section>\n" +
			fmt.Sprintf("<h1>%v</h1>\n", i) +
			item(p) +
			"</section>\n")
	}
	return html
}

func item(p api.HousePreview) string {
	return (fmt.Sprintf("<div>Title: %s</div>\n", p.Title) +
		fmt.Sprintf("<a href=\"%s\">link</a>\n", p.LinkURL) +
		fmt.Sprintf("<div>Price: %s</div>\n", p.Price) +
		fmt.Sprintf("<div>Type: %s</div>\n", p.Type) +
		fmt.Sprintf("<div>Address: %s</div>\n", p.Address) +
		fmt.Sprintf("<div>Agent: %s</div>\n", p.AgentName) +
		fmt.Sprintf("<div>Created: %s</div>\n", p.CreateTime) +

		fmt.Sprintf("<div>Tags: %s</div>\n", arrTransform(p.Tags, tagElement)) +
		fmt.Sprintf("<div>Image: %s</div>\n", arrTransform(p.ImageURLs, imageElement)))
}

func arrTransform(arr []string, transform func(s string) string) string {
	res := ""
	for _, x := range arr {
		res += transform(x)
	}
	return res
}

func tagElement(tag string) string {
	return tag + ","
}

func imageElement(src string) string {
	return fmt.Sprintf("<image src=\"%s\"/>", src)
}
