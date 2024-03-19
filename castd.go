package cast

func ToIntD(i any, defaultValue int) int {
	if val := ToInt(i); val == 0 {
		return defaultValue
	} else {
		return val
	}
}

func ToUIntD(i any, defaultValue uint) uint {
	if val := ToUint(i); val == 0 {
		return defaultValue
	} else {
		return val
	}
}

func ToInt64D(i any, defaultValue int64) int64 {
	if val := ToInt64(i); val == 0 {
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

func ToUInt8D(i any, defaultValue uint8) uint8 {
	if val := ToUint8(i); val == 0 {
		return defaultValue
	} else {
		return val
	}
}

func ToStringD(i any, defaultValue string) string {
	if val := ToString(i); val == "" {
		return defaultValue
	} else {
		return val
	}
}
