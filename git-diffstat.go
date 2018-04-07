package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var title = struct {
	add string
	delete string
	path string
}{
	"add",
	"delete",
	"path",
}

const spaces string = "  "

type result struct {
	path string
	add string
	delete string
}

func maxLen(max int, results []result, f func(result) string) int {
	for _, r := range results {
		l := len(f(r))
		if l > max {
			max = l
		}
	}
	return max
}

const usageMessage = "" +
`Usage of 'git-loc':

> git-loc HEAD^
path       add  delete
----------------------
README.md   14       0
`

func usage() {
	fmt.Fprintln(os.Stderr, usageMessage)
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	git := new(Git)

	args := []string{"--numstat"}
	args = append(args, flag.Args()...)

	lines, err := git.Diff(args...)
	if err != nil {
		log.Fatalln(err)
	}

	results := make([]result, len(lines))
	for i, line := range lines {
		s := strings.Fields(line)
		if len(s) != 3 {
			log.Fatalf("The number of fields is not 3, actual=%d", len(s))
		}

		results[i].add, results[i].delete, results[i].path = s[0], s[1], s[2]
	}

	width := struct {
		add int
		delete int
		path int
	}{
		maxLen(len(title.add), results, func(r result) string {return r.add}),
		maxLen(len(title.delete), results, func(r result) string {return r.delete}),
		maxLen(len(title.path), results, func(r result) string {return r.path}),
	}

	fmt.Printf("%-[4]*[1]s%[7]s%[5]*[2]s%[7]s%[6]*[3]s\n", title.path, title.add, title.delete, width.path, width.add, width.delete, spaces)
	fmt.Printf("%s\n", strings.Repeat("-", width.add+width.delete+width.path+len(spaces)*2))
	for _, r := range results {
		fmt.Printf("%-[4]*[1]s%[7]s%[5]*[2]s%[7]s%[6]*[3]s\n", r.path, r.add, r.delete, width.path, width.add, width.delete, spaces)
	}
}