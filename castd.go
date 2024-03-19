package cast

func ToIntD(i any, defaultValue int) int {
	if val := ToInt(i); val == 0 {
		return defaultValue
	} else {
		return val
	}
}

func ToUInt64D(i any, defaultValue uint64) uint64 {
	if val := ToUint64(i); val == 0 {
		return defaultValue
	} else {
		return val
	}
}

func ToInt8D(i any, defaultValue int8) int8 {
	if val := ToInt8(i); val == 0 {
		return defaultValue
	} else {
		return val
	}
}
