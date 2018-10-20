package starterlib
import "github.com/go-modules-by-example-staging/starterlib/hello"
import "fmt"
func WiseHello() {
        hello.Hello("World")
        fmt.Println("Don't communicate by sharing memory, share memory by communicating.")
}
