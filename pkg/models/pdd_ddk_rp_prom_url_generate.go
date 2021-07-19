package models

type DDKRpPromUrlGenerateRequest struct {
	Amount           *int64  `json:"amount,omitempty"`            //初始金额（单位分），有效金额枚举值：300、500、700、1100和1600，默认300
	ChannelType      *int    `json:"channel_type,omitempty"`      //营销工具类型
	CustomParameters *string `json:"custom_parameters,omitempty"` //自定义参数，为链接打上自定义标签；自定义参数最长限制64个字节；格式为： {"uid":"11111","sid":"22222"}
	DiyOneYuanParam  *struct {
		GoodsSign *string `json:"goods_sign,omitempty"` //商品goodsSign，支持通过goodsSign查询商品
	} `json:"diy_one_yuan_param,omitempty"` //一元购自定义参数，json格式
	DiyRedPacketParam *struct {
		AmountProbability *[]int64 `json:"amount_probability,omitempty"`  //红包金额列表
		DisText           *bool    `json:"dis_text,omitempty"`            //设置玩法，false-现金红包, true-现金券
		NotShowBackground *bool    `json:"not_show_background,omitempty"` //推广页设置，false-红包开启页, true-红包领取页
		OptID             *int     `json:"opt_id,omitempty"`              //优先展示类目
		RangeItems        *struct {
			RangeFrom *int64 `json:"range_from,omitempty"`
			RangeID   *int   `json:"range_id,omitempty"`
			RangeTo   *int64 `json:"range_to,omitempty"`
		} `json:"range_items,omitempty"`
	} `json:"diy_red_packet_param,omitempty"`
	GenerateQQApp     *bool     `json:"generate_qq_app,omitempty"`
	GenerateWeApp     *bool     `json:"generate_we_app,omitempty"`
	GenerateSchemaUrl *bool     `json:"generate_schema_url,omitempty"`
	GenerateShortUrl  *bool     `json:"generate_short_url,omitempty"`
	PIDList           *[]string `json:"p_id_list,omitempty"`
	ScratchCardAmount *int64    `json:"scratch_card_amount,omitempty"`
}

type DDKRpPromotionUrlGenerateResult struct {
	BaseResponse
	RpPromotionUrlGenerateResponse struct {
		ResourceList []struct {
			Desc string `json:"desc"`
			Url  string `json:"url"`
		} `json:"resource_list"`
		UrlList []struct {
			MobileShortUrl           string    `json:"mobile_short_url"`
			MobileUrl                string    `json:"mobile_url"`
			MultiGroupMobileShortUrl string    `json:"multi_group_mobile_short_url"`
			MultiGroupMobileUrl      string    `json:"multi_group_mobile_url"`
			SchemaUrl                string    `json:"schema_url"`
			ShortUrl                 string    `json:"short_url"`
			Url                      string    `json:"url"`
			QQAppInfo                QQAppInfo `json:"qq_app_info"`
			WeAppInfo                WeAppInfo `json:"we_app_info"`
		} `json:"url_list"`
	} `json:"rp_promotion_url_generate_response"`
}
