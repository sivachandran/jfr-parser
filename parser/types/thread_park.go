// Code generated by gen/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/grafana/jfr-parser/parser/types/def"
	"github.com/grafana/jfr-parser/util"
)

type BindThreadPark struct {
	Temp   ThreadPark
	Fields []BindFieldThreadPark
}

type BindFieldThreadPark struct {
	Field         *def.Field
	uint64        *uint64
	ThreadRef     *ThreadRef
	StackTraceRef *StackTraceRef
	ClassRef      *ClassRef
}

func NewBindThreadPark(typ *def.Class, typeMap *def.TypeMap) *BindThreadPark {
	res := new(BindThreadPark)
	res.Fields = make([]BindFieldThreadPark, 0, len(typ.Fields))
	for i := 0; i < len(typ.Fields); i++ {
		switch typ.Fields[i].Name {
		case "startTime":
			if typ.Fields[i].Equals(&def.Field{Name: "startTime", Type: typeMap.T_LONG, ConstantPool: false, Array: false}) {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i], uint64: &res.Temp.StartTime})
			} else {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i]}) // skip changed field
			}
		case "duration":
			if typ.Fields[i].Equals(&def.Field{Name: "duration", Type: typeMap.T_LONG, ConstantPool: false, Array: false}) {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i], uint64: &res.Temp.Duration})
			} else {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i]}) // skip changed field
			}
		case "eventThread":
			if typ.Fields[i].Equals(&def.Field{Name: "eventThread", Type: typeMap.T_THREAD, ConstantPool: true, Array: false}) {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i], ThreadRef: &res.Temp.EventThread})
			} else {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i]}) // skip changed field
			}
		case "stackTrace":
			if typ.Fields[i].Equals(&def.Field{Name: "stackTrace", Type: typeMap.T_STACK_TRACE, ConstantPool: true, Array: false}) {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i], StackTraceRef: &res.Temp.StackTrace})
			} else {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i]}) // skip changed field
			}
		case "parkedClass":
			if typ.Fields[i].Equals(&def.Field{Name: "parkedClass", Type: typeMap.T_CLASS, ConstantPool: true, Array: false}) {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i], ClassRef: &res.Temp.ParkedClass})
			} else {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i]}) // skip changed field
			}
		case "timeout":
			if typ.Fields[i].Equals(&def.Field{Name: "timeout", Type: typeMap.T_LONG, ConstantPool: false, Array: false}) {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i], uint64: &res.Temp.Timeout})
			} else {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i]}) // skip changed field
			}
		case "until":
			if typ.Fields[i].Equals(&def.Field{Name: "until", Type: typeMap.T_LONG, ConstantPool: false, Array: false}) {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i], uint64: &res.Temp.Until})
			} else {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i]}) // skip changed field
			}
		case "address":
			if typ.Fields[i].Equals(&def.Field{Name: "address", Type: typeMap.T_LONG, ConstantPool: false, Array: false}) {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i], uint64: &res.Temp.Address})
			} else {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i]}) // skip changed field
			}
		case "contextId":
			if typ.Fields[i].Equals(&def.Field{Name: "contextId", Type: typeMap.T_LONG, ConstantPool: false, Array: false}) {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i], uint64: &res.Temp.ContextId})
			} else {
				res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i]}) // skip changed field
			}
		default:
			res.Fields = append(res.Fields, BindFieldThreadPark{Field: &typ.Fields[i]}) // skip unknown new field
		}
	}
	return res
}

type ThreadPark struct {
	StartTime   uint64
	Duration    uint64
	EventThread ThreadRef
	StackTrace  StackTraceRef
	ParkedClass ClassRef
	Timeout     uint64
	Until       uint64
	Address     uint64
	ContextId   uint64
}

func (this *ThreadPark) Parse(data []byte, bind *BindThreadPark, typeMap *def.TypeMap) (pos int, err error) {
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
				switch bind.Fields[bindFieldIndex].Field.Type {
				case typeMap.T_THREAD:
					if bind.Fields[bindFieldIndex].ThreadRef != nil {
						*bind.Fields[bindFieldIndex].ThreadRef = ThreadRef(v32_)
					}
				case typeMap.T_STACK_TRACE:
					if bind.Fields[bindFieldIndex].StackTraceRef != nil {
						*bind.Fields[bindFieldIndex].StackTraceRef = StackTraceRef(v32_)
					}
				case typeMap.T_CLASS:
					if bind.Fields[bindFieldIndex].ClassRef != nil {
						*bind.Fields[bindFieldIndex].ClassRef = ClassRef(v32_)
					}
				}
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
					if bind.Fields[bindFieldIndex].uint64 != nil {
						*bind.Fields[bindFieldIndex].uint64 = v64_
					}
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

	*this = bind.Temp
	return pos, nil
}
