package filter_perf_test

import (
	"fmt"
	"log"
	"testing"

	"robpike.io/filter"
)

func BenchmarkGetIds1(b *testing.B) {
	log.Print("b.N = ", b.N)

	l := createElements(uint(b.N))
	b.ResetTimer()

	ids := getIds1(l)
	if len(ids) != b.N {
		b.Error("fail to get ids")
	}
}

func BenchmarkGetIds2(b *testing.B) {
	log.Print("b.N = ", b.N)

	l := createElements(uint(b.N))
	b.ResetTimer()

	ids := getIds2(l)
	if len(ids) != b.N {
		b.Error("fail to get ids")
	}
}

// Following is unit tests
func TestCreateElements(t *testing.T) {
	l := createElements(10)
	if len(l) != 10 || l[9].Id != 9 {
		t.Error("fail to create elements")
	}
}

func TestGetIds1(t *testing.T) {
	l := createElements(10)
	ids := getIds1(l)
	for i := 0; i < len(l); i++ {
		if ids[i] != uint(i) {
			t.Error("fail to get ids")
			break
		}
	}
}

func TestGetIds2(t *testing.T) {
	l := createElements(100)
	ids := getIds2(l)
	for i := 0; i < len(l); i++ {
		if ids[i] != uint(i) {
			t.Error("fail to get ids")
			break
		}
	}
}

// Following are not testing codes
type element struct {
	Id   uint
	Name string
}

func createElements(size uint) []element {
	result := make([]element, size)

	for i := 0; i < len(result); i++ {
		result[i].Id = uint(i)
		result[i].Name = fmt.Sprintf("name %v", i)
	}

	return result
}

// Use robpick/filter to get id of elements
func getIds1(l []element) []uint {
	return filter.Apply(l, func(e element) uint {
		return e.Id
	}).([]uint)
}

// Use loop to get id of elements
func getIds2(l []element) []uint {
	result := make([]uint, len(l))
	for i, e := range l {
		result[i] = e.Id
	}
	return result
}
