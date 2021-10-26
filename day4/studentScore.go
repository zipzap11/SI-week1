package main

import (
	"fmt"
	"strconv"
)

 

type Student struct {

 name  []string

 score []int

}

 

func (s Student) Avarage() float64 {
	var total float64
	for _, score := range s.score {
		total += float64(score)
	}
	return total / float64(len(s.score))
}

 

func (s Student) Min() (min int, name string) { 
	// with map implementation
	tmp := make(map[string]int)
	name = s.name[0]
	min = s.score[0]  
	for i := 0; i < len(s.name); i++ {
		tmp[s.name[i]] = s.score[i]
		if s.score[i] < min {
			name = s.name[i]
			min = s.score[i]
		}
	}
	min = tmp[name]
	return

	// without map
	// minName := s.name[0] 
	// minScore := s.score[0]
	// for i := 0; i < len(s.name); i++ {
	// 	if s.score[i] < minScore {
	// 		minScore = s.score[i]
	// 		minName = s.name[i]
	// 	} 
	// }
	// return minScore, minName 
}

 

func (s Student) Max() (max int, name string) {
	// with map
	tmp := make(map[string]int)
	name = s.name[0]
	max = s.score[0]  
	for i := 0; i < len(s.name); i++ {
		tmp[s.name[i]] = s.score[i]
		if s.score[i] > max {
			name = s.name[i]
			max = s.score[i]
		}
	}
	return


	// without map
	// name = s.name[0]
	// score = s.score[0]
	// for i := 0; i < len(s.name); i++ {
	// 	if s.score[i] > maxScore {
	// 		score = s.score[i]
	// 		name = s.name[i]
	// 	}
	// }
	// return 
}

 

func main() {

 var a = Student{}

 

 for i := 1; i < 6; i++ {

   var name string

   fmt.Print("Input " + strconv.Itoa(i) + " Student’s Name : ")

   fmt.Scan(&name)

   a.name = append(a.name, name)

   var score int

   fmt.Print("Input " + name + " Score : ")

   fmt.Scan(&score)

   a.score = append(a.score, score)

 }

 

 fmt.Println("\n\nAvarage Score Students is", a.Avarage())

 scoreMax, nameMax := a.Max()

 fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")

 scoreMin, nameMin := a.Min()

 fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")

}