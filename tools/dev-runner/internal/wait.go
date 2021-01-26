package internal

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func WaitForService(service string, timeout time.Duration) error {
	return WaitForServices([]string{service}, timeout)
}

func WaitForServices(services []string, timeout time.Duration) error {
	var depChan = make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(len(services))
	go func() {
		for _, s := range services {
			go func(s string) {
				defer wg.Done()
				for {
					if _, err := net.Dial("tcp", s); err == nil {
						return
					}
					time.Sleep(1 * time.Second)
				}
			}(s)
		}
		wg.Wait()
		close(depChan)
	}()
	select {
	case <-depChan:
		return nil
	case <-time.After(timeout):
		return fmt.Errorf("services aren't ready in %s", timeout)
	}
}
