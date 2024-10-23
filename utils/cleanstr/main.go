// package main

// import (
// 	"fmt"
// 	"os"
// )
// func main(){
// 	if len(os.Args) != 2 {
// 		fmt.Println("not enough argument")
// 		return
// 	}
// 	res := ""
// 	myArg := os.Args[1]
// 	str := span(myArg)
// 	for i := 0 ; i < len(str) ; i++ {
// 		if i != len(str)-1 {
// 			res += str[i] + " "
// 		} else {
// 			res += str[i]
// 		}
// 	}
// 	fmt.Print(res + "\n")

// }
// func span(s string)[]string {
// 	slc := []string{}
// 	res := ""
// 	for i := 0 ; i < len(s) ; i++ {
// 		if string(s[i]) != " " {
// 			res += string(s[i])
// 		} else if string(s[i]) == " " && res != "" {
// 			slc = append(slc, res)
// 			res = ""
// 		}
// 	}
// 	if res != "" {
// 		slc = append(slc, res)
// 	}
// 	return slc
// }
package main

import(
  "os"
  "github.com/01-edu/z01"
)
func splt(s string) []string {
  slc := []string{}
  res := ""
  for i := 0 ; i < len(s) ; i++ {
    if string(s[i]) != " "{
      res += string(s[i])
    } else if string(s[i]) == " " && res != ""{
      slc = append(slc , res)
      res = ""
    }
  }
  if res != "" {
    slc = append(slc ,res)
  }
  return slc
}

func main(){
  if len(os.Args) != 2 {
    z01.PrintRune('\n')
	return
  }
 
  myArg := os.Args[1:]
  if myArg[0] == "" {
    z01.PrintRune('\n')
  }
  slc := splt(myArg[0])
  res := ""
  for i := 0 ; i < len(slc) ; i++ {
    if i != len(slc)-1 {
      res += slc[i] + " "
    } else {
      res += slc[i]
    }
  }
  for _ , ch := range res {
    z01.PrintRune(ch)
   
  }
  z01.PrintRune('\n')
}