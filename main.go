package main

import (
	"fmt"
)

var (
	buildCommit  string
	buildTime    string
	buildVersion string
)

func main() {
	fmt.Printf("buildCommit=%s, buildTime=%s, buildVersion=%s", buildCommit, buildTime, buildVersion)
}