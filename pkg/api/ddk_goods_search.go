// https://open.pinduoduo.com/application/document/api?id=pdd.ddk.goods.search
package api

import (
	"github.com/truanter/pinduoduo-go/pkg/config"
	"github.com/truanter/pinduoduo-go/pkg/models"
)

type DDKGoodsSearch service

func (*DDKGoodsSearch) GetMethod() string {
	return config.APITypeDDKGoodsSearch
}

func (d *DDKGoodsSearch) Do(data models.DDKGoodsSearchRequest) (resp models.DDKGoodsSearchResult, err error) {
	err = d.client.do(d, data, &resp)
	return
}
