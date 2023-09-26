package httpapi

type ClientConfig struct {
	Version string      `json:"version"`
	Login   LoginConfig `json:"login"`
}

type LoginConfig struct {
	Providers []*ProviderConfig `json:"providers"`
}

type ProviderConfig struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}