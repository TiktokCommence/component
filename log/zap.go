package log

import (
	"github.com/TiktokCommence/component/log/config"
	"github.com/TiktokCommence/component/log/writer"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger(c config.Config) (*zap.Logger, error) {
	logLevel := map[string]zapcore.Level{
		config.DebugLevel: zapcore.DebugLevel,
		config.InfoLevel:  zapcore.InfoLevel,
		config.WarnLevel:  zapcore.WarnLevel,
		config.ErrorLevel: zapcore.ErrorLevel,
	}
	//获取writer
	var builders []WriterBuilder
	if c.LogStdout { //写入控制台
		builders = append(builders, writer.NewStdoutBuilder())
	}
	if c.FileConfig != nil { //写入文件
		builders = append(builders, writer.NewFileWriterBuilder(c.FileConfig))
	}
	if c.KafkaConfig != nil { //写入kafka
		builders = append(builders, writer.NewKafkaBuilder(c.KafkaConfig))
	}
	w, err := getWriteSyncer(builders...)
	if err != nil {
		return nil, err
	}
	encoder := getEncoder(c.LogFormat)
	level, ok := logLevel[c.LogLevel] // 日志打印级别
	if !ok {
		level = logLevel["info"]
	}
	writeSyncer := zapcore.NewMultiWriteSyncer(w...)
	core := zapcore.NewCore(encoder, writeSyncer, level)
	//在warn级别及以上打印堆栈信息
	//caller需要往上跳一级
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.WarnLevel)) // zap.Addcaller() 输出日志打印文件和行数如： logger/logger_test.go:33
	// 1. zap.ReplaceGlobals 函数将当前初始化的 logger 替换到全局的 logger,
	// 2. 使用 logger 的时候 直接通过 zap.S().Debugf("xxx") or zap.L().Debug("xxx")
	// 3. 使用 zap.S() 和 zap.L() 提供全局锁，保证一个全局的安全访问logger的方式
	return logger, nil
}

// getEncoder 编码器(如何写入日志)
func getEncoder(format string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // log 时间格式 例如: 2021-09-11t20:05:54.852+0800
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 输出level序列化为全大写字符串，如 INFO DEBUG ERROR
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	if format == config.JsonFormat {
		return zapcore.NewJSONEncoder(encoderConfig) // 以json格式写入
	}
	return zapcore.NewConsoleEncoder(encoderConfig) // 以logfmt格式写入
}

func getWriteSyncer(builder ...WriterBuilder) ([]zapcore.WriteSyncer, error) {
	var syncers []zapcore.WriteSyncer
	for _, b := range builder {
		w, err := b.Build()
		if err != nil {
			return nil, err
		}
		syncers = append(syncers, zapcore.AddSync(w))
	}
	return syncers, nil
}
