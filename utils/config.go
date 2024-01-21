package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var (
	linkFile string
	ccFile   string
)

type Config struct {
	User     User     `yaml:"user"`
	Browser  Browser  `yaml:"browser"`
	Telegram Telegram `yaml:"telegram"`
}

type User struct {
	Email       string `yaml:"email"`
	Address     string `yaml:"address"`
	FirstName   string `yaml:"firstName"`
	LastName    string `yaml:"lastName"`
	Company     string `yaml:"company"`
	PostCode    int    `yaml:"postCode"`
	City        string `yaml:"city"`
	PhoneCode   string `yaml:"phoneCode"`
	PhoneNumber string `yaml:"phoneNumber"`
}

type Browser struct {
	ChromeDriver string `yaml:"chromeDriver"`
	ChromePath   string `yaml:"chromePath"`
	UserAgent    string `yaml:"userAgent"`
	Proxy        string `yaml:"proxy"`
	Port         int    `yaml:"port"`
	SkipCaptcha  bool   `yaml:"skipCaptcha"`
	MaxWindows   int    `yaml:"maxWindows"`
	MaxTabs      int    `yaml:"maxTabs"`
	LoadTime     int    `yaml:"loadTime"`
}

type CC struct {
	CCNUM string    `yaml:"number"`
	YEAR  uint   `yaml:"year"`
	MONTH uint   `yaml:"month"`
	NAME  string `yaml:"name"`
	CVV   string `yaml:"cvv"`
}

type Telegram struct {
	AppID    int32  `yaml:"appID"`
	AppHash  string `yaml:"appHash"`
	Channels []int64  `yaml:"channels"`
}

func StartConfig() error {
	// Get the current working directory
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Ensure config folder exists
	configPath := filepath.Join(pwd, "config")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create the config folder
		err = os.MkdirAll(configPath, 0755) // Use 0755 for appropriate permissions
		if err != nil {
			return err
		}
	}

	// Set config file path
	configFile := filepath.Join(configPath, "y-z-a.yaml")

	// Ensure config file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("Config file not found, created:", configFile)
		// Create a new file with default configuration
		err = createDefaultConfigFile(configFile)
		if err != nil {
			return err
		}
	}

	dataPath := filepath.Join(pwd, "data")
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		// Create the config folder
		err = os.MkdirAll(dataPath, 0755) // Use 0755 for appropriate permissions
		if err != nil {
			return err
		}
	}

	// Set links file path
	linkFile = filepath.Join(dataPath, "links.txt")
	// Ensure links file exists
	if _, err := os.Stat(linkFile); os.IsNotExist(err) {
		// Create a new file with default links
		err = os.WriteFile(linkFile, nil, 0644)
		if err != nil {
			return err
		}
	}

	// Set cc file path
	ccFile = filepath.Join(dataPath, "credit-cards.json")
	// Ensure cc file exists
	if _, err := os.Stat(ccFile); os.IsNotExist(err) {
		// Create a new file with default cc
		err = createDefaultCCFile(ccFile)
		if err != nil {
			return err
		}
	}

	// Ensure screenshot folder exists
	screenshotPath := filepath.Join(pwd, "screenshot")
	if _, err := os.Stat(screenshotPath); os.IsNotExist(err) {
		// Create the screenshot folder
		err = os.MkdirAll(screenshotPath, 0755) // Use 0755 for appropriate permissions
		if err != nil {
			return err
		}
	}

	// Load configuration
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")
	viper.SetConfigName("y-z-a")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	return err
}

func LoadConfig() (Config, error) {
	config := Config{}
	err := viper.Unmarshal(&config)
	return config, err
}

// createDefaultConfigFile creates a new config file with default values
func createDefaultConfigFile(configFilePath string) error {
	defaultConfig := Config{
		User: User{
			Email:       "test@email.com",
			Address:     "test address",
			FirstName:   "test fisrt name",
			LastName:    "test last name",
			Company:     "shopify",
			PostCode:    10001,
			City:        "test city",
			PhoneCode:   "+1",
			PhoneNumber: "7073608450",
		},
		Browser: Browser{
			ChromeDriver: "",
			ChromePath:   "",
			UserAgent:    "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
			Proxy:        "http://localhost",
			Port:         80,
			SkipCaptcha:  false,
			MaxWindows:   2,
			MaxTabs:      10,
			LoadTime:     15,
		},
		Telegram: Telegram{
			AppID:   0,
			AppHash: "",
			Channels: []int64{
				-1001397199427,
			},
		},
	}

	// Marshal the default configuration to YAML
	yamlData, err := yaml.Marshal(&defaultConfig)
	if err != nil {
		return err
	}

	// Create the config file and write the YAML data
	return os.WriteFile(configFilePath, yamlData, 0644)
}

// createDefaultCCFile creates a new config file with default values
func createDefaultCCFile(CCFilePath string) error {
	defaultCC := []CC{
		{
			CCNUM: "5110200003199389",
			YEAR:  2025,
			MONTH: 12,
			NAME:  "Test Test",
			CVV:   "123",
		},
	}

	// Marshal the default cc to JSON
	jsonData, err := json.Marshal(&defaultCC)
	if err != nil {
		return err
	}

	// Create the cc file and write the JSON data
	return os.WriteFile(CCFilePath, jsonData, 0644)
}

func LoadLinks() ([]string, error) {
	if linkFile == "" {
		return nil, errors.New("there is no links.txt file")
	}

	linksData, err := os.ReadFile(linkFile)
	if err != nil {
		return nil, errors.New("failed to read links.txt file")
	}

	links := strings.Split(string(linksData), "\n")

	return links, nil
}

func LoadCC() ([]CC, error) {
	if ccFile == "" {
		return nil, errors.New("there is no credit-card.json file")
	}

	ccData, err := os.ReadFile(ccFile)
	if err != nil {
		return nil, errors.New("failed to read credit-card.json file")
	}

	var creditCards []CC
	if err := json.Unmarshal(ccData, &creditCards); err != nil {
		return nil, errors.New("failed to list credit-card.json")
	}

	return creditCards, nil
}
