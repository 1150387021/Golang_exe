package main

import (
	"awesomeProject1/Bitcoine/Hash/hashTableArray"
	"fmt"
)

func main1() {
	fmt.Println(hashTableArray.MySHA("abcd1", 100))
	fmt.Println(hashTableArray.MySHA("abcd2", 100))
}

func main2() {
	fmt.Println(hashTableArray.MySHA256("abcd1", 100))
	fmt.Println(hashTableArray.MySHA256("abcd", 100))
	fmt.Println(hashTableArray.MySHA256("abcd2", 100))
}
func main() {
	mytable, _ := hashTableArray.NewHashTable(100, hashTableArray.MySHA)
	mytable.Insert("abcd1")
	mytable.Insert("abcd12")
	mytable.Insert("abcd3")
	mytable.Insert("abcd4")
	pos := mytable.Find("abcd3")
	fmt.Println(mytable.GetValue(pos))
	pos1 := mytable.Find("abcd12")
	fmt.Println(mytable.GetValue(pos1))

}
