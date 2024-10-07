package conf

func NewConfig(opts ...Option) *Config {
	config := &Config{}
	for _, opt := range opts {
		opt(config)
	}
	return config
}

type Config struct {
	IsPublicApi bool                   `json:"is_public_api"`
	ApiDesc     string                 `json:"api_desc"`
	MockData    map[string]interface{} `json:"mock_data"`

	IsCloudFunc bool        `json:"is_cloud_func"`
	Request     interface{} `json:"request"`
	Response    interface{} `json:"response"`
}
type Option func(*Config)

func WithRequestAndResponse(request, response interface{}) Option {
	return func(c *Config) {
		c.Request = request
		c.Response = response
	}
}

func WithCloudFunc(isCloudFunc bool) Option {
	return func(c *Config) {
		c.IsCloudFunc = isCloudFunc
	}
}

func WithIsPublicApi(isPublic bool) Option {
	return func(c *Config) {
		c.IsPublicApi = isPublic
	}
}
func WithApiDesc(apiDesc string) Option {
	return func(c *Config) {
		c.ApiDesc = apiDesc
	}
}
func WithMockData(mockData map[string]interface{}) Option {
	return func(c *Config) {
		c.MockData = mockData
	}
}
