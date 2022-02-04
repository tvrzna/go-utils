package config

import (
	"bufio"
	"errors"
	"io"
	"os"
	"reflect"
	"strings"
)

// Loads config from file content on defined path.
func LoadConfigFromFile(c interface{}, filePath string, ignoreMissing bool) error {
	var reader io.Reader

	file, err := os.Open(filePath)
	if err != nil && !ignoreMissing {
		return errors.New("Could not open file " + filePath)
	} else if err != nil && ignoreMissing {
		reader = strings.NewReader("\n")
	} else {
		reader = file
		defer file.Close()
	}

	configMap, err := readProperties(reader)
	if err != nil {
		return err
	}

	return loadConfig(c, configMap)
}

// Loads config from text content.
func LoadConfigFromText(c interface{}, text string) error {
	configMap, err := readProperties(strings.NewReader(text))
	if err != nil {
		return err
	}

	return loadConfig(c, configMap)
}

// Loads field tag parameters and sets values as expected.
func loadConfig(c interface{}, configMap map[string]string) error {
	configValue := reflect.ValueOf(c)
	configType := reflect.Indirect(configValue).Type()

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)

		configParam := field.Tag.Get("config")
		parserName := field.Tag.Get("parser")
		defaultValue := field.Tag.Get("default")
		if parserName != "" && configParam != "" {
			settingValue, exists := configMap[strings.ToUpper(configParam)]
			if !exists {
				settingValue = defaultValue
			}
			parser := configValue.MethodByName(parserName)
			if parser.Kind() != reflect.Invalid {
				val := parser.Call([]reflect.Value{reflect.ValueOf(settingValue), reflect.ValueOf(defaultValue)})[0]
				configValue.Elem().Field(i).Set(val)
			} else {
				return errors.New("Parser '" + parserName + "' was not found.")
			}
		}
	}

	return nil
}

func readProperties(reader io.Reader) (result map[string]string, err error) {
	result = make(map[string]string)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "#") && strings.Index(line, "=") >= 0 {
			splitIndex := strings.Index(line, "=")
			key := strings.ReplaceAll(line[:splitIndex], "export ", "")
			value := line[splitIndex+1:]
			if strings.Index(value, "#") >= 0 {
				value = value[:strings.Index(value, "#")]
			}
			key = strings.ToUpper(strings.TrimSpace(key))
			value = strings.TrimSpace(value)
			result[key] = value
		}
	}
	return result, scanner.Err()
}
