package global

import "gin/extend/setting"
//在读取了文件的配置信息后，还是不够的，因为我们需要将配置信息和应用程序关联起来，
//我们才能够去使用它，因此在项目目录下的 global 目录下新建 setting.go 文件
var (
	ServerSetting  *setting.ServerSetting
	AppSetting     *setting.AppSetting
	LogSetting     *setting.LogSetting
	MySQLSetting   *setting.MySQLSetting
	RedisSetting   *setting.RedisSetting
	RabbiMQSetting *setting.RabbitMQSetting
)
