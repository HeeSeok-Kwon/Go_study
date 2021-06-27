package main

import "fmt"

func main() {
// 값 복사
// a := 2
// b := a
// a = 10
// fmt.Println(a, b)

// 주소 확인
// a := 2
// b := 5
// fmt.Println(&a, &b)

// Pointer 변수 생성
// a := 2
// b := &a
// fmt.Println(&a, b)
// fmt.Println(a, b)
// fmt.Println(b)
// fmt.Println(*b)
// a = 5
// fmt.Println(*b)

// Pointer를 통한 접근
// a := 2
// b := &a
// *b = 20
// fmt.Println(a)

// 정리
a := 2
b := &a
*b = 202020
fmt.Println(a)
}
