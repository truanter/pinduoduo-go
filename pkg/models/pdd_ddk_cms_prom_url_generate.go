package models

type DDKCmsPromUrlGenerateRequest struct {
	GenerateShortUrl  *bool     `json:"generate_short_url,omitempty"`  //是否生成短链接，true-是，false-否
	PIdList           *[]string `json:"p_id_list"`                     //推广位列表，例如：["60005_612"]
	GenerateMobile    *bool     `json:"generate_mobile,omitempty"`     //是否生成手机跳转链接。true-是，false-否，默认false
	MultiGroup        *bool     `json:"multi_group,omitempty"`         //单人团多人团标志。true-多人团，false-单人团 默认false
	CustomParameters  *string   `json:"custom_parameters,omitempty"`   //自定义参数，为链接打上自定义标签。自定义参数最长限制64个字节。
	GenerateSchemaUrl *bool     `json:"generate_schema_url,omitempty"` //是否返回 schema URL
	GenerateWeApp     *bool     `json:"generate_we_app,omitempty"`     //是否生成拼多多福利券微信小程序推广信息
	KeyWord           *string   `json:"keyword,omitempty"`             // 搜索关键词
	ChannelType       *int      `json:"channel_type,omitempty"`        //0, "1.9包邮"；1, "今日爆款"； 2, "品牌清仓"； 3, "默认商城"；4,"PC端专属商城"；非必填 ,默认是3,
}

type DDKCmsPromUrlGenerateResult struct {
	BaseResponse
	CmsPromotionUrlGenerateResponse struct {
		UrlList []CmsPromUrlGenerateInfo `json:"url_list"`
		Total   int                      `json:"total"` //数量
	} `json:"cms_promotion_url_generate_response"`
}

type CmsPromUrlGenerateInfo struct {
	Url                      string    `json:"url"`                          //商城推广链接
	Sign                     string    `json:"sign"`                         //CPSsign
	ShortUrl                 string    `json:"short_url"`                    //商城推广短链接
	MobileUrl                string    `json:"mobile_url"`                   //商城推广移动链接
	MobileShortUrl           string    `json:"mobile_short_url"`             //商城推广移动链接
	MultiGroupUrl            string    `json:"multi_group_url"`              //商城推广多人团链接
	MultiGroupShortUrl       string    `json:"multi_group_short_url"`        //商城推广多人团短链接
	MultiGroupMobileUrl      string    `json:"multi_group_mobile_url"`       //商城推广多人团移动链接
	MultiGroupMobileShortUrl string    `json:"multi_group_mobile_short_url"` //商城推广多人团移动短链接
	WeAppInfo                WeAppInfo `json:"we_app_info"`                  //拼多多福利券微信小程序信息
	MultiUrlList             UrlList   `json:"multi_url_list"`               //双人团链接列表
	SingleUrlList            UrlList   `json:"single_url_list"`              //单人团链接列表
}
