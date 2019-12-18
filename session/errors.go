package session

import "errors"

var ERR_SESSION_NOT_EXISTS = errors.New("session not exists")

var ERR_KEY_NOT_EXISTS_IN_SESSION = errors.New("key not exists in session")