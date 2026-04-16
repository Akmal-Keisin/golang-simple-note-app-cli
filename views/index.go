package views

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"simple-note-app/helpers"
	"strconv"
	"strings"
)

var menuOption = []string{
	"Create Note",
	"Read Note",
	"Update Note",
	"Delete Note",
	"Exit",
}

func (n *NoteView) Index(ctx context.Context, flashMessage ...string) {
	helpers.CallClear()
	if len(flashMessage) > 0 {
		fmt.Println(flashMessage[0])
	}

	fmt.Println("NOTE APP - CLI")

	for i, menu := range menuOption {
		fmt.Printf("%d. %s\n", i+1, menu)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select Menu Number: ")
	menuNumber, _ := reader.ReadString('\n')

	// Check if the input is a number
	number, err := strconv.Atoi(strings.TrimSpace((menuNumber)))
	if err != nil {
		fmt.Println("Please enter a valid number")
		return
	}

	// Adjust the number to zero based index
	number = number - 1

	// Check if the number is in the range of menu options
	if number < 0 || number > len(menuOption)-1 {
		fmt.Println("Please enter number from menu list")
	}

	// Check if the input is in the menu list
	if item := menuOption[number]; item == "" {
		fmt.Println("Please enter number from menu list")
		return
	}

	switch number {
	case 0:
		n.Create(ctx)
	case 1:
		n.Read(ctx)
	case 2:
		n.Edit(ctx)
	case 3:
		n.Delete(ctx)
	case 4:
		fmt.Println("Exiting...")
		os.Exit(0)
	}
}
