package main

import (
	"log"
	"time"
)

func main() {
	// goto
	goto goto1
	log.Fatal("goto 失敗！")

goto1:
	log.Print("goto 成功！")

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
			// log.Print(x)
			<-time.After(2 * time.Second)
		}
		close(chh)
	}(ch)

Loop:
	for {
		select {
		case s, ok := <-ch:
			log.Print(s)
			if !ok {
				break Loop
			}
		case <-time.After(1 * time.Second):
			log.Print("まだ")
		default:
			log.Print("default")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func chan_str(ch chan string, str string) {
	ch <- str
}
