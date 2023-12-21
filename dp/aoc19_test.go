package dp

import (
	"fmt"
	"hello/testhelper"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

type Part struct {
	x, m, a, s int
}

type Condition struct {
	output     string
	property   string
	comparator string
	q          int
}

type Eval struct {
	name      string
	condition []Condition
}

func Parse(lines []string) ([]Part, map[string]Eval) {
	pattern := regexp.MustCompile("((x|m|a|s)(>|<)(\\d+):(\\w+))|(\\w+)")
	isParts := false
	parts := make([]Part, 0)
	evals := make(map[string]Eval, 0)
	for _, l := range lines {
		if len(l) == 0 {
			isParts = true
			continue
		}

		if isParts {
			trimmed := l[1 : len(l)-1]
			vals := strings.Split(trimmed, ",")
			part := Part{}
			for _, val := range vals {
				fmt.Println(val)
				val_n, _ := strconv.Atoi(val[2:])

				switch val[0] {
				case 'x':
					part.x = val_n
				case 'm':
					part.m = val_n
				case 'a':
					part.a = val_n
				case 's':
					part.s = val_n
				}
			}
			parts = append(parts, part)
		} else {
			indexLBrace := strings.Index(l, "{")
			indexRBrace := len(l) - 1
			name := l[0:indexLBrace]
			cond_l := l[indexLBrace+1 : indexRBrace]
			conds := strings.Split(cond_l, ",")
			conditions := make([]Condition, 0)

			for _, cond := range conds {

				matches := pattern.FindStringSubmatch(cond)
				fmt.Println(pattern.NumSubexp())
				for _, m := range matches {
					fmt.Printf("%s - %s -%d\n", name, m, len(m))
				}
				expr := Condition{}
				if len(matches[6]) > 0 {
					expr.output = matches[6]
				} else {
					expr.property = matches[2]
					expr.comparator = matches[3]
					expr.q, _ = strconv.Atoi(matches[4])
					expr.output = matches[5]
				}
				conditions = append(conditions, expr)
			}
			evals[name] = Eval{name, conditions}

		}
	}
	return parts, evals
}

func TestRead(t *testing.T) {
	contents, _ := readFile("aoc19testdata.txt")
	lines := strings.Split(contents, "\n")
	testhelper.AssertInteger(t, 17, len(lines))
}

func TestParse(t *testing.T) {
	contents, _ := readFile("aoc19testdata.txt")
	lines := strings.Split(contents, "\n")
	parts, evals := Parse(lines)
	testhelper.AssertInteger(t, 5, len(parts))
	testhelper.AssertInteger(t, 11, len(evals))
	testhelper.AssertInteger(t, 1013, parts[4].s)
	testhelper.AssertInteger(t, 2655, parts[0].m)
	testhelper.AssertInteger(t, 2006, evals["px"].condition[0].q)
	testhelper.AssertCorrectMessage(t, "<", evals["px"].condition[0].comparator)
	testhelper.AssertCorrectMessage(t, "qkq", evals["px"].condition[0].output)
	testhelper.AssertCorrectMessage(t, "A", evals["px"].condition[1].output)
	testhelper.AssertCorrectMessage(t, "rfg", evals["px"].condition[2].output)
	testhelper.AssertCorrectMessage(t, "px", evals["px"].name)
}
