package task

import ("fmt" 
"os")

func main(){
	command := os.Args[1]
	if command != "task" {
		return
	}

	fmt.Println(command)
}