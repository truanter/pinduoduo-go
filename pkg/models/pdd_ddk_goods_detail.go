package models

type DDKGoodsDetailRequest struct {
	Pid              *string `json:"pid,omitempty"`               //自定义参数
	CustomParameters *string `json:"custom_parameters,omitempty"` //自定义参数，为链接打上自定义标签
	GoodsImgType     *int    `json:"goods_img_type,omitempty"`    //商品主图类型：1-场景图，2-白底图，默认为0
	ZsDuoID          *int64  `json:"zs_duo_id,omitempty"`         //	招商多多客ID
	SearchId         *string `json:"search_id,omitempty"`         //搜索id，建议填写，提高收益。来自pdd.ddk.goods.recommend.get、pdd.ddk.goods.search、pdd.ddk.top.goods.list.query等接口
	GoodsSign        *string `json:"goods_sign"`
}

type DDKGoodsDetailResult struct {
	BaseResponse
	GoodsDetailResponse struct {
		GoodsDetails []GoodsDetailInfo `json:"goods_details"` //多多进宝商品对象列表
	} `json:"goods_detail_response"`
}

type GoodsDetailInfo struct {
	ActivityPromotionRate       int64    `json:"activity_promotion_rate,omitempty"`        //活动佣金比例，千分比（特定活动期间的佣金比例）
	ActivityTags                []int    `json:"activity_tags,omitempty"`                  //商品活动标记数组，例：[4,7]，4-秒杀，7-百亿补贴，24-品牌高佣，20-行业精选，21-金牌商家，10044-潜力爆品，10475-爆品上新，其他的值请忽略
	BrandName                   string   `json:"brand_name,omitempty"`                     //商品品牌词信息
	CategoryID                  int64    `json:"category_id"`                              //商品类目ID，使用pdd.goods.cats.get接口获取
	CategoryName                string   `json:"category_name"`                            //商品类目名
	MallCouponID                int64    `json:"mall_coupon_id"`                           //店铺券id
	MallCouponDiscountPct       int      `json:"mall_coupon_discount_pct"`                 //店铺券折扣
	MallCouponMinOrderAmount    int64    `json:"mall_coupon_min_order_amount"`             //最小使用金额
	MallCouponMaxDiscountAmount int      `json:"mall_coupon_max_discount_amount"`          //最大使用金额
	MallCouponTotalQuantity     int64    `json:"mall_coupon_total_quantity"`               //店铺券总量
	MallCouponRemainQuantity    int64    `json:"mall_coupon_remain_quantity"`              //店铺券余量
	MallCouponStartTime         int64    `json:"mall_coupon_start_time"`                   //店铺券开始使用时间
	MallCouponEndTime           int64    `json:"mall_coupon_end_time"`                     //店铺券结束使用时间
	MallID                      int64    `json:"mall_id"`                                  //店铺id
	MallName                    string   `json:"mall_name"`                                //店铺名称
	OptIds                      []int64  `json:"opt_ids"`                                  //商品标签ID
	CatIds                      []int64  `json:"cat_ids"`                                  //商品一~四级类目ID列表
	CatID                       int64    `json:"cat_id"`                                   //商品类目ID，使用pdd.goods.cats.get接口获取
	GoodsEvalCount              int      `json:"goods_eval_count"`                         //商品评价数
	OptName                     string   `json:"opt_name"`                                 //商品标签名称
	CouponDiscount              int64    `json:"coupon_discount"`                          //优惠券面额，单位为分
	CouponRemainQuantity        int      `json:"coupon_remain_quantity"`                   //优惠券剩余数量
	CouponMinOrderAmount        int      `json:"coupon_min_order_amount"`                  //优惠券门槛金额，单位为分
	CouponTotalQuantity         int      `json:"coupon_total_quantity"`                    //优惠券总数量
	CouponStartTime             int64    `json:"coupon_start_time"`                        //优惠券生效时间，UNIX时间戳
	CouponEndTime               int64    `json:"coupon_end_time"`                          //优惠券生效时间，UNIX时间戳
	PromotionRate               int      `json:"promotion_rate"`                           //佣金比例，千分比
	OptID                       int64    `json:"opt_id"`                                   //商品标签ID，使用pdd.goods.opt.get接口获取
	GoodsID                     int64    `json:"goods_id"`                                 //参与多多进宝的商品ID
	GoodsName                   string   `json:"goods_name"`                               //参与多多进宝的商品标题
	GoodsDesc                   string   `json:"goods_desc"`                               //参与多多进宝的商品描述
	GoodsImageUrl               string   `json:"goods_image_url"`                          //多多进宝商品主图
	GoodsGalleryUrls            []string `json:"goods_gallery_urls"`                       //商品轮播图
	MinGroupPrice               int64    `json:"min_group_price"`                          //最低价sku的拼团价，单位为分
	MinNormalPrice              int64    `json:"min_normal_price"`                         //最低价sku的单买价，单位为分
	SalesTip                    string   `json:"sales_tip"`                                //模糊销量
	DescTxt                     string   `json:"desc_txt"`                                 //描述分
	ServTxt                     string   `json:"serv_txt"`                                 //物流分
	LgstTxt                     string   `json:"lgst_txt"`                                 //服务分
	CreateAt                    int64    `json:"create_at"`                                //创建时间（unix时间戳）
	ServiceTags                 []int    `json:"service_tags"`                             //服务标签: 4-送货入户并安装,5-送货入户,6-电子发票,9-坏果包赔,11-闪电退款,12-24小时发货,13-48小时发货,17-顺丰包邮,18-只换不修,19-全国联保,20-分期付款,24-极速退款,25-品质保障,26-缺重包退,27-当日发货,28-可定制化,29-预约配送,1000001-正品发票,1000002-送货入户并安装
	CltCpnBatchSn               string   `json:"clt_cpn_batch_sn"`                         //店铺收藏券id
	CltCpnStartTime             int64    `json:"clt_cpn_start_time"`                       //店铺收藏券起始时间
	CltCpnEndTime               int64    `json:"clt_cpn_end_time"`                         //店铺收藏券截止时间
	CltCpnQuantity              int64    `json:"clt_cpn_quantity"`                         //店铺收藏券总量
	CltCpnRemainQuantity        int64    `json:"clt_cpn_remain_quantity"`                  //店铺收藏券剩余量
	CltCpnDiscount              int64    `json:"clt_cpn_discount"`                         //店铺收藏券面额，单位为分
	CltCpnMinAmt                int64    `json:"clt_cpn_min_amt"`                          //店铺收藏券使用门槛价格，单位为分
	VideoUrls                   []string `json:"video_urls"`                               //商品视频url
	ZsDuoID                     int64    `json:"zs_duo_id"`                                //招商多多客ID
	OnlySceneAuth               bool     `json:"only_scene_auth"`                          //快手专享
	PlanType                    int      `json:"plan_type"`                                //推广计划类型 3:定向 4:招商
	MerchantType                int      `json:"merchant_type"`                            //店铺类型，1-个人，2-企业，3-旗舰店，4-专卖店，5-专营店，6-普通店
	MallCps                     int      `json:"mall_cps"`                                 //该商品所在店铺是否参与全店推广，0：否，1：是
	ExtraCouponAmount           int64    `json:"extra_coupon_amount,omitempty"`            //额外优惠券，单位为分
	GoodsThumbnailUrl           string   `json:"goods_thumbnail_url"`                      //商品缩略图
	HasMallCoupon               bool     `json:"has_mall_coupon"`                          //是否有店铺券
	GoodsSign                   string   `json:"goods_sign"`                               //good sign
	HasCoupon                   bool     `json:"has_coupon"`                               //商品是否有优惠券 true-有，false-没有
	MallImgUrl                  string   `json:"mall_img_url,omitempty"`                   //店铺logo图
	PredictPromotionRate        int64    `json:"predict_promotion_rate"`                   //比价行为预判定佣金，需要用户备案
	ShareRate                   int      `json:"share_rate,omitempty"`                     //招商分成服务费比例，千分比
	SubsidyAmount               int      `json:"subsidy_amount,omitempty"`                 //优势渠道专属商品补贴金额，单位为分。
	SubsidyDuoAmountTenMillion  int      `json:"subsidy_duo_amount_ten_million,omitempty"` //千万补贴给渠道的收入补贴，不允许直接给下级代理展示，单位为分
	UnifiedTags                 []string `json:"unified_tags,omitempty"`                   //优惠标签列表
}
