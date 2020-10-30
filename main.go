package main

import (
	"fmt"
)

const prompt = `Please enter the number to operate:
1. Download gincmf mod
2. Make install
3. Change administrator password
4. Fix the program
0. Exit the CLI`

func main() {
	fmt.Print("Welccome back to gincmf CLI \n")
Exit:
	for {
		fmt.Println(prompt)
		var num int
		var err error
		_, err = fmt.Scanf("%d \n", &num)
		if err != nil {
			fmt.Println(err)
		}
		switch num {
		case 1:
			fmt.Println("Getting all mod.......")
			// TODO: 获取所有模块列表+打印
			fmt.Println("Please enter the mod number to download:")
			var num int // 模块的编号
			var err error
			_, err = fmt.Scanf("%d \n", &num)
			if err != nil {
				fmt.Println(err)
			}
			// TODO:

		case 2:

		case 3:

		case 4:

		case 0:
			break Exit
		}
	}
}
