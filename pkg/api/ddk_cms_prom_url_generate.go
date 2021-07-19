package api

import (
	"github.com/truanter/pinduoduo-go/pkg/config"
	"github.com/truanter/pinduoduo-go/pkg/models"
)

type DDKCmsPromUrlGenerate service

func (*DDKCmsPromUrlGenerate) GetMethod() string {
	return config.APITypeDDKCmsPromUrlGenerate
}

func (d *DDKCmsPromUrlGenerate) Do(data models.DDKCmsPromUrlGenerateRequest) (resp models.DDKCmsPromUrlGenerateResult, err error) {
	err = d.client.do(d, data, &resp)
	return
}
