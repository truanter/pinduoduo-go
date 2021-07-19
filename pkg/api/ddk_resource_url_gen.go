package api

import (
	"github.com/truanter/pinduoduo-go/pkg/config"
	"github.com/truanter/pinduoduo-go/pkg/models"
)

type DDKResourceUrlGen service

func (*DDKResourceUrlGen) GetMethod() string {
	return config.APITypeDDKResourceUrlGen
}

func (d *DDKResourceUrlGen) Do(data models.DDkResourceUrlGenRequest) (resp models.DDKResourceUrlGenResult, err error) {
	err = d.client.do(d, data, &resp)
	return
}
