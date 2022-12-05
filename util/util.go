// util.go
package util

import (
	"fmt"
	"regexp"
	"rye/env"
	"strconv"
	"strings"
)

func PrintHeader() {
	fmt.Println("=-===============-===-===-=============-=")   // Output: -3
	fmt.Println(" _/|\\\\_-~*>%,_  Rejy ZERO  _,%<*~-_//|\\_") // Output: -3
	fmt.Println("=-===============-===-===-=============-=")   // Output: -3
}

func IndexOfAt(s, sep string, n int) int {
	idx := strings.Index(s[n:], sep)
	if idx > -1 {
		idx += n
	}
	return idx
}

func IsTruthy(o env.Object) bool {
	switch oo := o.(type) {
	case env.Integer:
		return oo.Value > 0
	case env.String:
		return len(oo.Value) > 0
	default:
		return false
	}
}

func Dict2Context(ps *env.ProgramState, s1 env.Dict) env.RyeCtx {
	ctx := env.NewEnv(ps.Ctx)
	for k, v := range s1.Data {
		word := ps.Idx.IndexWord(k)
		switch v1 := v.(type) {
		case env.Integer:
			ctx.Set(word, v1)
		case env.String:
			ctx.Set(word, v1)
		case string:
			ctx.Set(word, env.String{v1})
		}
	}
	return *ctx
}

func StringToFieldsWithQuoted(str string, sepa string, quote string) env.Block {
	quoted := false
	spl := strings.FieldsFunc(str, func(r rune) bool {
		if string(r) == quote {
			quoted = !quoted
		}
		return !quoted && string(r) == sepa
	})
	lst := make([]env.Object, len(spl))
	for i := 0; i < len(spl); i++ {
		//fmt.Println(spl[i])
		// TODO -- detect numbers and turn them to integers or floats, later we can also detect other types of values
		// val, _ := loader.LoadString(spl[i], false)
		numeric, _ := regexp.MatchString("[0-9]+", spl[i])
		// fmt.Println(numeric)
		pass := false
		if numeric {
			num, err := strconv.Atoi(spl[i])
			if err == nil {
				lst[i] = env.Integer{int64(num)}
				pass = true
			} else {
				// fmt.Println(err.Error())
			}
		}
		if !pass {
			clean := regexp.MustCompile(`^"(.*)"$`).ReplaceAllString(spl[i], `$1`)
			val := env.String{clean}
			lst[i] = val
		}
	}
	return *env.NewBlock(*env.NewTSeries(lst))
}

func FormatJson(val env.Object, e env.Idxs) string {
	// TODO -- this is currently made just for block of strings and integers
	var r strings.Builder
	switch b := val.(type) {
	case env.Block:
		r.WriteString("[ ")
		for i := 0; i < b.Series.Len(); i += 1 {
			if b.Series.Get(i) != nil {
				if i > 0 {
					r.WriteString(", ")
				}
				r.WriteString(b.Series.Get(i).Probe(e))
			}
		}
		r.WriteString(" ]")

	}
	return r.String()
}

func FormatCsv(val env.Object, e env.Idxs) string {
	// TODO -- this is currently made just for block of strings and integers
	var r strings.Builder
	switch b := val.(type) {
	case env.Block:
		for i := 0; i < b.Series.Len(); i += 1 {
			if b.Series.Get(i) != nil {
				if i > 0 {
					r.WriteString(",")
				}
				r.WriteString(b.Series.Get(i).Probe(e))
			}
		}
	}
	return r.String()
}

func FormatSsv(val env.Object, e env.Idxs) string {
	// TODO -- this is currently made just for block of strings and integers
	var r strings.Builder
	switch b := val.(type) {
	case env.Block:
		for i := 0; i < b.Series.Len(); i += 1 {
			if b.Series.Get(i) != nil {
				if i > 0 {
					r.WriteString(" ")
				}
				r.WriteString(b.Series.Get(i).Probe(e))
			}
		}
	}
	return r.String()
}

func SplitEveryString(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}

func SplitEveryList(s []env.Object, chunkSize int) [][]env.Object {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return append(make([][]env.Object, 0), s)
	}
	var chunks [][]env.Object = make([][]env.Object, 0, (len(s)-1)/chunkSize+1)
	var chunk []env.Object = make([]env.Object, 0, chunkSize)
	currentLen := 0
	// currentStart := 0
	for i := range s {
		chunk = append(chunk, s[i])
		currentLen++
		if currentLen == chunkSize {
			chunks = append(chunks, chunk)
			currentLen = 0
			//		currentStart = i + 1
			chunk = make([]env.Object, 0, chunkSize)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func IntersectStrings(a string, b string) string {
	res := ""
	for _, ch := range a {
		if strings.Index(b, string(ch)) > -1 && strings.Index(res, string(ch)) == -1 {
			res = res + string(ch)
		}
	}

	return res
}

func IntersectLists(a []env.Object, b []env.Object) []env.Object {
	set := make([]env.Object, 0)

	for _, v := range a {
		//		if ContainsGen(b, v) {
		set = append(set, v)
		//}
	}

	return set
}
