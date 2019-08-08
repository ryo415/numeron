package main

import( 
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"bufio"
	"os"
)

func random_num() string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(9)

	return strconv.Itoa(num)
}

func check_dup(num_seq string, num string) bool {
	length := len(num_seq)

	//	fmt.Println(string(num_seq),":",num)

	for i:=0;i<length;i++ {
		if string(num_seq[i]) == num {
	//		fmt.Println("false")
			return false
		}
	}
	//fmt.Println("true")
	return true
}

func initial_num(num int) string {
	var num_seq string
	var add_num string

	for i:=0; i<num; i++ {
		for true {
			add_num = random_num()
			if check_dup(num_seq, add_num) {
				num_seq += add_num
				break
			}
		}
	}
	return num_seq
}

func check_num(num_seq string, input string) (int, int) {
	num := len(num_seq)
	var eat, bite int

	for i:=0;i<num;i++ {
		for j:=0;j<num;j++ {
			if num_seq[i] == input[j] {
				if i == j {
					eat += 1
				} else {
					bite += 1
				}
			}
		}
	}

	return eat, bite
}

func main() {
	var num int
	var num_seq string
	var input string
	var eat,bite int

	fmt.Println("数字の個数 9以下")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num,_ = strconv.Atoi(scanner.Text())
	num_seq = initial_num(num)

	fmt.Println("--- Start ---")

	fmt.Println(num_seq)

	for true {
		fmt.Print("> ")
		scanner.Scan()
		input = scanner.Text()
		eat, bite = check_num(num_seq, input)
		if eat == num {
			fmt.Println("congraturation!!")
			break
		} else {
			fmt.Println("EAT:",eat)
			fmt.Println("BITE:",bite)
		}
	}

}
