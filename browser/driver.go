package browser

import "github.com/tebeka/selenium"

var (
	driver  selenium.WebDriver
	service *selenium.Service
)

func Close() {
	if driver != nil {
		driver.Close()
	}
	if service != nil {
		service.Stop()
	}
}
