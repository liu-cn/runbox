package runner

import (
	"fmt"
	"reflect"
	"strings"
)

type FuncParam struct {
	Code          string `json:"code,omitempty"`
	Desc          string `json:"desc,omitempty"`
	Mode          string `json:"mode,omitempty"`
	Type          string `json:"type,omitempty"`
	Value         string `json:"value,omitempty"`
	Options       string `json:"options,omitempty"`
	Required      string `json:"required,omitempty"`
	MockData      string `json:"mock_data,omitempty"`
	InputMode     string `json:"input_mode,omitempty"`
	TextLimit     string `json:"text_limit,omitempty"`
	NumberLimit   string `json:"number_limit,omitempty"`
	SelectOptions string `json:"select_options,omitempty"`
	FileSizeLimit string `json:"file_size_limit,omitempty"`
	FileTypeLimit string `json:"file_type_limit,omitempty"`
}

func NewConfig(opts ...Option) *Config {
	config := &Config{}
	for _, opt := range opts {
		opt(config)
	}
	return config
}

type Config struct {
	ApiDesc     string `json:"api_desc"`
	IsPublicApi bool   `json:"is_public_api"`

	ChineseName string      `json:"chinese_name"`
	EnglishName string      `json:"english_name"`
	Classify    string      `json:"classify"`
	Tags        string      `json:"tags"`
	Params      []FuncParam `json:"params"`

	Request  interface{} `json:"request,omitempty"`
	Response interface{} `json:"response,omitempty"`
}

func getRunnerTag(runnerTag string) FuncParam {
	funcP := FuncParam{}
	split := strings.Split(runnerTag, ";")
	mp := make(map[string]string)
	for _, s := range split {
		vals := strings.Split(s, ":")
		if len(vals) == 2 {
			mp[vals[0]] = vals[1]
		}
	}
	valueOf := reflect.ValueOf(&funcP).Elem()
	typeOf := reflect.TypeOf(funcP)
	for i := 0; i < valueOf.NumField(); i++ {
		field := typeOf.Field(i)
		if !field.IsExported() {
			continue
		}

		value := valueOf.Field(i)
		tag := field.Tag.Get("json")
		if v := mp[strings.Split(tag, ",")[0]]; v != "" {
			if value.CanSet() {
				value.SetString(v)
			}
		}
		//name := strings.ToLower(field.Name)
		//v := mp[name]
		//if v != "" {
		//f := valueOf.FieldByName(field.Name)
		//	if f.CanSet() {
		//		f.SetString(v)
		//	}
		//}
	}
	return funcP
}

func (c *Config) getParams(p interface{}, mode string) (params []FuncParam, err error) {
	if p == nil {
		return nil, nil
	}
	val := reflect.ValueOf(p)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input is not a struct")
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		par := FuncParam{
			Mode: mode,
			Code: typ.Field(i).Name,
		}
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			par.Code = strings.Split(jsonTag, ",")[0]
		}
		//runnerTag := field.Tag.Get("runner")
		runnerTag := field.Tag.Get("runner")
		if runnerTag == "" {

		} else {
			p1 := par
			par = getRunnerTag(runnerTag)
			if par.Mode == "" {
				par.Mode = p1.Mode
			}
			if par.Code == "" {
				par.Code = p1.Code
			}
		}

		if par.Type == "" {
			switch field.Type.Kind() {
			case reflect.String:
				par.Type = "string"
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				par.Type = "number"
			default:
				par.Type = "string"
			}
		}

		//pp := getRunnerTag(runnerTag)

		params = append(params, par)
	}
	return params, nil
}

func (c *Config) GetParams() ([]FuncParam, error) {
	var list []FuncParam
	if c.Request != nil {
		params, err := c.getParams(c.Request, "in")
		if err != nil {
			return nil, err
		}
		list = append(list, params...)
	}
	if c.Response != nil {
		params, err := c.getParams(c.Response, "out")
		if err != nil {
			return nil, err
		}
		list = append(list, params...)
	}
	return list, nil
}

type Option func(*Config)

func WithPublicApi() Option {
	return func(c *Config) {
		c.IsPublicApi = true
	}
}
func WithApiDesc(apiDesc string) Option {
	return func(c *Config) {
		c.ApiDesc = apiDesc
	}
}
