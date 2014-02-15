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

func (box *Bool) SetAllowNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(bool))
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

// Panics unless box is Full
func (box *Bool) Get() bool {
	if box.status != Full {
		panic("Tried to get from box that was not full")
	}

	return box.value
}

func (box *Bool) Status() byte {
	return box.status
}

func (box *Bool) IsFull() bool {
	return box.status == Full
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

func (box *Float32) SetAllowNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(float32))
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

// Panics unless box is Full
func (box *Float32) Get() float32 {
	if box.status != Full {
		panic("Tried to get from box that was not full")
	}

	return box.value
}

func (box *Float32) Status() byte {
	return box.status
}

func (box *Float32) IsFull() bool {
	return box.status == Full
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

func (box *Float64) SetAllowNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(float64))
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

// Panics unless box is Full
func (box *Float64) Get() float64 {
	if box.status != Full {
		panic("Tried to get from box that was not full")
	}

	return box.value
}

func (box *Float64) Status() byte {
	return box.status
}

func (box *Float64) IsFull() bool {
	return box.status == Full
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

func (box *Int8) SetAllowNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(int8))
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

// Panics unless box is Full
func (box *Int8) Get() int8 {
	if box.status != Full {
		panic("Tried to get from box that was not full")
	}

	return box.value
}

func (box *Int8) Status() byte {
	return box.status
}

func (box *Int8) IsFull() bool {
	return box.status == Full
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

func (box *Int16) SetAllowNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(int16))
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

// Panics unless box is Full
func (box *Int16) Get() int16 {
	if box.status != Full {
		panic("Tried to get from box that was not full")
	}

	return box.value
}

func (box *Int16) Status() byte {
	return box.status
}

func (box *Int16) IsFull() bool {
	return box.status == Full
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

func (box *Int32) SetAllowNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(int32))
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

// Panics unless box is Full
func (box *Int32) Get() int32 {
	if box.status != Full {
		panic("Tried to get from box that was not full")
	}

	return box.value
}

func (box *Int32) Status() byte {
	return box.status
}

func (box *Int32) IsFull() bool {
	return box.status == Full
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

func (box *Int64) SetAllowNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(int64))
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

// Panics unless box is Full
func (box *Int64) Get() int64 {
	if box.status != Full {
		panic("Tried to get from box that was not full")
	}

	return box.value
}

func (box *Int64) Status() byte {
	return box.status
}

func (box *Int64) IsFull() bool {
	return box.status == Full
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

func (box *String) SetAllowNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(string))
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

// Panics unless box is Full
func (box *String) Get() string {
	if box.status != Full {
		panic("Tried to get from box that was not full")
	}

	return box.value
}

func (box *String) Status() byte {
	return box.status
}

func (box *String) IsFull() bool {
	return box.status == Full
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

func (box *Time) SetAllowNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(time.Time))
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

// Panics unless box is Full
func (box *Time) Get() time.Time {
	if box.status != Full {
		panic("Tried to get from box that was not full")
	}

	return box.value
}

func (box *Time) Status() byte {
	return box.status
}

func (box *Time) IsFull() bool {
	return box.status == Full
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

func (box *UInt8) SetAllowNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(uint8))
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

// Panics unless box is Full
func (box *UInt8) Get() uint8 {
	if box.status != Full {
		panic("Tried to get from box that was not full")
	}

	return box.value
}

func (box *UInt8) Status() byte {
	return box.status
}

func (box *UInt8) IsFull() bool {
	return box.status == Full
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

func (box *UInt16) SetAllowNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(uint16))
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

// Panics unless box is Full
func (box *UInt16) Get() uint16 {
	if box.status != Full {
		panic("Tried to get from box that was not full")
	}

	return box.value
}

func (box *UInt16) Status() byte {
	return box.status
}

func (box *UInt16) IsFull() bool {
	return box.status == Full
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

func (box *UInt32) SetAllowNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(uint32))
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

// Panics unless box is Full
func (box *UInt32) Get() uint32 {
	if box.status != Full {
		panic("Tried to get from box that was not full")
	}

	return box.value
}

func (box *UInt32) Status() byte {
	return box.status
}

func (box *UInt32) IsFull() bool {
	return box.status == Full
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

func (box *UInt64) SetAllowNil(v interface{}, nilStatus byte) {
	if v != nil {
		box.Set(v.(uint64))
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

// Panics unless box is Full
func (box *UInt64) Get() uint64 {
	if box.status != Full {
		panic("Tried to get from box that was not full")
	}

	return box.value
}

func (box *UInt64) Status() byte {
	return box.status
}

func (box *UInt64) IsFull() bool {
	return box.status == Full
}
