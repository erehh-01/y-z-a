package utils

import "github.com/tebeka/selenium"

func CheckError(driver selenium.WebDriver, err error) error {
	if err != nil {
		driver.Quit()
		return err
	}

	return nil
}
