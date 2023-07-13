package conf

import (
	"strings"
)

type ConfigStruct struct {
	Port          int               `toml:"port"`
	Tokens        []string          `toml:"tokens"`
	Gateway       string            `toml:"gateway"`
	Bot           map[string]string `toml:"bot"`
	SimulateRoles int               `toml:"simulate-roles"`
	RateLimit     int               `toml:"rate-limit"`
	CoolDown      int               `toml:"cool-down"`
	Timeout       int               `toml:"timeout"`
	Key           string            `toml:"key"`

}

type ModelDef struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	OwnedBy string `json:"owned_by"`
}

type ModelsResp struct {
	Object string     `json:"object"`
	Data   []ModelDef `json:"data"`
}

var Conf ConfigStruct

var Models ModelsResp


func Setup() {
	//v, err := os.ReadFile("config.toml")
	//if err != nil {
	//	panic(err)
	//}
	//err = toml.Unmarshal(v, &Conf)
	//if err != nil {
	//	panic(err)
	//}
	if Conf.Port == 0 {
		Conf.Port = 8080

	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	if Conf.Bot == nil {
		Conf.Bot = map[string]string{
			"gpt-3.5-turbo":           "chinchilla",
			"gpt-4":                   "beaver",
			"gpt-3.5-turbo-0301":      "a2",
			"gpt-4-32k":               "a2_100k",
			"gpt-4-0314":              "a2_2",
			"gpt-Sage":                "capybara",
			"gpt-ChatGPT":             "chinchilla",
			"gpt-GPT-4":               "beaver",
			"gpt-Claude-instant":      "a2",
			"gpt-Claude-instant-100k": "a2_100k",
			"gpt-Claude+":             "a2_2",
			"Sage":                    "capybara",
			"ChatGPT":                 "chinchilla",
			"GPT-4":                   "beaver",
			"Claude-instant":          "a2",
			"Claude-instant-100k":     "a2_100k",
			"Claude+":                 "a2_2",
		}
	}

	Models.Object = ""

	for key := range Conf.Bot {
		Models.Data = append(Models.Data, ModelDef{
			ID:      key,
			Object:  "",
			Created: 0,
			OwnedBy: "",
		})
	}
}