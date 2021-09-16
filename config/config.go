package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config = new(config)

func init() {
	readConfig()
}

func readConfig() {
	v := viper.New()
	viper.AutomaticEnv()
	v.SetConfigName("app")
	v.SetConfigType("toml")
	v.AddConfigPath(".")

	v.SetDefault("log.info_file", "./log")
	v.SetDefault("log.error_file", "./log")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("v.ReadConfig error:%v", err)
		panic(err)
	}
	if err := v.Unmarshal(Config); err != nil {
		fmt.Printf("v.Unmarshal error:%v", err)
		panic(err)
	}

	v.OnConfigChange(func(in fsnotify.Event) {
		//配置文件变化
		conf := new(config)
		if err := v.Unmarshal(conf); err != nil {
			//todo 这个panic需要recovery
			panic(err)
		}

		Config = conf
	})
	v.WatchConfig()
}

type config struct {
	Debug           bool
	SwaggerFlag     string `mapstructure:"swagger_flag"`
	JwtFlag         bool   `mapstructure:"jwt_flag"`
	IrisLoggerLevel string `mapstructure:"iris_logger_level"`

	Mysql struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string `mapstructure:"db_name"`
	}

	Postgres struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string `mapstructure:"db_name"`
	}

	Redis struct {
		Addr     string
		Password string
		DB       int
	}

	Mongo struct {
		Host       string
		Port       int
		User       string
		Password   string
		AuthSource string
	}

	Web struct {
		Port     int
		PageSize int `mapstructure:"page_size"`
	}

	Secret struct {
		Jwt string
	}

	Log struct {
		InfoFile  string `mapstructure:"info_file"`
		ErrorFile string `mapstructure:"error_file"`
	}
}
