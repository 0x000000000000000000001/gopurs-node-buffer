package Node_Buffer_Immutable

import (
	"gopurs/output/gopurs_runtime"
)

var ShowImpl = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	return a
})

var EqImpl = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{}
	})
})

var CompareImpl = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{}
	})
})

var ComparePartsImpl = gopurs_runtime.Func(func(src gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(target gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(targetStart gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(targetEnd gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(sourceStart gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(sourceEnd gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{}
					})
				})
			})
		})
	})
})

var Alloc = gopurs_runtime.Func(func(size gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})

var FromArray = gopurs_runtime.Func(func(octets gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})

var Size = gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})

var ToArray = gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})

var ToArrayBuffer = gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})

var FromArrayBuffer = gopurs_runtime.Func(func(ab gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})

var FromStringImpl = gopurs_runtime.Func(func(str gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(encoding gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{}
	})
})

var ReadImpl = gopurs_runtime.Func(func(ty gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(offset gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(buf gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})
	})
})

var ReadStringImpl = gopurs_runtime.Func(func(enc gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(start gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(end gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{}
			})
		})
	})
})

var GetAtOffsetImpl = gopurs_runtime.Func(func(offset gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{}
	})
})

var ToStringImpl = gopurs_runtime.Func(func(enc gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{}
	})
})

var ToStringSubImpl = gopurs_runtime.Func(func(enc gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(start gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(end gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{}
			})
		})
	})
})

var SliceImpl = gopurs_runtime.Func(func(start gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(end gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})
	})
})

var Concat = gopurs_runtime.Func(func(buffs gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})

var ConcatToLength = gopurs_runtime.Func(func(buffs gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(totalLength gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{}
	})
})
