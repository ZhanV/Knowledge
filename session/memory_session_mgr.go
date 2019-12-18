package session

import (
	uuid "github.com/satori/go.uuid"
	"sync"
)

type MemorySessionMgr struct {
	SessionMap map[string]Session
	RwLock sync.RWMutex
}

func NewMemorySessionMgr() *MemorySessionMgr {
	return &MemorySessionMgr{
		SessionMap: make(map[string]Session, 8) ,
		RwLock:     sync.RWMutex{},
	}
}

func (m *MemorySessionMgr) Init(addr string, options ...string) error{

	return nil
}

func NewMemorySession(id string) *MemorySession {
	return &MemorySession{
		SessionId: id,
		Data:      make(map[string]interface{},8),
		RwLock:    sync.RWMutex{},
	}
}

func (m *MemorySessionMgr) Get(key string) (Session, error)  {

	m.RwLock.RLock()
	defer m.RwLock.RUnlock()

	session , ok := m.SessionMap[key]
	if !ok {
		return session,ERR_KEY_NOT_EXISTS_IN_SESSION
	}
	return session,nil
}

func (m *MemorySessionMgr) CreateSession() (session Session , err error)  {
	m.RwLock.Lock()
	m.RwLock.Unlock()
	id :=uuid.NewV4()

	session = NewMemorySession(id.String())

	m.SessionMap[id.String()] = session
	return
}