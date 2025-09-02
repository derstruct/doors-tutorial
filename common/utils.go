package common

import (
	"context"

	"github.com/derstruct/doors-tutorial/driver"
	"github.com/doors-dev/doors"
)

func GetSession(r doors.R) *driver.Session {
	c, err := r.GetCookie("session")
	if err != nil {
		return nil
	}
	s, found := driver.Sessions.Get(c.Value)
	if !found {
		return nil
	}
	return &s
}

type sessionKey struct{}

func StoreSession(ctx context.Context, session *driver.Session) {
	// save to the pages global "thread safe" storage
	doors.InstanceSave(ctx, sessionKey{}, session)
}

func LoadSession(ctx context.Context) *driver.Session {
	session, ok := doors.InstanceLoad(ctx, sessionKey{}).(*driver.Session)
	if !ok {
		return nil
	}
	return session
}

// helper just to check
func IsAuthorized(ctx context.Context) bool {
	return LoadSession(ctx) != nil
}
