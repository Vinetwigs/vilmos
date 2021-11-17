package interpreter

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	type args struct {
		maxSize int
	}
	tests := []struct {
		name    string
		args    args
		want    *Stack
		wantErr bool
	}{
		{
			name: "Without max size",
			args: args{
				maxSize: -1,
			},
			want: &Stack{
				maxSize: -1,
			},
			wantErr: false,
		},
		{
			name: "With max size",
			args: args{
				maxSize: 10,
			},
			want: &Stack{
				maxSize: 10,
			},
			wantErr: false,
		},
		{
			name: "With invalid max size",
			args: args{
				maxSize: -2,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStack(tt.args.maxSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type args struct {
		val int32
	}
	tests := []struct {
		name    string
		stack   *Stack
		args    args
		wantErr bool
	}{
		{
			name: "No max size",
			stack: &Stack{
				items:   []int32{},
				maxSize: -1,
			},
			args: args{
				val: 10,
			},
			wantErr: false,
		},
		{
			name: "With max size",
			stack: &Stack{
				items:   []int32{10, 20, 30},
				maxSize: 5,
			},
			args: args{
				val: 40,
			},
			wantErr: false,
		},
		{
			name: "With error",
			stack: &Stack{
				items:   []int32{10, 20},
				maxSize: 2,
			},
			args: args{
				val: 30,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.stack.Push(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Stack.Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name  string
		stack *Stack
		want  int32
	}{
		{
			name: "Empty stack",
			stack: &Stack{
				items: []int32{},
			},
			want: 0,
		},
		{
			name: "Stack not empty",
			stack: &Stack{
				items:   []int32{10, -3, 0, 4},
				maxSize: 0,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.Peek(); got != tt.want {
				t.Errorf("Stack.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name    string
		stack   *Stack
		want    int32
		wantErr bool
	}{
		{
			name: "Empty stack",
			stack: &Stack{
				items:   []int32{},
				maxSize: 0,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Stack with one element",
			stack: &Stack{
				items: []int32{10},
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "Stack with multiple elements",
			stack: &Stack{
				items: []int32{10, 20},
			},
			want:    20,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.stack.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Size(t *testing.T) {
	tests := []struct {
		name  string
		stack *Stack
		want  int
	}{
		{
			name: "Empty stack",
			stack: &Stack{
				items: []int32{},
			},
			want: 0,
		},
		{
			name: "Stack with elements",
			stack: &Stack{
				items:   []int32{10, 20, 30},
				maxSize: 0,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.Size(); got != tt.want {
				t.Errorf("Stack.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		stack *Stack
		want  bool
	}{
		{
			name: "Empty stack",
			stack: &Stack{
				items:   []int32{},
				maxSize: 0,
			},
			want: true,
		},
		{
			name: "Not empty stack",
			stack: &Stack{
				items:   []int32{10},
				maxSize: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.IsEmpty(); got != tt.want {
				t.Errorf("Stack.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Cycle(t *testing.T) {
	tests := []struct {
		name  string
		stack *Stack
		want  *Stack
	}{
		{
			name: "Cycle test 1",
			stack: &Stack{
				items: []int32{10, 20, 30, 40, 50},
			},
			want: &Stack{
				items: []int32{50, 10, 20, 30, 40},
			},
		},
		{
			name: "Cycle test 2",
			stack: &Stack{
				items: []int32{50, 40, 30, 20, 10},
			},
			want: &Stack{
				items: []int32{10, 50, 40, 30, 20},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.Cycle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Cycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_RCycle(t *testing.T) {
	tests := []struct {
		name  string
		stack *Stack
		want  *Stack
	}{
		{
			name: "RCycle test 1",
			stack: &Stack{
				items: []int32{10, 20, 30, 40, 50},
			},
			want: &Stack{
				items: []int32{20, 30, 40, 50, 10},
			},
		},
		{
			name: "RCycle test 2",
			stack: &Stack{
				items: []int32{50, 40, 30, 20, 10},
			},
			want: &Stack{
				items: []int32{40, 30, 20, 10, 50},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.RCycle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.RCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Reverse(t *testing.T) {
	tests := []struct {
		name  string
		stack *Stack
		want  *Stack
	}{
		{
			name: "Reverse test 1",
			stack: &Stack{
				items: []int32{10},
			},
			want: &Stack{
				items: []int32{10},
			},
		},
		{
			name: "Reverse test 2",
			stack: &Stack{
				items:   []int32{10, 20, 30, 40, 50},
				maxSize: 0,
			},
			want: &Stack{
				items:   []int32{50, 40, 30, 20, 10},
				maxSize: 0,
			},
		},
		{
			name: "Reverse stack empty",
			stack: &Stack{
				items: []int32{},
			},
			want: &Stack{
				items: []int32{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.Reverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_GetItemAt(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		stack   *Stack
		args    args
		want    int32
		wantErr bool
	}{
		{
			name: "Get item in stack",
			stack: &Stack{
				items: []int32{10, 20, 30, 40, 50},
			},
			args: args{
				i: 3,
			},
			want:    40,
			wantErr: false,
		},
		{
			name: "Negative argument",
			stack: &Stack{
				items: []int32{10, 20, 30, 40, 50},
			},
			args: args{
				i: -1,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Exceeded index argument",
			stack: &Stack{
				items: []int32{10, 20, 30, 40, 50},
			},
			args: args{
				i: 5,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.stack.GetItemAt(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.GetItemAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Stack.GetItemAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Clear(t *testing.T) {
	tests := []struct {
		name  string
		stack *Stack
		want  *Stack
	}{
		{
			name: "Stack with elements",
			stack: &Stack{
				items: []int32{10, 20, 30},
			},
			want: &Stack{
				items: nil,
			},
		},
		{
			name: "Stack without elements",
			stack: &Stack{
				items: []int32{},
			},
			want: &Stack{
				items: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.Clear(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Clear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Output(t *testing.T) {
	tests := []struct {
		name  string
		stack *Stack
		want  bool
	}{
		{
			name: "Output test",
			stack: &Stack{
				items: []int32{10, 20},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.Output(); got != tt.want {
				t.Errorf("Stack.Output() = %v, want %v", got, tt.want)
			}
		})
	}
}
