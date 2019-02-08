package main

import "strings"
import "fmt"

func numUniqueEmails(emails []string) int {
	seen := make(map[string]bool, 0)
	for _, email := range emails {
		at := strings.Index(email, "@")
		localName := email[:at]
		localName = strings.Replace(localName, ".", "", -1)
		localName = localName[:strings.Index(localName, "+")]
		domain := email[at:]
		seen[fmt.Sprintf("%s@%s", localName, domain)] = true
	}
	return len(seen)
}

func main() {

}
