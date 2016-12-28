package eiffel

type (
	Service interface {
		Start() bool
		Shutdown()
	}

	ServiceConfig map[string]Service
)

func (e *Eiffel) InitService(s ServiceConfig) {
	for k, v := range s {
		e.serviceList = append(e.serviceList, k)
		e.services[k] = v
	}
}
