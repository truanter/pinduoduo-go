package api

import (
	"github.com/truanter/pinduoduo-go/pkg/config"
	"github.com/truanter/pinduoduo-go/pkg/models"
)

type DDKGoodsRecommendGet service

func (*DDKGoodsRecommendGet) GetMethod() string {
	return config.APITypeDDKGoodsRecommendGet
}

func (d *DDKGoodsRecommendGet) Do(data models.DDKGoodsRecommendRequest) (resp models.DDKGoodsRecommendResult, err error) {
	err = d.client.do(d, data, &resp)
	return
}
