package main

func q1() int{
    x := 5
    y := &x // y stores the address of x 
    *y++    
    return x
}

