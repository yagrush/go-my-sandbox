package main

import (
	"log/slog"
	"time"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	// goto
	goto goto1
	slog.Error("goto 失敗！")

goto1:
	slog.Debug("goto 成功！")

	// chan
	ch := make(chan string)
	strs := []string{
		"hoge",
		"hige",
		"xxx",
	}

	go func(chh chan string) {
		for _, s := range strs {
			chan_str(chh, s)
			// x := <-ch
			// slog.Debug(x)
			<-time.After(2 * time.Second)
		}
		close(chh)
	}(ch)

Loop:
	for {
		select {
		case s, ok := <-ch:
			slog.Debug(s)
			if !ok {
				break Loop
			}
		case <-time.After(1 * time.Second):
			slog.Debug("まだ")
		default:
			slog.Debug("default")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func chan_str(ch chan string, str string) {
	ch <- str
}
