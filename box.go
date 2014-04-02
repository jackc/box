package box

import (
	"time"
)

const (
	Undefined = iota
	Unknown   = iota
	Empty     = iota
	Full      = iota
)

type Bool struct {
	value  bool
	status byte
}

func NewBool(v bool) (box Bool) {
	box.Set(v)
	return box
}

func (box *Bool) Set(v bool) {
	box.value = v
	box.status = Full
}

func (box *Bool) GetCoerceNil() interface{} {
	if box.status == Full {
		return box.value
	} else {
		return nil
	}
}

func (box *Bool) SetCoerceNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(bool))
	} else {
		box.status = nilStatus
	}
}

func (box *Bool) GetCoerceZero() bool {
	if box.status == Full {
		return box.value
	} else {
		var zero bool
		return zero
	}
}

func (box *Bool) SetCoerceZero(v bool, nilStatus byte) {
	var zero bool

	if v != zero {
		box.Set(v)
	} else {
		box.status = nilStatus
	}
}

func (box *Bool) SetUndefined() {
	box.status = Undefined
}

func (box *Bool) SetUnknown() {
	box.status = Unknown
}

func (box *Bool) SetEmpty() {
	box.status = Empty
}

func (box *Bool) MustGet() bool {
	if box.status != Full {
		panic("called MustGet on a box that was not full")
	}

	return box.value
}

func (box *Bool) Get() (bool, bool) {
	if box.status != Full {
		var zeroVal bool
		return zeroVal, false
	}

	return box.value, true
}

func (box *Bool) Status() byte {
	return box.status
}

type Float32 struct {
	value  float32
	status byte
}

func NewFloat32(v float32) (box Float32) {
	box.Set(v)
	return box
}

func (box *Float32) Set(v float32) {
	box.value = v
	box.status = Full
}

func (box *Float32) GetCoerceNil() interface{} {
	if box.status == Full {
		return box.value
	} else {
		return nil
	}
}

func (box *Float32) SetCoerceNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(float32))
	} else {
		box.status = nilStatus
	}
}

func (box *Float32) GetCoerceZero() float32 {
	if box.status == Full {
		return box.value
	} else {
		var zero float32
		return zero
	}
}

func (box *Float32) SetCoerceZero(v float32, nilStatus byte) {
	var zero float32

	if v != zero {
		box.Set(v)
	} else {
		box.status = nilStatus
	}
}

func (box *Float32) SetUndefined() {
	box.status = Undefined
}

func (box *Float32) SetUnknown() {
	box.status = Unknown
}

func (box *Float32) SetEmpty() {
	box.status = Empty
}

func (box *Float32) MustGet() float32 {
	if box.status != Full {
		panic("called MustGet on a box that was not full")
	}

	return box.value
}

func (box *Float32) Get() (float32, bool) {
	if box.status != Full {
		var zeroVal float32
		return zeroVal, false
	}

	return box.value, true
}

func (box *Float32) Status() byte {
	return box.status
}

type Float64 struct {
	value  float64
	status byte
}

func NewFloat64(v float64) (box Float64) {
	box.Set(v)
	return box
}

func (box *Float64) Set(v float64) {
	box.value = v
	box.status = Full
}

func (box *Float64) GetCoerceNil() interface{} {
	if box.status == Full {
		return box.value
	} else {
		return nil
	}
}

func (box *Float64) SetCoerceNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(float64))
	} else {
		box.status = nilStatus
	}
}

func (box *Float64) GetCoerceZero() float64 {
	if box.status == Full {
		return box.value
	} else {
		var zero float64
		return zero
	}
}

func (box *Float64) SetCoerceZero(v float64, nilStatus byte) {
	var zero float64

	if v != zero {
		box.Set(v)
	} else {
		box.status = nilStatus
	}
}

func (box *Float64) SetUndefined() {
	box.status = Undefined
}

func (box *Float64) SetUnknown() {
	box.status = Unknown
}

func (box *Float64) SetEmpty() {
	box.status = Empty
}

func (box *Float64) MustGet() float64 {
	if box.status != Full {
		panic("called MustGet on a box that was not full")
	}

	return box.value
}

func (box *Float64) Get() (float64, bool) {
	if box.status != Full {
		var zeroVal float64
		return zeroVal, false
	}

	return box.value, true
}

func (box *Float64) Status() byte {
	return box.status
}

type Int8 struct {
	value  int8
	status byte
}

func NewInt8(v int8) (box Int8) {
	box.Set(v)
	return box
}

func (box *Int8) Set(v int8) {
	box.value = v
	box.status = Full
}

