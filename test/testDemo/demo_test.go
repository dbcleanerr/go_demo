package testDemo

import "testing"

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	want := 1 + 2

	if got != want {
		t.Errorf("expcted:%#v, got:%#v", want, got)
	}
}

func TestSquare(t *testing.T) {
	// 用 run 来执行子测试
	t.Run("case1", func(t *testing.T) {
		got := Square(5)
		want := 5 * 5

		if got != want {
			t.Errorf("expcted:%#v, got:%#v", want, got)
		}
	})

	t.Run("case2", func(t *testing.T) {
		got := Square(10)
		want := 10 * 10

		if got != want {
			t.Errorf("expcted:%#v, got:%#v", want, got)
		}
	})
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Square(10)
	}
}

// 性能比较测试

func benchmarkSum(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Sum(1, n)
	}
}

func BenchmarkSum100(b *testing.B) {
	benchmarkSum(b, 100)
}

func BenchmarkSum1000(b *testing.B) {
	benchmarkSum(b, 1000)
}

func BenchmarkSum10000(b *testing.B) {
	benchmarkSum(b, 10000)
}

func BenchmarkSum100000(b *testing.B) {
	benchmarkSum(b, 100000)
}

func BenchmarkSum1000000(b *testing.B) {
	benchmarkSum(b, 1000000)
}
