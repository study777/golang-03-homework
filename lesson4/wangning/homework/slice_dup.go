package main

import "fmt"

func dup(s []string) []string {
	n := len(s)
	j := -1
	for i := 1; i < n-1; i++ {
	//	fmt.Println(s[i],s[i+1])
		if s[i] == s[i+1] {
			s = append(s[:i], s[i+1:]...)
			i = j
		} else {
			j+=1
			i=j
		}
	n = len(s)
	//fmt.Println(s)
	//fmt.Println(n)
	}
	return s
}

func dup2(s []string) []string {
	var s1 []string
	s1 = append(s1, s[0])
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			s1 = append(s1, s[i])

		}
		//	fmt.Println(s1)
	}
	return s1
}

func dup3(s []string) []string {
		

//		fmt.Println("=========3=======")
		b := 0 
		for x,y :=0,1; y < len(s);x,y= x+1,y+1 {

			if s[x] == s[y] {
//				fmt.Println(b,x,y)
				if x ==  b {
					b +=1
//				fmt.Println(s)	
				} else {
					for i:=x;i>=b;i-- {
					s[i+1]=s[i]
					}
				b +=1
				//fmt.Println(s)
				//fmt.Println(x,y)
				}
				
				
			}	
		}
		return s[b:]
}

func main() {

	s := []string{"Tom", "Tom", "Tom", "Tom", "Tom", "Tom", "Make", "Make", "Tom", "JOM", "Make", "Make", "Tom", "Tom"}
	s2 := []string{"Tom", "Tom", "Tom", "Tom", "Tom", "Tom", "Make", "Make", "Tom", "JOM", "Make", "Make", "Tom", "Tom"}
	s3 := []string{"Tom", "Tom", "Tom", "Tom", "Tom", "Tom", "Make", "Make", "Tom", "JOM", "Make", "Make", "Tom", "Tom"}
	fmt.Println(dup(s))
	fmt.Println(dup2(s2))
	fmt.Println(dup3(s3))

}
