package utils

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func ConvertToTimestampProto(t time.Time) (*timestamp.Timestamp, error) {
	// Convert time.Time to seconds since epoch
	seconds := t.Unix()

	// Create a Timestamp using the seconds
	timestampProto := &timestamp.Timestamp{
		Seconds: seconds,
	}

	return timestampProto, nil
}
