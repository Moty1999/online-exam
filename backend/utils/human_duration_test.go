package utils

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	type args struct {
		d string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr bool
	}{
		{
			name:    "5h20m",
			args:    args{"5h20m"},
			want:    time.Hour*5 + time.Minute*20,
			wantErr: false,
		},
		{
			name:    "1d5h20m",
			args:    args{"1d5h20m"},
			want:    time.Hour*24 + time.Hour*5 + time.Minute*20,
			wantErr: false,
		},
		{
			name: "1d",
			args: args{
				"1d",
			},
			want:    time.Hour * 24,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDuration(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseDuration() got = %v, want %v", got, tt.want)
			}
		})
	}
}
