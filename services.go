package eiffel

type (
	Service interface {
	}

	ServiceConfig map[string]Service
)

func (e *Eiffel) InitService(s ServiceConfig) {
	for k, v := range s {
		e.services[k] = v
	}
}
