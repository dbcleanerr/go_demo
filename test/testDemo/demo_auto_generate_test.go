package testDemo

import "testing"

// TestAdd2 得用 goland 自动生成的
func TestAdd2(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				a: 11,
				b: 22,
			},
			want: 33,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add2(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare2(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Square2(tt.args.a); got != tt.want {
				t.Errorf("Square2() = %v, want %v", got, tt.want)
			}
		})
	}
}
