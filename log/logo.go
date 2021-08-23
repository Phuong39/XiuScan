package log

import "fmt"

func PrintLogo(version, author string) {
	logo := " __   ___        _____                 \n \\ \\" +
		" / (_)      / ____|                \n  \\ V / _ _  " +
		" _| (___   ___ __ _ _ __  \n   > < | | | | |\\___ \\" +
		" / __/ _` | '_ \\ \n  / . \\| | |_| |____) | (_| (_" +
		"| | | | |\n /_/ \\_\\_|\\__,_|_____/ \\___\\__,_|_" +
		"| |_|"
	fmt.Println(logo)
	fmt.Printf("version: %s   author: %s\n", version, author)
	fmt.Println("A Super Java Vulnerability Scanner")
}
