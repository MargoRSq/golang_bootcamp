package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var amount = 17
var wg = &sync.WaitGroup{}

func getBody(url string) (body string) {
	response, _ := http.Get(url)
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	body = string(bodyBytes)
	return
}

func crawlWeb(inCh <-chan string, ctx context.Context) <-chan string {
	guard := make(chan struct{}, 8)

	outCh := make(chan string, len(inCh))
	defer close(outCh)
	for url := range inCh {
		// time.Sleep(time.Duration(1 * int(time.Second)))
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				// break
				return
			default:
				// mut.Lock()
				// cors++
				// fmt.Printf("%d goroutine started\n", cors)
				// mut.Unlock()

				// time.Sleep(time.Duration(3 * int(time.Second)))

				guard <- struct{}{}
				outCh <- getBody(url)

				<-guard
				// fmt.Println("goroutine ended")
			}
		}(url)
	}
	wg.Wait()
	return outCh
}

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		fmt.Println("Program killed !")
		cancel()
	}()

	inCh := make(chan string, amount)
	for i := 0; i < amount; i++ {
		inCh <- "https://httpbin.org/user-agent"
	}

	close(inCh)
	outCh := crawlWeb(inCh, ctx)

	// for res := range outCh {
	// 	fmt.Println(res)
	// }
	fmt.Println(len(outCh))

}
