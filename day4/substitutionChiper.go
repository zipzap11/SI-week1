package main

import "fmt"

 

type student struct {

   name       string

   nameEncode string

   score      int

}

 

type Chiper interface {

   Encode() string

   Decode() string

}

// a b c d e f g h i j  k  l  m  n  o  p  q  r  s  t  u  v  w  x  y  z
// 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26
// rizky irapb
func (s *student) Encode() string {
   var nameEncode = ""
   // your code here
   for _ , ch := range s.name {
		asc := ch - 97
		nameEncode += string(122 - asc)
   }
   return nameEncode

}

 

func (s *student) Decode() string {

   var nameDecode = ""
	for _, ch := range s.nameEncode {
		asc := 122 - ch
		nameDecode += string(asc + 97)
	}
   // your code here
   return nameDecode

}

 

func main() {

   var menu int

   var a = student{}

   var c Chiper = &a

   fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")

   fmt.Scan(&menu)

   if menu == 1 {

       fmt.Print("\nInput Student’s Name : ")

       fmt.Scan(&a.name)

       fmt.Print("\nEncode of Student’s Name " + a.name + " is : " + c.Encode())

   } else if menu == 2 {

       fmt.Print("\nInput Student’s Encode Name : ")

       fmt.Scan(&a.nameEncode)

       fmt.Print("\nDecode of Student’s Name " + a.nameEncode + " is : " + c.Decode())

   } else {

       fmt.Println("Wrong input name menu")

   }

}