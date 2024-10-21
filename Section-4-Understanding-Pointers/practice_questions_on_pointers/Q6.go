package main


func q6() int{
    var p *int // declares a pointer p but does not initialize it.
    p = new(int) // allocates memory for an integer and assigns the address to p
    *p = 10 // sets the value at that memory location to 10
    return (*p) // prints the value pointed to by p, which is 10.
}

