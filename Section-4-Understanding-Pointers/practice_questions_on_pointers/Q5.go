package main


func q5() (int, int){
    a := 1
    b := 2
	swap(&a, &b)
	return a, b
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
