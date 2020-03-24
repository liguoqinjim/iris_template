package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(config)

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
	if err := v.Unmarshal(Conf); err != nil {
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

		Conf = conf
	})
	v.WatchConfig()
}

type config struct {
	SwaggerFlag     string `mapstructure:"swagger_flag"`
	JwtFlag         bool   `mapstructure:"jwt_flag"`
	IrisLoggerLevel string `mapstructure:"iris_logger_level"`

	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string `mapstructure:"db_name"`
	}

	Web struct {
		Port     int
		PageSize int
	}

	Secret struct {
		Jwt string
	}

	Log struct {
		InfoFile  string `mapstructure:"info_file"`
		ErrorFile string `mapstructure:"error_file"`
	}
}
