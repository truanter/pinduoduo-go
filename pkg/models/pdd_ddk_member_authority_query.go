package models

type DDKMemberAuthorityQueryRequest struct {
	Pid              *string `json:"pid,omitempty"`
	CustomParameters *string `json:"custom_parameters,omitempty"`
}

type DDKMemberAuthorityQueryResult struct {
	BaseResponse
	AuthorityQueryResponse struct {
		Bind int64 `json:"bind"`
	} `json:"authority_query_response"`
}
