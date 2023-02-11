package main

import "testing"

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divider  int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "basic scenario",
			args: args{
				dividend: 6,
				divider:  3,
			},
			want: 2,
		},
		{
			name: "divide by 0",
			args: args{
				dividend: 20,
				divider:  0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide(tt.args.dividend, tt.args.divider)
			if (err != nil) != tt.wantErr {
				t.Errorf("divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("divide() got = %v, want %v", got, tt.want)
			}
		})
	}
}
