package types

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
)

// UUID type is a wrapper for https://github.com/google/uuid
// It fix sql Scan and Value methods
type UUID struct {
	uuid.UUID
}

// NewUUID is UUID constructor
func NewUUID() *UUID {
	return &UUID{uuid.New()}
}

// ParseUUIDString func
func ParseUUIDString(s string) (*UUID, error) {
	id, err := uuid.Parse(s)
	return &UUID{id}, err
}

// MustParseUUIDString func
func MustParseUUIDString(s string) *UUID {
	id, err := uuid.Parse(s)
	if err != nil {
		panic(err)
	}
	return &UUID{id}
}

// Nil method
func (u *UUID) Nil() bool {
	return u.UUID == uuid.Nil
}

// Scan method
func (u *UUID) Scan(value interface{}) (err error) {
	var base uuid.UUID
	switch value.(type) {
	case []byte:
		base, err = uuid.FromBytes(value.([]byte))
	case string:
		base, err = uuid.Parse(value.(string))
	case uuid.UUID:
		base = value.(uuid.UUID)
	case *uuid.UUID:
		base = *(value.(*uuid.UUID))
	default:
		err = fmt.Errorf("can not convert '%v' to UUID", value)
	}
	if err == nil {
		u.UUID = base
	}
	return
}

// Value method
func (u UUID) Value() (driver.Value, error) {
	return u.UUID.MarshalBinary()
}

// ContainsUUID method
func ContainsUUID(l []UUID, u UUID) bool {
	for _, i := range l {
		if i == u {
			return true
		}
	}
	return false
}
