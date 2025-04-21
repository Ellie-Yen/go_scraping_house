package api

var DEFAULT_HEADERS = [][]string{
	{"Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
	{"Accept-Language", "en-GB,en;q=0.9"},
	{"Connection", "keep-alive"},
	{"DNT", "1"},
	{"Referer", "https://rent.591.com.tw/list/?region=1&price=15000$_25000$&other=lift&sort=posttime_desc"},
	{"Sec-Fetch-Dest", "document"},
	{"Sec-Fetch-Mode", "navigate"},
	{"Sec-Fetch-Site", "same-origin"},
	{"Sec-Fetch-User", "?1"},
	{"Upgrade-Insecure-Requests", "1"},
	{"User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36"},
	{"sec-ch-ua", "\"Google Chrome\";v=\"135\", \"Not-A.Brand\";v=\"8\", \"Chromium\";v=\"135\""},
	{"sec-ch-ua-mobile", "?0"},
	{"sec-ch-ua-platform", "\"macOS\""},
}

var DEFAULT_COOKIES = [][]string{
	{"webp", "1"},
	{"T591_TOKEN", "6mvmp00FCRG-XmQqDNDSx"},
	{"urlJumpIp", "1"},
	{"_gcl_au", "1.1.380720769.1744461733"},
	{"_ga", "GA1.1.1688184761.1744461733"},
	{"__lt__cid", "ec0a2173-8b11-4046-840d-1afcc5a002fd"},
	{"__lt__sid", "a635d837-565c57aa"},
	{"T591_TOKEN", "3o4o3j4i4fd9nhnpbm7hs7lemr"},
	{"PHPSESSID", "dkh597s1ekhmq88un4ndjk5m28"},
	{"__loc__", "MTIzLjE5Mi4yMDUuMTYx"},
	{"591_new_session", "eyJpdiI6ImNRUkxORUovcXQzK1lhckVTUyt0dWc9PSIsInZhbHVlIjoiWDNkRVFRSVpzNitFdkJGeFVGemQvcWpkZlZkZm1OWDJzR1FxZnQvcUlOaDFnU2RjZFBJZkg3TjArNi9KSGpkaEJTYTlGSGRhZGJkVS9RUFd6RXIrT1RXQlpOMElyZXl5UEVQdE5mOTFCeHNmRjNSb1JlanFubUNNYWhZK1JaSXoiLCJtYWMiOiI5NzYyMjI4NDM3ZTQwMTNlZmM4MDU4YzY5MTQwZDBlNDc5NjE0OTVlNDY5MDE5ZDEwYTYyNjE4YzRjNzk2NGI5IiwidGFnIjoiIn0%3D"},
	{"timeDifference", "-1"},
	{"__one_id__", "01JP00EGMTE2EYT44DR4722DDZ"},
	{"_ga_HDSPSZ773Q", "GS1.1.1744461733.1.1.1744461762.31.0.0"},
}

const BASE_URL = "https://rent.591.com.tw"
