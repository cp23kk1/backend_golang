package enum

import "database/sql/driver"

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)

func (self *Role) Scan(value interface{}) error {
	*self = Role(value.([]byte))
	return nil
}

func (self Role) Value() (driver.Value, error) {
	return string(self), nil
}
