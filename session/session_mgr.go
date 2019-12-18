package session

type SessionMgr interface {
	Init(addr string, options ...string) error
	Get(key string) (Session, error)
	CreateSession() (Session,error)
}

