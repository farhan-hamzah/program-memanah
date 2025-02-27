package main
import "fmt"

func main(){
    var p1, p2, p3, h1, h2, h3, terbesar int
    fmt.Scan(&p1, &p2, &p3)
    h1 = p1
    h2 = p2
    h3 = p3
    for p1 != 0 && p2 != 0 && p3 != 0{
        fmt.Scan(&p1, &p2, &p3)
        h1 = h1 + p1
        h2 = h2 + p2
        h3 = h3 + p3
    }
    if h1 > h2 && h1 > h3{
        terbesar = h1
    }else if h2 > h1 && h2 > h3{
        terbesar = h2
    }else{
        terbesar = h3
    }
    fmt.Print(h1, h2, h3, terbesar)
}