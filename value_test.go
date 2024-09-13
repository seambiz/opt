package opt

import (
	"fmt"
	"reflect"
	"testing"
)

type (
	customInt       int
	customInt8      int8
	customInt16     int16
	customInt32     int32
	customInt64     int64
	customUint      uint
	customUint8     uint8
	customUint16    uint16
	customUint32    uint32
	customUint64    uint64
	customFloat32   float32
	customFloat64   float64
	customBool      bool
	customByteSlice []byte
	customString    string
	customStruct    struct{}
)

func TestToDriverValue(t *testing.T) {
	s := "test"
	var ss *string
	u := uint(1)

	doTest[int64](t, customInt(0))
	doTest[int64](t, customInt8(0))
	doTest[int64](t, customInt16(0))
	doTest[int64](t, customInt32(0))
	doTest[int64](t, customInt64(0))
	doTest[int64](t, customUint(0))
	doTest[int64](t, customUint8(0))
	doTest[int64](t, customUint16(0))
	doTest[int64](t, customUint32(0))
	doTest[int64](t, customUint64(0))
	doTest[float64](t, customFloat32(0))
	doTest[float64](t, customFloat64(0))
	doTest[bool](t, customBool(false))
	doTest[[]byte](t, customByteSlice(""))
	doTest[string](t, customString(""))
	doTest[customStruct](t, customStruct{})
	doTest[string](t, &s)
	doTest[any](t, ss)
	doTest[any, *string](t, nil)
	doTest[any, *uint](t, nil)
	doTest[int64](t, &u)
}

func doTest[E any, T any](t *testing.T, v T) {
	t.Helper()

	t.Run(fmt.Sprintf("%T", v), func(t *testing.T) {
		drVal, err := ToDriverValue(v)
		if err != nil {
			t.Fatalf("error getting driver.Value: %v", err)
		}

		var zero E
		if reflect.TypeOf(drVal) != reflect.TypeOf(zero) {
			t.Fatalf("Expected value type %T but got %T (%v - %v)", zero, drVal, zero, drVal)
		}

		var vv interface{}
		vv = v
		if reflect.ValueOf(v).Kind() == reflect.Pointer {
			if reflect.ValueOf(v).IsNil() {
				vv = nil
			} else {
				vv = reflect.ValueOf(v).Elem()
			}
		}
		if fmt.Sprintf("%v", drVal) != fmt.Sprintf("%v", vv) {
			t.Fatalf("Expected value %v but got %v", drVal, vv)
		}
	})
}
