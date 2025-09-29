package api

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/Ellie-Yen/go_scraping_house/utils"
	"github.com/PuerkitoBio/goquery"
)

func houseList(priceMin int, priceMax int, page int, saveRaw bool) (*HouseListResponse, error) {
	// Create a new request
	url := fmt.Sprintf("%s/list?region=1&price=%v$_%v$&other=lift&sort=posttime_desc", BASE_URL, priceMin, priceMax)
	if page > 0 {
		url += fmt.Sprintf("&page=%v", page)
	}
	fmt.Println(url)
	req := &Request{
		Method:  "GET",
		Url:     url,
		Headers: DEFAULT_HEADERS,
		Cookies: DEFAULT_COOKIES,
		Body:    nil,
	}
	body, err := Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	html := string(body)
	if saveRaw {
		utils.SaveFile(fmt.Sprintf("resp_%v.html", page), html)
	}
	listings, err := parse(html)
	if err != nil {
		log.Fatalf("Error parsing HTML: %v", err)
	}
	return listings, nil
}

// parse extracts all property listings from the HTML
func parse(html string) (*HouseListResponse, error) {
	// Create a goquery document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("error loading HTML: %v", err)
	}

	response := &HouseListResponse{
		TotalCount:    0,
		HousePreviews: []HousePreview{},
	}

	// Extract total count from the rent-list object
	totalCount := doc.Text()
	r := regexp.MustCompile(`total:\s*"(\d+)"`)
	matches := r.FindStringSubmatch(totalCount)
	if len(matches) > 1 {
		count, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing total count: %v", err)
		}
		response.TotalCount = count
	}

	// Find all listing items
	doc.Find(".item").Each(func(i int, s *goquery.Selection) {
		item := HousePreview{}

		// Extract title
		item.Title = strings.TrimSpace(s.Find(".item-info-title a").AttrOr("title", ""))

		// Extract link URL
		item.LinkURL = s.Find(".item-info-title a").AttrOr("href", "")

		// Extract price
		priceText := s.Find(".price-info .price.font-arial").Text()
		item.Price = strings.TrimSpace(priceText)

		// Extract property type
		item.Type = strings.TrimSpace(s.Find(".ic-house.house-home").Parent().Text())

		// Extract address
		item.Address = strings.TrimSpace(s.Find(".ic-house.house-place").Parent().Text())

		// Extract agent info and metadata
		s.Find(".role-name span").Each(func(j int, meta *goquery.Selection) {
			text := strings.TrimSpace(meta.Text())

			if j == 0 {
				// First span is agent name
				item.AgentName = text
			}
		})

		// Extract tags
		s.Find(".item-info-tag .tag").Each(func(_ int, tag *goquery.Selection) {
			tagText := strings.TrimSpace(tag.Text())
			if tagText != "" {
				item.Tags = append(item.Tags, tagText)
			}
		})

		// Extract image URLs
		s.Find(".image-list li img").Each(func(_ int, img *goquery.Selection) {
			// Get the real image URL from data-src attribute
			dataSrc := img.AttrOr("data-src", "")
			if dataSrc != "" {
				item.ImageURLs = append(item.ImageURLs, dataSrc)
			}
		})

		// use image url to get create time, use the oldest one
		if len(item.ImageURLs) > 0 {
			r := regexp.MustCompile(`(\d{4}/\d{2}/\d{2})`)
			for _, img := range item.ImageURLs {
				matches := r.FindAllStringSubmatch(img, -1)
				if len(matches) <= 0 {
					continue
				}

				t := matches[0][0]
				if len(item.CreateTime) == 0 || t < item.CreateTime {
					item.CreateTime = t
				}

			}
		}

		response.HousePreviews = append(response.HousePreviews, item)
	})

	return response, nil
}
