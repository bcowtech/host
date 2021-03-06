package host

import "github.com/bcowtech/host/internal"

func Startup(app interface{}) *Starter {
	return internal.NewStarter(app)
}

func RegisterHostService(starter *Starter, service HostService) {
	internal.RegisterHostService(starter, service)
}

func StdHostServiceInstance() HostService {
	return internal.StdHostServiceInstance()
}
