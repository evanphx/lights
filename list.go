package lights

import "fmt"

type ListCommand struct {
}

func (o *ListCommand) Execute(args []string) error {
	bridge, err := Globals.Bridge()
	if err != nil {
		return err
	}

	lights, err := bridge.GetAllLights()
	if err != nil {
		return err
	}

	for _, light := range lights {
		fmt.Printf("%s: %s\n", light.Id, light.Name)
	}

	return nil
}
