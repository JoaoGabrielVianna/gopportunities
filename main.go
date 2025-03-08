package main

import (
	"github.com/joaogabrielvianna/gopportunities/config"
	c "github.com/joaogabrielvianna/gopportunities/config"
	r "github.com/joaogabrielvianna/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// initialize config
	err := c.Init()
	if err != nil {
		logger.Errorf("config inialization error: %v", err)
		return
	}

	r.Initialize()
}
