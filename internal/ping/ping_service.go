package ping

import (
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/ping/impl"
)

type Ping interface {
	Ping() string
}

func ProvidePingService() Ping {
	return impl.ProvidePingService()
}