func (box *Int8) GetCoerceNil() interface{} {
	if box.status == Full {
		return box.value
	} else {
		return nil
	}
}

func (box *Int8) SetCoerceNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(int8))
	} else {
		box.status = nilStatus
	}
}

func (box *Int8) GetCoerceZero() int8 {
	if box.status == Full {
		return box.value
	} else {
		var zero int8
		return zero
	}
}

func (box *Int8) SetCoerceZero(v int8, nilStatus byte) {
	var zero int8

	if v != zero {
		box.Set(v)
	} else {
		box.status = nilStatus
	}
}

func (box *Int8) SetUndefined() {
	box.status = Undefined
}

func (box *Int8) SetUnknown() {
	box.status = Unknown
}

func (box *Int8) SetEmpty() {
	box.status = Empty
}

func (box *Int8) MustGet() int8 {
	if box.status != Full {
		panic("called MustGet on a box that was not full")
	}

	return box.value
}

func (box *Int8) Get() (int8, bool) {
	if box.status != Full {
		var zeroVal int8
		return zeroVal, false
	}

	return box.value, true
}

func (box *Int8) Status() byte {
	return box.status
}

type Int16 struct {
	value  int16
	status byte
}

func NewInt16(v int16) (box Int16) {
	box.Set(v)
	return box
}

func (box *Int16) Set(v int16) {
	box.value = v
	box.status = Full
}

func (box *Int16) GetCoerceNil() interface{} {
	if box.status == Full {
		return box.value
	} else {
		return nil
	}
}

func (box *Int16) SetCoerceNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(int16))
	} else {
		box.status = nilStatus
	}
}

func (box *Int16) GetCoerceZero() int16 {
	if box.status == Full {
		return box.value
	} else {
		var zero int16
		return zero
	}
}

func (box *Int16) SetCoerceZero(v int16, nilStatus byte) {
	var zero int16

	if v != zero {
		box.Set(v)
	} else {
		box.status = nilStatus
	}
}

func (box *Int16) SetUndefined() {
	box.status = Undefined
}

func (box *Int16) SetUnknown() {
	box.status = Unknown
}

func (box *Int16) SetEmpty() {
	box.status = Empty
}

func (box *Int16) MustGet() int16 {
	if box.status != Full {
		panic("called MustGet on a box that was not full")
	}

	return box.value
}

func (box *Int16) Get() (int16, bool) {
	if box.status != Full {
		var zeroVal int16
		return zeroVal, false
	}

	return box.value, true
}

func (box *Int16) Status() byte {
	return box.status
}

type Int32 struct {
	value  int32
	status byte
}

func NewInt32(v int32) (box Int32) {
	box.Set(v)
	return box
}

func (box *Int32) Set(v int32) {
	box.value = v
	box.status = Full
}

func (box *Int32) GetCoerceNil() interface{} {
	if box.status == Full {
		return box.value
	} else {
		return nil
	}
}

func (box *Int32) SetCoerceNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(int32))
	} else {
		box.status = nilStatus
	}
}

func (box *Int32) GetCoerceZero() int32 {
	if box.status == Full {
		return box.value
	} else {
		var zero int32
		return zero
	}
}

func (box *Int32) SetCoerceZero(v int32, nilStatus byte) {
	var zero int32

	if v != zero {
		box.Set(v)
	} else {
		box.status = nilStatus
	}
}

func (box *Int32) SetUndefined() {
	box.status = Undefined
}

func (box *Int32) SetUnknown() {
	box.status = Unknown
}

func (box *Int32) SetEmpty() {
	box.status = Empty
}

func (box *Int32) MustGet() int32 {
	if box.status != Full {
		panic("called MustGet on a box that was not full")
	}

	return box.value
}

func (box *Int32) Get() (int32, bool) {
	if box.status != Full {
		var zeroVal int32
		return zeroVal, false
	}

	return box.value, true
}

func (box *Int32) Status() byte {
	return box.status
}

type Int64 struct {
	value  int64
	status byte
}

func NewInt64(v int64) (box Int64) {
	box.Set(v)
	return box
}

func (box *Int64) Set(v int64) {
	box.value = v
	box.status = Full
}

func (box *Int64) GetCoerceNil() interface{} {
	if box.status == Full {
		return box.value
	} else {
		return nil
	}
}

func (box *Int64) SetCoerceNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(int64))
	} else {
		box.status = nilStatus
	}
}

func (box *Int64) GetCoerceZero() int64 {
	if box.status == Full {
		return box.value
	} else {
		var zero int64
		return zero
	}
}

