package FusionExport

import (
    "fmt"
    "bytes"
)

type ExportConfig struct {
    configs map[string]string
}

func NewExportConfig () ExportConfig {
    config := ExportConfig{}
    config.configs = make(map[string]string)
    return config
}

func (config *ExportConfig) Set (name, value string) {
    config.configs[name] = value
}

func (config *ExportConfig) Get (name string) string {
    return config.configs[name]
}

func (config *ExportConfig) Remove (name string) {
    delete(config.configs, name)
}

func (config *ExportConfig) Has (name string) bool {
    if _, ok := config.configs[name]; ok {
        return true
    }

    return false
}

func (config *ExportConfig) Clear (name string) {
    config.configs = make(map[string]string)
}

func (config *ExportConfig) Count () int {
    return len(config.configs)
}

func (config *ExportConfig) ConfigNames () []string {
    var keys []string
    for key := range config.configs {
        keys = append(keys, key)
    }
    return keys
}

func (config *ExportConfig) ConfigValues () []string {
    var values []string
    for _, val := range values {
        values = append(values, val)
    }
    return values
}

func (config *ExportConfig) Clone () ExportConfig {
    newConfig := NewExportConfig()
    for k, v := range config.configs {
        newConfig.Set(k, v)
    }
    return newConfig
}

func (config *ExportConfig) GetFormattedConfigs () string {
    var buffer bytes.Buffer

    for key, value := range config.configs {
        formattedConfigValue := config.getFormattedConfigValue(key, value)
        keyValuePair := fmt.Sprintf("\"%s\": %s, ", key, formattedConfigValue)
        buffer.WriteString(keyValuePair)
    }

    configsAsJSON := buffer.String()
    if len(configsAsJSON) > 2 {
        configsAsJSON = configsAsJSON[:len(configsAsJSON) - 2]
    }

    return fmt.Sprintf("{%s}", configsAsJSON)
}

func (config *ExportConfig) getFormattedConfigValue (name, value string) string {
    switch name {
        case "chartConfig":
        case "asyncCapture":
        case "exportAsZip":
            return value
        default:
            return fmt.Sprintf("\"%s\"", value)
    }
    
    return value
}
