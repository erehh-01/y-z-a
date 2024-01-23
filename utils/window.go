package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

func Window(driver selenium.WebDriver, config Config, link string) (selenium.WebDriver, error) {
	var err error

	if config.Browser.LoadTime != 0 {
		load := time.Duration(config.Browser.LoadTime) * time.Second
		driver.SetPageLoadTimeout(load)
	}

	err = driver.Get(link)
	if err != nil {
		return nil, err
	}

	_, err = driver.ExecuteScript(
		`
			var elementToRemove = document.querySelector('div.recommendation-modal__container');
			if (elementToRemove) {
				elementToRemove.remove();
			}
			`,
		nil,
	)
	if err != nil {
		return nil, err
	}

	_, err = driver.ExecuteScript(
		`
		var elementToRemove = document.querySelector('div.needsclick');
		if (elementToRemove) {
			elementToRemove.remove();
		}
		`,
		nil,
	)
	if err != nil {
		return nil, err
	}

	inputs, err := driver.FindElements(selenium.ByTagName, "input")
	if err != nil {
		return nil, err
	}

	isFilled := make(map[string]bool)

	for _, p := range inputs {
		name, err := p.GetAttribute("name")
		if err != nil {
			continue
		}

		if isFilled[name] {
			continue
		}

		isFilled[name] = true

		switch name {
		case "firstName":
			err = p.Clear()
			if err != nil {
				return nil, err
			}
			err = p.SendKeys(fmt.Sprint(config.User.FirstName))
			if err != nil {
				return nil, err
			}
			//firstName := strings.Split(fmt.Sprint(config.User.FirstName), "")
			//err = p.Clear()
			//if err != nil {
			//	return err
			//}
			//for _, c := range firstName {
			//	err = p.SendKeys(c)
			//	if err != nil {
			//		return err
			//	}
			//}
		case "lastName":
			err = p.Clear()
			if err != nil {
				return nil, err
			}
			err = p.SendKeys(fmt.Sprint(config.User.LastName))
			if err != nil {
				return nil, err
			}
			//lastName := strings.Split(fmt.Sprint(config.User.LastName), "")
			//err = p.Clear()
			//if err != nil {
			//	return err
			//}
			//for _, c := range lastName {
			//	err = p.SendKeys(c)
			//	if err != nil {
			//		return err
			//	}
			//}
		case "email":
			err = p.Clear()
			if err != nil {
				return nil, err
			}
			err = p.SendKeys(fmt.Sprint(config.User.Email))
			if err != nil {
				return nil, err
			}
			//email := strings.Split(fmt.Sprint(config.User.Email), "")
			//err = p.Clear()
			//if err != nil {
			//	return err
			//}
			//for _, c := range email {
			//	err = p.SendKeys(c)
			//	if err != nil {
			//		return err
			//	}
			//}
		case "address1":
			err = p.Clear()
			if err != nil {
				return nil, err
			}
			err = p.SendKeys(fmt.Sprint(config.User.Address))
			if err != nil {
				return nil, err
			}
			//address := strings.Split(fmt.Sprint(config.User.Address), "")
			//err = p.Clear()
			//if err != nil {
			//	return err
			//}
			//for _, c := range address {
			//	err = p.SendKeys(c)
			//	if err != nil {
			//		return err
			//	}
			//}
		case "company":
			err = p.Clear()
			if err != nil {
				return nil, err
			}
			err = p.SendKeys(fmt.Sprint(config.User.Company))
			if err != nil {
				return nil, err
			}
			//company := strings.Split(fmt.Sprint(config.User.Company), "")
			//err = p.Clear()
			//if err != nil {
			//	return err
			//}
			//for _, c := range company {
			//	err = p.SendKeys(c)
			//	if err != nil {
			//		return err
			//	}
			//}
		case "postalCode":
			err = p.Clear()
			if err != nil {
				return nil, err
			}
			err = p.SendKeys(fmt.Sprint(config.User.PostCode))
			if err != nil {
				return nil, err
			}
			//postcode := strings.Split(fmt.Sprint(config.User.PostCode), "")
			//err = p.Clear()
			//if err != nil {
			//	return err
			//}
			//for _, c := range postcode {
			//	err = p.SendKeys(c)
			//	if err != nil {
			//		return err
			//	}
			//}
		case "phone":
			phone := strings.Split(config.User.PhoneCode+config.User.PhoneNumber, "")
			err = p.Clear()
			if err != nil {
				return nil, err
			}
			for _, c := range phone {
				err = p.SendKeys(c)
				if err != nil {
					return nil, err
				}
			}
		case "city":
			err = p.Clear()
			if err != nil {
				return nil, err
			}
			err = p.SendKeys(fmt.Sprint(config.User.City))
			if err != nil {
				return nil, err
			}
			//city := strings.Split(fmt.Sprint(config.User.City), "")
			//err = p.Clear()
			//if err != nil {
			//	return err
			//}
			//for _, c := range city {
			//	err = p.SendKeys(c)
			//	if err != nil {
			//		return err
			//	}
			//}
		default:
			continue
		}
	}

	time.Sleep(100 * time.Millisecond)

	return driver, nil
}

