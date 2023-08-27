package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("before test")

	m.Run()

	fmt.Println("after test")
}

func BenchmarkSayHelloAziz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Aziz")
	}
}

func BenchmarkSayHelloSubBenchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Aziz")
	}

	b.Run("Aziz", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Aziz")
		}
	})

	b.Run("Albasyir", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Albasyir")
		}
	})
}

func BenchmarkTableImplementation(b *testing.B) {
	tests := []struct {
		x int
		y int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
	}

	for _, test := range tests {
		b.Run(fmt.Sprintf("%d x %d = %d", test.x, test.y, test.x*test.y), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				DimentionCalculation(test.x, test.y)
			}
		})
	}
}

func BenchmarkSayHelloMoreData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("ThisWithMoreDataSoThisWillBeHeavyThanAzizString")
	}
}

func TestSayHello(t *testing.T) {
	assert.Equal(t, "Hello, Aziz!", SayHello("Aziz"), "Result must be 'Hello, Aziz!'")
}

func TestBrokenSayHello(t *testing.T) {
	t.Skip("This test is broken")
	assert.Equal(t, "Hello, Aziz !", SayHello("Aziz"), "This expected error")
	require.Equal(t, "Hello, Aziz !", SayHello("Aziz"), "This expected error")
}

func TestSkipUTest(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can't run on mac")
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Aziz", func(t *testing.T) {
		assert.Equal(t, "Hello, Aziz!", SayHello("Aziz"), "Result must be 'Hello, Aziz!'")
	})

	t.Run("Albasyir", func(t *testing.T) {
		assert.Equal(t, "Hello, Albasyir!", SayHello("Albasyir"), "Result must be 'Hello, Aziz!'")
	})
}

func TestTableTest(t *testing.T) {
	tests := []struct {
		x int
		y int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d x %d = %d", test.x, test.y, test.x*test.y), func(t *testing.T) {
			assert.Equal(t, test.x*test.y, DimentionCalculation(test.x, test.y))
		})
	}
}
