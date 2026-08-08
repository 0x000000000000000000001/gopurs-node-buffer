package Node_Buffer

import (
	"gopurs/output/gopurs_runtime"
)

var AllocUnsafeImpl = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})

var AllocUnsafeSlowImpl = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
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
				return gopurs_runtime.Value{}
			})
		})
	})
})

var WriteStringInternal = gopurs_runtime.Func(func(encoding gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(offset gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(length gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{}
				})
			})
		})
	})
})

var SetAtOffsetImpl = gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(offset gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(buff gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})
	})
})

var CopyImpl = gopurs_runtime.Func(func(srcStart gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(srcEnd gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(src gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(targStart gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(targ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{}
				})
			})
		})
	})
})

var FillImpl = gopurs_runtime.Func(func(octet gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(start gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(end gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(buf gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{}
			})
		})
	})
})

var PoolSize = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})

var SetPoolSizeImpl = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})

var Swap16Impl = gopurs_runtime.Func(func(buf gopurs_runtime.Value) gopurs_runtime.Value {
	return buf
})

var Swap32Impl = gopurs_runtime.Func(func(buf gopurs_runtime.Value) gopurs_runtime.Value {
	return buf
})

var Swap64Impl = gopurs_runtime.Func(func(buf gopurs_runtime.Value) gopurs_runtime.Value {
	return buf
})

var TranscodeImpl = gopurs_runtime.Func(func(buf gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(from gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(to gopurs_runtime.Value) gopurs_runtime.Value {
			return buf
		})
	})
})
