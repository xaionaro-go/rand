package fastrand

import (
	"time"
)

func init() {
	now := time.Now()
	nano := now.UnixNano()
	uint32nPosition = uint32(nano) + 1
	uint32nPositionSafe = uint32(nano) + 2
	uint64nPosition = uint64(nano) + 3
	uint64nPositionSafe = uint64(nano) + 4
}
