package vcard

import (
	"testing"
)

func Test_getOutput(t *testing.T) {
	type args struct {
		v any
	}

	type testcase struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}

	tests := []testcase{
		{
			name: "error",
			args: args{
				v: "hello",
			},
			wantOut: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := getOutput(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("getOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut != tt.wantOut {
				t.Errorf("getOutput() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_getParsed(t *testing.T) {
	type testcase struct {
		name    string
		tmpl    string
		wantErr bool
	}

	tests := []testcase{
		{
			name:    "error",
			tmpl:    `foo {{*if boo}}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getParsed(tt.tmpl)
			if (err != nil) != tt.wantErr {
				t.Errorf("getParsed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
