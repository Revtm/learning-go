package main

import "fmt"

func main(){
  var nama [5]string
  nama[0] = "Budi1"
  nama[1] = "Budi2"
  nama[2] = "Budi3"
  nama[3] = "Budi4"
  nama[4] = "Budi5"

  var umur = [5]uint8 {1,2,3,4,5}
  var sliceUmur = umur[1:3]
  fmt.Println(nama)
  fmt.Println(umur)
  fmt.Println(sliceUmur)
  fmt.Println(len(sliceUmur))
  fmt.Println(cap(sliceUmur))

  sliceUmur[0] = 100
  fmt.Println(nama)
  fmt.Println(umur)
  fmt.Println(sliceUmur)

  // jika mengubah slice, maka akan mengubah array asli
  // []string -> slice
}
