package vcard

import (
	"encoding/json"
	"testing"
)

func TestDate_UnmarshalJSON(t *testing.T) {
	type args struct {
		bs []byte
	}

	type testcase struct {
		name    string
		args    args
		want    Date
		wantErr bool
	}
	tests := []testcase{
		{
			name: "Valid Date",
			args: args{
				bs: []byte(`"2025-01-01"`),
			},
			want:    "2025-01-01",
			wantErr: false,
		},
		{
			name: "Invalid Date",
			args: args{
				bs: []byte(`"2025-27-963"`),
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
				got Date
				err error
			)

			if err = json.Unmarshal(tt.args.bs, &got); (err != nil) != tt.wantErr {
				t.Fatalf("Date.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Fatalf("Date.UnmarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
