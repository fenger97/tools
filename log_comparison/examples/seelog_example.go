package examples

import (
	"github.com/cihub/seelog"
)

type seelogger struct {
}

func NewSeelogger() Logger {
	sl := &seelogger{}
	logger, err := seelog.LoggerFromConfigAsString(`
		<seelog>
			<outputs formatid="main">
				<console />
			</outputs>
			<formats>
				<format id="main" format="%Date/%Time [%LEVEL] %Msg%n" />
			</formats>
		</seelog>
		`)

	if err != nil {
		panic(err)
	}

	seelog.ReplaceLogger(logger)
	return sl
}

func (sl *seelogger) Start() {
	defer seelog.Flush()
	seelog.Debug("调试消息")
	seelog.Info("信息消息")
	seelog.Warn("警告消息")
	seelog.Error("错误消息")
}
