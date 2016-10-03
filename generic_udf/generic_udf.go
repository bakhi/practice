package main

import "sensorbee/bql/udf"

// SensorBee provides a helper function to register a regular Go function
// as a UDF without implementing the UDF interface explicitly

func Inc(v int) int {
	return v + 1
}

// Inc can be transformed into a UDf by ConvertGeneric or MustConvertGeneric function defined in the
// gopkg.in/sensorbee/sensorbee.v0/bql/duf package. By combining it with
// RegisterGlobalUDF, the Inc function can easily be registered as a UDF:
func init() {
	udf.MustRegisterGlobalUDF("inc", udf.MustConvertGeneric(Inc))
}

// [NOTE] A UDF implementation and registration should actually be separated to different packages.
// See Development Flow of Components in Go for details.

// Although this approach is handy, there is some small overhead compared to a UDF implemented in the regular way.
// Most of such overhead comes from type checking and conversions.

// Functions passed to ConverGeneric need to satisfy some restrictions on
// form of their argument and return value types. Each restriction is described in the following subsections.
