package Node_Buffer_Immutable

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"

	"gopurs/output/gopurs_runtime"
)

func nodeBufferImmutable_getBytes(val interface{}) []byte {
    v := val.(gopurs_runtime.Value)
	return (*(*any)(v.UnsafePtr)).([]byte)
}

func nodeBufferImmutable_boxBytes(b []byte) gopurs_runtime.Value {
	return gopurs_runtime.Any(b)
}

var ShowImpl = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	b := nodeBufferImmutable_getBytes(a)
	if len(b) > 50 {
		return gopurs_runtime.Str(fmt.Sprintf("<Buffer % x ...>", b[:50]))
	}
	return gopurs_runtime.Str(fmt.Sprintf("<Buffer % x>", b))
})

func EqImpl(a interface{}, b interface{}) interface{} {
	b1 := nodeBufferImmutable_getBytes(a)
	b2 := nodeBufferImmutable_getBytes(b)
	return bytes.Equal(b1, b2)
}

func CompareImpl(a interface{}, b interface{}) interface{} {
	cmp := bytes.Compare(nodeBufferImmutable_getBytes(a), nodeBufferImmutable_getBytes(b))
	if cmp < 0 {
		return -1
	} else if cmp > 0 {
		return 1
	}
	return 0
}

func ComparePartsImpl(src interface{}, target interface{}, targetStart interface{}, targetEnd interface{}, sourceStart interface{}, sourceEnd interface{}) interface{} {
	s := nodeBufferImmutable_getBytes(src)
	t := nodeBufferImmutable_getBytes(target)
	ts := int(targetStart.(gopurs_runtime.Value).IntVal)
	te := int(targetEnd.(gopurs_runtime.Value).IntVal)
	ss := int(sourceStart.(gopurs_runtime.Value).IntVal)
	se := int(sourceEnd.(gopurs_runtime.Value).IntVal)
	if ts < 0 { ts = 0 }
	if te > len(t) { te = len(t) }
	if ss < 0 { ss = 0 }
	if se > len(s) { se = len(s) }
	if ts > te { ts = te }
	if ss > se { ss = se }
	
	cmp := bytes.Compare(s[ss:se], t[ts:te])
	if cmp < 0 { return -1 }
	if cmp > 0 { return 1 }
	return 0
}

var Alloc = gopurs_runtime.Func(func(size gopurs_runtime.Value) gopurs_runtime.Value {
	sz := int(size.IntVal)
	return nodeBufferImmutable_boxBytes(make([]byte, sz))
})

var FromArray = gopurs_runtime.Func(func(octets gopurs_runtime.Value) gopurs_runtime.Value {
	arr := *(*[]gopurs_runtime.Value)(octets.UnsafePtr)
	b := make([]byte, len(arr))
	for i, v := range arr {
		b[i] = byte(v.IntVal)
	}
	return nodeBufferImmutable_boxBytes(b)
})

var Size = gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
	b := nodeBufferImmutable_getBytes(buff)
	return gopurs_runtime.Int(int64(len(b)))
})

var ToArray = gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
	b := nodeBufferImmutable_getBytes(buff)
	arr := make([]gopurs_runtime.Value, len(b))
	for i, v := range b {
		arr[i] = gopurs_runtime.Int(int64(v))
	}
	return gopurs_runtime.Array(arr)
})

var ToArrayBuffer = gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
	return buff
})

var FromArrayBuffer = gopurs_runtime.Func(func(ab gopurs_runtime.Value) gopurs_runtime.Value {
	return ab
})

func FromStringImpl(str interface{}, encoding interface{}) interface{} {
	s := gopurs_runtime.StrValue(str.(gopurs_runtime.Value))
	enc := gopurs_runtime.StrValue(encoding.(gopurs_runtime.Value))
	var b []byte
	switch enc {
	case "hex":
		b, _ = hex.DecodeString(s)
	case "base64":
		b, _ = base64.StdEncoding.DecodeString(s)
	default:
		b = []byte(s)
	}
	return nodeBufferImmutable_boxBytes(b)
}

