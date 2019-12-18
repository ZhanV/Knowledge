package session

import (
	"sync"
)

type MemorySession struct {
	SessionId string
	Data      map[string]interface{}
	RwLock    sync.RWMutex
}

func (s *MemorySession) Get(key string) (val interface{}, err error) {

	s.RwLock.RLocker()
	defer s.RwLock.RUnlock()

	val, ok := s.Data[key]
	if !ok {
		return val, ERR_SESSION_NOT_EXISTS
	}

	return

}

func (s *MemorySession) Set(key string, val interface{}) error {

	s.RwLock.Lock()
	defer s.RwLock.Unlock()

	s.Data[key] = val

	return nil
}

func (s *MemorySession) Del(key string) error {

	s.RwLock.Lock()
	defer s.RwLock.Unlock()

	delete(s.Data, key)
	return nil
}

func (s *MemorySession) Save() error {
	return nil
}
