// main.go
package main
import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Fprintf(os.Stderr, "Usage: %s <version1> <version2>\n", os.Args[0])
        os.Exit(1)
    }

    v1 := os.Args[1]
    v2 := os.Args[2]

    res, err := VersionCompare(v1, v2)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }

    switch res {
    case RESULT_EQ:
        fmt.Printf("%s == %s\n", v1, v2)
    case RESULT_LT:
        fmt.Printf("%s < %s\n", v1, v2)
    case RESULT_GT:
        fmt.Printf("%s > %s\n", v1, v2)
    }
} 