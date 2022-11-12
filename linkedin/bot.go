package linkedin

import "github.com/tebeka/selenium"

type bot struct {
	data *data
	wd   selenium.WebDriver
}

func MakeBot(dataDir string) (*bot, error) {
	data, err := makeData(dataDir)
	if err != nil {
		return nil, err
	}
	return &bot{data: data}, nil
}

type Creds struct {
	Username, Password string
}
