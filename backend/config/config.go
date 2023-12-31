package config

type Configuration struct {
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	MySQL MySQL `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	PGSQL PGSQL `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	Jwt   JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

type App struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"` // 环境值
	Port          int    `mapstructure:"port" json:"port" yaml:"port"`
	AppName       string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl        string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
	DbType        string `mapstructure:"db_type" json:"db_type" yaml:"db_type"` // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	RouterPrefix  string `mapstructure:"router_prefix" json:"router_prefix" yaml:"router_prefix"`
	UseMultipoint bool   `mapstructure:"router_prefix" json:"use_multipoint" yaml:"use_multipoint"` // 多点登录拦截
}
