package api

import (
	"github.com/truanter/pinduoduo-go/pkg/config"
	"github.com/truanter/pinduoduo-go/pkg/models"
)

type DDKGoodsDetail service

func (*DDKGoodsDetail) GetMethod() string {
	return config.APITypeDDKGoodsDetail
}

func (d *DDKGoodsDetail) Do(data models.DDKGoodsDetailRequest) (resp models.DDKGoodsDetailResult, err error) {
	err = d.client.do(d, data, &resp)
	return
}
