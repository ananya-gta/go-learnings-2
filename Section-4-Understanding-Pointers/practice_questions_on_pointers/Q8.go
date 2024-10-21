package main


func q8() (int, int){
    a := 1
    b := 2
    update(&a, b)
    return a, b
}

func update(x *int, y int) {
    *x += y
}
