
package main

import "fmt"
import "unicode/utf8"

import (
  "unicode"
  "regexp"
)



func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func ReverseFastButIncorrect(s string) string {
    size := len(s)
    buf := make([]byte, size)
    for start := 0; start < size; {
        r, n := utf8.DecodeRuneInString(s[start:])
        start += n
        utf8.EncodeRune(buf[size-start:], r)
    }
    return string(buf)
}



func ReverseGrapheme(str string) string {

  buf := []rune("")
  checked := false
  index := 0
  ret := "" 

    for _, c := range str {

        if !unicode.Is(unicode.M, c) {

            if len(buf) > 0 {
                ret = string(buf) + ret
            }

            buf = buf[:0]
            buf = append(buf, c)

            if checked == false {
                checked = true
            }

        } else if checked == false {
            ret = string(append([]rune(""), c)) + ret
        } else {
            buf = append(buf, c)
        }

        index += 1
    }

    return string(buf) + ret
}

func ReverseGrapheme2(str string) string {
    re := regexp.MustCompile("\\PM\\pM*|.")
    slice := re.FindAllString(str, -1)
    length := len(slice)
    ret := ""

    for i := 0; i < length; i += 1 {
        ret += slice[length-1-i]
    }

    return ret
}


// http://mortoray.com/2013/11/27/the-string-type-is-broken/
// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func main() {
    s := "Les Mise\u0301rables"
	// s = "noël"
	//fmt.Printf(Reverse(s))
	//fmt.Printf(ReverseFastButIncorrect(s))

	fmt.Printf(ReverseGrapheme(s))
}
