package config

import (
	"os"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v3"
)

const (
	defaultConfig = "hackertea/config.yaml"
)

type Config struct {
	Style   Style `yaml:"style"`
	Workers int   `yaml:"workers"`
}

type Style struct {
	ListItem ListItem      `yaml:"listItem"`
	Visited  AdaptiveColor `yaml:"visited"`
	Window   Window        `yaml:"window"`
	Tab      Tab           `yaml:"tab"`
}

type ListItem struct {
	NormalTitle   AdaptiveColor `yaml:"normalTitle"`
	NormalDesc    AdaptiveColor `yaml:"normalDesc"`
	SelectedTitle ListItemStyle `yaml:"selectedTitle"`
	SelectedDesc  AdaptiveColor `yaml:"selectedDesc"`
	DimmedTitle   AdaptiveColor `yaml:"dimmedTitle"`
	DimmedDesc    AdaptiveColor `yaml:"dimmedDesc"`
	FilterMatch   ListItemStyle `yaml:"filterMatch"`
}

type ListItemStyle struct {
	BorderForeground AdaptiveColor
	Foreground       AdaptiveColor
}

type Window struct {
	Border string        `yaml:"border"`
	Color  AdaptiveColor `yaml:"color"`
}

type Tab struct {
	Color AdaptiveColor `yaml:"color"`
}

type AdaptiveColor struct {
	Dark  string `yaml:"dark"`
	Light string `yaml:"light"`
}

// LoadConfig loads the configuration file.
// It first searches for the configuration file using the XDG Base Directory Specification.
// If the configuration file is found, it is loaded and parsed using the getConfig function.
// If the configuration file is not found, a default configuration is written to a new file
// using the initConfig function, and the default configuration is returned.
func LoadConfig() (*Config, error) {
	configFilePath, err := xdg.SearchConfigFile(defaultConfig)
	if err == nil {
		return getConfig(configFilePath)
	}

	return initConfig(defaultConfig)
}

// getConfig reads a configuration file from a specified path and decodes
// it into a Config.
func getConfig(configPath string) (*Config, error) {
	cFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	defer cFile.Close()

	cfg := &Config{}
	if err = yaml.NewDecoder(cFile).Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// initConfig creates a new configuration file at the specified path and returns a Config struct with default values.
// The default values are set in the basicConfig function.
func initConfig(configPath string) (*Config, error) {
	configFilePath, err := xdg.ConfigFile(configPath)
	if err != nil {
		return nil, err
	}

	confB, _ := yaml.Marshal(basicConfig())
	if err = os.WriteFile(configFilePath, confB, 0o755); err != nil {
		return nil, err
	}

	return basicConfig(), nil
}

// basicConfig returns the default configuration.
func basicConfig() *Config {
	return &Config{
		Style: Style{
			ListItem: ListItem{
				NormalTitle: AdaptiveColor{
					Dark:  "#E6EBE9",
					Light: "#7D56F4",
				},
				NormalDesc: AdaptiveColor{
					Dark:  "#4f4f4f",
					Light: "#7D56F4",
				},
				SelectedTitle: ListItemStyle{
					BorderForeground: AdaptiveColor{
						Dark:  "#2A8f69",
						Light: "#7D56F4",
					},
					Foreground: AdaptiveColor{
						Dark:  "#2A8f69",
						Light: "#7D56F4",
					},
				},
				SelectedDesc: AdaptiveColor{
					Dark:  "#4f4f4f",
					Light: "#7D56F4",
				},
				DimmedTitle: AdaptiveColor{
					Dark:  "#874BFD",
					Light: "#7D56F4",
				},
				DimmedDesc: AdaptiveColor{
					Dark:  "#874BFD",
					Light: "#7D56F4",
				},
				FilterMatch: ListItemStyle{
					BorderForeground: AdaptiveColor{
						Dark:  "#2A8f69",
						Light: "#7D56F4",
					},
					Foreground: AdaptiveColor{
						Dark:  "#2A8f69",
						Light: "#7D56F4",
					},
				},
			},
			Visited: AdaptiveColor{
				Dark:  "#777777",
				Light: "#777777",
			},
			Window: Window{
				Border: "normal",
				Color: AdaptiveColor{
					Dark:  "#2A8f69",
					Light: "#7D56F4",
				},
			},
			Tab: Tab{
				Color: AdaptiveColor{
					Dark:  "#2A8f69",
					Light: "#7D56F4",
				},
			},
		},
		Workers: 10,
	}
}
