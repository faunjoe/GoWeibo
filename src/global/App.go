package global

import (
	"ini"
	"time"
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
}

// App 初始化app
var App = &app{}

func init() {
	cfg, err := ini.InsensitiveLoad("env.ini")
	App.LaunchTime = time.Now()
}
