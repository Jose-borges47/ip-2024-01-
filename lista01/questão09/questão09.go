package main
import "fmt"

func main() {
    var a, b, c, delta int
  fmt.Println("digite A, B e C")
  fmt.Scan(&a, &b, &c)
  delta = (b*b - 4 * a * c)
  fmt.Print("O VALOR DE DELTA É =", delta)
}