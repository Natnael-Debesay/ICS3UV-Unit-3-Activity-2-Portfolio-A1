// Author: Natnael Debesay
// Version: 1.0.0
// Date: 2025-11-18
// Fileoverview: This program displays the "Happy Birthday" song with the prompted name from the user.

package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

func main() {
	// variables
	var userName string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is your name? ")
	userName, _ = reader.ReadString('\n')
	userName = strings.TrimSpace(userName)

	// output
	fmt.Println()
	fmt.Printf("Happy Birthday to you. ")
	fmt.Printf("Happy Birthday to you. ")
	fmt.Printf("%s", "Happy Birthday, dear " + userName + ". ");
	fmt.Printf("Happy Birthday to you.")

	fmt.Println("\nDone.")
}
