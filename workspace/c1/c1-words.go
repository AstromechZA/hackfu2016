package main

import (
    "fmt"
    "os"
    "path/filepath"
    "io/ioutil"
    "strings"
    "errors"
    "regexp"
)

type LetterEntry struct {
    x, y int
}


var LetterEntries = [11]LetterEntry{
    LetterEntry{x: 85, y: 8},
    LetterEntry{x: 124, y: 11},
    LetterEntry{x: 1984, y: 8},
    LetterEntry{x: 3, y: 5},
    LetterEntry{x: 901, y: 1},
    LetterEntry{x: 3, y: 13},
    LetterEntry{x: 8546, y: 12},
    LetterEntry{x: 5, y: 2},
    LetterEntry{x: 3, y: 4},
    LetterEntry{x: 85, y: 10},
    LetterEntry{x: 3437, y: 7},
}

func GetLines(path string) *[]string {
    bytes, err := ioutil.ReadFile(path)
    if err != nil {
        panic(err.Error())
    }
    lines := strings.Split(string(bytes), "\n")
    return &lines
}

func FilterLines(lines *[]string, f func(string) bool) *[]string {
    ols := make([]string, 0)
    for _, line := range *lines {
        if f(line) {
            ols = append(ols, line)
        }
    }
    return &ols
}

var badCharRe = regexp.MustCompile("[^a-zA-Z ]")
var multiSpaceRe = regexp.MustCompile("\\s+")

func FetchWordEntry(lines *[]string, xx int, yy int) (string, error) {
    if len(*lines) <= xx {
        return "", errors.New("Not enough lines")
    }

    line := (*lines)[xx]
    line = strings.TrimSpace(line)
    line = multiSpaceRe.ReplaceAllString(line, " ")
    line = badCharRe.ReplaceAllString(line, "")
    words := strings.Split(line, " ")
    if len(words) <= yy {
        return "", errors.New("Not enough characters")
    }

    return string(words[yy]), nil
}

func main() {

    thisAbsPath, _ := filepath.Abs(os.Args[0])
    bookDir := filepath.Join(thisAbsPath, "..", "Books/*")

    files, err := filepath.Glob(bookDir)
    if err != nil {
        panic("failed to list dir")
    }

    for _, f := range files {
        lines := GetLines(f)
        password := ""

        var err error
        var v string
        for _, e := range LetterEntries {
            v, err = FetchWordEntry(lines, e.x - 1, e.y - 1)
            if err == nil {
                password = password + v
            } else {
                break
            }
        }
        if err == nil {
            fmt.Println(f)
            fmt.Println("'", password, "'")
        }
    }
}
