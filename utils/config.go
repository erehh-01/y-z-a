package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var (
	linkFile string
	ccFile   string
)

type Config struct {
	User     User     `json:"user"`
	Browser  Browser  `json:"browser"`
	Telegram Telegram `json:"telegram"`
}

type User struct {
	Email       string `json:"email"`
	Address     string `json:"address"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Company     string `json:"company"`
	PostCode    int    `json:"postCode"`
	City        string `json:"city"`
	PhoneCode   string `json:"phoneCode"`
	PhoneNumber string `json:"phoneNumber"`
}

type Browser struct {
	ChromeDriver string `json:"chromeDriver"`
	ChromePath   string `json:"chromePath"`
	UserAgent    string `json:"userAgent"`
	Proxy        string `json:"proxy"`
	Port         int    `json:"port"`
	SkipCaptcha  bool   `json:"skipCaptcha"`
	MaxWindows   int    `json:"maxWindows"`
	MaxTabs      int    `json:"maxTabs"`
	LoadTime     int    `json:"loadTime"`
}

type CC struct {
	CCNUM string `json:"number"`
	YEAR  uint   `json:"year"`
	MONTH uint   `json:"month"`
	NAME  string `json:"name"`
	CVV   uint   `json:"cvv"`
}

type Telegram struct {
	AppID   int32  `json:"appID"`
	AppHash string `json:"appHash"`
	
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
	configFile := filepath.Join(configPath, "y-z-a.json")
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
	viper.AddConfigPath(filepath.Clean(configPath))
	viper.SetConfigType("json")
	viper.SetConfigName("y-z-a")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	return err
}

func LoadConfig() (Config, error) {
	config := Config{}
	configPath := viper.ConfigFileUsed()
	if configPath == "" {
		return config, fmt.Errorf("there is no config file")
	}

	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return config, fmt.Errorf("config file is empty")
	}

	if err = json.Unmarshal(configBytes, &config); err != nil {
		return config, fmt.Errorf("error in config file")
	}

	return config, nil
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
			SkipCaptcha:  true,
			MaxWindows:   2,
			MaxTabs:      10,
			LoadTime:     15,
		},
		Telegram: Telegram{
			AppID:   0,
			AppHash: "",
		},
	}

	// Marshal the default configuration to json
	jsonData, err := json.Marshal(&defaultConfig)
	if err != nil {
		return err
	}

	// Create the config file and write the json data
	return os.WriteFile(configFilePath, jsonData, 0644)
}

// createDefaultCCFile creates a new config file with default values
func createDefaultCCFile(CCFilePath string) error {
	defaultCC := []CC{
		{
			CCNUM: "5110200003199389",
			YEAR:  2025,
			MONTH: 12,
			NAME:  "Test Test",
			CVV:   123,
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
