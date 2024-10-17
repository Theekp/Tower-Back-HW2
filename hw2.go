package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ibst struct { //дерево для интовых значений
	value int
	right *ibst
	left  *ibst
}

type sbst struct { //дерево для стринговых значений
	value string
	left  *sbst
	right *sbst
}

func main() {
	var flag int
	var str string
	fmt.Println("Your Tree: void")
	fmt.Print("Your type of values, int(1)/string(0): ")
	fmt.Scanln(&flag)
	if flag == 1 {

		var tree *ibst
		fmt.Print("Enter a list of your integers: ")
		fmt.Scanln(&str)
		fields := strings.Fields(str)

		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			tree = ins1(tree, num)
		}

		for flag == 1 {
			var a int
			fmt.Print("What operation you want to do, insert(1)/IsExist(2)/delete(3): ")
			fmt.Scanln(&a)
			switch a {
			case 1:
				var b int
				fmt.Print("Enter an integer: ")
				fmt.Scanln(&b)
				tree = ins1(tree, b)
			case 2:
				var b int
				fmt.Print("Enter an integer: ")
				fmt.Scanln(&b)
				res := exist1(tree, b)
				if res != nil {
					fmt.Println("Integer %d is exist", b)
				} else {
					fmt.Println("Integer %d is not exist", b)
				}
			case 3:
				var b int
				fmt.Print("Enter an integer: ")
				fmt.Scanln(&b)
				tree = del1(tree, b)
			}
			fmt.Print("if you want continue - 1, stop - 0: ")
			fmt.Scanln(&flag)
		}
	} else {
		var tree *sbst
		for flag != 1 {
			fmt.Println("Enter a list of your strings")
		}
	}

}

func ins1(root *ibst, value int) *ibst {
	if root == nil {
		return &ibst{value: value}
	}
	if value < root.value {
		root.left = ins1(root.left, value)
	} else {
		root.right = ins1(root.right, value)
	}
	return root
}

func ins2(root *sbst, value int) *sbst {
	if root == nil {
		return &sbst{value: value}
	}
	if value < root.value {
		root.left = ins2(root.left, value)
	} else {
		root.right = ins2(root.right, value)
	}
	return root
}

func del1(root *ibst, value int) *ibst {

}

func exist1(root *ibst, value int) *ibst {
	if root == nil || root.value == value {
		return root
	}

	if value < root.value {
		return exist1(root.left, value)
	} else {
		return exist1(root.right, value)
	}
}

func del2() {

}

func exist2() {

}
