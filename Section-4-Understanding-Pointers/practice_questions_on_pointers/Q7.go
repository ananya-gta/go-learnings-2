package main


func q7() int {
    x := 1
    f(&x)
    return (x)
}

func f(n *int) {
    n = nil
}
