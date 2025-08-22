package common

import (
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
