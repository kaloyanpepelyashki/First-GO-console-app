package taskmanager

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// START TASK MANAGER
func RunTaskManager() {
	reader := bufio.NewReader(os.Stdin)
	clearConsole()
	printTaskManagerOptions()

	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)

	if strings.Contains(answer, "1") {

		runCreateNewTask()

	} else if answer == "2" {

		if len(tasksArray) > 0 {

			printAllTasks()

		} else {

			clearConsole()
			fmt.Println("\n\n No tasks to show")
		}

	} else if answer == "0" {

		clearConsole()
		fmt.Println("==========================\n Closing....\n==========================")
		os.Exit(0)

	} else if answer == "3" {

		completeTask()

	} else if answer == "4" {

		printCompletedTasks()

	} else {

		fmt.Println("! Invalid option !")
		answer = ""
		fmt.Println("================\n Press any key to continue...")
		_, _, _ = reader.ReadRune()
		RunTaskManager()
	}

}
