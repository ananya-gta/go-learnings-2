package main


func q4() int{
    x := 5
    y := &x
    z := *y
    z++
    return x
}

