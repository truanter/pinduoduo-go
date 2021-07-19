package api

import (
	"github.com/truanter/pinduoduo-go/pkg/config"
	"github.com/truanter/pinduoduo-go/pkg/models"
)

type DDKRpPromUrlGenerate service

func (*DDKRpPromUrlGenerate) GetMethod() string {
	return config.APITypeDDKRpPromUrlGenerate
}

func (d *DDKRpPromUrlGenerate) Do(data models.DDKRpPromUrlGenerateRequest) (resp models.DDKRpPromotionUrlGenerateResult, err error) {
	err = d.client.do(d, data, &resp)
	return
}
