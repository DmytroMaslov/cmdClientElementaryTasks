package tools

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

const EMPTY = ""
const JSON_SYFIX = ".json"
const TASK = 0
const RESP = 1
const REASON = 2
const FIRST_ARGUMENT = 0

type Data struct {
	File string
	Task string
}

func GetData() Data {
	data := getDataCmd()
	if data.File != EMPTY {
		if !isJSON(data.File) {
			fmt.Println("Argumemt isn`t Json file")
			data.File = EMPTY
		}
		if !IsPresent(data.File) {
			fmt.Println("Can't find file")
			data.File = EMPTY
		}
	}
	if data.File == EMPTY {
		ok := true
		for ok {
			file, err := GetFileName()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				ok = false
				data.File = file
			}
		}
	}
	return data
}

func getDataCmd() Data {
	fileNameF := flag.String("file", EMPTY, "File contains Json")
	taskNameF := flag.String("Task", EMPTY, "TASK number")
	flag.Parse()
	data := new(Data)
	data.File = *fileNameF
	data.Task = *taskNameF
	if data.File == EMPTY {
		data.File = flag.Arg(FIRST_ARGUMENT)
	}

	return *data
}

func GetFileName() (file string, err error) {
	fmt.Println("Please enter filename:")
	scaner := bufio.NewScanner(os.Stdin)
	scaner.Scan()
	file = scaner.Text()
	if !isJSON(file) {
		return file, errors.New("Not JSON file")
	}
	if !IsPresent(file) {
		return file, errors.New("Can't find file")
	}
	return file, nil

}

func PrintResult(tasks map[string][]string, err error) {
	if err != nil {
		fmt.Printf("Error:%s", err.Error())
	} else {
		for _, v := range tasks {
			fmt.Printf("TASK:%s\nResp:\n%s\nReason:\n%s\n", v[TASK], v[RESP], v[REASON])
		}
		fmt.Println("++++++++++++++++++++++++")
	}

}

func isJSON(str string) bool {
	return strings.HasSuffix(str, JSON_SYFIX)
}
