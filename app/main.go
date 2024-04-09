package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	// doGoto()
	// doGoroutine()
	doErrorGroup()
}

func doGoto() {
	// goto
	goto goto1
	slog.Error("goto 失敗！")

goto1:
	slog.Debug("goto 成功！")
}

func doGoroutine() {
	// chan
	ch := make(chan string)
	strs := []string{
		"hoge",
		"hige",
		"xxx",
	}

	go func(chh chan string) {
		defer close(chh)
		for _, s := range strs {
			chanStr(chh, s)
			// x := <-ch
			// slog.Debug(x)
			<-time.After(2 * time.Second)
		}
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

func chanStr(ch chan string, str string) {
	ch <- str
}

func doErrorGroup() {
	strs := []string{
		"hoge",
		"hige",
		"xxx",
		"aaa",
		"bbb",
		"ccc",
	}

	ctx, cancel := context.WithCancel(context.Background())
	errg, ctx := errgroup.WithContext(ctx)

	errg.Go(func() error {
		return func(ctx context.Context) error {
		Loop:
			for _, s := range strs {
				select {
				case <-ctx.Done():
					slog.Debug("done!!")
					break Loop
				case <-time.After(1 * time.Second):
					slog.Debug(s)
				}
				// if i >= d {
				// 	return fmt.Errorf("%s found! error", s)
				// }
			}
			return nil
		}(ctx)
	})
	<-time.After(3 * time.Second)
	cancel()

	if err := errg.Wait(); err != nil {
		slog.Error(fmt.Sprintf("error: %v", err))
		return
	}
}
