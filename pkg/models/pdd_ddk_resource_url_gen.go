package models

type DDkResourceUrlGenRequest struct {
	Pid          string  `json:"pid"`
	ResourceType *int    `json:"resource_type"`
	Url          *string `json:"url"`
}

type DDKResourceUrlGenResult struct {
	BaseResponse
	ResourceUrlResponse struct {
		SingleUrlList ResourceUrlGenUrlList `json:"single_url_list"`
		MultiUrlList  ResourceUrlGenUrlList `json:"multi_url_list"`
		Sign          string                `json:"sign"`
		WeAppInfo     WeAppInfo             `json:"we_app_info"`
	} `json:"resource_url_response"`
}

type ResourceUrlGenUrlList struct {
	URL                  string `json:"url"`
	ShortUrl             string `json:"short_url"`
	MobileUrl            string `json:"mobile_url"`
	MobileShortUrl       string `json:"mobile_short_url"`
	WeAppWebViewUrl      string `json:"we_app_web_view_url"`
	WeAppWebViewShortUrl string `json:"we_app_web_view_short_url"`
	WeAppPagePath        string `json:"we_app_page_path"`
}
