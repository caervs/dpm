package parser

type Command struct {
	Name       string   `json:"name"`
	Image      string   `json:"image"`
	Entrypoint string   `json:"entrypoint"`
	Context    string   `json:"context"`
	Volumes    []string `json:"volumes"`
	// Possile additions
	// * priveleged
	// * env vars
}

type CommandSet map[string]Command

type Environment struct {
	Commands CommandSet `json:"commands"`
}

type Identifier string

type DPMConfig struct {
	// Maps env ID to env name or empty string if env is nameless
	Environments map[Identifier]string `json:"environments"`
}
