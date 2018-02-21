package FusionExport

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type TypingsElementSchema struct {
	Type      string `json:"type"`
	Converter string `json:"converter"`
}

type MetaElementSchema struct {
	IsBase64Required bool `json:"isBase64Required"`
}

type ExportConfig struct {
	configs map[string]interface{}
	typings map[string]TypingsElementSchema
	meta    map[string]MetaElementSchema
}

func NewExportConfig() ExportConfig {
	config := ExportConfig{}
	config.configs = make(map[string]interface{})

	buffer, err := ioutil.ReadFile("metadata/fusionexport-typings.json")
	warn(err)
	json.Unmarshal(buffer, &config.typings)

	buffer, err = ioutil.ReadFile("metadata/fusionexport-meta.json")
	warn(err)
	json.Unmarshal(buffer, &config.meta)

	return config
}

func booleanConverter(val interface{}) (bool, error) {
	if reflect.TypeOf(val).String() == "bool" {
		return val.(bool), nil
	} else if reflect.TypeOf(val).String() == "string" {
		value := val.(string)
		value = strings.ToLower(value)

		if value == "true" || value == "1" {
			return true, nil
		} else if value == "false" || value == "0" {
			return false, nil
		}

		return false, errors.New("cannot convert " + value + " to bool")
	} else if reflect.TypeOf(val).String() == "int" {
		value := val.(int)

		if value == 1 {
			return true, nil
		} else if value == 0 {
			return false, nil
		}

		return false, errors.New("cannot convert " + strconv.Itoa(value) + " to bool")
	}

	return false, errors.New("cannot convert value to bool")
}

func numberConverter(val interface{}) (int, error) {
	if reflect.TypeOf(val).String() == "int" {
		return val.(int), nil
	} else if reflect.TypeOf(val).String() == "string" {
		value, err := strconv.Atoi(val.(string))
		if err != nil {
			return 0, err
		}
		return value, nil
	}

	return 0, errors.New("cannot convert value to bool")
}

func (config *ExportConfig) tryConvertType(name string, value interface{}) (interface{}, error) {
	if elm, ok := config.typings[name]; ok {
		converter := elm.Converter

		if converter == "BooleanConverter" {
			return booleanConverter(value)
		} else if converter == "NumberConverter" {
			return numberConverter(value)
		}

		return value, nil
	}

	return nil, errors.New(name + " is not supported")
}

func (config *ExportConfig) checkType(name string, value interface{}) error {
	if elm, ok := config.typings[name]; ok {
		expectedType := elm.Type

		if expectedType == "string" && reflect.TypeOf(value).String() != "string" {
			return errors.New(name + " must be of type string")
		} else if expectedType == "boolean" && reflect.TypeOf(value).String() != "bool" {
			return errors.New(name + " must be of type bool")
		} else if expectedType == "integer" && reflect.TypeOf(value).String() != "int" {
			return errors.New(name + " must be of type int")
		}
	} else {
		return errors.New(name + " is not supported")
	}

	return nil
}

func (config *ExportConfig) Set(name string, value interface{}) error {
	var err error
	value, err = config.tryConvertType(name, value)
	if err != nil {
		return err
	}

	err = config.checkType(name, value)
	if err != nil {
		return err
	}

	config.configs[name] = value
	return nil
}

func (config *ExportConfig) Get(name string) interface{} {
	return config.configs[name]
}

func (config *ExportConfig) Remove(name string) {
	delete(config.configs, name)
}

func (config *ExportConfig) Has(name string) bool {
	if _, ok := config.configs[name]; ok {
		return true
	}

	return false
}

func (config *ExportConfig) Clear(name string) {
	config.configs = make(map[string]interface{})
}

func (config *ExportConfig) Count() int {
	return len(config.configs)
}

func (config *ExportConfig) ConfigNames() []string {
	var keys []string
	for key := range config.configs {
		keys = append(keys, key)
	}
	return keys
}

func (config *ExportConfig) ConfigValues() []string {
	var values []string
	for _, val := range values {
		values = append(values, val)
	}
	return values
}

func (config *ExportConfig) Clone() ExportConfig {
	newConfig := NewExportConfig()
	for k, v := range config.configs {
		newConfig.Set(k, v)
	}
	return newConfig
}

func (config *ExportConfig) GetFormattedConfigs() (string, error) {
	formattedConfigs, err := config.formatConfigs()
	if err != nil {
		return "", err
	}

	json, err := json.Marshal(formattedConfigs)
	if err != nil {
		return "", err
	}

	return string(json), nil
}

func (config *ExportConfig) formatConfigs() (map[string]interface{}, error) {
	formattedConfigs := make(map[string]interface{})

	if tmpl, ok := config.configs["templateFilePath"]; ok {
		var tb TemplateBundler

		if res, k := config.configs["resourceFilePath"]; k {
			tb = TemplateBundler{
				Template:  tmpl.(string),
				Resources: res.(string),
			}
		} else {
			tb = TemplateBundler{
				Template: tmpl.(string),
			}
		}

		tb.Process()

		formattedConfigs["templateFilePath"] = tb.GetTemplatePathInZip()
		formattedConfigs["resourceFilePath"] = tb.GetResourcesZip()
	}

	if val, ok := config.configs["chartConfig"]; ok {
		if strings.HasSuffix(val.(string), ".json") {
			data, err := ioutil.ReadFile(val.(string))
			if err != nil {
				return nil, err
			}
			formattedConfigs["chartConfig"] = string(data)
		} else {
			formattedConfigs["chartConfig"] = val.(string)
		}
	}

	for key, val := range config.configs {
		switch key {
		case "templateFilePath":
		case "resourceFilePath":
		case "chartConfig":
			break
		default:
			formattedConfigs[key] = val
		}
	}

	for key, val := range formattedConfigs {
		if metaVal, ok := config.meta[key]; ok && metaVal.IsBase64Required {
			data, err := ioutil.ReadFile(val.(string))
			if err != nil {
				return nil, err
			}

			formattedConfigs[key] = base64.StdEncoding.EncodeToString(data)
		}
	}

	formattedConfigs["clientName"] = "GO"

	return formattedConfigs, nil
}
