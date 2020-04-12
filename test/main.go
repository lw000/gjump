// test project main.go
package main

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
	tyhttp "tuyue/tuyue_common/network/http"
	"tuyue/tuyue_web_v2/test/config"
)

var (
	successCount uint64 = 1
	failCount    uint64 = 1

	cfg *config.Config
	wg  *sync.WaitGroup
)

func test(w *sync.WaitGroup) {
	w.Add(1)
	defer func() {
		w.Done()
	}()

	_, body, er := tyhttp.DoHttpGet(cfg.Url, nil, time.Second*time.Duration(10))
	if er != nil {
		atomic.AddUint64(&failCount, 1)
		log.Println(er)
		return
	}
	atomic.AddUint64(&successCount, 1)
	log.Println(string(body))
}

func main() {
	defer func() {
		time.Sleep(time.Millisecond * time.Duration(100))
	}()

	cfg = config.NewConfig()
	err := cfg.Load("./conf/conf.json")
	if err != nil {
		return
	}

	wg = &sync.WaitGroup{}
	start := time.Now()
	for i := 0; i < cfg.Count; i++ {
		go test(wg)
	}
	wg.Wait()

	end := time.Now()
	log.Printf("success:[%d], fail:[%d], [%v]", successCount, failCount, end.Sub(start))
}
