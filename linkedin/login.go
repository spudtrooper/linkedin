package linkedin

import (
	"fmt"
	"log"
	"time"

	goutilselenium "github.com/spudtrooper/goutil/selenium"
	"github.com/tebeka/selenium"
)

//go:generate genopts --function Login seleniumVerbose seleniumHead
func (b *bot) Login(creds Creds, optss ...LoginOption) (func(), error) {
	opts := MakeLoginOptions(optss...)

	wd, cancel, err := goutilselenium.MakeWebDriver(goutilselenium.MakeWebDriverOptions{
		Verbose:  opts.SeleniumVerbose(),
		Headless: !opts.SeleniumHead(),
	})
	if err != nil {
		return cancel, err
	}

	if err := b.login(wd, creds); err != nil {
		return cancel, err
	}

	return cancel, nil
}

func (b *bot) login(wd selenium.WebDriver, creds Creds) error {
	log.Printf("logging in...")

	if err := wd.Get("https://www.linkedin.com/"); err != nil {
		return err
	}

	findCredsInputs := func() (selenium.WebElement, selenium.WebElement, error) {
		usernameInput, err := wd.FindElement(selenium.ByID, "session_key")
		if err != nil {
			return nil, nil, err
		}
		if usernameInput == nil {
			return nil, nil, nil
		}
		passwordInput, err := wd.FindElement(selenium.ByID, "session_password")
		if err != nil {
			return nil, nil, err
		}
		if passwordInput == nil {
			return nil, nil, nil
		}
		return usernameInput, passwordInput, nil
	}
	var usernameInput, passwordInput selenium.WebElement
	wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		log.Printf("waiting for username/password inputs...")
		u, p, err := findCredsInputs()
		if err != nil {
			return false, err
		}
		if u == nil || p == nil {
			return false, nil
		}
		usernameInput = u
		passwordInput = p
		return true, nil
	})
	if usernameInput == nil {
		return fmt.Errorf("no username input")
	}
	if passwordInput == nil {
		return fmt.Errorf("no password input")
	}

	if err := usernameInput.Clear(); err != nil {
		return err
	}
	usernameInput.SendKeys(creds.Username)

	if err := passwordInput.Clear(); err != nil {
		return err
	}
	passwordInput.SendKeys(creds.Password)

	loginBtn, err := wd.FindElement(selenium.ByClassName, "sign-in-form__submit-button")
	if err != nil {
		return err
	}
	if err := loginBtn.Click(); err != nil {
		return err
	}

	time.Sleep(2 * time.Second)

	b.wd = wd

	log.Printf("logged in")

	return nil
}