func (box *Int64) SetCoerceZero(v int64, nilStatus byte) {
	var zero int64

	if v != zero {
		box.Set(v)
	} else {
		box.status = nilStatus
	}
}

func (box *Int64) SetUndefined() {
	box.status = Undefined
}

func (box *Int64) SetUnknown() {
	box.status = Unknown
}

func (box *Int64) SetEmpty() {
	box.status = Empty
}

func (box *Int64) MustGet() int64 {
	if box.status != Full {
		panic("called MustGet on a box that was not full")
	}

	return box.value
}

func (box *Int64) Get() (int64, bool) {
	if box.status != Full {
		var zeroVal int64
		return zeroVal, false
	}

	return box.value, true
}

func (box *Int64) Status() byte {
	return box.status
}

type String struct {
	value  string
	status byte
}

func NewString(v string) (box String) {
	box.Set(v)
	return box
}

func (box *String) Set(v string) {
	box.value = v
	box.status = Full
}

func (box *String) GetCoerceNil() interface{} {
	if box.status == Full {
		return box.value
	} else {
		return nil
	}
}

func (box *String) SetCoerceNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(string))
	} else {
		box.status = nilStatus
	}
}

func (box *String) GetCoerceZero() string {
	if box.status == Full {
		return box.value
	} else {
		var zero string
		return zero
	}
}

func (box *String) SetCoerceZero(v string, nilStatus byte) {
	var zero string

	if v != zero {
		box.Set(v)
	} else {
		box.status = nilStatus
	}
}

func (box *String) SetUndefined() {
	box.status = Undefined
}

func (box *String) SetUnknown() {
	box.status = Unknown
}

func (box *String) SetEmpty() {
	box.status = Empty
}

func (box *String) MustGet() string {
	if box.status != Full {
		panic("called MustGet on a box that was not full")
	}

	return box.value
}

func (box *String) Get() (string, bool) {
	if box.status != Full {
		var zeroVal string
		return zeroVal, false
	}

	return box.value, true
}

func (box *String) Status() byte {
	return box.status
}

type Time struct {
	value  time.Time
	status byte
}

func NewTime(v time.Time) (box Time) {
	box.Set(v)
	return box
}

func (box *Time) Set(v time.Time) {
	box.value = v
	box.status = Full
}

func (box *Time) GetCoerceNil() interface{} {
	if box.status == Full {
		return box.value
	} else {
		return nil
	}
}

func (box *Time) SetCoerceNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(time.Time))
	} else {
		box.status = nilStatus
	}
}

func (box *Time) GetCoerceZero() time.Time {
	if box.status == Full {
		return box.value
	} else {
		var zero time.Time
		return zero
	}
}

func (box *Time) SetCoerceZero(v time.Time, nilStatus byte) {
	var zero time.Time

	if v != zero {
		box.Set(v)
	} else {
		box.status = nilStatus
	}
}

func (box *Time) SetUndefined() {
	box.status = Undefined
}

func (box *Time) SetUnknown() {
	box.status = Unknown
}

func (box *Time) SetEmpty() {
	box.status = Empty
}

func (box *Time) MustGet() time.Time {
	if box.status != Full {
		panic("called MustGet on a box that was not full")
	}

	return box.value
}

func (box *Time) Get() (time.Time, bool) {
	if box.status != Full {
		var zeroVal time.Time
		return zeroVal, false
	}

	return box.value, true
}

func (box *Time) Status() byte {
	return box.status
}

type UInt8 struct {
	value  uint8
	status byte
}

func NewUInt8(v uint8) (box UInt8) {
	box.Set(v)
	return box
}

func (box *UInt8) Set(v uint8) {
	box.value = v
	box.status = Full
}

func (box *UInt8) GetCoerceNil() interface{} {
	if box.status == Full {
		return box.value
	} else {
		return nil
	}
}

func (box *UInt8) SetCoerceNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(uint8))
	} else {
		box.status = nilStatus
	}
}

func (box *UInt8) GetCoerceZero() uint8 {
	if box.status == Full {
		return box.value
	} else {
		var zero uint8
		return zero
	}
}

func (box *UInt8) SetCoerceZero(v uint8, nilStatus byte) {
	var zero uint8

	if v != zero {
		box.Set(v)
	} else {
		box.status = nilStatus
	}
}

func (box *UInt8) SetUndefined() {
	box.status = Undefined
}

func (box *UInt8) SetUnknown() {
	box.status = Unknown
}

func (box *UInt8) SetEmpty() {
	box.status = Empty
}

func (box *UInt8) MustGet() uint8 {
	if box.status != Full {
		panic("called MustGet on a box that was not full")
	}

	return box.value
}

