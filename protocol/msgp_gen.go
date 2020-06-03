package protocol

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"github.com/algorand/msgp/msgp"
)

// The following msgp objects are implemented in this file:
// ConsensusVersion
//         |-----> MarshalMsg
//         |-----> CanMarshalMsg
//         |-----> (*) UnmarshalMsg
//         |-----> (*) CanUnmarshalMsg
//         |-----> Msgsize
//         |-----> MsgIsZero
//
// Error
//   |-----> MarshalMsg
//   |-----> CanMarshalMsg
//   |-----> (*) UnmarshalMsg
//   |-----> (*) CanUnmarshalMsg
//   |-----> Msgsize
//   |-----> MsgIsZero
//
// HashID
//    |-----> MarshalMsg
//    |-----> CanMarshalMsg
//    |-----> (*) UnmarshalMsg
//    |-----> (*) CanUnmarshalMsg
//    |-----> Msgsize
//    |-----> MsgIsZero
//
// NetworkID
//     |-----> MarshalMsg
//     |-----> CanMarshalMsg
//     |-----> (*) UnmarshalMsg
//     |-----> (*) CanUnmarshalMsg
//     |-----> Msgsize
//     |-----> MsgIsZero
//
// Tag
//  |-----> MarshalMsg
//  |-----> CanMarshalMsg
//  |-----> (*) UnmarshalMsg
//  |-----> (*) CanUnmarshalMsg
//  |-----> Msgsize
//  |-----> MsgIsZero
//
// TxType
//    |-----> MarshalMsg
//    |-----> CanMarshalMsg
//    |-----> (*) UnmarshalMsg
//    |-----> (*) CanUnmarshalMsg
//    |-----> Msgsize
//    |-----> MsgIsZero
//

// MarshalMsg implements msgp.Marshaler
func (z ConsensusVersion) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

func (_ ConsensusVersion) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(ConsensusVersion)
	if !ok {
		_, ok = (z).(*ConsensusVersion)
	}
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ConsensusVersion) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = ConsensusVersion(zb0001)
	}
	o = bts
	return
}

func (_ *ConsensusVersion) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*ConsensusVersion)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ConsensusVersion) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}

// MsgIsZero returns whether this is a zero value
func (z ConsensusVersion) MsgIsZero() bool {
	return z == ""
}

// MarshalMsg implements msgp.Marshaler
func (z Error) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

func (_ Error) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(Error)
	if !ok {
		_, ok = (z).(*Error)
	}
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Error) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = Error(zb0001)
	}
	o = bts
	return
}

func (_ *Error) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Error)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Error) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}

// MsgIsZero returns whether this is a zero value
func (z Error) MsgIsZero() bool {
	return z == ""
}

// MarshalMsg implements msgp.Marshaler
func (z HashID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

func (_ HashID) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(HashID)
	if !ok {
		_, ok = (z).(*HashID)
	}
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *HashID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = HashID(zb0001)
	}
	o = bts
	return
}

func (_ *HashID) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*HashID)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z HashID) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}

// MsgIsZero returns whether this is a zero value
func (z HashID) MsgIsZero() bool {
	return z == ""
}

// MarshalMsg implements msgp.Marshaler
func (z NetworkID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

func (_ NetworkID) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(NetworkID)
	if !ok {
		_, ok = (z).(*NetworkID)
	}
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NetworkID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = NetworkID(zb0001)
	}
	o = bts
	return
}

func (_ *NetworkID) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*NetworkID)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z NetworkID) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}

// MsgIsZero returns whether this is a zero value
func (z NetworkID) MsgIsZero() bool {
	return z == ""
}

// MarshalMsg implements msgp.Marshaler
func (z Tag) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

func (_ Tag) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(Tag)
	if !ok {
		_, ok = (z).(*Tag)
	}
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Tag) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = Tag(zb0001)
	}
	o = bts
	return
}

func (_ *Tag) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Tag)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Tag) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}

// MsgIsZero returns whether this is a zero value
func (z Tag) MsgIsZero() bool {
	return z == ""
}

// MarshalMsg implements msgp.Marshaler
func (z TxType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

func (_ TxType) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(TxType)
	if !ok {
		_, ok = (z).(*TxType)
	}
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TxType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = TxType(zb0001)
	}
	o = bts
	return
}

func (_ *TxType) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*TxType)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z TxType) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}

// MsgIsZero returns whether this is a zero value
func (z TxType) MsgIsZero() bool {
	return z == ""
}
