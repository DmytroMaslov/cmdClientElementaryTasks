package main

import "tools"

func main() {

	data := tools.GetData()

	if data.Task != tools.EMPTY {
		tools.PrintResult(tools.RunTask(data.Task, data.File))

	}
	if data.Task == tools.EMPTY {
		tools.PrintResult(tools.RunAllTasks(data.File))
	}
}
