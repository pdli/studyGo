package main
import ("fmt")

import ("calculation")

func main() {
    fmt.Println("Hello World! This is my first Go Program\n")

    var x int
    x = 10

    fmt.Println("x xyz , x", x, "xzy")

    a := [5]  int16 {1, 2, 3, 4, 5}

    slice_a := a[1:3]

    slice_a[1] = 100

    slice_b := [] int16 {111, 222}

//    #    slice_a = append(slice_a, 12, 13)

    slice_a = append(slice_a, slice_b...)

    fmt.Println("a is :", a, "slice_a is :", slice_a)

    fmt.Println("slice_b is: ", slice_b)

    sum := calculation.Do_add(slice_a[0], slice_b[0])
    fmt.Println("Sum is: ", sum)

}

