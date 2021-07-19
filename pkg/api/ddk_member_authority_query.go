package api

import (
	"github.com/truanter/pinduoduo-go/pkg/config"
	"github.com/truanter/pinduoduo-go/pkg/models"
)

type DDKMemberAuthorityQuery service

func (*DDKMemberAuthorityQuery) GetMethod() string {
	return config.APITypeDDKMemberAuthorityQuery
}

func (d *DDKMemberAuthorityQuery) Do(data models.DDKMemberAuthorityQueryRequest) (resp models.DDKMemberAuthorityQueryResult, err error) {
	err = d.client.do(d, data, &resp)
	return
}
