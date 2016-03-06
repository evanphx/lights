package lights

import "fmt"

type OnCommand struct {
	Id string `short:"i" description:"id of light to turn off"`
}

func (o *OnCommand) Execute(args []string) error {
	bridge, err := Globals.Bridge()
	if err != nil {
		return err
	}

	lights, err := bridge.GetAllLights()
	if err != nil {
		return err
	}

	for _, light := range lights {
		if light.Id == o.Id {
			light.On()
			return nil
		}
	}

	return fmt.Errorf("Unable to find light '%s'", o.Id)
}
