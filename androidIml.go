package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var excludeFolderJunks = map[string]bool{
	"idea-classes":      true,
	"idea-test-classes": true,
	"surefire":          true,
	"surefire-reports":  true,
	"maven-archiver":    true,
	"test-classes":      true,
}

func main() {
	maybeExcludeFolderJunkRegexp := regexp.MustCompile(`<excludeFolder\s+url="file://\$MODULE_DIR\$/target/([\w\-]+)"`)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		maybeExcludeFolderJunks := maybeExcludeFolderJunkRegexp.FindStringSubmatch(line)
		var printLine bool
		if len(maybeExcludeFolderJunks) == 2 {
			maybeExcludeFolderJunk := maybeExcludeFolderJunks[1]
			printLine = !excludeFolderJunks[maybeExcludeFolderJunk]
		} else {
			printLine = true
		}
		if printLine {
			fmt.Println(line)
		}
	}
}
