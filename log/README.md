# 如何使用日志组件?

```go
conf := config.NewConfig(
    	//以下都是可选项
		config.WithLogFormat(config.FmtFormat),// 日志的格式【包含json和正常的】
		config.WithLogLevel(config.InfoLevel),// 日志的级别【包含info,warn,error,debug】
		config.WithLogStdout(true),//是否输出到控制台中去
    	config.WithFileConfig(&config.FileConfig{})//文件相关配置
    	/*
    		// FileConfig 将日志写入文件
			type FileConfig struct {
				LogPath           string // 输出日志文件路径
				LogFileName       string // 输出日志文件名称
				LogFileMaxSize    int    // 【日志分割】单个日志文件最多存储量 单位(mb)
				LogFileMaxBackups int    // 【日志分割】日志备份文件最多数量
				LogMaxAge         int    // 日志保留时间，单位: 天 (day)
				LogCompress       bool   // 是否压缩日志
			}
    	*/
    	config.WithKafkaConfig(&config.KafkaConfig{})//kafka相关配置
    	/*
    		// KafkaConfig 将日志发送到kafka，由专门的日志收集系统来读取
            type KafkaConfig struct {
                BrokersAddr []string //brokers的地址
                TopicName   string	//topic的name
            }
    	*/
	)
//默认配置是
/*
var defaultConfig = Config{
	LogLevel:  "info",
	LogFormat: "log_fmt",
	LogStdout: true,
}
*/

zapLogger, err := NewZapLogger(conf)

//zapLogger是*zap.Logger类型

//如果你想适配kratos的接口的话
//引用  kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"这个包
//然后使用kzap.NewLogger(NewZapLogger(zapLogger))来生成log.Logger(kratos中日志接口)


```

其他的可以参考[zap_test.go](./zap_test.go)