package vcf

import (
	"bytes"
	"errors"
)

var (
	ErrEmptyFullName = errors.New("FullName field is empty")
)

type Contact struct {
	// FullName represents the display name for the contact. This is the only required field
	FullName string `json:"fullName,omitempty"`

	Prefix     string `json:"prefix,omitempty"`
	FirstName  string `json:"firstName,omitempty"`
	MiddleName string `json:"middleName,omitempty"`
	LastName   string `json:"lastName,omitempty"`
	Suffix     string `json:"suffix,omitempty"`

	Organization string `json:"organization,omitempty"`
	Title        string `json:"title,omitempty"`

	PhoneNumber string    `json:"phoneNumber,omitempty"`
	PhoneType   PhoneType `json:"phoneType,omitempty"` // Only required if PhoneNumber is set

	Email string `json:"email,omitempty"`

	StreetAddress   string `json:"streetAddress,omitempty"`
	ExtendedAddress string `json:"extendedAddress,omitempty"`
	POBox           string `json:"poBox,omitempty"`
	City            string `json:"city,omitempty"`
	State           string `json:"state,omitempty"`
	PostalCode      string `json:"postalCode,omitempty"`
	Country         string `json:"country,omitempty"`

	URL      URL  `json:"url,omitempty"`
	Birthday Date `json:"birthday,omitempty"`

	PhotoURL  URL       `json:"photoURL,omitempty"`
	PhotoType PhotoType `json:"photoType,omitempty"` // Only required if PhotoURL is set

	Note string `json:"note,omitempty"`
}

func (c *Contact) Validate() (err error) {
	var errs []error
	if len(c.FullName) == 0 {
		errs = append(errs, ErrEmptyFullName)
	}

	if len(c.PhoneNumber) > 0 {
		if err = c.PhoneType.Validate(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(c.PhotoURL) > 0 {
		if err = c.PhotoURL.Validate(); err != nil {
			errs = append(errs, err)
		}

		if err = c.PhotoType.Validate(); err != nil {
			errs = append(errs, err)
		}
	}

	if err = c.URL.Validate(); err != nil {
		errs = append(errs, err)
	}

	if err = c.Birthday.Validate(); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func (c *Contact) VCF() (out string, err error) {
	if err = c.Validate(); err != nil {
		return
	}

	buf := bytes.NewBuffer(nil)
	if err = parsed.Execute(buf, c); err != nil {
		return
	}

	out = buf.String()
	return
}
