package pfcpType

import (
	"encoding/binary"
	"fmt"
)

type UserPlaneInactivityTimer struct {
	UserPlaneInactivityTimerdata []byte
}

func (t *UserPlaneInactivityTimer) MarshalBinary() ([]byte, error) {
	return t.UserPlaneInactivityTimerdata, nil
}

func (t *UserPlaneInactivityTimer) UnmarshalBinary(data []byte) error {
	if len(data) != 4 {
		return fmt.Errorf("UserPlaneInactivityTimer: expected 4 bytes, got %d", len(data))
	}
	t.UserPlaneInactivityTimerdata = make([]byte, 4)
	copy(t.UserPlaneInactivityTimerdata, data)
	return nil
}

// Seconds returns the inactivity timer value in seconds.
func (t *UserPlaneInactivityTimer) Seconds() uint32 {
	if len(t.UserPlaneInactivityTimerdata) != 4 {
		return 0
	}
	return binary.BigEndian.Uint32(t.UserPlaneInactivityTimerdata)
}
