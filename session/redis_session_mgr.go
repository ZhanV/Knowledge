package session

import (
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type RedisSessionMgr struct {
	Addr string
	Pwd string
	pool *redis.Pool
	SessionMap map[string]Session
	RwLock sync.RWMutex
}

func NewRedisSessionMgr() *RedisSessionMgr {
	return &RedisSessionMgr{
		SessionMap: make(map[string]Session,8),
		RwLock:     sync.RWMutex{},
	}
}

func (m *RedisSessionMgr) Init(addr string , options ...string)  error{
	m.pool = &redis.Pool{
		MaxIdle:64,
		MaxActive:256,
		IdleTimeout:240*time.Second,
		Dial: func() (conn redis.Conn, e error) {
			c, err := redis.Dial("tcp",addr)
			if err!= nil {
				return  nil,err
			}
			return c,err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_,err := c.Do("PING")
			return err
		},
	}

	m.Addr = addr
	if len(options) > 0 {
		m.Pwd = options[0]
	}
	return nil
}

func (m *RedisSessionMgr) Get(key string) (Session, error)  {

	m.RwLock.RLock()
	defer m.RwLock.RUnlock()

	session , ok := m.SessionMap[key]
	if !ok {
		return session,ERR_KEY_NOT_EXISTS_IN_SESSION
	}
	return session,nil
}

func (m *RedisSessionMgr) CreateSession() (session Session , err error)  {
	m.RwLock.Lock()
	m.RwLock.Unlock()
	id :=uuid.NewV4()

	session = NewRedisSession(id.String(), m.pool)

	m.SessionMap[id.String()] = session
	return
}
