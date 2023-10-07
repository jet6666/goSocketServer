package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func f(a any) {
	switch a.(type) {
	case int:
		fmt.Println("int ")
	case string:
		fmt.Println("is string ")
	}
}

//全大写
type AppConfig struct {
	*RedisConfig `mapstructure:"redis"`
	*MySQLConfig `mapstructure:"mysql"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}

type MySQLConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}

//https://blog.csdn.net/mrtwenty/article/details/97621402
//YAML ： https://www.ruanyifeng.com/blog/2016/07/yaml.html
// YAML y读取 https://blog.csdn.net/qq_51898139/article/details/126482375
func main() {
	f("11111")
	f(123)
	for _, args := range os.Args {
		log.Println("parasm=", args)
	}
	if len(os.Args) < 2 {
		panic("start failed: config parameter missing ")
	}
	configFile := os.Args[1]
	config := viper.New()
	config.SetConfigFile(configFile)
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	log.Println(config.Get("mysql.host"))
	log.Println(config.Get("redis"))
	redisConf :=config.Get("redis")
	if redisConf ==nil {
		panic("redis config err ")
	}

	//第二种方式，直接map 起来
	mapConf := new (AppConfig)
	if err := config.Unmarshal(mapConf) ;err !=nil {
		log.Println("unable to marshal " ,err.Error())
		return
	}
	log.Println(mapConf.RedisConfig.Port)
	log.Println(mapConf.MySQLConfig.Port)
}
