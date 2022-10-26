package impl

type PingImpl struct{}

func (p PingImpl) Ping() string {
	return "pong"
}

func ProvidePingService() PingImpl {
	return PingImpl{}
}
