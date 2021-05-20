package setting

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

// Setting 全局配置
type Setting struct {
	v *viper.Viper
}

func (s *Setting) ReadSection(k string, v interface{}) (err error) {
	err = s.v.UnmarshalKey(k, v)
	return err
}

func NewSetting() (s *Setting, err error) {
	v := viper.New()
	switch gin.Mode() {
	case gin.DebugMode:
		v.SetConfigName("app.dev")
	case gin.TestMode:
		v.SetConfigName("app.test")
	case gin.ReleaseMode:
		v.SetConfigName("app.release")
	}
	v.AddConfigPath("./config")
	v.SetConfigType("yaml")
	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{v}, nil
}

//ServerSetting Server配置
type ServerSetting struct {
	Port string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

// AppSetting app配置
type AppSetting struct {
	PageSize int
}

// LogSetting 日志配置
type LogSetting struct {
	FileName string
	Ext string
	SavePath string
}

// MySQLSetting MySQL配置
type MySQLSetting struct {
	Host string
	Port string
	Database string
	Username string
	Password string
	Prefix string
}

// RedisSetting Redis配置
type RedisSetting struct {
	Host string
	Port string
}

// RabbitMQSetting RabbitMQ配置
type RabbitMQSetting struct {
	Host string
	Port string
	Username string
	Password string
}
