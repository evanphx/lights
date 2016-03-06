package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type BridgeConfig struct {
	Name     string `json:"name"`
	Ip       string `json:"ip"`
	Username string `json:"username"`
}

type Config struct {
	Bridges map[string]*BridgeConfig `json:"bridges"`
}

func AddBridge(name, ip, username string) error {
	CurrentConfig.Bridges[name] = &BridgeConfig{
		Name:     name,
		Ip:       ip,
		Username: username,
	}

	return WriteConfig()
}

func Bridge(name string) (*BridgeConfig, bool) {
	b, ok := CurrentConfig.Bridges[name]
	return b, ok
}

var CurrentConfig Config

func LoadConfig() error {
	path := filepath.Join(os.Getenv("HOME"), ".lights.json")

	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil
	}

	defer f.Close()

	return json.NewDecoder(f).Decode(&CurrentConfig)
}

func WriteConfig() error {
	path := filepath.Join(os.Getenv("HOME"), ".lights.json")

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewEncoder(f).Encode(&CurrentConfig)
}

func init() {
	CurrentConfig.Bridges = make(map[string]*BridgeConfig)
	err := LoadConfig()
	if err != nil {
		panic(err)
	}
}
