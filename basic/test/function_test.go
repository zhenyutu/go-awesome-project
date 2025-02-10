package mytestcase

import (
	"fmt"
	"testing"
)

func TestTriangle(t *testing.T) {
	tests := []struct {
		a int
		b int
		c int
	}{
		{3, 4, 5.0},
		{5, 12, 13.0},
	}

	for _, test := range tests {
		var actual = calcTriangle(test.a, test.b)
		if actual != test.c {
			t.Errorf("calcTriangle(%d, %d) = %d, want %d", test.a, test.b, actual, test.c)
		}
	}
}
func TestLoop(t *testing.T) {
	t.Log("Loop:", Loop(uint64(32)))
}

func TestFactorial(t *testing.T) {
	t.Log("Factorial:", Factorial(uint64(32)))
}

func BenchmarkLoop(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Loop(uint64(40))
	}
}

func BenchmarkFactorial(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Factorial(uint64(40))
	}
}

func ExampleLoop() {
	res := Loop(uint64(32))
	fmt.Println(res)

	//Output:
	//12400865694432886784
}
