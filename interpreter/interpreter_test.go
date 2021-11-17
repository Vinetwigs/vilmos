package interpreter

import (
	"image"
	"os"
	"reflect"
	"testing"
)

const (
	IMG_WIDTH  = 1000
	IMG_HEIGHT = 800
)

var (
	filledStack = &Stack{
		items:   []int32{10, 20, 30, 40, 50},
		maxSize: 20,
	}
	stack, _   = NewStack(-1)
	imgFile, _ = os.Open("..\\examples\\tests\\load_image.png")
	img, _, _  = image.Decode(imgFile)
	r, g, b, _ = img.At(0, 0).RGBA()
)

func TestNewInterpreter(t *testing.T) {
	type args struct {
		debug           bool
		maxSize         int
		instructionSize int
	}
	tests := []struct {
		name string
		args args
		want *Interpreter
	}{
		{
			name: "New Interpreter test 1",
			args: args{
				debug:           false,
				maxSize:         -1,
				instructionSize: 0,
			},
			want: &Interpreter{
				image: nil,
				stack: stack,
				pc: image.Point{
					X: 0,
					Y: 0,
				},
				width:           0,
				height:          0,
				isDebug:         false,
				instructionSize: 0,
				openedFile:      nil,
			},
		},
		{
			name: "New Interpreter test 2",
			args: args{
				debug:           false,
				maxSize:         -1,
				instructionSize: 0,
			},
			want: &Interpreter{
				image: nil,
				stack: stack,
				pc: image.Point{
					X: 0,
					Y: 0,
				},
				width:           0,
				height:          0,
				isDebug:         false,
				instructionSize: 0,
				openedFile:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInterpreter(tt.args.debug, tt.args.maxSize, tt.args.instructionSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInterpreter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterpreter_LoadImage(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		i       *Interpreter
		args    args
		wantErr bool
	}{
		{
			name: "Load without errors",
			i: &Interpreter{
				image: nil,
				stack: stack,
				pc: image.Point{
					X: 0,
					Y: 0,
				},
				width:           0,
				height:          0,
				isDebug:         false,
				instructionSize: 0,
				openedFile:      nil,
			},
			args: args{
				path: "..\\examples\\tests\\load_image.png",
			},
			wantErr: false,
		},
		{
			name: "Load with open error",
			i: &Interpreter{
				image: nil,
				stack: stack,
				pc: image.Point{
					X: 0,
					Y: 0,
				},
				width:           0,
				height:          0,
				isDebug:         false,
				instructionSize: 0,
				openedFile:      nil,
			},
			args: args{
				path: "..\\examples\\test\\load_image.png",
			},
			wantErr: true,
		},
		{
			name: "Load with extension error",
			i: &Interpreter{
				image: nil,
				stack: stack,
				pc: image.Point{
					X: 0,
					Y: 0,
				},
				width:           0,
				height:          0,
				isDebug:         false,
				instructionSize: 0,
				openedFile:      nil,
			},
			args: args{
				path: "..\\examples\\tests\\load_image.txt",
			},
			wantErr: true,
		},
		{
			name: "Load with decode error",
			i: &Interpreter{
				image: nil,
				stack: stack,
				pc: image.Point{
					X: 0,
					Y: 0,
				},
				width:           0,
				height:          0,
				isDebug:         false,
				instructionSize: 0,
				openedFile:      nil,
			},
			args: args{
				path: "..\\examples\\tests\\load_image_invalid.png",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.LoadImage(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Interpreter.LoadImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInterpreter_Run(t *testing.T) {
	tests := []struct {
		name    string
		i       *Interpreter
		wantErr bool
	}{
		{
			name: "Running without debugging",
			i: &Interpreter{
				image: img,
				stack: stack,
				pc: image.Point{
					X: 0,
					Y: 0,
				},
				width:           IMG_WIDTH,
				height:          IMG_HEIGHT,
				isDebug:         false,
				instructionSize: 200,
				openedFile:      nil,
			},
			wantErr: true,
		},
		{
			name: "Running with debugger",
			i: &Interpreter{
				image: img,
				stack: stack,
				pc: image.Point{
					X: 0,
					Y: 0,
				},
				width:           IMG_WIDTH,
				height:          IMG_HEIGHT,
				isDebug:         true,
				instructionSize: 200,
				openedFile:      nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Interpreter.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInterpreter_Step(t *testing.T) {
	tests := []struct {
		name  string
		i     *Interpreter
		want  bool
		want1 string
	}{
		{
			name: "Single step test 1",
			i: &Interpreter{
				image: img,
				stack: stack,
				pc: image.Point{
					X: 0,
					Y: 0,
				},
				width:           IMG_WIDTH,
				height:          IMG_HEIGHT,
				isDebug:         false,
				instructionSize: 200,
				openedFile:      &os.File{},
			},
			want:  true,
			want1: "Pushed 0 into the stack",
		},
		{
			name: "Single step test 2",
			i: &Interpreter{
				image: img,
				stack: stack,
				pc: image.Point{
					X: 200,
					Y: 0,
				},
				width:           IMG_WIDTH - 1,
				height:          IMG_HEIGHT - 1,
				isDebug:         true,
				instructionSize: 200,
				openedFile:      &os.File{},
			},
			want:  true,
			want1: "Pushed 108 into the stack",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.i.Step()
			if got != tt.want {
				t.Errorf("Interpreter.Step() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Interpreter.Step() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInterpreter_readPixel(t *testing.T) {
	tests := []struct {
		name string
		i    *Interpreter
		want *Pixel
	}{
		{
			name: "Read pixel test",
			i: &Interpreter{
				image: img,
				stack: stack,
				pc: image.Point{
					X: 0,
					Y: 0,
				},
				width:           IMG_WIDTH,
				height:          IMG_HEIGHT,
				isDebug:         false,
				instructionSize: 0,
				openedFile:      &os.File{},
			},
			want: &Pixel{
				R: 0,
				G: 0,
				B: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.readPixel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interpreter.readPixel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rgbaToPixel(t *testing.T) {
	type args struct {
		r   uint32
		g   uint32
		b   uint32
		in3 uint32
	}
	tests := []struct {
		name string
		args args
		want *Pixel
	}{
		{
			name: "rgb to Pixel structure test",
			args: args{
				r:   r,
				g:   g,
				b:   b,
				in3: 0,
			},
			want: &Pixel{
				R: 0,
				G: 0,
				B: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rgbaToPixel(tt.args.r, tt.args.g, tt.args.b, tt.args.in3); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rgbaToPixel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_popOrErr(t *testing.T) {
	type args struct {
		i *Interpreter
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Stack pop test",
			args: args{
				i: &Interpreter{
					image:           nil,
					stack:           filledStack,
					pc:              image.Point{},
					width:           IMG_WIDTH,
					height:          IMG_HEIGHT,
					isDebug:         false,
					instructionSize: 0,
					openedFile:      &os.File{},
				},
			},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := popOrErr(tt.args.i); got != tt.want {
				t.Errorf("popOrErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pushOrErr(t *testing.T) {
	type args struct {
		i   *Interpreter
		val int32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Push stack test",
			args: args{
				i: &Interpreter{
					image: img,
					stack: filledStack,
					pc: image.Point{
						X: 0,
						Y: 0,
					},
					width:           IMG_WIDTH,
					height:          IMG_HEIGHT,
					isDebug:         false,
					instructionSize: 0,
					openedFile:      &os.File{},
				},
				val: 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pushOrErr(tt.args.i, tt.args.val)
		})
	}
}

func Test_processPixel(t *testing.T) {
	type args struct {
		pixel *Pixel
		i     *Interpreter
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "OUTPUT_INT",
			args: args{
				pixel: &Pixel{R: 0, G: 0, B: 1},
				i: &Interpreter{
					image: img,
					stack: filledStack,
					pc: image.Point{
						X: 0,
						Y: 0,
					},
					width:           IMG_WIDTH,
					height:          IMG_HEIGHT,
					isDebug:         true,
					instructionSize: 200,
					openedFile:      nil,
				},
			},
			want: "Popped 20 from the stack and printed it in the console",
		},
		{
			name: "SUM",
			args: args{
				pixel: &Pixel{R: 0, G: 206, B: 209},
				i: &Interpreter{
					image: img,
					stack: filledStack,
					pc: image.Point{
						X: 0,
						Y: 0,
					},
					width:           IMG_WIDTH,
					height:          IMG_HEIGHT,
					isDebug:         true,
					instructionSize: 200,
					openedFile:      nil,
				},
			},
			want: "Popped 40, popped 30 and then pushed into the stack their sum (70)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processPixel(tt.args.pixel, tt.args.i); got != tt.want {
				t.Errorf("processPixel() = %v, want %v", got, tt.want)
			}
		})
	}
}
