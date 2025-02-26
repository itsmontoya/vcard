package vcf

import (
	"encoding/json"
	"testing"
)

func TestPhotoType_UnmarshalJSON(t *testing.T) {
	type args struct {
		bs []byte
	}

	type testcase struct {
		name    string
		args    args
		want    PhotoType
		wantErr bool
	}
	tests := []testcase{
		{
			name: "jpeg",
			args: args{
				bs: []byte(`"JPEG"`),
			},
			want:    "JPEG",
			wantErr: false,
		},
		{
			name: "png",
			args: args{
				bs: []byte(`"PNG"`),
			},
			want:    "PNG",
			wantErr: false,
		},
		{
			name: "gif",
			args: args{
				bs: []byte(`"GIF"`),
			},
			want:    "GIF",
			wantErr: false,
		},
		{
			name: "bmp",
			args: args{
				bs: []byte(`"BMP"`),
			},
			want:    "BMP",
			wantErr: false,
		},
		{
			name: "svg",
			args: args{
				bs: []byte(`"SVG"`),
			},
			want:    "SVG",
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
				got PhotoType
				err error
			)

			if err = json.Unmarshal(tt.args.bs, &got); (err != nil) != tt.wantErr {
				t.Errorf("PhotoType.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("PhotoType.UnmarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
