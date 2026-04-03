package main

import (
	"context"
	"sync"
)

func scoreboardManager(ctx context.Context, in <-chan func(map[string]int)) {
	scoreboard := map[string]int{}
	for {
		select {
		case <-ctx.Done():
			return
		case f := <-in:
			f(scoreboard)
		}
	}
}

type ChannelScoreboardManager chan func(map[string]int)

func NewChannelScoreboardManager(ctx context.Context) ChannelScoreboardManager {
	ch := make(ChannelScoreboardManager)
	go scoreboardManager(ctx, ch)
	return ch
}

func (csm ChannelScoreboardManager) Update(name string, val int) {
	csm <- func(m map[string]int) {
		m[name] = val
	}
}

func (csm ChannelScoreboardManager) Read(name string) (int, bool) {
	type Result struct {
		out int
		ok  bool
	}
	resultCh := make(chan Result)
	csm <- func(m map[string]int) {
		out, ok := m[name]
		resultCh <- Result{out, ok}
	}
	result := <-resultCh
	return result.out, result.ok
}

// mutex score board
type MutexScoreboardManager struct {
	l          sync.RWMutex
	scoreboard map[string]int
}

func NewMutexScoreboardManager() *MutexScoreboardManager {
	return &MutexScoreboardManager{
		scoreboard: map[string]int{},
	}
}

func (msm *MutexScoreboardManager) Update(name string, val int) {
	msm.l.Lock()
	defer msm.l.Unlock()
	msm.scoreboard[name] = val
}

func (msm *MutexScoreboardManager) Read(name string) (int, bool) {
	msm.l.RLock()
	defer msm.l.RUnlock()
	val, ok := msm.scoreboard[name]
	return val, ok
}
