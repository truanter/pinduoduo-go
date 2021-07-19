package models

type BaseResponse struct {
	ErrorResponse ErrorResponse `json:"error_response,omitempty"`
}

type WeAppInfo struct {
	WeAppIconUrl      string `json:"we_app_icon_url"`
	BannerUrl         string `json:"banner_url"`
	Desc              string `json:"desc"`
	SourceDisplayName string `json:"source_display_name"`
	PagePath          string `json:"page_path"`
	UserName          string `json:"user_name"`
	Title             string `json:"title"`
	AppId             string `json:"app_id"`
}

type QQAppInfo struct {
	AppId             string `json:"app_id"`
	BannerUrl         string `json:"banner_url"`
	Desc              string `json:"desc"`
	PagePath          string `json:"page_path"`
	QqAppIconUrl      string `json:"qq_app_icon_url"`
	SourceDisplayName string `json:"source_display_name"`
	UserName          string `json:"user_name"`
	Title             string `json:"title"`
}

type UrlList struct {
	MobileShortUrl string `json:"mobile_short_url"`
	MobileUrl      string `json:"mobile_url"`
	SchemaUrl      string `json:"schema_url"`
	ShortUrl       string `json:"short_url"`
	Url            string `json:"url"`
}
