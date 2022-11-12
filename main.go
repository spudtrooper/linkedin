package main

import (
	"flag"
	"strings"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/goutil/io"
	"github.com/spudtrooper/linkedin/linkedin"
)

var (
	url             = flag.String("url", "", "url to fetch")
	urlCSV          = flag.String("url_csv", "", "CSV with urls one line at a time")
	username        = flag.String("username", "", "username to use")
	password        = flag.String("password", "", "password to use")
	seleniumVerbose = flag.Bool("selenium_verbose", false, "verbose selenium logging")
	seleniumHead    = flag.Bool("selenium_head", false, "Take screenshots withOUT headless chrome")
	data            = flag.String("data", ".data", "directory to store data")
)

func urlsFromFlags() ([]string, error) {
	if *url != "" {
		return []string{*url}, nil
	}
	if *urlCSV != "" {
		lines, err := io.ReadLines(*urlCSV)
		if err != nil {
			return nil, err
		}
		var res []string
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" {
				res = append(res, line)
			}
		}
		return res, nil
	}
	return nil, errors.Errorf("no urls")
}

func realMain() error {
	flags.RequireString(username, "username")
	flags.RequireString(password, "password")
	urls, err := urlsFromFlags()
	if err != nil {
		return err
	}
	if len(urls) == 0 {
		return errors.Errorf("no urls")
	}

	bot, err := linkedin.MakeBot(*data)
	if err != nil {
		return err
	}

	logout, err := bot.Login(linkedin.Creds{
		Username: *username,
		Password: *password,
	}, linkedin.LoginSeleniumVerbose(*seleniumVerbose), linkedin.LoginSeleniumHead(*seleniumHead))
	if err != nil {
		return err
	}
	defer logout()

	if err := bot.Connect(urls); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()
	check.Err(realMain())
}
