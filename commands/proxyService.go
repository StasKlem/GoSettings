package commands

type ProxyService interface {
	GetUsername() string
	GetPassword() string
	GetIP() string
	GetPort() string
}
