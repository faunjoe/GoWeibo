package global

import (
	"GoWeibo/util"
	"time"

	"github.com/go-ini/ini"
)

const (
	// Debug 开发版本
	Debug = "Debug"
	// ReleaseDebug 外测版本
	ReleaseDebug = "ReleaseDebug"
	// Release 正式版
	Release = "Release"
)

type app struct {
	Name       string
	Version    string
	LaunchTime time.Time
	Date       time.Time
	Uptime     time.Duration
	Env        string
	Port       string
}

// App 初始化app
var App = &app{}

func init() {
	cfg, err := ini.InsensitiveLoad("env.ini")
	util.Err(err)
	secGlobal, errG := cfg.GetSection("global")
	util.Err(errG)
	secListen, errL := cfg.GetSection("listen")
	util.Err(errL)

	App.Name = secGlobal.Key("appName").String()
	App.Version = secGlobal.Key("appVersion").String()
	App.Env = secGlobal.Key("appVersion").String()
	App.LaunchTime = time.Now()
	App.Port = secListen.Key("port").String()
}
