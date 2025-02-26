package vcf

import (
	"encoding/json"
	"testing"
)

func TestPhoneType_UnmarshalJSON(t *testing.T) {
	type args struct {
		bs []byte
	}

	type testcase struct {
		name    string
		args    args
		want    PhoneType
		wantErr bool
	}

	tests := []testcase{
		{
			name: "cell",
			args: args{
				bs: []byte(`"CELL"`),
			},
			want:    "CELL",
			wantErr: false,
		},
		{
			name: "home",
			args: args{
				bs: []byte(`"HOME"`),
			},
			want:    "HOME",
			wantErr: false,
		},
		{
			name: "work",
			args: args{
				bs: []byte(`"WORK"`),
			},
			want:    "WORK",
			wantErr: false,
		},
		{
			name: "fax",
			args: args{
				bs: []byte(`"FAX"`),
			},
			want:    "FAX",
			wantErr: false,
		},
		{
			name: "pager",
			args: args{
				bs: []byte(`"PAGER"`),
			},
			want:    "PAGER",
			wantErr: false,
		},
		{
			name: "voice",
			args: args{
				bs: []byte(`"VOICE"`),
			},
			want:    "VOICE",
			wantErr: false,
		},
		{
			name: "text",
			args: args{
				bs: []byte(`"TEXT"`),
			},
			want:    "TEXT",
			wantErr: false,
		},
		{
			name: "video",
			args: args{
				bs: []byte(`"VIDEO"`),
			},
			want:    "VIDEO",
			wantErr: false,
		},
		{
			name: "bbs",
			args: args{
				bs: []byte(`"BBS"`),
			},
			want:    "BBS",
			wantErr: false,
		},
		{
			name: "modem",
			args: args{
				bs: []byte(`"MODEM"`),
			},
			want:    "MODEM",
			wantErr: false,
		},
		{
			name: "msg",
			args: args{
				bs: []byte(`"MSG"`),
			},
			want:    "MSG",
			wantErr: false,
		},
		{
			name: "car",
			args: args{
				bs: []byte(`"CAR"`),
			},
			want:    "CAR",
			wantErr: false,
		},
		{
			name: "isdn",
			args: args{
				bs: []byte(`"ISDN"`),
			},
			want:    "ISDN",
			wantErr: false,
		},
		{
			name: "pcs",
			args: args{
				bs: []byte(`"PCS"`),
			},
			want:    "PCS",
			wantErr: false,
		},
		{
			name: "foo",
			args: args{
				bs: []byte(`"FOO"`),
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				got PhoneType
				err error
			)

			if err = json.Unmarshal(tt.args.bs, &got); (err != nil) != tt.wantErr {
				t.Errorf("PhoneType.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("PhoneType.UnmarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
