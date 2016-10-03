package plugin

import (
	"sensorbee/bql/udf"
	"strings"
	""
)

func init() {
	udf.MustRegisterGlobalUDF("my_inc", udf.MustConvertGeneric(udfs.Inc))
	udf.MustRegisterGlobalUDF("my_join", &udfs.Join{})
	udf.MustRegisterGlobalUDF("my_join2", udf.MustConvertGeneric(strings.Join))
}
