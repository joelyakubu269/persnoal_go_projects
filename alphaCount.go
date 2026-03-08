package main

import "fmt"

func main() {
    text:= "my name is joel"
    fmt.Println(wordCount(text))
}
func wordCount(s string) int {
    count:=0
    for _,r:= range s {
        if r==' '{
            continue
        }
        count++
    }
    return count
}
second code
package main

import "fmt"

func main() {
    text:= "my name is joel"
    fmt.Println(alphaCount(text))
}
func alphaCount(s string) int {
    count:=0
    for _,r:= range {
        if r>='A' && r<='Z' || r>='a'&& r<='z'{
            count++
        }
    }
    return count
}
third code
package main

import "fmt"

func main() {
    text:= "my name is joel"
    fmt.Println(alphaCount(text))
}
func wordCount(s string) int {
    count:=0
    for _,r:= range s {
        if r==' ' {
            continue
        }
        count++
    }
    return count
}