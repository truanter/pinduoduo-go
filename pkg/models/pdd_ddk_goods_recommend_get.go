package models

type DDKGoodsRecommendResult struct {
	BaseResponse
	GoodsBasicDetailResponse struct {
		List     []GoodsRecommendInfo `json:"list"`
		ListID   string               `json:"list_id"`
		SearchID string               `json:"search_id"`
		Total    string               `json:"total"`
	} `json:"goods_basic_detail_response"`
}
type DDKGoodsRecommendRequest struct {
	ActivityTags     *[]int    `json:"activity_tags,omitempty"`     //活动商品标记数组，例：[4,7]，4-秒杀，7-百亿补贴，10851-千万补贴，10913-招商礼金商品，31-品牌黑标，10564-精选爆品-官方直推爆款，10584-精选爆品-团长推荐，24-品牌高佣，其他的值请忽略
	CatID            *int64    `json:"cat_id,omitempty"`            //猜你喜欢场景的商品类目，20100-百货，20200-母婴，20300-食品，20400-女装，20500-电器，20600-鞋包，20700-内衣，20800-美妆，20900-男装，21000-水果，21100-家纺，21200-文具,21300-运动,21400-虚拟,21500-汽车,21600-家装,21700-家具,21800-医药;
	CustomParameters *string   `json:"custom_parameters,omitempty"` //自定义参数，为链接打上自定义标签；自定义参数最长限制64个字节；格式为： {"uid":"11111","sid":"22222"} ，其中 uid 为用户唯一标识，可自行加密后传入，每个用户仅且对应一个标识，必填； sid 为上下文信息标识，例如sessionId等，非必填。该json字符串中也可以加入其他自定义的key。
	GoodsImgType     *int      `json:"goods_img_type,omitempty"`    //商品主图类型：1-场景图，2-白底图，默认为0
	GoodsSignList    *[]string `json:"goods_sign_list,omitempty"`   //商品goodsSign列表，相似商品推荐场景时必传，仅取数组的第一位
	ListID           *string   `json:"list_id,omitempty"`           //翻页时建议填写前页返回的list_id值
	PID              *string   `json:"pid"`                         //推广位id
	Offset           *int      `json:"offset,omitempty"`            //从多少位置开始请求；默认值 ： 0
	Limit            *int      `json:"limit,omitempty"`             //请求数量；默认值 ： 400
	ChannelType      *int      `json:"channel_type,omitempty"`      //频道类型；0, "1.9包邮", 1, "今日爆款", 2, "品牌清仓", 3, "默认商城", 非必填 ,默认是1
}

type GoodsRecommendInfo struct {
	QrCodeImageUrl       string  `json:"qr_code_image_url"`       //二维码主图
	LockEdit             string  `json:"lock_edit"`               //编辑锁定
	Broker               string  `json:"broker"`                  //代理人
	Rank                 string  `json:"rank"`                    //顺序
	SaleNumToday         int     `json:"sale_num_today"`          //今日成交量
	SaleNum24            int     `json:"sale_num24"`              //成交量
	ShareDesc            string  `json:"share_desc"`              //分享描述
	CatID                int64   `json:"cat_id"`                  //商品类目id
	GoodsEvalCount       int     `json:"goods_eval_count"`        //商品评价数量
	MarketFee            int     `json:"market_fee"`              //市场服务费
	GoodsRate            int     `json:"goods_rate"`              //商品等级
	CouponPrice          int     `json:"coupon_price"`            //优惠券金额 分
	PromotionRate        int     `json:"promotion_rate"`          //佣金比例,千分比
	CouponEndTime        int64   `json:"coupon_end_time"`         //优惠券失效时间,UNIX时间戳
	CouponStartTime      int64   `json:"coupon_start_time"`       //优惠券生效时间,UNIX时间戳
	CouponRemainQuantity int     `json:"coupon_remain_quantity"`  //优惠券剩余数量
	CouponTotalQuantity  int     `json:"coupon_total_quantity"`   //优惠券总数量
	CouponDiscount       int     `json:"coupon_discount"`         //优惠券面额,单位为分
	CouponMinOrderAmount int     `json:"coupon_min_order_amount"` //优惠券门槛价格,单位为分
	CouponID             int64   `json:"coupon_id"`               //优惠券id
	HasCoupon            bool    `json:"has_coupon"`              //商品是否带券,true-带券,false-不带券
	GoodsType            int     `json:"goods_type"`              //商品类型
	CatIds               []int64 `json:"cat_ids"`                 //商品一~四级类目ID列表
	OptIds               []int64 `json:"opt_ids"`                 //商品一~四级标签类目ID列表
	OptName              string  `json:"opt_name"`                //商品标签名
	OptID                int64   `json:"opt_id"`                  //商品标签类目ID,使用pdd.goods.opt.get获取
	CategoryName         string  `json:"category_name"`           //分类名称
	CategoryID           int64   `json:"category_id"`             //类目id
	MerchantType         int     `json:"merchant_type"`           //商家类型
	MallName             string  `json:"mall_name"`               //店铺名称
	MallID               int64   `json:"mall_id"`                 //商家id
	MinNormalPrice       int     `json:"min_normal_price"`        //最小单买价格，单位分
	MinGroupPrice        int     `json:"min_group_price"`         //最小成团价格，单位分
	GoodsFactPrice       int     `json:"goods_fact_price"`        //商品实际价格
	GoodsMarkPrice       int     `json:"goods_mark_price"`        //商品标准价格
	SoldQuantity         int     `json:"sold_quantity"`           //销售量
	GoodsGalleryUrls     string  `json:"goods_gallery_urls"`      //商品详情图列表
	GoodsImageURL        string  `json:"goods_image_url"`         //商品主图
	GoodsThumbnailUrl    string  `json:"goods_thumbnail_url"`     //商品缩略图
	GoodsDesc            string  `json:"goods_desc"`              //商品描述
	GoodsName            string  `json:"goods_name"`              //商品名称
	GoodsID              int64   `json:"goods_id"`                //商品id
	CreateAt             int64   `json:"create_at"`               //创建时间
	DescTxt              string  `json:"desc_txt"`
	ServTxt              string  `json:"serv_txt"`
	LgstTxt              string  `json:"lgst_txt"`
	SearchID             string  `json:"search_id"`
	GoodsSign            string  `json:"goods_sign"`
}
