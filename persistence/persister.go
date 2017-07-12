package persistence

import "github.com/fermayo/dpm/api"

type Persister interface {
	StoreConfig(*api.DPMConfig) error
	GetConfig() (*api.DPMConfig, error)
	StoreEnvironment(api.Identifier, *Environemnt) error
	GetEnvironment(api.Identifier) (*Environment, error)
	RemoveEnvironment(api.Identifier) error
}
