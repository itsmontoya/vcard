package vcard

import (
	"encoding/json"
	"testing"
)

func TestURL_UnmarshalJSON(t *testing.T) {
	type args struct {
		bs []byte
	}

	type testcase struct {
		name    string
		args    args
		want    URL
		wantErr bool
	}
	tests := []testcase{
		{
			name: "Valid url",
			args: args{
				bs: []byte(`"https://google.com"`),
			},
			want:    "https://google.com",
			wantErr: false,
		},
		{
			name: "Invalid URL",
			args: args{
				bs: []byte{'"', 0x7f, '"'},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "wrong input type",
			args: args{
				bs: []byte(`[]`),
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				got URL
				err error
			)

			if err = json.Unmarshal(tt.args.bs, &got); (err != nil) != tt.wantErr {
				t.Fatalf("URL.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Fatalf("URL.UnmarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
