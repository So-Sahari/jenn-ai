package state

import (
	"sync"
)

type AppState struct {
	mu             sync.RWMutex
	platform       string
	model          string
	conversationID int
	maxTokens      int
	topP           float64
	topK           int
	temperature    float64
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

func (s *AppState) SetConversationID(conversationID int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.conversationID = conversationID
}

func (s *AppState) GetConversationID() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.conversationID
}

func (s *AppState) SetMaxTokens(maxTokens int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.maxTokens = maxTokens
}

func (s *AppState) GetMaxTokens() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.maxTokens
}

func (s *AppState) SetTopK(topK int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.topK = topK
}

func (s *AppState) GetTopK() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.topK
}

func (s *AppState) SetTopP(topP float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.topP = topP
}

func (s *AppState) GetTopP() float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.topP
}

func (s *AppState) SetTemperature(temperature float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.temperature = temperature
}

func (s *AppState) GetTemperature() float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.temperature
}
