package types

type (
	// LoggerOption 日志配置模块
	LoggerOption struct {
		Level      string `default:"info"`
		Name       string
		Filename   string
		MaxSize    int  `default:"100"`
		MaxAge     int  `default:"7"`
		MaxBackups int  `default:"30"`
		LocalTime  bool `default:"true"`
		Compress   bool
	}

	// RouterOption 服务运行时配置
	RouterOption struct {
		Port            int  `default:"8080"`
		MergeLog        bool `default:"false"`
		DumpResponse    bool `default:"false"`
		LogStackIfPanic bool `default:"false"`
	}

	// DatabaseOption mysql数据库配置
	DatabaseOption struct {
		Host               string `default:"localhost"`
		Port               int    `default:"3306"`
		Username           string
		Password           string
		Database           string
		MaxOpenConnections int `default:"5"`
		MaxIdleConnections int `default:"3"`
	}

	// Option 配置入口
	Option struct {
		Title       string
		Description string
		Logger      LoggerOption
		Router      RouterOption
		Database    DatabaseOption
	}
)
