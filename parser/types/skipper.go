// Code generated by gen/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/grafana/jfr-parser/parser/types/def"
	"github.com/grafana/jfr-parser/util"
)

type BindSkipConstantPool struct {
	Temp   SkipConstantPool
	Fields []BindFieldSkipConstantPool
}

type BindFieldSkipConstantPool struct {
	Field *def.Field
}

func NewBindSkipConstantPool(typ *def.Class, typeMap *def.TypeMap) *BindSkipConstantPool {
	res := new(BindSkipConstantPool)
	res.Fields = make([]BindFieldSkipConstantPool, 0, len(typ.Fields))
	for i := 0; i < len(typ.Fields); i++ {
		switch typ.Fields[i].Name {
		default:
			res.Fields = append(res.Fields, BindFieldSkipConstantPool{Field: &typ.Fields[i]}) // skip unknown new field
		}
	}
	return res
}

type SkipConstantPoolRef uint32
type SkipConstantPoolList struct {
}

type SkipConstantPool struct {
}

func (this *SkipConstantPoolList) Parse(data []byte, bind *BindSkipConstantPool, typeMap *def.TypeMap) (pos int, err error) {
	v32_, err := util.ParseVarInt(data, &pos)
	if err != nil {
		return 0, err
	}
	n := int(v32_)
	for i := 0; i < n; i++ {
		v32_, err = util.ParseVarInt(data, &pos)
		if err != nil {
			return 0, err
		}
		for bindFieldIndex := 0; bindFieldIndex < len(bind.Fields); bindFieldIndex++ {
			bindArraySize := 1
			if bind.Fields[bindFieldIndex].Field.Array {
				v32_, err := util.ParseVarInt(data, &pos)
				if err != nil {
					return 0, err
				}
				bindArraySize = int(v32_)
			}
			for bindArrayIndex := 0; bindArrayIndex < bindArraySize; bindArrayIndex++ {
				if bind.Fields[bindFieldIndex].Field.ConstantPool {
					v32_, err := util.ParseVarInt(data, &pos)
					if err != nil {
						return 0, err
					}
					_ = v32_
				} else {
					bindFieldTypeID := bind.Fields[bindFieldIndex].Field.Type
					switch bindFieldTypeID {
					case typeMap.T_STRING:
						s_, err := util.ParseString(data, &pos)
						if err != nil {
							return 0, err
						}
						// skipping
						_ = s_
					case typeMap.T_INT:
						v32_, err := util.ParseVarInt(data, &pos)
						if err != nil {
							return 0, err
						}
						// skipping
						_ = v32_
					case typeMap.T_LONG:
						v64_, err := util.ParseVarLong(data, &pos)
						if err != nil {
							return 0, err
						}
						// skipping
						_ = v64_
					case typeMap.T_BOOLEAN:
						b_, err := util.ParseByte(data, &pos)
						if err != nil {
							return 0, err
						}
						// skipping
						_ = b_
					case typeMap.T_FLOAT:
						v32_, err := util.ParseVarInt(data, &pos)
						if err != nil {
							return 0, err
						}
						// skipping
						_ = v32_
					default:
						bindFieldType := typeMap.IDMap[bind.Fields[bindFieldIndex].Field.Type]
						if bindFieldType == nil || len(bindFieldType.Fields) == 0 {
							return 0, fmt.Errorf("unknown type %d", bind.Fields[bindFieldIndex].Field.Type)
						}
						bindSkipObjects := 1
						if bind.Fields[bindFieldIndex].Field.Array {
							v32_, err := util.ParseVarInt(data, &pos)
							if err != nil {
								return 0, err
							}
							bindSkipObjects = int(v32_)
						}
						for bindSkipObjectIndex := 0; bindSkipObjectIndex < bindSkipObjects; bindSkipObjectIndex++ {
							for bindskipFieldIndex := 0; bindskipFieldIndex < len(bindFieldType.Fields); bindskipFieldIndex++ {
								bindSkipFieldType := bindFieldType.Fields[bindskipFieldIndex].Type
								if bindFieldType.Fields[bindskipFieldIndex].ConstantPool {
									_, err := util.ParseVarInt(data, &pos)
									if err != nil {
										return 0, err
									}
								} else if bindSkipFieldType == typeMap.T_STRING {
									_, err := util.ParseString(data, &pos)
									if err != nil {
										return 0, err
									}
								} else if bindSkipFieldType == typeMap.T_INT {
									_, err := util.ParseVarInt(data, &pos)
									if err != nil {
										return 0, err
									}
								} else if bindSkipFieldType == typeMap.T_FLOAT {
									_, err := util.ParseVarInt(data, &pos)
									if err != nil {
										return 0, err
									}
								} else if bindSkipFieldType == typeMap.T_LONG {
									_, err := util.ParseVarLong(data, &pos)
									if err != nil {
										return 0, err
									}
								} else if bindSkipFieldType == typeMap.T_BOOLEAN {
									_, err := util.ParseByte(data, &pos)
									if err != nil {
										return 0, err
									}
								} else {
									return 0, fmt.Errorf("nested objects not implemented. ")
								}
							}
						}
					}
				}
			}
		}

	}
	return pos, nil
}
