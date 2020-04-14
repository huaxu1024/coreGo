package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var c = 1

func doSome(i int) error {
	c++
	fmt.Println(c)
	if c > 3 {
		return errors.New("err occur")
	}
	return nil
}

func speakMemo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx.Done")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("exec default func")
			err := doSome(3)
			if err != nil {
				fmt.Printf("cancelFunc()")
			}
		}
	}
}

func contextGo() {
	rootContext := context.Background()
	//ctx, cancelFunc := context.WithTimeout(rootContext, time.Duration(1)*time.Second)
	ctx, cancelFunc := context.WithDeadline(rootContext, time.Now().Add(time.Duration(1)*time.Second))
	defer cancelFunc()
	go speakMemo(ctx)
	time.Sleep(time.Second * 5)
}
