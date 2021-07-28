package types

type PropertyTypeEnum string

const (
	PropertyTypeEnumSystem PropertyTypeEnum = "system" // 系统定义类型
	PropertyTypeEnumCustom PropertyTypeEnum = "custom" // 自定义类型
)

const (
	PropertySystemClassEnumInt32  = "int32"
	PropertySystemClassEnumInt64  = "int64"
	PropertySystemClassEnumUint32 = "uint32"
	PropertySystemClassEnumUint64 = "uint64"
	PropertySystemClassEnumFloat  = "float"
	PropertySystemClassEnumDouble = "double"
	PropertySystemClassEnumString = "string"
	PropertySystemClassEnumList   = "list"
	PropertySystemClassEnumDict   = "map"
)

func AllPropertySystemClassEnum() []string {
	return []string{
		PropertySystemClassEnumInt64,
		PropertySystemClassEnumInt32,
		PropertySystemClassEnumUint32,
		PropertySystemClassEnumUint64,
		PropertySystemClassEnumFloat,
		PropertySystemClassEnumDouble,
		PropertySystemClassEnumString,
		PropertySystemClassEnumList,
		PropertySystemClassEnumDict,
	}
}
