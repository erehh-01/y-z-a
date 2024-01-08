package browser

import (
	"errors"
	"fmt"

	"github.com/dj-yacine-flutter/y-z-a/utils"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func Chrome(headless bool) (selenium.WebDriver, error) {
	var err error

	conf, err := utils.LoadConfig()
	if err != nil {
		return nil, err
	}

	if conf.Browser.ChromePath == "" {
		return nil, errors.New("to use the chrome browser put the path in the config file")
	}

	if conf.Browser.ChromeDriver == "" {
		return nil, errors.New("missing driver path in the config file")
	}

	caps := selenium.Capabilities{
		"browserName":    "chrome",
		"browserVersion": "119.0",
		"se:noVncPort":   7900,
		"se:vncEnabled":  true,
	}

	chromeCaps := chrome.Capabilities{
		Path: conf.Browser.ChromePath,
		Args: []string{
			"--silent",
			"--ignore-certificate-errors",
			"--disk-cache-size=1",
			"--media-cache-size=1",
			"--disable-extensions",
			"--disable-infobars",
			"--disable-background-networking",
			"--disable-client-side-phishing-detection",
			"--disable-component-extensions-with-background-pages",
			"--disable-features=InterestFeedContentSuggestions",
			"--disable-features=Translate",
			"--mute-audio",
			"--no-default-browser-check",
			"--no-first-run",
			"--ash-no-nudges",
			"--disable-search-engine-choice-screen",
			"--disable-features=CalculateNativeWinOcclusion",
			"--disable-features=LazyFrameLoading",
			"--disable-notifications",
			"--block-new-web-content",
			"--disable-prompt-on-repost",
			"--noerrdialogs",
			"--disable-save-password-bubble",
			"--disable-domain-reliability",
			"--no-pings",
			//"--in-process-gpu",
			"--disable-default-apps",
			"--disable-hang-monitor",
			"--disable-popup-blocking",
			"--enable-scripts-block-automatic-downloads",
			"--disable-plugins-discovery",
			"--disable-features=CastMediaRouteProvider",
			"--disable-remote-playback",
			"--disable-features=AppCast",
			"--disable-features=CastReceiverMediaCast",
			"--disable-cast",
			"--disable-app-list",
			"--disable-media-router",
			"--disable-prompt-on-repost",
			"--disable-sync",
			"--disable-web-resources",
			"--disable-crash-reporter",
			"--disable-oopr-debug-crash-dump",
			"--disable-blink-features=AutomationControlled",
			"--enable-automation",
			"--enable-quic",
			"--enable-tcp-fast-open",
			"--allow-running-insecure-content",
			//"--start-maximized",
			"--FontRenderHinting[none]",
			"--no-crash-upload",
			fmt.Sprintf("--user-agent=%s", conf.Browser.UserAgent),
			"--disable-blink-features",
			"--disable-features",
		},
		W3C: true,
	}

	if headless {
		chromeCaps.Args = append(chromeCaps.Args, "--headless")
		chromeCaps.Args = append(chromeCaps.Args, "--headless=new")
		chromeCaps.Args = append(chromeCaps.Args, "--disable-gpu")
		chromeCaps.Args = append(chromeCaps.Args, "--no-sandbox")
		chromeCaps.Args = append(chromeCaps.Args, "--disable-dev-shm-usage")
	}

	caps.AddChrome(chromeCaps)

	service, err = selenium.NewChromeDriverService(conf.Browser.ChromeDriver, 4444)
	if err != nil {
		service.Stop()
		return nil, err
	}

	driver, err = selenium.NewRemote(caps, "")
	if err != nil {
		driver.Close()
		return nil, err
	}

	return driver, nil
}
