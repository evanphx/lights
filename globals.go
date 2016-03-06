package lights

import (
	"fmt"

	"github.com/evanphx/lights/config"
	"github.com/evanphx/lights/hue"
)

type GlobalOptions struct {
	Name string `short:"b" default:"default" description:"bridge to talk with"`
}

var Globals GlobalOptions

func (g *GlobalOptions) Bridge() (*hue.Bridge, error) {
	conf, ok := config.Bridge(g.Name)
	if !ok {
		return nil, fmt.Errorf("No configuration for bridge '%s', register it!\n", g.Name)
	}

	return hue.NewBridge(conf.Ip, conf.Username), nil
}
