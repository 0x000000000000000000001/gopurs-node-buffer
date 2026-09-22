package Node_Buffer

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"math"

	"gopurs/output/gopurs_runtime"
)

func nodeBuffer_getBytes(val gopurs_runtime.Value) []byte {
	return (*(*any)(val.UnsafePtr)).([]byte)
}

func nodeBuffer_boxBytes(b []byte) gopurs_runtime.Value {
	return gopurs_runtime.Any(b)
}

var AllocUnsafeImpl = gopurs_runtime.Func(func(size gopurs_runtime.Value) gopurs_runtime.Value {
	sz := int(size.IntVal)
	return nodeBuffer_boxBytes(make([]byte, sz))
})

var AllocUnsafeSlowImpl = gopurs_runtime.Func(func(size gopurs_runtime.Value) gopurs_runtime.Value {
	sz := int(size.IntVal)
	return nodeBuffer_boxBytes(make([]byte, sz))
})

var FreezeImpl = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	return a
})

var ThawImpl = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	return a
})

var WriteInternal = gopurs_runtime.Func(func(ty gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(offset gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(buf gopurs_runtime.Value) gopurs_runtime.Value {
				b := nodeBuffer_getBytes(buf)
				off := int(offset.IntVal)
				
				// Read strings from `*string`
				t := gopurs_runtime.StrValue(ty)
				
				// The value could be a Float or Int, handle carefully.
				// In PureScript, `value` parameter for `WriteInternal` is `Number`.
				// `Number` is `TypeFloat`, so `value.IntVal` holds math.Float64bits
				val := math.Float64frombits(uint64(value.IntVal))
				
				switch t {
				case "UInt8", "Int8":
					b[off] = byte(val)
				case "UInt16LE", "Int16LE":
					binary.LittleEndian.PutUint16(b[off:], uint16(val))
				case "UInt16BE", "Int16BE":
					binary.BigEndian.PutUint16(b[off:], uint16(val))
				case "UInt32LE", "Int32LE":
					binary.LittleEndian.PutUint32(b[off:], uint32(val))
				case "UInt32BE", "Int32BE":
					binary.BigEndian.PutUint32(b[off:], uint32(val))
				case "FloatLE":
					binary.LittleEndian.PutUint32(b[off:], math.Float32bits(float32(val)))
				case "FloatBE":
					binary.BigEndian.PutUint32(b[off:], math.Float32bits(float32(val)))
				case "DoubleLE":
					binary.LittleEndian.PutUint64(b[off:], math.Float64bits(val))
				case "DoubleBE":
					binary.BigEndian.PutUint64(b[off:], math.Float64bits(val))
				}
				
				return gopurs_runtime.Value{} // Effect Unit
			})
		})
	})
})

var WriteStringInternal = gopurs_runtime.Func(func(encoding gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(offset gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(length gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
					enc := gopurs_runtime.StrValue(encoding)
					off := int(offset.IntVal)
					l := int(length.IntVal)
					val := gopurs_runtime.StrValue(value)
					b := nodeBuffer_getBytes(buff)
					
					var decoded []byte
					switch enc {
					case "hex":
						decoded, _ = hex.DecodeString(val)
					case "base64":
						decoded, _ = base64.StdEncoding.DecodeString(val)
					default:
						decoded = []byte(val)
					}
					
					if l > len(decoded) {
						l = len(decoded)
					}
					copy(b[off:], decoded[:l])
					return gopurs_runtime.Int(int64(l)) // returns Int (bytes written)
				})
			})
		})
	})
})

var SetAtOffsetImpl = gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(offset gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
			b := nodeBuffer_getBytes(buff)
			off := int(offset.IntVal)
			if off >= 0 && off < len(b) {
				b[off] = byte(value.IntVal)
			}
			return gopurs_runtime.Value{}
		})
	})
})

var CopyImpl = gopurs_runtime.Func(func(srcStart gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(srcEnd gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(src gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(targStart gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(targ gopurs_runtime.Value) gopurs_runtime.Value {
					ss := int(srcStart.IntVal)
					se := int(srcEnd.IntVal)
					ts := int(targStart.IntVal)
					s := nodeBuffer_getBytes(src)
					t := nodeBuffer_getBytes(targ)
					
					if ss < 0 { ss = 0 }
					if se > len(s) { se = len(s) }
					if ss > se { ss = se }
					
					n := copy(t[ts:], s[ss:se])
					return gopurs_runtime.Int(int64(n)) // returns Int
				})
			})
		})
	})
})

var FillImpl = gopurs_runtime.Func(func(octet gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(start gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(end gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(buf gopurs_runtime.Value) gopurs_runtime.Value {
				b := nodeBuffer_getBytes(buf)
				val := byte(octet.IntVal)
				s := int(start.IntVal)
				e := int(end.IntVal)
				
				if s < 0 { s = 0 }
				if e > len(b) { e = len(b) }
				if s > e { s = e }
				
				for i := s; i < e; i++ {
					b[i] = val
				}
				
				return gopurs_runtime.Value{}
			})
		})
	})
})

var PoolSize = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Int(8192)
})

var SetPoolSizeImpl = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})

var Swap16Impl = gopurs_runtime.Func(func(buf gopurs_runtime.Value) gopurs_runtime.Value {
	b := nodeBuffer_getBytes(buf)
	for i := 0; i < len(b)-1; i += 2 {
		b[i], b[i+1] = b[i+1], b[i]
	}
	return buf
})

var Swap32Impl = gopurs_runtime.Func(func(buf gopurs_runtime.Value) gopurs_runtime.Value {
	b := nodeBuffer_getBytes(buf)
	for i := 0; i < len(b)-3; i += 4 {
		b[i], b[i+1], b[i+2], b[i+3] = b[i+3], b[i+2], b[i+1], b[i]
	}
	return buf
})

var Swap64Impl = gopurs_runtime.Func(func(buf gopurs_runtime.Value) gopurs_runtime.Value {
	b := nodeBuffer_getBytes(buf)
	for i := 0; i < len(b)-7; i += 8 {
		b[i], b[i+1], b[i+2], b[i+3], b[i+4], b[i+5], b[i+6], b[i+7] = b[i+7], b[i+6], b[i+5], b[i+4], b[i+3], b[i+2], b[i+1], b[i]
	}
	return buf
})

var TranscodeImpl = gopurs_runtime.Func(func(buf gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(from gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(to gopurs_runtime.Value) gopurs_runtime.Value {
			return buf
		})
	})
})
