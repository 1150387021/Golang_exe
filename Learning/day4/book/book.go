package book

import "fmt"

func Printtype(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "type is int")
		case string:
			fmt.Println(arg, "type is string")
		case float64:
			fmt.Println(arg, "type is float64")
		default:
			fmt.Println(arg, "is an unknow type")

			var s []string
			s = append(s, []string{"nihoa", "小猫", "加油"}...)
			fmt.Println(s)

		}
	}
}
