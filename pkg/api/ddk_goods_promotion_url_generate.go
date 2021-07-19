package api

import (
	"github.com/truanter/pinduoduo-go/pkg/config"
	"github.com/truanter/pinduoduo-go/pkg/models"
)

type DDKGoodsPromotionUrlGenerate service

func (*DDKGoodsPromotionUrlGenerate) GetMethod() string {
	return config.APITypeDDKGoodsPromotionUrlGenerate
}

func (d *DDKGoodsPromotionUrlGenerate) Do(data models.DDKGoodsPromotionUrlGenerateRequest) (resp models.DDKGoodsPromotionUrlGenerateResult, err error) {
	err = d.client.do(d, data, &resp)
	return
}
