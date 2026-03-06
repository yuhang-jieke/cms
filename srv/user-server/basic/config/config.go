package config

type Mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}
type Redis struct {
	Host     string
	Port     int
	Password string
	Database int
}
type Nacos struct {
	Addr      string
	Prot      int
	Namespace string
	DataId    string
	Group     string
}
type AppConfig struct {
	Mysql
	Redis
	Nacos
}
