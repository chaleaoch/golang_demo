package entity

type Alert struct {
	ID      int64
	Message string
	Level   string
}

type MIp struct {
	Ip       string
	Username string
	Password string
	Port     string
}
