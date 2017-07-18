package main

import (
	"fmt"
	"flag"
	"tools"
	"bufio"
	"os"
	"strings"
	"errors"
)
const EMPTY = ""

	type flags struct {
		file string
		task string
	}

func main(){
	fileNameF := flag.String("file",EMPTY,"File contains Json")
	taskNameF := flag.String("task", EMPTY, "Task number")
	flag.Parse()
	flags := new (flags)
	flags.file = *fileNameF
	flags.task = *taskNameF
	fileNameForAll := flag.Arg(0)

	if (flags.file == EMPTY && fileNameForAll !=EMPTY){
		flags.file = fileNameForAll
		//add validation
	}
	if (flags.file == EMPTY) {
		ok := true
		for ok {
			file, err := GetFileName()
			if err != nil{
				fmt.Println(err.Error())
			} else{
				ok = false
				flags.file = file
			}

		}

	}
	if(flags.task!=EMPTY){
		PrintResult(tools.RunTask(flags.task, flags.file))


	}
	if (flags.task ==EMPTY){
		PrintResult(tools.RunAllTasks(flags.file))
	}
}

func GetFileName () (file string, err error) {
		fmt.Println("Please enter filename:")
		scaner := bufio.NewScanner(os.Stdin)
		scaner.Scan()
		file = scaner.Text()
		v := strings.HasSuffix(file, ".json")
		if !v {
			return file, errors.New("Not JSON file")
		}
		ok := tools.IsPresent(file)
			if !ok {
				return file, errors.New("Can't find file")}
		return file, nil

	}

func PrintResult(tasks map[string][]string, err error) {
	if err != nil{
		fmt.Printf("Error:%s", err.Error())
	} else {
		for _, v := range tasks {
			fmt.Printf("Task:%s\nResp:\n%s\nReason:\n%s\n", v[0], v[1], v[2])
		}
		fmt.Println("++++++++++++++++++++++++")
	}

}


