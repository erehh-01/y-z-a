package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/tebeka/selenium"
)

func Shot(driver selenium.WebDriver) error {
	img, err := driver.Screenshot()
	if err != nil {
		driver.Quit()
		return err
	}

	uuid := uuid.New()
	err = os.WriteFile(filepath.Join("screenshot", fmt.Sprintf("%s.png", uuid.String())), img, 0644)
	if err != nil {
		driver.Quit()
		return err
	}

	return nil
}
