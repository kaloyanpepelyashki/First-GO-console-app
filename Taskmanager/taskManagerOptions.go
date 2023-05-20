package taskmanager

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// CREATE TASK interface
func createTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Let's begin\n Enter task title : ")

	title, _ := reader.ReadString('\n')

	fmt.Println(" Enter task description: ")

	description, _ := reader.ReadString('\n')

	fmt.Println(" Enter deadline: ")
	fmt.Println(" Format (dd/mm/yy) ")

	deadline, _ := reader.ReadString('\n')

	var task = Task{}
	taskId++
	task.CreateTask(taskId, title, description, deadline)
	tasksArray = append(tasksArray, task)

	fmt.Println("==========================\n\n New task added : ", title, "ID : ", task.id, "\n Description of the task : ", description, "\n Deadline of the task : ", deadline, "\n==========================")
	fmt.Println("================\n Press any key to continue...")
	_, _, _ = reader.ReadRune()
	RunTaskManager()
}

// CREATE A NEW TASK SCREEN
func runCreateNewTask() {
	reader := bufio.NewReader(os.Stdin)
	clearConsole()

	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("Do you wnat to create a new task?: ")
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)

	if strings.ToLower(answer) == "yes" {
		createTask()
	} else if strings.ToLower(answer) == "no" {
		fmt.Println("Okay, bye")
		fmt.Println(" ")
		fmt.Println("================\n Press any key to continue...")
		_, _, _ = reader.ReadRune()

		RunTaskManager()
	}

}

// COMPLETE TASK FUNCTION
func completeTask() {
	reader := bufio.NewReader(os.Stdin)
	clearConsole()
	fmt.Println("==========================\n\n Taks ID to be completed : ")

	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)

	var fitsCriteria bool = false

	var intValue, err = strconv.Atoi(answer)

	if err == nil {
		for _, task := range tasksArray {
			if task.id == intValue {
				fitsCriteria = true
				break
			}
		}

	} else {
		fmt.Println("Invalid value")
		fmt.Println("================\n Press any key to continue...")
		_, _, _ = reader.ReadRune()
		RunTaskManager()
	}

	if !fitsCriteria {
		fmt.Println("No element with this ID")
		fmt.Println(" ")
		fmt.Println("================\n Press any key to continue...")
		_, _, _ = reader.ReadRune()
		RunTaskManager()

	} else if fitsCriteria {
		for i := range tasksArray {
			if tasksArray[i].id == intValue {
				tasksArray[i].completed = true
				break
			}
		}
		fmt.Println("================\n Press any key to continue...")
		_, _, _ = reader.ReadRune()
		RunTaskManager()

	}

}

// PRINT ALL TASKS
func printAllTasks() {
	reader := bufio.NewReader(os.Stdin)
	clearConsole()

	fmt.Println("\n\n==========================\n Your tasks : ")

	for _, task := range tasksArray {

		if !task.completed && !task.deleted {
			fmt.Println("---------\n ID : ", task.id, "\n Title : ", task.title, "\n Description : ", task.description, "\n Completed : ", task.completed, "\n---------")
		}
	}

	fmt.Println("================\n Press any key to continue...")
	_, _, _ = reader.ReadRune()
	RunTaskManager()
}

// Printing the task manager Options
func printTaskManagerOptions() {
	fmt.Printf("\n\n==========================\nWelcome to the task manager\nOptions :\n 1) Create a new task\n 2) See all tasks\n 3) Complete a task\n 4) See completed tasks \n 0) Exit\n==========================\n")
}

// PRINT THE COMPLETED TASKS
func printCompletedTasks() {
	reader := bufio.NewReader(os.Stdin)
	clearConsole()
	var notCompleted = true
	for _, task := range tasksArray {
		if task.completed {
			notCompleted = false
			break
		}
	}
	fmt.Println("\n\n==========================\n Completed tasks :")
	if !notCompleted {

		for i := range tasksArray {
			if tasksArray[i].completed {
				fmt.Println("---------\n ID : ", tasksArray[i].id, "\n Title : ", tasksArray[i].title, "\n Description : ", tasksArray[i].description, "\n---------")
				break
			}
		}
		fmt.Println("================\n Press any key to continue...")
		_, _, _ = reader.ReadRune()
		RunTaskManager()
	} else {

		fmt.Println("\n\nNo completed tasks yet\n\n")
		fmt.Println("================\n Press any key to continue...")
		_, _, _ = reader.ReadRune()
		RunTaskManager()
	}

}

// CLEAR CONSOLE FUNC
func clearConsole() {

	// Clear console based on the operating system
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported operating system")
	}

}

// PRESS ANY KEY TO CONTINUE
func pressKeyToContinue() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("================\n Press any key to continue...")
	_, _, _ = reader.ReadRune()

}
