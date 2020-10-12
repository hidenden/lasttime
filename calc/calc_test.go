package calc

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		t.Fatal("sum(1,2) should be 3. But doesn't match")
	}
}

func ExampleHello() {
	fmt.Println("Hello")
	// Output: Hello
}

func ExampleSum() {
	i := sum(10, 1)
	fmt.Println(i)
	// Output: 11
}

func ExampleShufullWillBeFailed() {
	x := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range x {
		fmt.Printf("k=%s v=%d\n", k, v)
	}
	// Unordered output:
	// k=a v=1
	// k=c v=3
	// k=b v=2
}
