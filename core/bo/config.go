package bo

type Server struct {
	Jwt    Jwt
	Mysql  Mysql
	Redis  Redis
	System System
}

type Jwt struct {
	Key string `yaml:"key"`
}

type Mysql struct {
	Url          string `yaml:"url"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	MaxIdleConns int    `yaml:"max-idle-conns"`
	MaxOpenConns int    `yaml:"max-open-conns"`
	LogMode      bool   `yaml:"log-mode"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type System struct {
	Env  string `yaml:"env"`
	Port int    `yaml:"port"`
}
