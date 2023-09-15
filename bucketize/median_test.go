package bucketize

import "testing"

func TestMedian(t *testing.T) {
	vs := []float64{4, 3, 2, 1}

	{
		got := partition5(vs)
		if want := 1; got != want {
			t.Errorf("partition5(): got %v, want %v", got, want)
		}
	}

	{
		got := partition(vs, 0, 3, 1, 3)
		if want := 1; got != want {
			t.Errorf("partition(): got %v, want %v", got, want)
		}
	}
}
