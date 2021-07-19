package api

type GeneralAPI service

func (g *GeneralAPI) SetType(t string) {
	g.Method = t
}

func (g *GeneralAPI) GetMethod() string {
	return g.Method
}

func (g *GeneralAPI) Do(data interface{}) (resp interface{}, err error) {
	err = g.client.do(g, data, &resp)
	return
}