func Stream(driver selenium.WebDriver, config Config, cc CC) error {

	var err error
	_, err = driver.CurrentWindowHandle()
	if err != nil {
		fmt.Println("Window is closed")
		return nil
	}

	iframes, err := driver.FindElements(selenium.ByCSSSelector, "iframe.card-fields-iframe")
	if err != nil {
		return err
	}

	for _, iframe := range iframes {
		src, err := iframe.GetAttribute("src")
		if err != nil {
			continue
		}

		if strings.Contains(src, "/number?") {
			err = driver.SwitchFrame(iframe)
			if err != nil {
				return err
			}

			//iframePageSource, err := driver.PageSource()
			//if err != nil {
			//	return err
			//}
			//
			//_ = os.WriteFile("number.html", []byte(iframePageSource), 0644)

			input, err := driver.FindElement(selenium.ByCSSSelector, "input.input-placeholder-color--lvl-22")
			if err != nil {
				return err
			}

			ccNum := strings.Split(fmt.Sprint(cc.CCNUM), "")
			for _, c := range ccNum {
				err = input.SendKeys(c)
				if err != nil {
					return err
				}
			}
		}

		if strings.Contains(src, "/expiry?") {

			err = driver.SwitchFrame(iframe)
			if err != nil {
				return err
			}

			//iframePageSource, err := driver.PageSource()
			//if err != nil {
			//	return err
			//}
			//
			//_ = os.WriteFile("expiry.html", []byte(iframePageSource), 0644)

			input, err := driver.FindElement(selenium.ByCSSSelector, "input.input-placeholder-color--lvl-22")
			if err != nil {
				return err
			}
			ccMonth := strings.Split(fmt.Sprint(cc.MONTH), "")
			for _, c := range ccMonth {
				err = input.SendKeys(c)
				if err != nil {
					return err
				}
			}
			ccYear := strings.Split(fmt.Sprint(cc.YEAR), "")
			for _, c := range ccYear {
				err = input.SendKeys(c)
				if err != nil {
					return err
				}
			}
		}

		if strings.Contains(src, "/verification_value?") {
			err = driver.SwitchFrame(iframe)
			if err != nil {
				return err
			}

			//iframePageSource, err := driver.PageSource()
			//if err != nil {
			//	return err
			//}
			//
			//_ = os.WriteFile("verify.html", []byte(iframePageSource), 0644)

			input, err := driver.FindElement(selenium.ByCSSSelector, "input.input-placeholder-color--lvl-22")
			if err != nil {
				return err
			}
			//ccCVV := strings.Split(fmt.Sprint(cc.CVV), "")
			//for _, c := range ccCVV {
			//	err = input.SendKeys(c)
			//	if err != nil {
			//		return err
			//	}
			//}

			err = input.SendKeys(fmt.Sprint(cc.CVV))
			if err != nil {
				return err
			}
		}

		if strings.Contains(src, "/name?") {
			err = driver.SwitchFrame(iframe)
			if err != nil {
				return err
			}

			//iframePageSource, err := driver.PageSource()
			//if err != nil {
			//	return err
			//}
			//
			//_ = os.WriteFile("name.html", []byte(iframePageSource), 0644)

			input, err := driver.FindElement(selenium.ByCSSSelector, "input.input-placeholder-color--lvl-22")
			if err != nil {
				return err
			}
			err = input.SendKeys(fmt.Sprint(cc.NAME))
			if err != nil {
				return err
			}

			//ccName := strings.Split(fmt.Sprint(cc.NAME), "")
			//err = input.Clear()
			//if err != nil {
			//	return err
			//}
			//for _, c := range ccName {
			//	err = input.SendKeys(c)
			//	if err != nil {
			//		return err
			//	}
			//}

		}

		err = driver.SwitchFrame(nil)
		if err != nil {
			return err
		}
	}
	return nil
}
