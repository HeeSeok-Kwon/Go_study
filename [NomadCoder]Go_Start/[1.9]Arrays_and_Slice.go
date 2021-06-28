package main 

import "fmt"

func main() {
  // array
  // names := [5] string {"nico", "lynn", "dal"}
  // names[3] = "alala"
  // names[4] = "alala"
  // names[5] = "alala" // index error

  // slice
  names := [] string {"nico", "lynn", "dal"}
  // names[3] = "alala" // index error
  names = append(names, "flynn")
  // append는 names를 수정하지 않는다.
  // 새로운 값이 추가된 slice를 리턴한다.
  fmt.Println(names)
}
