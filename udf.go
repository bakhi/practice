package main

import (
	"fmt"
	"sensorbee/bql/udf"
	"sensorbee/core"
	"sensorbee/data"
)

// Any struct implementing the following interface can be used as a UDF:
type UDF interface {
	// Call calls the UDF.
	Call(*core.Context, ...data.Value) (data.Value, error)

	// Accept checks if the function accepts the given number of arguments
	// excluding core.Context.
	Accept(arity int) bool

	// IsAggregationParameter returns true if the k-th parameter expects
	// aggregated values. A UDF with Accept(n) == true is an aggregate
	// function if and only if this function returns true for one or more
	// values of k in the range 0, ..., n-1.
	IsAggregationParameter(k int) bool
}

// This interface defined in the gopkg.in/sensorbee/sensorbee.v0/bql/udf
// A UDf can be registered via the RegisterGlobalUDF or MustRegisterGlobalUDF functions from the same pacakge.
// MustRegisterGlobalUDF is the same as RegisteredUDF but panics on failutre instead of returnning an error.
// These functions are usually called from the init function in the UDF package's plugin subpackage.
// A typical implmentation of a UDF looks as follows:

//A UDF needs to implement three methods to satisfy udf.UDF interface: Call, Accept, and IsAggregationParameter

type MyUDF struct {
}


SELECT RSTREAM my_udf(arg1, arg2) FROM stream [RANGE 1 TUPLES];
//In this example, arg1 and arg2 are passed to the Call method:
func (m *MyUDF) Call(ctx *core.Context, args ...data.Value) (data.Value, error) {
	// When my_udf(arg, arg2) is called, len(args) is 2.
	// args[0] is arg1 and args[1] is arg2.
	// it is guaranteed that m.Accept(len(args)) is always true.
}

// The Accept method verifies if the UDF accepts the specific number of arguments. It can return true for
// multiple arities as long as it can receive the given number of arguments. If a UDF only accepts two arguments,
// the method is implemented as follows:
func (m *MyUDF) Accept(arity int) bool {
	return arity == 2
	// When a UDf aims to support variadic parameters (a.k.a. variable-length arguments) with two required
	// arguments (e.g. my_udf(arg1, arg2, optional1, optional2, ...)), the implementation would be:
	func (m *MyUDF) Accept(arity int) bool{
		return arity >=2
	}
}

// Returns whether the k-th argument (starting from 0) is an aggregation parameter. Aggregation parameters are
// passed as a data.Array containing all values of a field in each group.
// All of these methods can be called concurrently from multiple goroutines and they must be thread-safe
// The registered UDF is looked up based on its name and the number of argument passed to it.
SELECT RSTREAM my_udf(arg1, arg2) FROM stream [RANGE 1 TUPLES];
// In this SELECT, a UDF having the name my_udf is looked up first.
// After that, its Accept method is called with 2 and my_udf is actually selected if Accept(2) returned true
// IsAggregationParameter method is additionally called on each argument to see if the argument needs to be an aggregation parameter.
// Then, if there is no mismatch, my_udf is finally called.

func (m *MyUDF) IsAggregationParameter(k int) bool {

}

// [NOTE] A UDF does not have a schema at the moment, so any error regarding types of arguments will not be
// reported until the statement calling the UDF actually processes a tuple. 

func init() {
	// MyUDF can be used as my_udf in BQL statements.
	udf.MustRegisterGlobalUDF("my_udf", &MyUDF{})
}

func main() {
	fmt.Println("vim-go")
}
