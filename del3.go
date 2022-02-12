package del3

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestXxx(*testing.T)
func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
func BenchmarkXxx(*testing.B)
func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}
BenchmarkRandInt-8   	68453040	        17.8 ns/op
func BenchmarkBigLen(b *testing.B) {
	big := NewBig()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Len()
	}
}
func BenchmarkTemplateParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}
func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}

func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}
func ExamplePerm() {
	for _, value := range Perm(5) {
		fmt.Println(value)
	}
	// Unordered output: 4
	// 2
	// 1
	// 3
	// 0
}
func Example() { ... }
func ExampleF() { ... }
func ExampleT() { ... }
func ExampleT_M() { ... }

func Example_suffix() { ... }
func ExampleF_suffix() { ... }
func ExampleT_suffix() { ... }
func ExampleT_M_suffix() { ... }

func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	...
}
func TestFoo(t *testing.T) {
	// <setup code>
	t.Run("A=1", func(t *testing.T) { ... })
	t.Run("A=2", func(t *testing.T) { ... })
	t.Run("B=1", func(t *testing.T) { ... })
	// <tear-down code>
}

func TestGroupedParallel(t *testing.T) {
	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			...
		})
	}
}

func TestTeardownParallel(t *testing.T) {
	// This Run will not return until the parallel tests finish.
	t.Run("group", func(t *testing.T) {
		t.Run("Test1", parallelTest1)
		t.Run("Test2", parallelTest2)
		t.Run("Test3", parallelTest3)
	})
	// <tear-down code>
}

func TestMain(m *testing.M)
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}