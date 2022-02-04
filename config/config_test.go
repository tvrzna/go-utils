package config

import (
	"strconv"
	"testing"
)

type ExampleConfig struct {
	Name      string `config:"Name" parser:"ToString" default:"Default name"`
	Age       int    `config:"Age" parser:"ToInt" default:"18"`
	Available bool   `config:"Available" parser:"ToBool" default:"false"`
}

func (c *ExampleConfig) ToString(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func (c *ExampleConfig) ToInt(value, defaultValue string) int {
	res, _ := strconv.Atoi(c.ToString(value, defaultValue))
	return res
}

func (c *ExampleConfig) ToBool(value, defaultValue string) bool {
	return c.ToString(value, defaultValue) == "true"
}

func TestLoadConfigFromFile(t *testing.T) {
	c := ExampleConfig{}

	err := LoadConfigFromFile(&c, "config_test.txt", false)
	if err != nil {
		t.Fatal(err)
	}

	if c.Name != "Unknown" || c.Age != 20 || !c.Available {
		t.Error("TestLoadConfigFromFile: Wrong values on load from file")
	}
}

func TestLoadConfigFromFileNoFile(t *testing.T) {
	c := ExampleConfig{}

	err := LoadConfigFromFile(&c, "non-existing_config_test.txt", false)
	if err == nil {
		t.Error("TestLoadConfigFromFileNoFile: Wrong values on load from file")
	}
}

func TestLoadConfigFromFileDefaults(t *testing.T) {
	c := ExampleConfig{}

	err := LoadConfigFromFile(&c, "non-existing_config_test.txt", true)
	if err != nil {
		t.Fatal(err)
	}

	if c.Name != "Default name" || c.Age != 18 || c.Available {
		t.Error("TestLoadConfigFromFileDefaults: Wrong values on load from file")
	}
}

func TestLoadConfigFromText(t *testing.T) {
	c := ExampleConfig{}

	err := LoadConfigFromText(&c, `NAME=Person
age=333 # This is actual value
#age=80
export avAILable=true`)
	if err != nil {
		t.Fatal(err)
	}

	if c.Name != "Person" || c.Age != 333 || !c.Available {
		t.Error("TestLoadConfigFromText: Wrong values on load from file")
	}
}

func TestLoadConfigFromTextDefaults(t *testing.T) {
	c := ExampleConfig{}

	err := LoadConfigFromText(&c, "")
	if err != nil {
		t.Fatal(err)
	}

	if c.Name != "Default name" || c.Age != 18 || c.Available {
		t.Error("TestLoadConfigFromFileDefaults: Wrong values on load from file")
	}
}
