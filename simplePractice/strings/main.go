package main

import "strings"
import "fmt"

func main() {
	resultMap := AnalyzeLogin("  Peter-Parker-1987@example.com  ")
	fmt.Println("  Peter-Parker-1987@example.com  ")
	fmt.Println(resultMap["trimmed"])
	fmt.Println(resultMap["lowercase"])
	fmt.Println(resultMap["username"])
	fmt.Println(resultMap["domain"])
	fmt.Println(resultMap["replacedDash"])

}

func AnalyzeLogin(input string) map[string]string {
	results := make(map[string]string)
	trimmed := strings.TrimSpace(input)
	results["trimmed"] = trimmed
	toLower := strings.ToLower(trimmed)
	results["lowercase"] = toLower
	username, domain, found := strings.Cut(toLower, "@")
	if found == true {
		results["username"] = username
		results["domain"] = domain
	}
	replacedDash := strings.ReplaceAll(username, "-", "_")
	results["replacedDash"] = replacedDash
/*	initials := ""
	for _, r := range []rune(input) {

	}*/

	return results
}
