package sdk

import "github.com/truanter/pinduoduo-go/pkg/api"

func (c *Client) DDKGoodsSearch() *api.DDKGoodsSearch {
	return c.API.DDKGoodsSearch
}

func (c *Client) DDKGoodsPromotionUrlGenerate() *api.DDKGoodsPromotionUrlGenerate {
	return c.API.DDKGoodsPromotionUrlGenerate
}

func (c *Client) DDKResourceUrlGen() *api.DDKResourceUrlGen {
	return c.API.DDKResourceUrlGen
}

func (c *Client) DDKGoodsRecommendGet() *api.DDKGoodsRecommendGet {
	return c.API.DDKGoodsRecommendGet
}

func (c *Client) DDKGoodsDetail() *api.DDKGoodsDetail {
	return c.API.DDKGoodsDetail
}

func (c *Client) DDKRpPromUrlGenerate() *api.DDKRpPromUrlGenerate {
	return c.API.DDKRpPromUrlGenerate
}

func (c *Client) DDKMemberAuthorityQuery() *api.DDKMemberAuthorityQuery {
	return c.API.DDKMemberAuthorityQuery
}

func (c *Client) DDKCmsPromUrlGenerate() *api.DDKCmsPromUrlGenerate {
	return c.API.DDKCmsPromUrlGenerate
}

func (c *Client) GeneralAPI() *api.GeneralAPI {
	return c.API.GeneralAPI
}
