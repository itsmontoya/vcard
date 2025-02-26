package vcf

import (
	"encoding/json"
	"time"
)

type Date string

func (d Date) Validate() (err error) {
	if len(d) == 0 {
		return
	}

	// Ensure that the date is ISO 8601
	_, err = time.Parse("2006-01-02", string(d))
	return
}

func (d *Date) UnmarshalJSON(bs []byte) (err error) {
	var str string
	if err = json.Unmarshal(bs, &str); err != nil {
		return
	}

	date := Date(str)
	if err = date.Validate(); err != nil {
		return
	}

	*d = date
	return
}