func ReadImpl(ty interface{}, offset interface{}, buf interface{}) interface{} {
	b := nodeBufferImmutable_getBytes(buf)
	off := int(offset.(gopurs_runtime.Value).IntVal)
	t := gopurs_runtime.StrValue(ty.(gopurs_runtime.Value))
	var n float64
	switch t {
	case "UInt8":
		n = float64(b[off])
	case "UInt16LE":
		n = float64(binary.LittleEndian.Uint16(b[off:]))
	case "UInt16BE":
		n = float64(binary.BigEndian.Uint16(b[off:]))
	case "UInt32LE":
		n = float64(binary.LittleEndian.Uint32(b[off:]))
	case "UInt32BE":
		n = float64(binary.BigEndian.Uint32(b[off:]))
	case "Int8":
		n = float64(int8(b[off]))
	case "Int16LE":
		n = float64(int16(binary.LittleEndian.Uint16(b[off:])))
	case "Int16BE":
		n = float64(int16(binary.BigEndian.Uint16(b[off:])))
	case "Int32LE":
		n = float64(int32(binary.LittleEndian.Uint32(b[off:])))
	case "Int32BE":
		n = float64(int32(binary.BigEndian.Uint32(b[off:])))
	case "FloatLE":
		n = float64(math.Float32frombits(binary.LittleEndian.Uint32(b[off:])))
	case "FloatBE":
		n = float64(math.Float32frombits(binary.BigEndian.Uint32(b[off:])))
	case "DoubleLE":
		n = math.Float64frombits(binary.LittleEndian.Uint64(b[off:]))
	case "DoubleBE":
		n = math.Float64frombits(binary.BigEndian.Uint64(b[off:]))
	}
	return n
}

func ReadStringImpl(enc interface{}, start interface{}, end interface{}, buff interface{}) interface{} {
	b := nodeBufferImmutable_getBytes(buff)
	s := int(start.(gopurs_runtime.Value).IntVal)
	e := int(end.(gopurs_runtime.Value).IntVal)
	if s < 0 { s = 0 }
	if e > len(b) { e = len(b) }
	if s > e { s = e }
	sub := b[s:e]
	encoding := gopurs_runtime.StrValue(enc.(gopurs_runtime.Value))
	switch encoding {
	case "hex":
		return hex.EncodeToString(sub)
	case "base64":
		return base64.StdEncoding.EncodeToString(sub)
	default:
		return string(sub)
	}
}

func GetAtOffsetImpl(offset interface{}, buff interface{}) interface{} {
	b := nodeBufferImmutable_getBytes(buff)
	off := int(offset.(gopurs_runtime.Value).IntVal)
	if off < 0 || off >= len(b) {
		return gopurs_runtime.Box(any(nil))
	}
	return gopurs_runtime.Box(gopurs_runtime.Int(int64(b[off])))
}

func ToStringImpl(enc interface{}, buff interface{}) interface{} {
	b := nodeBufferImmutable_getBytes(buff)
	encoding := gopurs_runtime.StrValue(enc.(gopurs_runtime.Value))
	switch encoding {
	case "hex":
		return hex.EncodeToString(b)
	case "base64":
		return base64.StdEncoding.EncodeToString(b)
	default:
		return string(b)
	}
}

func ToStringSubImpl(enc interface{}, start interface{}, end interface{}, buff interface{}) interface{} {
	b := nodeBufferImmutable_getBytes(buff)
	s := int(start.(gopurs_runtime.Value).IntVal)
	e := int(end.(gopurs_runtime.Value).IntVal)
	if s < 0 { s = 0 }
	if e > len(b) { e = len(b) }
	if s > e { s = e }
	sub := b[s:e]
	encoding := gopurs_runtime.StrValue(enc.(gopurs_runtime.Value))
	switch encoding {
	case "hex":
		return hex.EncodeToString(sub)
	case "base64":
		return base64.StdEncoding.EncodeToString(sub)
	default:
		return string(sub)
	}
}

func SliceImpl(start interface{}, end interface{}, buff interface{}) interface{} {
	b := nodeBufferImmutable_getBytes(buff)
	s := int(start.(gopurs_runtime.Value).IntVal)
	e := int(end.(gopurs_runtime.Value).IntVal)
	if s < 0 { s = 0 }
	if e > len(b) { e = len(b) }
	if s > e { s = e }
	return nodeBufferImmutable_boxBytes(b[s:e])
}

var Concat = gopurs_runtime.Func(func(buffs gopurs_runtime.Value) gopurs_runtime.Value {
	arr := *(*[]gopurs_runtime.Value)(buffs.UnsafePtr)
	var bs [][]byte
	for _, v := range arr {
		bs = append(bs, nodeBufferImmutable_getBytes(v))
	}
	return nodeBufferImmutable_boxBytes(bytes.Join(bs, nil))
})

func ConcatToLength(buffs interface{}, totalLength interface{}) interface{} {
	arr := *(*[]gopurs_runtime.Value)(buffs.(gopurs_runtime.Value).UnsafePtr)
	total := int(totalLength.(gopurs_runtime.Value).IntVal)
	b := make([]byte, 0, total)
	for _, v := range arr {
		b = append(b, nodeBufferImmutable_getBytes(v)...)
		if len(b) >= total {
			break
		}
	}
	if len(b) > total {
		b = b[:total]
	}
	return nodeBufferImmutable_boxBytes(b)
}
