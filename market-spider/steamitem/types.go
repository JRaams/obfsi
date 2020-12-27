package steamitem

// PriceJob is being read from priceJobs.json and contains an url to fetch a list of items from
type PriceJob struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// SteamResponse is used to unmarshal a Steam market response to
type SteamResponse struct {
	Results []SteamItem `json:"results"`
}

// SteamItem is used to unmarshal individual Steam market response result items to
type SteamItem struct {
	Name     string `json:"name"`
	Listings int    `json:"sell_listings"`
	Price    int    `json:"sell_price"`
	Desc     Desc   `json:"asset_description"`
}

// Desc is a child entity of SteamItem
type Desc struct {
	IconURL   string `json:"icon_url"`
	NameColor string `json:"name_color"`
	Type      string `json:"type"`
}

// Asset contains data about a Steam market asset
type Asset struct {
	Name      string `json:"name"`
	NameColor string `json:"nameColor"`
	IconURL   string `json:"iconUrl"`
	Type      string `json:"type"`
}

// Price contains price and listing info about a specific Steam market asset at a point in time
type Price struct {
	ID        int    `json:"id"`
	AssetName string `json:"assets_name"`
	Time      int64  `json:"time"`
	Listings  int    `json:"listings"`
	Price     int    `json:"price"`
}
