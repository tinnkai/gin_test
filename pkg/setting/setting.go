package setting

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v3"
)

type App struct {
	JwtSecret       string   `yaml:"jwtSecret"`
	PageSize        int      `yaml:"pageSize"`
	PrefixUrl       string   `yaml:"prefixUrl"`
	RuntimeRootPath string   `yaml:"runtimeRootPath"`
	ImageSavePath   string   `yaml:"imageSavePath"`
	ImageMaxSize    int      `yaml:"imageMaxSize"`
	ImageAllowExts  []string `yaml:"imageAllowExts"`
	ExportSavePath  string   `yaml:"exportSavePath"`
	QrCodeSavePath  string   `yaml:"qrCodeSavePath"`
	FontSavePath    string   `yaml:"fontSavePath"`

	LogSavePath string `yaml:"logSavePath"`
	LogSaveName string `yaml:"logSaveName"`
	LogFileExt  string `yaml:"logFileExt"`
	TimeFormat  string `yaml:"timeFormat"`
}
type AppNode struct {
	App
}

var AppSetting = &AppNode{}

type Server struct {
	RunMode      string        `yaml:"runMode"`
	HttpPort     int           `yaml:"httpPort"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}
type ServerNode struct {
	Server
}

var ServerSetting = &ServerNode{}

type Database struct {
	Type        string `yaml:"type"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Name        string `yaml:"name"`
	TablePrefix string `yaml:"tablePrefix"`
}
type DatabaseNode struct {
	Database
}

var DatabaseSetting = &DatabaseNode{}

type Mongo struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
}
type MongoNode struct {
	Mongo
}

var MongoSetting = &MongoNode{}

type Redis struct {
	Host        string        `yaml:"host"`
	Password    string        `yaml:"password"`
	MaxIdle     int           `yaml:"maxIdle"`
	MaxActive   int           `yaml:"maxActive"`
	IdleTimeout time.Duration `yaml:"idleTimeout"`
}
type RedisNode struct {
	Redis
}

var RedisSetting = &RedisNode{}

type Session struct {
	SessionOn             string `yaml:"sessionOn"`
	SessionName           string `yaml:"sessionName"`
	SessionProvider       string `yaml:"sessionProvider"`
	Host                  string `yaml:"host"`
	SessionCookieLifetime int    `yaml:"sessionCookieLifetime"`
	SessionGcmaxLifetime  int    `yaml:"sessionGcmaxLifetime"`
	MaxIdleConns          int    `yaml:"maxIdleConns"`
	PoolSize              int    `yaml:"poolSize"`
	Db                    string `yaml:"db"`
}
type SessionNode struct {
	Session
}

var SessionSetting = &Session{}

type Log struct {
	Skip int `yaml:"skip"`
}
type LogNode struct {
	Log
}

var LogSetting = &LogNode{}

type Rediskey struct {
	AuthUserKey string `yaml:"authUserKey"`
}
type RediskeyNode struct {
	Rediskey
}

var RediskeySetting = &RediskeyNode{}

// Setup initialize the configuration instance
func Setup() {
	yamlFile, err := ioutil.ReadFile("conf/app.yaml")
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf("Yaml open file error: %v", err)
	}

	mapTo(yamlFile, AppSetting)
	mapTo(yamlFile, ServerSetting)
	mapTo(yamlFile, DatabaseSetting)
	mapTo(yamlFile, MongoSetting)
	mapTo(yamlFile, RedisSetting)
	mapTo(yamlFile, SessionSetting)
	mapTo(yamlFile, LogSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map
func mapTo(in []byte, out interface{}) {
	err := yaml.Unmarshal(in, out)
	if err != nil {
		log.Fatalf("setting map to error: %v", err)
	}
}
