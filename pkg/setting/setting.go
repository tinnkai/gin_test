package setting

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
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

	LogSavePath    string `yaml:"logSavePath"`
	LogSaveName    string `yaml:"logSaveName"`
	LogFileExt     string `yaml:"logFileExt"`
	DateFormat     string `yaml:"dateFormat"`
	DateTimeFormat string `yaml:"dateTimeFormat"`
}

var AppSetting = &App{}

type Server struct {
	RunMode      string        `yaml:"runMode"`
	HttpPort     int           `yaml:"httpPort"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

var ServerSetting = &Server{}

type Database struct {
	Type         string `yaml:"type"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Name         string `yaml:"name"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	TablePrefix  string `yaml:"tablePrefix"`
}

var DatabaseSetting = &Database{}

type DatabaseActivity struct {
	Type         string `yaml:"type"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Name         string `yaml:"name"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	TablePrefix  string `yaml:"tablePrefix"`
}

var DatabaseActivitySetting = &DatabaseActivity{}

type Mongo struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
}

var MongoSetting = &Mongo{}

type Redis struct {
	Host        string        `yaml:"host"`
	Password    string        `yaml:"password"`
	Db          int           `yaml:"db"`
	MaxIdle     int           `yaml:"maxIdle"`
	MaxActive   int           `yaml:"maxActive"`
	IdleTimeout time.Duration `yaml:"idleTimeout"`
}

var RedisSetting = &Redis{}

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

var SessionSetting = &Session{}

type Log struct {
	Skip int `yaml:"skip"`
}

var LogSetting = &Log{}

// ip 相关
type Ip struct {
	WhiteList []string `yaml:"whiteList"`
	BlackList []string `yaml:"blackList"`
}

var IpSetting = &Ip{}

// Setup initialize the configuration instance
func Setup() {
	// 针对不同的环境获取配置文件
	// 获取 CONFIGOR_ENV 环境变量(dev/test/pro) 不设置默认 pro
	configorEnv := os.Getenv("CONFIGOR_ENV")
	if configorEnv == "" {
		configorEnv = "pro"
	}
	configName := "app_" + configorEnv
	viper.SetConfigName(configName) // 配置文件的文件名，没有扩展名，如 .yaml, .toml 这样的扩展名
	viper.SetConfigType("yaml")     // 设置扩展名。在这里设置文件的扩展名。另外，如果配置文件的名称没有扩展名，则需要配置这个选项
	viper.AddConfigPath("conf/")    // 查找配置文件所在路径
	err := viper.ReadInConfig()     // 搜索并读取配置文件
	if err != nil {                 // 处理错误
		panic(fmt.Errorf("Fatal error open config file: %v", err))
	}
	// 设置
	SetConfig()
}

// 设置配置
func SetConfig() {
	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("databaseactivity", DatabaseActivitySetting)
	mapTo("mongo", MongoSetting)
	mapTo("redis", RedisSetting)
	mapTo("session", SessionSetting)
	mapTo("log", LogSetting)
	mapTo("ip", IpSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// 转换到指定结构体
func mapTo(key string, out interface{}) {
	err := viper.UnmarshalKey(key, out)
	if err != nil {
		panic(fmt.Errorf("setting unable to decode into struct: %v", err))
	}
}
