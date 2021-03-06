package scheduler

import (
	"crawler-concurrent/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}


func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c 
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	//send request down to worker chan
	// s.workerChan <- r

	//并发调度器，解决循环等待（参考done.go部分内容）
	go func() {
		s.workerChan <- r
	}()
}
