package linkedin

import (
	"log"

	"github.com/pkg/errors"
	"github.com/tebeka/selenium"
)

func (b *bot) Connect(urls []string) error {
	b.data.CreateQueue(urls)
	for {
		hasMore, err := b.data.HasMore()
		if err != nil {
			return err
		}
		if !hasMore {
			log.Println("!hasMore")
			break
		}
		url, done, err := b.data.Next()
		if err != nil {
			return err
		}
		if done {
			log.Println("done")
			break
		}
		if err := b.connect(url); err != nil {
			return err
		}
	}
	return nil
}

func (b *bot) connect(url string) error {
	url += "&allowUnsupportedBrowser=true"
	log.Printf("connecting to %s", url)
	b.data.Try(url)
	if err := b.wd.Get(url); err != nil {
		return err
	}
	var connectButtons []selenium.WebElement
	b.wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		log.Printf("waiting for buttons...")
		bs, err := b.wd.FindElements(selenium.ByTagName, "button")
		if err != nil {
			return false, err
		}
		for _, b := range bs {
			text, err := b.Text()
			if err != nil {
				return false, err
			}
			if text == "Connect" {
				log.Printf("found connect button")
				connectButtons = append(connectButtons, b)
			}
		}
		if len(connectButtons) > 0 {
			return true, nil
		}
		return false, nil
	})

	if len(connectButtons) == 0 {
		return errors.Errorf("no buttons")
	}

	for _, btn := range connectButtons {
		if err := btn.Click(); err != nil {
			return err
		}
		var sendButtons []selenium.WebElement
		b.wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			log.Printf("waiting for buttons...")
			bs, err := b.wd.FindElements(selenium.ByTagName, "button")
			if err != nil {
				return false, err
			}
			for _, b := range bs {
				text, err := b.Text()
				if err != nil {
					return false, err
				}
				if text == "Send" {
					log.Printf("found connect button")
					sendButtons = append(sendButtons, b)
				}
			}
			if len(sendButtons) > 0 {
				return true, nil
			}
			return false, nil
		})
		for _, btn := range sendButtons {
			if err := btn.Click(); err != nil {
				return err
			}
		}
	}
	return nil
}
