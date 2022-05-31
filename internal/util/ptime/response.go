package ptime

import (
	"encoding/json"
	"time"
)

// TimeResponse ...
type TimeResponse struct {
	Time time.Time
}

// UnmarshalJSON ...
func (t *TimeResponse) UnmarshalJSON(b []byte) error {
	if string(b) == "" || string(b) == "\"\"" {
		return nil
	}
	return json.Unmarshal(b, &t.Time)
}

// MarshalJSON ...
func (t TimeResponse) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return json.Marshal("")
	}
	return json.Marshal(t.Time.Format(dateLayoutFull))
}

// TimeResponseInit ...
func TimeResponseInit(t time.Time) *TimeResponse {
	return &TimeResponse{Time: t}
}
