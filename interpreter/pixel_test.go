package interpreter

import (
	"testing"
)

func TestPixel_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Pixel
		want string
	}{
		{
			name: "Pixel string test 1",
			p: &Pixel{
				R: 0,
				G: 0,
				B: 0,
			},
			want: "{R: 0 G: 0 B: 0}",
		},
		{
			name: "Pixel string test 2",
			p: &Pixel{
				R: 200,
				G: 103,
				B: 40,
			},
			want: "{R: 200 G: 103 B: 40}",
		},
		{
			name: "Pixel string test 3",
			p: &Pixel{
				R: 255,
				G: 255,
				B: 255,
			},
			want: "{R: 255 G: 255 B: 255}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Pixel.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPixel_Equals(t *testing.T) {
	type args struct {
		other Pixel
	}
	tests := []struct {
		name string
		p    *Pixel
		args args
		want bool
	}{
		{
			name: "Equals pixels",
			p: &Pixel{
				R: 10,
				G: 20,
				B: 30,
			},
			args: args{
				other: Pixel{
					R: 10,
					G: 20,
					B: 30,
				},
			},
			want: true,
		},
		{
			name: "Different pixels",
			p: &Pixel{
				R: 10,
				G: 20,
				B: 30,
			},
			args: args{
				other: Pixel{
					R: 10,
					G: 30,
					B: 20,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Equals(tt.args.other); got != tt.want {
				t.Errorf("Pixel.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
