/*package main

import "fmt"

func main() {
	var mounth int
	fmt.Println("请输入一个表示月份的数字：")
	fmt.Scanf("%d", &mounth)
	switch {
	case 3 <= mounth && mounth <= 5:
		fmt.Println("spring")
	case 6 <= mounth && mounth <= 8:
		fmt.Println("summer")
	case 9 <= mounth && mounth <= 11:
		fmt.Println("autumn")
	case mounth == 12 || 1 == mounth || mounth == 2:
		fmt.Println("winter")
	default:
		fmt.Println("请输入正常月份")
	}

}
*/
package main

import "fmt"

func main() {
	fmt.Printf(Season(3))
}

func Season(month int) string {
	switch month {
	case 12, 1, 2:
		return "Winter"
	case 3, 4, 5:
		return "Spring"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Autumn"
	}
	return "Season unknown"
}
