package starterlib
import "github.com/go-modules-by-example-staging/starterlib/v2/hello"
import "fmt"
func WiseHello(who string) {
        hello.Hello(who)
        fmt.Println("Don't communicate by sharing memory, share memory by communicating.")
}
