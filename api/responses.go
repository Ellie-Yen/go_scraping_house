package api

type HouseListResponse struct {
	HousePreviews []HousePreview
	TotalCount    int
}

// HousePreview represents a real estate listing with all its details
type HousePreview struct {
	Title      string
	Price      string
	Type       string
	Address    string
	Tags       []string
	ImageURLs  []string
	AgentName  string
	LinkURL    string
	CreateTime string
}
