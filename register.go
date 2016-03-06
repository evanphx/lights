package lights

import (
	"fmt"
	"log"
	"time"

	"github.com/evanphx/lights/config"
	"github.com/evanphx/lights/hue"
)

type RegisterCommand struct {
	Name     string `short:"b" default:"default" description:"name to give the bridge"`
	Username string `short:"u" default:"lights" description:"username to create as"`
}

func (o *RegisterCommand) Execute(args []string) error {
	locators, err := hue.DiscoverBridges(false)
	if err != nil {
		log.Fatal(err)
	}

	locator := locators[0] // find the first locator
	deviceType := o.Username

	fmt.Printf("Discovered a bridge, go press the link button\n")

	for i := 0; i < 10; i++ {
		fmt.Printf("Trying to create user...\n")
		// remember to push the button on your hue first
		bridge, err := locator.CreateUser(deviceType)
		if err != nil {
			return err
		}

		if bridge.Username != "" {
			config.AddBridge(o.Name, bridge.IpAddr, bridge.Username)
			fmt.Printf("Registered on bridge '%s'\n", o.Name)
			return nil
		}

		time.Sleep(5 * time.Second)
	}

	return fmt.Errorf("Unable to register with bridge")
}
