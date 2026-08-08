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

func getBytes(val gopurs_runtime.Value) []byte {
	return (*(*any)(val.UnsafePtr)).([]byte)
}

func boxBytes(b []byte) gopurs_runtime.Value {
	return gopurs_runtime.Any(b)
}

var ShowImpl = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	b := getBytes(a)
	if len(b) > 50 {
		return gopurs_runtime.Str(fmt.Sprintf("<Buffer % x ...>", b[:50]))
	}
	return gopurs_runtime.Str(fmt.Sprintf("<Buffer % x>", b))
})

var EqImpl = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
		b1 := getBytes(a)
		b2 := getBytes(b)
		return gopurs_runtime.Bool(bytes.Equal(b1, b2))
	})
})

var CompareImpl = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
		cmp := bytes.Compare(getBytes(a), getBytes(b))
		if cmp < 0 {
			return gopurs_runtime.Int(-1)
		} else if cmp > 0 {
			return gopurs_runtime.Int(1)
		}
		return gopurs_runtime.Int(0)
	})
})

var ComparePartsImpl = gopurs_runtime.Func(func(src gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(target gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(targetStart gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(targetEnd gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(sourceStart gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(sourceEnd gopurs_runtime.Value) gopurs_runtime.Value {
						s := getBytes(src)
						t := getBytes(target)
						ts := int(targetStart.IntVal)
						te := int(targetEnd.IntVal)
						ss := int(sourceStart.IntVal)
						se := int(sourceEnd.IntVal)
						if ts < 0 { ts = 0 }
						if te > len(t) { te = len(t) }
						if ss < 0 { ss = 0 }
						if se > len(s) { se = len(s) }
						if ts > te { ts = te }
						if ss > se { ss = se }
						
						cmp := bytes.Compare(s[ss:se], t[ts:te])
						if cmp < 0 { return gopurs_runtime.Int(-1) }
						if cmp > 0 { return gopurs_runtime.Int(1) }
						return gopurs_runtime.Int(0)
					})
				})
			})
		})
	})
})

var Alloc = gopurs_runtime.Func(func(size gopurs_runtime.Value) gopurs_runtime.Value {
	sz := int(size.IntVal)
	return boxBytes(make([]byte, sz))
})

var FromArray = gopurs_runtime.Func(func(octets gopurs_runtime.Value) gopurs_runtime.Value {
	arr := *(*[]gopurs_runtime.Value)(octets.UnsafePtr)
	b := make([]byte, len(arr))
	for i, v := range arr {
		b[i] = byte(v.IntVal)
	}
	return boxBytes(b)
})

var Size = gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
	b := getBytes(buff)
	return gopurs_runtime.Int(int64(len(b)))
})

var ToArray = gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
	b := getBytes(buff)
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

var FromStringImpl = gopurs_runtime.Func(func(str gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(encoding gopurs_runtime.Value) gopurs_runtime.Value {
		s := *(*string)(str.UnsafePtr)
		enc := *(*string)(encoding.UnsafePtr)
		var b []byte
		switch enc {
		case "hex":
			b, _ = hex.DecodeString(s)
		case "base64":
			b, _ = base64.StdEncoding.DecodeString(s)
		default:
			b = []byte(s)
		}
		return boxBytes(b)
	})
})

var ReadImpl = gopurs_runtime.Func(func(ty gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(offset gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(buf gopurs_runtime.Value) gopurs_runtime.Value {
			b := getBytes(buf)
			off := int(offset.IntVal)
			t := *(*string)(ty.UnsafePtr)
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
			return gopurs_runtime.Float(n)
		})
	})
})

var ReadStringImpl = gopurs_runtime.Func(func(enc gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(start gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(end gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
				b := getBytes(buff)
				s := int(start.IntVal)
				e := int(end.IntVal)
				if s < 0 { s = 0 }
				if e > len(b) { e = len(b) }
				if s > e { s = e }
				sub := b[s:e]
				encoding := *(*string)(enc.UnsafePtr)
				switch encoding {
				case "hex":
					return gopurs_runtime.Str(hex.EncodeToString(sub))
				case "base64":
					return gopurs_runtime.Str(base64.StdEncoding.EncodeToString(sub))
				default:
					return gopurs_runtime.Str(string(sub))
				}
			})
		})
	})
})

var GetAtOffsetImpl = gopurs_runtime.Func(func(offset gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
		b := getBytes(buff)
		off := int(offset.IntVal)
		if off < 0 || off >= len(b) {
			return gopurs_runtime.Box(any(nil)) // Nullable.Null
		}
		// Because gopurs-nullable has `NotNull(x interface{})`, its type in Go is just interface{}.
		// In gopurs, Box wraps it in `Value`. So we box our Value into any, and Box boxes that again, 
		// but `Box` simplifies any(Value) to just Value. 
		// Wait, `NotNull` returns Box(val). So we can just return Box(gopurs_runtime.Int(int64(b[off]))).
		return gopurs_runtime.Box(gopurs_runtime.Int(int64(b[off])))
	})
})

var ToStringImpl = gopurs_runtime.Func(func(enc gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
		b := getBytes(buff)
		encoding := *(*string)(enc.UnsafePtr)
		switch encoding {
		case "hex":
			return gopurs_runtime.Str(hex.EncodeToString(b))
		case "base64":
			return gopurs_runtime.Str(base64.StdEncoding.EncodeToString(b))
		default:
			return gopurs_runtime.Str(string(b))
		}
	})
})

var ToStringSubImpl = gopurs_runtime.Func(func(enc gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(start gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(end gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
				b := getBytes(buff)
				s := int(start.IntVal)
				e := int(end.IntVal)
				if s < 0 { s = 0 }
				if e > len(b) { e = len(b) }
				if s > e { s = e }
				sub := b[s:e]
				encoding := *(*string)(enc.UnsafePtr)
				switch encoding {
				case "hex":
					return gopurs_runtime.Str(hex.EncodeToString(sub))
				case "base64":
					return gopurs_runtime.Str(base64.StdEncoding.EncodeToString(sub))
				default:
					return gopurs_runtime.Str(string(sub))
				}
			})
		})
	})
})

var SliceImpl = gopurs_runtime.Func(func(start gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(end gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
			b := getBytes(buff)
			s := int(start.IntVal)
			e := int(end.IntVal)
			if s < 0 { s = 0 }
			if e > len(b) { e = len(b) }
			if s > e { s = e }
			return boxBytes(b[s:e])
		})
	})
})

var Concat = gopurs_runtime.Func(func(buffs gopurs_runtime.Value) gopurs_runtime.Value {
	arr := *(*[]gopurs_runtime.Value)(buffs.UnsafePtr)
	var bs [][]byte
	for _, v := range arr {
		bs = append(bs, getBytes(v))
	}
	return boxBytes(bytes.Join(bs, nil))
})

var ConcatToLength = gopurs_runtime.Func(func(buffs gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(totalLength gopurs_runtime.Value) gopurs_runtime.Value {
		arr := *(*[]gopurs_runtime.Value)(buffs.UnsafePtr)
		total := int(totalLength.IntVal)
		b := make([]byte, 0, total)
		for _, v := range arr {
			b = append(b, getBytes(v)...)
			if len(b) >= total {
				break
			}
		}
		if len(b) > total {
			b = b[:total]
		}
		return boxBytes(b)
	})
})
