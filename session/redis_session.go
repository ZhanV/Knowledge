package session

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"sync"
)

const (
	SessionFlagNone = iota
	SessionFlagModify
	SessionFlagLoad

)


type RedisSession struct {
	SessionId string
	Pool *redis.Pool
	Data map[string]interface{}
	RwLock sync.RWMutex
	flag int
}

func NewRedisSession(id string , pool *redis.Pool) *RedisSession {
	return &RedisSession{
		SessionId: id,
		Pool:      pool,
		Data:      make(map[string]interface{},8),
		RwLock:    sync.RWMutex{},
		flag:      SessionFlagNone,
	}
}

func (s *RedisSession) loadFromRedis()error{
	conn := s.Pool.Get()
	reply, err := conn.Do("GET",s.SessionId)
	if err != nil {
		return err
	}

	data , err := redis.String(reply,err)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(data),&s.Data)
	return  err
}


func (s *RedisSession) Get(key string) (val interface{}, err error) {

	s.RwLock.RLocker()
	defer s.RwLock.RUnlock()

	if s.flag == SessionFlagNone {
		err = s.loadFromRedis()
		if err != nil {
			return
		}
	}

	val, ok := s.Data[key]
	if !ok {
		return val, ERR_SESSION_NOT_EXISTS
	}

	return

}

func (s *RedisSession) Set(key string, val interface{}) error {

	s.RwLock.Lock()
	defer s.RwLock.Unlock()

	s.Data[key] = val
	s.flag = SessionFlagModify
	return nil
}

func (s *RedisSession) Del(key string) error {

	s.RwLock.Lock()
	defer s.RwLock.Unlock()

	delete(s.Data, key)
	s.flag = SessionFlagModify
	return nil
}

func (s *RedisSession) Save() error {
	s.RwLock.Lock()
	defer s.RwLock.Unlock()

	if s.flag == SessionFlagNone {
		return nil
	}

	data, err := json.Marshal(s.Data)
	if err != nil {
		return err
	}

	conn := s.Pool.Get()
	_, err = conn.Do("SET", s.SessionId, string(data))
	if err != nil {
		return err
	}
	s.flag = SessionFlagNone
	return nil
}
