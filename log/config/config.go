package config

const (
	InfoLevel  = "info"
	DebugLevel = "debug"
	WarnLevel  = "warn"
	ErrorLevel = "error"

	FmtFormat  = "log_fmt"
	JsonFormat = "json"
)

var defaultConfig = Config{
	LogLevel:  InfoLevel,
	LogFormat: FmtFormat,
	LogStdout: true,
}

// FileConfig 将日志写入文件
type FileConfig struct {
	LogPath           string // 输出日志文件路径
	LogFileName       string // 输出日志文件名称
	LogFileMaxSize    int    // 【日志分割】单个日志文件最多存储量 单位(mb)
	LogFileMaxBackups int    // 【日志分割】日志备份文件最多数量
	LogMaxAge         int    // 日志保留时间，单位: 天 (day)
	LogCompress       bool   // 是否压缩日志
}

// KafkaConfig 将日志发送到kafka，由专门的日志收集系统来读取
type KafkaConfig struct {
	BrokersAddr []string
	TopicName   string
}

type Config struct {
	LogLevel    string // 日志打印级别 debug  info  warning  error
	LogFormat   string // 输出日志格式	logfmt, json
	LogStdout   bool   // 是否输出到控制台，如果其他的writer都为nil，这个会被强制打开
	FileConfig  *FileConfig
	KafkaConfig *KafkaConfig
}
type Option func(*Config)

func NewConfig(opt ...Option) Config {
	c := defaultConfig
	for _, o := range opt {
		o(&c)
	}
	if c.FileConfig == nil && c.KafkaConfig == nil {
		c.LogStdout = true
	}
	return c
}

func WithLogLevel(level string) Option {
	return func(o *Config) {
		o.LogLevel = level
	}

}
func WithLogFormat(format string) Option {
	return func(o *Config) {
		o.LogFormat = format
	}
}
func WithLogStdout(stdout bool) Option {
	return func(o *Config) {
		o.LogStdout = stdout
	}
}
func WithFileConfig(fileCfg *FileConfig) Option {
	return func(o *Config) {
		o.FileConfig = fileCfg
	}
}
func WithKafkaConfig(kafkaCfg *KafkaConfig) Option {
	return func(o *Config) {
		o.KafkaConfig = kafkaCfg
	}
}
