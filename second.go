package main

import ("fmt")

func add(x float64,y float64) float64{
	return  x+ y

}

func mulitple(a,b string) (string,string){
	return a,b
}

func main(){
 //number declation and initialization
  num1 :=5.6
  num2 :=5.4

  fmt.Println(add(num1,num2))
// string 
  w1,w2:="hey","Nimisha"
 fmt.Println(mulitple(w1,w2))

 //conversion of int to float
  num3:=3
  float1 := float64(num3)
  fmt.Println(float1)
//pointers

 x := 15
//memory address
 a := &x
 fmt.Println(a)
 fmt.Println(*a)
 *a =5
  fmt.Println(x)
}