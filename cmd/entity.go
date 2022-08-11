package cmd

type SmartSample struct {
	Schema string `json:"$schema"`
	Global Global `json:"global"`
}

type Global struct {
	FabricVersion string     `json:"fabricVersion"`
	Tls           bool       `json:"tls"`
	Monitoring    Monitoring `json:"monitoring"`
}
type Monitoring struct {
	LogLevel string `json:"loglevel"`
}
