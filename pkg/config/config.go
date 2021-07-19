package config

const (
	Url = "https://gw-api.pinduoduo.com/api/router"
)

// DefaultParams public params and their default value
var DefaultParams = map[string]interface{}{
	"data_type": "JSON",
	"version":   "V1",
}

func GetHost() string {
	return Url
}
