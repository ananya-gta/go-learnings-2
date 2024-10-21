package main


func q2() int{
    a := 10
    b := &a
    a = 20
    return *b
}

