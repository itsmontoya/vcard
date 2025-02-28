package vcf

import "testing"

func TestContact_VCF(t *testing.T) {
	type testcase struct {
		name string
		c    Contact

		wantOut string
		wantErr bool
	}

	tests := []testcase{
		{
			name: "basic",
			c: Contact{
				FullName:        "John Doe",
				Prefix:          "",
				FirstName:       "John",
				MiddleName:      "",
				LastName:        "Doe",
				Suffix:          "",
				PhoneNumber:     "+13034567890",
				PhoneType:       "CELL",
				Email:           "",
				Organization:    "",
				StreetAddress:   "",
				ExtendedAddress: "",
				City:            "",
				State:           "",
				PostalCode:      "",
				Country:         "",
			},

			wantOut: `BEGIN:VCARD
VERSION:3.0
FN:John Doe
N:Doe;John;;;
TEL;TYPE=CELL:+13034567890
END:VCARD`,
			wantErr: false,
		},
		{
			name: "errors",
			c: Contact{
				FullName:        "",
				Prefix:          "",
				FirstName:       "John",
				MiddleName:      "",
				LastName:        "Doe",
				Suffix:          "",
				PhoneNumber:     "+13034567890",
				PhoneType:       "",
				Email:           "",
				Organization:    "",
				StreetAddress:   "",
				ExtendedAddress: "",
				City:            "",
				State:           "",
				PostalCode:      "",
				Country:         "",
				URL:             "82s091://",
				Birthday:        "20-20005-12",
				PhotoURL:        "82s091://",
				PhotoType:       "foo",
			},

			wantOut: ``,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := tt.c.VCF()
			if (err != nil) != tt.wantErr {
				t.Errorf("Contact.VCF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut != tt.wantOut {
				t.Errorf("Contact.VCF() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
