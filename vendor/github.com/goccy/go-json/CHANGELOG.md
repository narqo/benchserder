# v0.4.11 - 2021/4/3

* Improve decoder performance for interface type

# v0.4.10 - 2021/4/2

### Fix encoder

* Fixed a bug when encoding slice and map containing recursive structures
* Fixed a logic to determine if indirect reference

# v0.4.9 - 2021/3/29

### Add debug mode

If you use `json.MarshalWithOption(v, json.Debug())` and `panic` occurred in `go-json`, produces debug information to console.

### Support a new feature to compatible with encoding/json

- invalid UTF-8 is coerced to valid UTF-8 ( without performance down )

### Fix encoder

- Fixed handling of MarshalJSON of function type

### Fix decoding of slice of pointer type

If there is a pointer value, go-json will use it. (This behavior is necessary to achieve the ability to prioritize pre-filled values). However, since slices are reused internally, there was a bug that referred to the previous pointer value. Therefore, it is not necessary to refer to the pointer value in advance for the slice element, so we explicitly initialize slice element by `nil`.

# v0.4.8 - 2021/3/21

### Reduce memory usage at compile time

* go-json have used about 2GB of memory at compile time, but now it can compile with about less than 550MB.

### Fix any encoder's bug

* Add many test cases for encoder
* Fix composite type ( slice/array/map )
* Fix pointer types
* Fix encoding of MarshalJSON or MarshalText or json.Number type

### Refactor encoder

* Change package layout for reducing memory usage at compile
* Remove anonymous and only operation
* Remove root property from encodeCompileContext and opcode

### Fix CI

* Add Go 1.16
* Remove Go 1.13
* Fix `make cover` task

### Number/Delim/Token/RawMessage use the types defined in encoding/json by type alias

# v0.4.7 - 2021/02/22

### Fix decoder

* Fix decoding of deep recursive structure
* Fix decoding of embedded unexported pointer field
* Fix invalid test case
* Fix decoding of invalid value
* Fix decoding of prefilled value
* Fix not being able to return UnmarshalTypeError when it should be returned
* Fix decoding of null value
* Fix decoding of type of null string
* Use pre allocated pointer if exists it at decoding

### Reduce memory usage at compile

* Integrate int/int8/int16/int32/int64 and uint/uint8/uint16/uint32/uint64 operation to reduce memory usage at compile

### Remove unnecessary optype
