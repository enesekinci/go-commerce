package extensions

func ToString(content interface{}) string {
	if content == nil {
		return ""
	}
	return content.(string)
}

func ToUint(content interface{}) uint {
	if content == nil {
		return 0
	}
	return content.(uint)
}

func ToUint8(content interface{}) uint8 {
	if content == nil {
		return 0
	}
	return content.(uint8)
}

func ToUint16(content interface{}) uint16 {
	if content == nil {
		return 0
	}
	return content.(uint16)
}

func ToUint32(content interface{}) uint32 {
	if content == nil {
		return 0
	}
	return content.(uint32)
}

func ToUint64(content interface{}) uint64 {
	if content == nil {
		return 0
	}
	return content.(uint64)
}

func ToInt(content interface{}) int {
	if content == nil {
		return 0
	}
	return content.(int)
}

func ToInt8(content interface{}) int8 {
	if content == nil {
		return 0
	}
	return content.(int8)
}

func ToInt16(content interface{}) int16 {
	if content == nil {
		return 0
	}
	return content.(int16)
}

func ToInt32(content interface{}) int32 {
	if content == nil {
		return 0
	}
	return content.(int32)
}

func ToInt64(content interface{}) int64 {
	if content == nil {
		return 0
	}
	return content.(int64)
}

func ToFloat32(content interface{}) float32 {
	if content == nil {
		return 0
	}
	return content.(float32)
}

func ToFloat64(content interface{}) float64 {
	if content == nil {
		return 0
	}
	return content.(float64)
}

func ToBool(content interface{}) bool {
	if content == nil {
		return false
	}
	return content.(bool)
}

func ToSlice(content interface{}) []interface{} {
	if content == nil {
		return nil
	}
	return content.([]interface{})
}

func ToMap(content interface{}) map[string]interface{} {
	if content == nil {
		return nil
	}
	return content.(map[string]interface{})
}
