package state

import (
	"sync"
)

type AppState struct {
	mu       sync.RWMutex
	platform string
	model    string
}

var (
	state *AppState
	once  sync.Once
)

func GetState() *AppState {
	once.Do(func() {
		state = &AppState{}
	})
	return state
}

func (s *AppState) SetPlatform(platform string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.platform = platform
}

func (s *AppState) GetPlatform() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.platform
}

func (s *AppState) SetModel(model string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.model = model
}

func (s *AppState) GetModel() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.model
}
