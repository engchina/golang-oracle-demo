package model

import "time"

type DBTime struct {
	time.Time
}

func (t DBTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Format(`"2006-01-02 15:04:05"`)), nil
}

func (t *DBTime) UnmarshalJSON(b []byte) error {
	tm, err := time.Parse(`"2006-01-02 15:04:05"`, string(b))
	if err != nil {
		return err
	}
	t.Time = tm
	return nil
}
