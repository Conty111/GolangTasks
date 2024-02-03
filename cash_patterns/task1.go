/*
1. Cache patterns

Создать хранилище сессий пользователей.
Время работы каждой сессии (TTL) составляет 2 минуты.
Описание структур и функций.

type Session struct {
    ID        string
    UserID    string
    ExpiresAt time.Time
}

type SessionManager struct {
    sessions map[string]Session
    mutex    sync.RWMutex
}

func NewSessionManager() *SessionManager
func (sm *SessionManager) StartSession(userID string) string
func (sm *SessionManager) GetSession(sessionID string) (Session, bool)

*/

package cash_patterns

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Session struct {
	ID        string
	UserID    string
	ExpiresAt time.Time
}

type SessionManager struct {
	sessions map[string]Session
	mutex    sync.RWMutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]Session),
		mutex:    sync.RWMutex{},
	}
}

func (sm *SessionManager) StartSession(userID string) string {
	numid := 100 + rand.Intn(900)
	id := fmt.Sprintf("%s%d", userID, numid)
	sm.mutex.Lock()
	_, ok := sm.sessions[id]
	for ok {
		numid = 1000 + rand.Intn(9000)
		id = fmt.Sprintf("%s%d", userID, numid)
		_, ok = sm.sessions[id]
	}
	sm.mutex.Unlock()
	sm.mutex.RLock()
	sm.sessions[id] = Session{
		ID:        id,
		UserID:    userID,
		ExpiresAt: time.Now().Add(time.Minute * 2),
	}
	sm.mutex.RUnlock()
	return id
}

func (sm *SessionManager) GetSession(sessionID string) (Session, bool) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	val, ok := sm.sessions[sessionID]
	if !ok || time.Now().After(val.ExpiresAt) {
		sm.mutex.RLock()
		delete(sm.sessions, sessionID)
		sm.mutex.RUnlock()
		return val, false
	}
	return val, true
}
