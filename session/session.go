package session

type Session interface {
	Get(key string) (interface{},error)
	Set(key string ,val interface{}) error
	Del(key string) error
	Save() error
}
