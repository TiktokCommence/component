package writer

import (
	"component/log/config"
	"fmt"
	"github.com/natefinch/lumberjack"
	"io"
	"os"
	"path/filepath"
)

type FileWriterBuilder struct {
	conf *config.FileConfig
}

func NewFileWriterBuilder(conf *config.FileConfig) *FileWriterBuilder {
	return &FileWriterBuilder{
		conf: conf,
	}
}
func (f *FileWriterBuilder) Build() (io.Writer, error) {
	// 判断日志路径是否存在，如果不存在就创建
	if exist := isExist(f.conf.LogPath); !exist {
		if err := os.MkdirAll(f.conf.LogPath, os.ModePerm); err != nil {
			return nil, fmt.Errorf("failed to create file writer: %v", err)
		}
	}
	// 日志文件 与 日志切割 配置
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(f.conf.LogPath, f.conf.LogFileName), // 日志文件路径
		MaxSize:    f.conf.LogFileMaxSize,                             // 单个日志文件最大多少 mb
		MaxBackups: f.conf.LogFileMaxBackups,                          // 日志备份数量
		MaxAge:     f.conf.LogMaxAge,                                  // 日志最长保留时间
		Compress:   f.conf.LogCompress,                                // 是否压缩日志
	}
	return lumberJackLogger, nil
}

// isExist 判断文件或者目录是否存在
func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
