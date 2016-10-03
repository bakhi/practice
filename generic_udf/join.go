package udfs

import (
	"errors"
	"fmt"
	"sensorbee/core"
	"sensorbee/data"
	"strings"
)

type Join struct {
}

func (j *Join) Call(ctx *core.Context, args ...data.Value) (data.Value, error) {
	empty := data.String("")
	if len(args) == 1 {
		return empty, nil
	}

	switch args[0].Type() {
	case data.TypeString: // my_join("a", "b", "c", "sep") form
		var ss []string
		for _, v := range args {
			s, err := data.AsString(v)
			if err != nil {
				return empty, err
			}
			ss = append(ss, s)
		}
		return data.String(strings.Join(ss[:len(ss)-1], ss[len(ss)-1])), nil

	case data.TypeArray: // my_join(["a", "b", "c"], "sep") form
		if len(args) != 2 {
			return empty, errors.New("wrong number of arguments for my_join(array, sep)")
		}
		sep, err := data.AsString(args[1])
		if err != nil {
			return empty, err
		}

		a, _ := data.AsArray(args[0])
		var ss []string
		for _, v := range a {
			s, err := data.AsString(v)
			if err != nil {
				return empty, err
			}
			ss = append(ss, s)
		}
		return data.String(strings.Join(ss, sep)), nil
	default:
		return empty, erros.New("the first argument must be a string or an array")
	}
}

func (j *Join) Accept(arity int) bool {
	return arity >= 1
}

func (j *Join) IsAggregationParameter(k int) bool {
	return false
}

func main() {
	fmt.Println("vim-go")
}
