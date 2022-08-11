package cmd

import (
	"encoding/json"

	"github.com/fadellh/smart-cli/usecase"
)

type SmartSample struct {
	Schema string          `json:"$schema"`
	Global Global          `json:"global"`
	Orgs   json.RawMessage `json:"orgs"`
}

type Global struct {
	FabricVersion string              `json:"fabricVersion"`
	Tls           bool                `json:"tls"`
	Monitoring    Monitoring          `json:"monitoring"`
	Channels      []usecase.Channel   `json:"channels"`
	ChainCodes    []usecase.ChainCode `json:"chaincodes"`
}
type Monitoring struct {
	LogLevel string `json:"loglevel"`
}
