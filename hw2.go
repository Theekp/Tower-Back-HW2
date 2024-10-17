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
	_, err := fmt.Scanln(&flag)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if flag == 1 {

		var tree *ibst
		fmt.Print("Enter a list of your integers: ")
		_, err := fmt.Scanln(&str)
		if err != nil {
			fmt.Println("Error:", err)
		}
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
			_, err := fmt.Scanln(&a)
			if err != nil {
				fmt.Println("Error:", err)
			}
			switch a {
			case 1:
				var b int
				fmt.Print("Enter an integer: ")
				_, err := fmt.Scanln(&b)
				if err != nil {
					fmt.Println("Error:", err)
				}
				tree = ins1(tree, b)
			case 2:
				var b int
				fmt.Print("Enter an integer: ")
				_, err := fmt.Scanln(&b)
				if err != nil {
					fmt.Println("Error:", err)
				}
				res := exist1(tree, b)
				if res != nil {
					fmt.Printf("Integer %d is exist\n", b)
				} else {
					fmt.Printf("Integer %d is not exist\n", b)
				}
			case 3:
				var b int
				fmt.Print("Enter an integer: ")
				_, err := fmt.Scanln(&b)
				if err != nil {
					fmt.Println("Error:", err)
				}
				tree = del1(tree, b)
			default:
				fmt.Println("Invalid operation")
			}
			fmt.Print("if you want continue - 1, stop - 0: ")
			_, err1 := fmt.Scanln(&flag)
			if err1 != nil {
				fmt.Println("Error:", err1)
			}
		}
	} else {

		var tree *sbst
		fmt.Print("Enter a list of your strings: ")
		_, err := fmt.Scanln(&str)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fields := strings.Fields(str)

		for _, field := range fields {
			tree = ins2(tree, field)
		}

		for flag != 1 {
			var a int
			fmt.Print("What operation you want to do, insert(1)/IsExist(2)/delete(3): ")
			_, err := fmt.Scanln(&a)
			if err != nil {
				fmt.Println("Error:", err)
			}
			switch a {
			case 1:
				var b string
				fmt.Print("Enter an string: ")
				_, err := fmt.Scanln(&b)
				if err != nil {
					fmt.Println("Error:", err)
				}
				tree = ins2(tree, b)
			case 2:
				var b string
				fmt.Print("Enter an string: ")
				_, err := fmt.Scanln(&b)
				if err != nil {
					fmt.Println("Error:", err)
				}
				res := exist2(tree, b)
				if res != nil {
					fmt.Printf("String %v is exist\n", b)
				} else {
					fmt.Printf("String %v is not exist\n", b)
				}
			case 3:
				var b string
				fmt.Print("Enter an string: ")
				_, err := fmt.Scanln(&b)
				if err != nil {
					fmt.Println("Error:", err)
				}
				tree = del2(tree, b)
			}
			fmt.Print("if you want continue - 1, stop - 0: ")
			_, err1 := fmt.Scanln(&flag)
			if err1 != nil {
				fmt.Println("Error:", err1)
			}
			switch flag {
			case 1:
				flag = 0
			case 0:
				flag = 1
			default:
				fmt.Println("Error")
			}
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

func ins2(root *sbst, value string) *sbst {
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
	if root == nil {
		return root
	}
	if value < root.value {
		root.left = del1(root.left, value)
	} else if value > root.value {
		root.right = del1(root.right, value)
	} else {
		if root.left == nil && root.right == nil {
			root = nil
		} else if root.left == nil {
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else {
			min1 := min1(root.right)
			root.value = min1.value
			root.right = del1(root.right, min1.value)
		}
	}
	return root
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

func del2(root *sbst, value string) *sbst {
	if root == nil {
		return root
	}
	if value < root.value {
		root.left = del2(root.left, value)
	} else if value > root.value {
		root.right = del2(root.right, value)
	} else {
		if root.left == nil && root.right == nil {
			root = nil
		} else if root.left == nil {
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else {
			min2 := min2(root.right)
			root.value = min2.value
			root.right = del2(root.right, min2.value)
		}
	}
	return root
}

func exist2(root *sbst, value string) *sbst {
	if root == nil || root.value == value {
		return root
	}

	if value < root.value {
		return exist2(root.left, value)
	} else {
		return exist2(root.right, value)
	}
}

func min1(root *ibst) *ibst {
	current := root
	for current.left != nil {
		current = current.left
	}
	return current
}

func min2(root *sbst) *sbst {
	current := root
	for current.left != nil {
		current = current.left
	}
	return current
}
