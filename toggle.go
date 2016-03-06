package lights

import "fmt"

type ToggleCommand struct {
	Id string `short:"i" description:"id of light to toggle"`
}

func (o *ToggleCommand) Execute(args []string) error {
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
			if light.State.On {
				light.Off()
			} else {
				light.On()
			}

			return nil
		}
	}

	return fmt.Errorf("Unable to find light '%s'", o.Id)
}
