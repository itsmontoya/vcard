package vcf

import (
	"encoding/json"
	"net/url"
)

type URL string

func (u URL) Validate() (err error) {
	_, err = url.Parse(string(u))
	return
}

func (u *URL) UnmarshalJSON(bs []byte) (err error) {
	var str string
	if err = json.Unmarshal(bs, &str); err != nil {
		return
	}

	url := URL(str)
	if err = url.Validate(); err != nil {
		return
	}

	*u = url
	return
}
