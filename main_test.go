package main

import "testing"

func Test_getSNIServerName(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{buf: []uint8{22, 3, 1, 1, 2, 1, 0, 0, 254, 3, 3, 247, 68, 239, 62, 76, 12, 172, 30, 61, 229, 84, 250, 148, 64, 71, 81, 45, 164, 200, 3, 109, 97, 229, 137, 70, 33, 240, 4, 89, 224, 176, 190, 0, 0, 110, 192, 47, 192, 43, 192, 48, 192, 44, 0, 158, 192, 39, 0, 103, 192, 40, 0, 107, 192, 36, 192, 20, 192, 10, 0, 165, 0, 163, 0, 161, 0, 159, 0, 106, 0, 105, 0, 104, 0, 57, 0, 56, 0, 55, 0, 54, 192, 50, 192, 46, 192, 42, 192, 38, 192, 15, 192, 5, 0, 157, 0, 61, 0, 53, 192, 35, 192, 19, 192, 9, 0, 164, 0, 162, 0, 160, 0, 64, 0, 63, 0, 62, 0, 51, 0, 50, 0, 49, 0, 48, 192, 49, 192, 45, 192, 41, 192, 37, 192, 14, 192, 4, 0, 156, 0, 60, 0, 47, 0, 255, 1, 0, 0, 103, 0, 0, 0, 15, 0, 13, 0, 0, 10, 103, 111, 111, 103, 108, 101, 46, 99, 111, 109, 0, 11, 0, 4, 3, 0, 1, 2, 0, 10, 0, 28, 0, 26, 0, 23, 0, 25, 0, 28, 0, 27, 0, 24, 0, 26, 0, 22, 0, 14, 0, 13, 0, 11, 0, 12, 0, 9, 0, 10, 0, 35, 0, 0, 0, 13, 0, 32, 0, 30, 6, 1, 6, 2, 6, 3, 5, 1, 5, 2, 5, 3, 4, 1, 4, 2, 4, 3, 3, 1, 3, 2, 3, 3, 2, 1, 2, 2, 2, 3, 51, 116, 0, 0,}},
			name:"google",
			want:"google.com",
		},
		{
			args: args{buf: []uint8{22,3,1,1,1,1,1,1,1,1,1,1}},
			name:"empty",
			want:"",
		},
		{
			args: args{buf: []uint8{22,3,1,1,1}},
			name:"no panic",
			want:"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSNIServerName(tt.args.buf); got != tt.want {
				t.Errorf("getSNIServerName() = %v, want %v", got, tt.want)
			}
		})
	}
}
