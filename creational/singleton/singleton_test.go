package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	resetSingleton()
	i1 := GetInstance()
	i2 := GetInstance()

	if i1 != i2 {
		t.Fatalf("Different addresses ( %p != %p )", i1, i2)
	}
}

func TestGetInstanceWithIf(t *testing.T) {
	resetSingleton()
	i1 := GetInstanceWithIf()
	i2 := GetInstanceWithIf()

	if i1 != i2 {
		t.Fatalf("Different addresses ( %p != %p )", i1, i2)
	}
}

func BenchmarkIsolatedGetInstance(b *testing.B) {
	var is *singleton
	resetSingleton()
	for i := 0; i < b.N; i++ {
		is = GetInstance()
	}
	b.Log(is)

}

func BenchmarkIsolatedGetInstanceWithIf(b *testing.B) {
	var is *singleton
	resetSingleton()
	for i := 0; i < b.N; i++ {
		is = GetInstanceWithIf()
	}
	b.Log(is)
}
