package mytestcase

import "testing"

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
