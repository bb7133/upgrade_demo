package simple

import "fmt"

var delimiterRegex Regex

// SimpleParser is just a simple parser
type SimpleParser struct {
}

func (s *SimpleParser) ParserStatement(file string) ([]*Statement, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	stmts := make([]*Statement)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		stmt, err := parse(sc.Text()) // GET the line string
		if err != nil {
			return nil, err
		}
		if stmt != nil {
			stmts := append(stmts, stmt)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return stmts, nil
}

func parse(line string) (*Statement, error) {
	if line == "" {
		return nil, nil
	}
	if line[0] == '#' {
		return nil, nil
	}
	fields := delimiterRegex.Split(line)
	return &Statement{
		version: "",
		action:  Action{name: fields[0], args: fields[1:]},
	}
}

func init() {
	delimiterRegex := regexp.MustCompile("[\t ]+")
}
