package grep

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type options struct {
	lineNumbers     bool
	onlyFilenames   bool
	caseInsensitive bool
	invertMatch     bool
	matchFullLines  bool
	outputFilename  bool
}

func searchFile(path, pattern string, opts *options) []string {
	res := []string{}
	lineNumber := 0

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
	}
	return res
}

func Search(pattern string, flags []string, files []string) []string {
	opts := parseFlags(flags)

	if len(files) > 1 {
		opts.outputFilename = true
	}

	if opts.caseInsensitive {
		pattern = strings.ToLower(pattern)
	}

	// if opts.matchFullLines {
	// 	match = func(line, pattern string) bool {
	// 		ret := pattern == line

	// 		if opts.invertMatch {
	// 			return !ret
	// 		}

	// 		return ret
	// 	}
	// } else {
	// 	match = func(line, pattern string) bool {
	// 		ret := strings.Index(line, pattern) != -1
	// 		if opts.invertMatch {
	// 			return !ret
	// 		}

	// 		return ret
	// 	}
	// }

	res := []string{}

	for _, path := range files {
		res = append(res, searchFile(path, pattern, opts)...)

		// file, err := os.Open(path)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer file.Close()

		// scanner := bufio.NewScanner(file)

		// found := false
		// i := 0
		// for scanner.Scan() {
		// 	i++
		// 	line := strings.TrimSpace(scanner.Text())
		// 	if opts.caseInsensitive {
		// 		line = strings.ToLower(line)
		// 	}
		// 	if match(line, pattern) {
		// 		found = true
		// 		if opts.onlyFileNames {
		// 			break
		// 		}

		// 		res := scanner.Text()

		// 		if opts.lineNumbers {
		// 			res = fmt.Sprintf("%d:%s", i, res)
		// 		}
		// 		if multipleFiles {
		// 			res = fmt.Sprintf("%s:%s", path, res)
		// 		}

		// 		output = append(output, res)
		// 	}
		// }

		// if opts.onlyFileNames && found {
		// 	output = append(output, path)
		// }

		// if err := scanner.Err(); err != nil {
		// 	log.Fatal(err)
		// }
	}

	return res
}

func parseFlags(flags []string) *options {
	o := options{}

	for _, f := range flags {
		switch f {
		case "-n":
			o.lineNumbers = true
		case "-l":
			o.onlyFilenames = true
		case "-i":
			o.caseInsensitive = true
		case "-v":
			o.invertMatch = true
		case "-x":
			o.matchFullLines = true
		}
	}

	return &o
}
