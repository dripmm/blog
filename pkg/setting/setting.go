package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	PageSize int
	JwtSecret string
	PrefixUrl string
	RuntimeRootPath string
	ImageSavePath string
	ImageMaxSize int
	ImageAllowExts []string
	ExportSavePath string
	QrCodeSavePath string
	FontSavePath string
	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string
}

var AppSetting = &App{}

type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host string
	Password string
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

func Setup() {
	// loading configuration files
	Cfg, err := ini.Load("conf/app.ini")
	CheckErr(err)

	// Mapping to app configuration
	err = Cfg.Section("app").MapTo(AppSetting)
	CheckErr(err)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	// Mapping to server configration
	err = Cfg.Section("server").MapTo(ServerSetting)
	CheckErr(err)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	// Mapping to database configration
	err = Cfg.Section("database").MapTo(DatabaseSetting)
	CheckErr(err)

	// Mapping to redis configuation
	err = Cfg.Section("redis").MapTo(RedisSetting)
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal("Fatal to error:", err)
	}
}
