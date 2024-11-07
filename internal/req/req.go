package req

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/lcmetzger/stress_test/internal/report"
)

func MakeRequests(rep *report.Report) error {
	var wg sync.WaitGroup
	var mu sync.Mutex

	rep.TotalRequest = 0

	requestChan := make(chan struct{}, rep.Concurrency)

	for range rep.Requests {
		wg.Add(1)
		requestChan <- struct{}{}

		go func() {
			defer wg.Done()
			defer func() { <-requestChan }()

			resp, err := http.Get(rep.Url)
			if err != nil {
				fmt.Println("Erro ao fazer a requisição:", err)
				return
			}
			defer resp.Body.Close()

			mu.Lock()
			rep.TotalRequest++
			rep.MapRequest[resp.StatusCode]++
			mu.Unlock()
		}()
	}

	wg.Wait()

	return nil

}
