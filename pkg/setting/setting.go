package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"time"
)

type App struct {
	PageSize    int
	LogPathFile string
	LogLevel    string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	Timeout     time.Duration
	MaxPoolSize uint64
}

var DatabaseSetting = &Database{}

var cfg *ini.File

// Init the configuration file app.ini
func Init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		fmt.Println("parse app.ini error")
		panic(err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	DatabaseSetting.Timeout = DatabaseSetting.Timeout * time.Second

}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		fmt.Println("cfg section err")
		panic(err)
	}
	fmt.Printf("%s setting:%v\n", section, v)
}
