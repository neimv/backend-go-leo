package color

import "fmt"

const (
	ColorBlack  string = "\u001b[30m"
	ColorRed    string = "\u001b[31m"
	ColorGreen  string = "\u001b[32m"
	ColorYellow string = "\u001b[33m"
	ColorBlue   string = "\u001b[34m"
	ColorReset  string = "\u001b[0m"
)

// Function to output text in color black
func PrintBlack(message string) {
	fmt.Println(string(ColorBlack), message, string(ColorReset))
}

// Function to output text in color red
func PrintRed(message string) {
	fmt.Println(string(ColorRed), message, string(ColorReset))
}

// Function to output text in color green
func PrintGreen(message string) {
	fmt.Println(string(ColorGreen), message, string(ColorReset))
}

// Function to output text in color Yellow
func PrintYellow(message string) {
	fmt.Println(string(ColorYellow), message, string(ColorReset))
}

// Function to output text in color Blue
func PrintBlue(message string) {
	fmt.Println(string(ColorBlue), message, string(ColorReset))
}
