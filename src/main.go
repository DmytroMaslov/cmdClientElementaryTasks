package main

import (
	"fmt"
	"flag"
	"tools"
	"bufio"
	"os"
)
	type flags struct {
		file string
		task string
	}

func main(){
	fileNameF := flag.String("file","","File contains Json")
	taskNameF := flag.String("task", "", "Task number")
	flag.Parse()
	flags := new (flags)
	flags.file = *fileNameF
	flags.task = *taskNameF
	fileNameForAll := flag.Arg(0)
	if (flags.file == "" && fileNameForAll !=""){
		flags.file = fileNameForAll
		//add validation
	}
	if (flags.file == "" && fileNameForAll ==""){
		fmt.Println("Please enter filename:")
		scaner := bufio.NewScanner(os.Stdin)
		scaner.Scan()
		flags.file = scaner.Text()
		fmt.Println(len (flags.file))
		//add validation
	}
	if(flags.task!=""){
		fmt.Println(tools.RunTask(flags.task, flags.file))
	}
	if (flags.task ==""){
		fmt.Println(tools.RunAllTasks(flags.file))
	}
}