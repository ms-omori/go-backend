package pkg

import (
	"testing"
)

// TestFibonacci is a test function for the Fibonacci function.
// 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 ...
func TestFibonacci(t *testing.T) {

	tests := []struct {
		arg  int
		want int
	}{
		{
			arg:  0,
			want: 0,
		},
		{
			arg:  1,
			want: 1,
		},
		{
			arg:  2,
			want: 1,
		},
		{
			arg:  3,
			want: 2,
		},
		{
			arg:  4,
			want: 3,
		},
		{
			arg:  5,
			want: 5,
		},
		{
			arg:  20,
			want: 6765,
		},
		{
			arg:  50,
			want: 12586269025,
		},
	}

	for _, tt := range tests {
		t.Run("test fibonacci", func(t *testing.T) {
			if got := Fibonacci(tt.arg); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}

}

// ベンチマークテスト
func BenchmarkFibonacci(b *testing.B) {
	for n := 0; n < 40; n++ {
		Fibonacci(n)
	}
}
