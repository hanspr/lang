package lang

import (
	"bufio"
	"os"
	"strings"
)

type Lang struct {
	name     string
	path     string
	paths    []string
	sentence map[string]string
}

func NewLang(name, filename string) *Lang {
	l := new(Lang)

	l.name = name
	l.sentence = make(map[string]string)
	if filename == "" || strings.Contains(filename, ".lang") == false {
		return l
	}
	l.path = filename
	l.LoadLangFile(filename)
	return l
}

func (l *Lang) GetLangName() string {
	return l.name
}

func (l *Lang) ChangeLangName(name, filename string) bool {
	if name == l.name {
		return false
	}
	l.name = name
	l.paths = nil
	if strings.Contains(filename, ".lang") == false {
		filename = ""
	}
	l.path = filename
	l.ReloadLang()
	return true
}

func (l *Lang) ReloadLang() {
	for k, _ := range l.sentence {
		delete(l.sentence, k)
	}
	l.LoadLangFiles()
	return
}

func (l *Lang) LoadLangFiles() {
	if l.path != "" {
		l.LoadLangFile(l.path)
	}
	for _, filename := range l.paths {
		l.LoadLangFile(filename)
	}
	return
}

func (l *Lang) AppendLangFile(name, filename string) bool {
	if filename == "" || strings.Contains(filename, ".lang") == false {
		return false
	}
	l.paths = append(l.paths, filename)
	l.LoadLangFile(filename)
	return true
}

func (l *Lang) LoadLangFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" && strings.Contains(text, "|") {
			line := strings.Split(text, "|")
			l.sentence[line[0]] = line[1]
		}
	}
	return
}

func (l *Lang) Translate(s string) string {
	V, ok := l.sentence[s]
	if ok == false {
		return s
	}
	if V == "" {
		return s
	}
	return V
}
