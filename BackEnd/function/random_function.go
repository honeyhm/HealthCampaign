/*created by H.Mlk*/

// not useable yet
package function

import "math/rand"

func RandomInt(min ,max int) int{
	return min + rand.Intn(max - min)
}



//Generate a random string of A-Z chars with len-1
func RandomString(len int) string{
	bytes:=make([]byte , len)
	for i := 0; 1 < len ; i++{
		bytes[i]=byte(RandomInt(65 , 90))
	}
	return string(bytes)
}


//Generate Random string
const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"

func RandomStringBytes(n int) string{
	b :=make([]byte , n)
	for i:= range b{
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