func (box *UInt8) Get() (uint8, bool) {
	if box.status != Full {
		var zeroVal uint8
		return zeroVal, false
	}

	return box.value, true
}

func (box *UInt8) Status() byte {
	return box.status
}

type UInt16 struct {
	value  uint16
	status byte
}

func NewUInt16(v uint16) (box UInt16) {
	box.Set(v)
	return box
}

func (box *UInt16) Set(v uint16) {
	box.value = v
	box.status = Full
}

func (box *UInt16) GetCoerceNil() interface{} {
	if box.status == Full {
		return box.value
	} else {
		return nil
	}
}

func (box *UInt16) SetCoerceNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(uint16))
	} else {
		box.status = nilStatus
	}
}

func (box *UInt16) GetCoerceZero() uint16 {
	if box.status == Full {
		return box.value
	} else {
		var zero uint16
		return zero
	}
}

func (box *UInt16) SetCoerceZero(v uint16, nilStatus byte) {
	var zero uint16

	if v != zero {
		box.Set(v)
	} else {
		box.status = nilStatus
	}
}

func (box *UInt16) SetUndefined() {
	box.status = Undefined
}

func (box *UInt16) SetUnknown() {
	box.status = Unknown
}

func (box *UInt16) SetEmpty() {
	box.status = Empty
}

func (box *UInt16) MustGet() uint16 {
	if box.status != Full {
		panic("called MustGet on a box that was not full")
	}

	return box.value
}

func (box *UInt16) Get() (uint16, bool) {
	if box.status != Full {
		var zeroVal uint16
		return zeroVal, false
	}

	return box.value, true
}

func (box *UInt16) Status() byte {
	return box.status
}

type UInt32 struct {
	value  uint32
	status byte
}

func NewUInt32(v uint32) (box UInt32) {
	box.Set(v)
	return box
}

func (box *UInt32) Set(v uint32) {
	box.value = v
	box.status = Full
}

func (box *UInt32) GetCoerceNil() interface{} {
	if box.status == Full {
		return box.value
	} else {
		return nil
	}
}

func (box *UInt32) SetCoerceNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(uint32))
	} else {
		box.status = nilStatus
	}
}

func (box *UInt32) GetCoerceZero() uint32 {
	if box.status == Full {
		return box.value
	} else {
		var zero uint32
		return zero
	}
}

func (box *UInt32) SetCoerceZero(v uint32, nilStatus byte) {
	var zero uint32

	if v != zero {
		box.Set(v)
	} else {
		box.status = nilStatus
	}
}

func (box *UInt32) SetUndefined() {
	box.status = Undefined
}

func (box *UInt32) SetUnknown() {
	box.status = Unknown
}

func (box *UInt32) SetEmpty() {
	box.status = Empty
}

func (box *UInt32) MustGet() uint32 {
	if box.status != Full {
		panic("called MustGet on a box that was not full")
	}

	return box.value
}

func (box *UInt32) Get() (uint32, bool) {
	if box.status != Full {
		var zeroVal uint32
		return zeroVal, false
	}

	return box.value, true
}

func (box *UInt32) Status() byte {
	return box.status
}

type UInt64 struct {
	value  uint64
	status byte
}

func NewUInt64(v uint64) (box UInt64) {
	box.Set(v)
	return box
}

func (box *UInt64) Set(v uint64) {
	box.value = v
	box.status = Full
}

func (box *UInt64) GetCoerceNil() interface{} {
	if box.status == Full {
		return box.value
	} else {
		return nil
	}
}

func (box *UInt64) SetCoerceNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(uint64))
	} else {
		box.status = nilStatus
	}
}

func (box *UInt64) GetCoerceZero() uint64 {
	if box.status == Full {
		return box.value
	} else {
		var zero uint64
		return zero
	}
}

func (box *UInt64) SetCoerceZero(v uint64, nilStatus byte) {
	var zero uint64

	if v != zero {
		box.Set(v)
	} else {
		box.status = nilStatus
	}
}

func (box *UInt64) SetUndefined() {
	box.status = Undefined
}

func (box *UInt64) SetUnknown() {
	box.status = Unknown
}

func (box *UInt64) SetEmpty() {
	box.status = Empty
}

func (box *UInt64) MustGet() uint64 {
	if box.status != Full {
		panic("called MustGet on a box that was not full")
	}

	return box.value
}

func (box *UInt64) Get() (uint64, bool) {
	if box.status != Full {
		var zeroVal uint64
		return zeroVal, false
	}

	return box.value, true
}

func (box *UInt64) Status() byte {
	return box.status
}
