// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: expression.proto

package tipb

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ExprType int32

const (
	// Values are encoded bytes.
	ExprType_Null    ExprType = 0
	ExprType_Int64   ExprType = 1
	ExprType_Uint64  ExprType = 2
	ExprType_Float32 ExprType = 3
	ExprType_Float64 ExprType = 4
	ExprType_String  ExprType = 5
	ExprType_Bytes   ExprType = 6
	// Mysql specific types.
	ExprType_MysqlBit      ExprType = 101
	ExprType_MysqlDecimal  ExprType = 102
	ExprType_MysqlDuration ExprType = 103
	ExprType_MysqlEnum     ExprType = 104
	ExprType_MysqlHex      ExprType = 105
	ExprType_MysqlSet      ExprType = 106
	ExprType_MysqlTime     ExprType = 107
	// Encoded value list.
	ExprType_ValueList ExprType = 151
	// Column reference. value is int64 column ID.
	ExprType_ColumnRef ExprType = 201
	// Unary operations, children count 1.
	ExprType_Not    ExprType = 1001
	ExprType_Neg    ExprType = 1002
	ExprType_BitNeg ExprType = 1003
	// Comparison operations.
	ExprType_LT     ExprType = 2001
	ExprType_LE     ExprType = 2002
	ExprType_EQ     ExprType = 2003
	ExprType_NE     ExprType = 2004
	ExprType_GE     ExprType = 2005
	ExprType_GT     ExprType = 2006
	ExprType_NullEQ ExprType = 2007
	// Bit operations.
	ExprType_BitAnd    ExprType = 2101
	ExprType_BitOr     ExprType = 2102
	ExprType_BitXor    ExprType = 2103
	ExprType_LeftShift ExprType = 2104
	ExprType_RighShift ExprType = 2105
	// Arithmatic.
	ExprType_Plus   ExprType = 2201
	ExprType_Minus  ExprType = 2202
	ExprType_Mul    ExprType = 2203
	ExprType_Div    ExprType = 2204
	ExprType_IntDiv ExprType = 2205
	ExprType_Mod    ExprType = 2206
	// Logic operations.
	ExprType_And ExprType = 2301
	ExprType_Or  ExprType = 2302
	ExprType_Xor ExprType = 2303
	// Aggregate functions.
	ExprType_Count       ExprType = 3001
	ExprType_Sum         ExprType = 3002
	ExprType_Avg         ExprType = 3003
	ExprType_Min         ExprType = 3004
	ExprType_Max         ExprType = 3005
	ExprType_First       ExprType = 3006
	ExprType_GroupConcat ExprType = 3007
	// Math functions.
	ExprType_Abs   ExprType = 3101
	ExprType_Pow   ExprType = 3102
	ExprType_Round ExprType = 3103
	// String functions.
	ExprType_Concat         ExprType = 3201
	ExprType_ConcatWS       ExprType = 3202
	ExprType_Left           ExprType = 3203
	ExprType_Length         ExprType = 3204
	ExprType_Lower          ExprType = 3205
	ExprType_Repeat         ExprType = 3206
	ExprType_Replace        ExprType = 3207
	ExprType_Upper          ExprType = 3208
	ExprType_Strcmp         ExprType = 3209
	ExprType_Convert        ExprType = 3210
	ExprType_Cast           ExprType = 3211
	ExprType_Substring      ExprType = 3212
	ExprType_SubstringIndex ExprType = 3213
	ExprType_Locate         ExprType = 3214
	ExprType_Trim           ExprType = 3215
	// Control flow functions.
	ExprType_If     ExprType = 3301
	ExprType_NullIf ExprType = 3302
	ExprType_IfNull ExprType = 3303
	// Time functions.
	ExprType_Date        ExprType = 3401
	ExprType_DateAdd     ExprType = 3402
	ExprType_DateSub     ExprType = 3403
	ExprType_Year        ExprType = 3411
	ExprType_YearWeek    ExprType = 3412
	ExprType_Month       ExprType = 3421
	ExprType_Week        ExprType = 3431
	ExprType_Weekday     ExprType = 3432
	ExprType_WeekOfYear  ExprType = 3433
	ExprType_Day         ExprType = 3441
	ExprType_DayName     ExprType = 3442
	ExprType_DayOfYear   ExprType = 3443
	ExprType_DayOfMonth  ExprType = 3444
	ExprType_DayOfWeek   ExprType = 3445
	ExprType_Hour        ExprType = 3451
	ExprType_Minute      ExprType = 3452
	ExprType_Second      ExprType = 3453
	ExprType_Microsecond ExprType = 3454
	ExprType_Extract     ExprType = 3461
	// Other functions;
	ExprType_Coalesce ExprType = 3501
	ExprType_Greatest ExprType = 3502
	ExprType_Least    ExprType = 3503
	// Json functions;
	ExprType_JsonExtract      ExprType = 3601
	ExprType_JsonType         ExprType = 3602
	ExprType_JsonArray        ExprType = 3603
	ExprType_JsonObject       ExprType = 3604
	ExprType_JsonMerge        ExprType = 3605
	ExprType_JsonValid        ExprType = 3606
	ExprType_JsonSet          ExprType = 3607
	ExprType_JsonInsert       ExprType = 3608
	ExprType_JsonReplace      ExprType = 3609
	ExprType_JsonRemove       ExprType = 3610
	ExprType_JsonContains     ExprType = 3611
	ExprType_JsonUnquote      ExprType = 3612
	ExprType_JsonContainsPath ExprType = 3613
	// Other expressions.
	ExprType_In      ExprType = 4001
	ExprType_IsTruth ExprType = 4002
	ExprType_IsNull  ExprType = 4003
	ExprType_ExprRow ExprType = 4004
	ExprType_Like    ExprType = 4005
	ExprType_RLike   ExprType = 4006
	ExprType_Case    ExprType = 4007
)

var ExprType_name = map[int32]string{
	0:    "Null",
	1:    "Int64",
	2:    "Uint64",
	3:    "Float32",
	4:    "Float64",
	5:    "String",
	6:    "Bytes",
	101:  "MysqlBit",
	102:  "MysqlDecimal",
	103:  "MysqlDuration",
	104:  "MysqlEnum",
	105:  "MysqlHex",
	106:  "MysqlSet",
	107:  "MysqlTime",
	151:  "ValueList",
	201:  "ColumnRef",
	1001: "Not",
	1002: "Neg",
	1003: "BitNeg",
	2001: "LT",
	2002: "LE",
	2003: "EQ",
	2004: "NE",
	2005: "GE",
	2006: "GT",
	2007: "NullEQ",
	2101: "BitAnd",
	2102: "BitOr",
	2103: "BitXor",
	2104: "LeftShift",
	2105: "RighShift",
	2201: "Plus",
	2202: "Minus",
	2203: "Mul",
	2204: "Div",
	2205: "IntDiv",
	2206: "Mod",
	2301: "And",
	2302: "Or",
	2303: "Xor",
	3001: "Count",
	3002: "Sum",
	3003: "Avg",
	3004: "Min",
	3005: "Max",
	3006: "First",
	3007: "GroupConcat",
	3101: "Abs",
	3102: "Pow",
	3103: "Round",
	3201: "Concat",
	3202: "ConcatWS",
	3203: "Left",
	3204: "Length",
	3205: "Lower",
	3206: "Repeat",
	3207: "Replace",
	3208: "Upper",
	3209: "Strcmp",
	3210: "Convert",
	3211: "Cast",
	3212: "Substring",
	3213: "SubstringIndex",
	3214: "Locate",
	3215: "Trim",
	3301: "If",
	3302: "NullIf",
	3303: "IfNull",
	3401: "Date",
	3402: "DateAdd",
	3403: "DateSub",
	3411: "Year",
	3412: "YearWeek",
	3421: "Month",
	3431: "Week",
	3432: "Weekday",
	3433: "WeekOfYear",
	3441: "Day",
	3442: "DayName",
	3443: "DayOfYear",
	3444: "DayOfMonth",
	3445: "DayOfWeek",
	3451: "Hour",
	3452: "Minute",
	3453: "Second",
	3454: "Microsecond",
	3461: "Extract",
	3501: "Coalesce",
	3502: "Greatest",
	3503: "Least",
	3601: "JsonExtract",
	3602: "JsonType",
	3603: "JsonArray",
	3604: "JsonObject",
	3605: "JsonMerge",
	3606: "JsonValid",
	3607: "JsonSet",
	3608: "JsonInsert",
	3609: "JsonReplace",
	3610: "JsonRemove",
	3611: "JsonContains",
	3612: "JsonUnquote",
	3613: "JsonContainsPath",
	4001: "In",
	4002: "IsTruth",
	4003: "IsNull",
	4004: "ExprRow",
	4005: "Like",
	4006: "RLike",
	4007: "Case",
}
var ExprType_value = map[string]int32{
	"Null":             0,
	"Int64":            1,
	"Uint64":           2,
	"Float32":          3,
	"Float64":          4,
	"String":           5,
	"Bytes":            6,
	"MysqlBit":         101,
	"MysqlDecimal":     102,
	"MysqlDuration":    103,
	"MysqlEnum":        104,
	"MysqlHex":         105,
	"MysqlSet":         106,
	"MysqlTime":        107,
	"ValueList":        151,
	"ColumnRef":        201,
	"Not":              1001,
	"Neg":              1002,
	"BitNeg":           1003,
	"LT":               2001,
	"LE":               2002,
	"EQ":               2003,
	"NE":               2004,
	"GE":               2005,
	"GT":               2006,
	"NullEQ":           2007,
	"BitAnd":           2101,
	"BitOr":            2102,
	"BitXor":           2103,
	"LeftShift":        2104,
	"RighShift":        2105,
	"Plus":             2201,
	"Minus":            2202,
	"Mul":              2203,
	"Div":              2204,
	"IntDiv":           2205,
	"Mod":              2206,
	"And":              2301,
	"Or":               2302,
	"Xor":              2303,
	"Count":            3001,
	"Sum":              3002,
	"Avg":              3003,
	"Min":              3004,
	"Max":              3005,
	"First":            3006,
	"GroupConcat":      3007,
	"Abs":              3101,
	"Pow":              3102,
	"Round":            3103,
	"Concat":           3201,
	"ConcatWS":         3202,
	"Left":             3203,
	"Length":           3204,
	"Lower":            3205,
	"Repeat":           3206,
	"Replace":          3207,
	"Upper":            3208,
	"Strcmp":           3209,
	"Convert":          3210,
	"Cast":             3211,
	"Substring":        3212,
	"SubstringIndex":   3213,
	"Locate":           3214,
	"Trim":             3215,
	"If":               3301,
	"NullIf":           3302,
	"IfNull":           3303,
	"Date":             3401,
	"DateAdd":          3402,
	"DateSub":          3403,
	"Year":             3411,
	"YearWeek":         3412,
	"Month":            3421,
	"Week":             3431,
	"Weekday":          3432,
	"WeekOfYear":       3433,
	"Day":              3441,
	"DayName":          3442,
	"DayOfYear":        3443,
	"DayOfMonth":       3444,
	"DayOfWeek":        3445,
	"Hour":             3451,
	"Minute":           3452,
	"Second":           3453,
	"Microsecond":      3454,
	"Extract":          3461,
	"Coalesce":         3501,
	"Greatest":         3502,
	"Least":            3503,
	"JsonExtract":      3601,
	"JsonType":         3602,
	"JsonArray":        3603,
	"JsonObject":       3604,
	"JsonMerge":        3605,
	"JsonValid":        3606,
	"JsonSet":          3607,
	"JsonInsert":       3608,
	"JsonReplace":      3609,
	"JsonRemove":       3610,
	"JsonContains":     3611,
	"JsonUnquote":      3612,
	"JsonContainsPath": 3613,
	"In":               4001,
	"IsTruth":          4002,
	"IsNull":           4003,
	"ExprRow":          4004,
	"Like":             4005,
	"RLike":            4006,
	"Case":             4007,
}

func (x ExprType) Enum() *ExprType {
	p := new(ExprType)
	*p = x
	return p
}
func (x ExprType) String() string {
	return proto.EnumName(ExprType_name, int32(x))
}
func (x *ExprType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExprType_value, data, "ExprType")
	if err != nil {
		return err
	}
	*x = ExprType(value)
	return nil
}
func (ExprType) EnumDescriptor() ([]byte, []int) { return fileDescriptorExpression, []int{0} }

// Evaluators should implement evaluation functions for every expression type.
type Expr struct {
	Tp               ExprType `protobuf:"varint,1,opt,name=tp,enum=tipb.ExprType" json:"tp"`
	Val              []byte   `protobuf:"bytes,2,opt,name=val" json:"val,omitempty"`
	Children         []*Expr  `protobuf:"bytes,3,rep,name=children" json:"children,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Expr) Reset()                    { *m = Expr{} }
func (m *Expr) String() string            { return proto.CompactTextString(m) }
func (*Expr) ProtoMessage()               {}
func (*Expr) Descriptor() ([]byte, []int) { return fileDescriptorExpression, []int{0} }

func (m *Expr) GetTp() ExprType {
	if m != nil {
		return m.Tp
	}
	return ExprType_Null
}

func (m *Expr) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

func (m *Expr) GetChildren() []*Expr {
	if m != nil {
		return m.Children
	}
	return nil
}

// ByItem type for group by and order by.
type ByItem struct {
	Expr             *Expr  `protobuf:"bytes,1,opt,name=expr" json:"expr,omitempty"`
	Desc             bool   `protobuf:"varint,2,opt,name=desc" json:"desc"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ByItem) Reset()                    { *m = ByItem{} }
func (m *ByItem) String() string            { return proto.CompactTextString(m) }
func (*ByItem) ProtoMessage()               {}
func (*ByItem) Descriptor() ([]byte, []int) { return fileDescriptorExpression, []int{1} }

func (m *ByItem) GetExpr() *Expr {
	if m != nil {
		return m.Expr
	}
	return nil
}

func (m *ByItem) GetDesc() bool {
	if m != nil {
		return m.Desc
	}
	return false
}

func init() {
	proto.RegisterType((*Expr)(nil), "tipb.Expr")
	proto.RegisterType((*ByItem)(nil), "tipb.ByItem")
	proto.RegisterEnum("tipb.ExprType", ExprType_name, ExprType_value)
}
func (m *Expr) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Expr) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintExpression(dAtA, i, uint64(m.Tp))
	if m.Val != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintExpression(dAtA, i, uint64(len(m.Val)))
		i += copy(dAtA[i:], m.Val)
	}
	if len(m.Children) > 0 {
		for _, msg := range m.Children {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintExpression(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ByItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ByItem) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Expr != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExpression(dAtA, i, uint64(m.Expr.Size()))
		n1, err := m.Expr.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	dAtA[i] = 0x10
	i++
	if m.Desc {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Expression(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Expression(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintExpression(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Expr) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovExpression(uint64(m.Tp))
	if m.Val != nil {
		l = len(m.Val)
		n += 1 + l + sovExpression(uint64(l))
	}
	if len(m.Children) > 0 {
		for _, e := range m.Children {
			l = e.Size()
			n += 1 + l + sovExpression(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ByItem) Size() (n int) {
	var l int
	_ = l
	if m.Expr != nil {
		l = m.Expr.Size()
		n += 1 + l + sovExpression(uint64(l))
	}
	n += 2
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovExpression(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozExpression(x uint64) (n int) {
	return sovExpression(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Expr) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpression
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Expr: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Expr: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tp", wireType)
			}
			m.Tp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpression
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tp |= (ExprType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpression
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthExpression
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Val = append(m.Val[:0], dAtA[iNdEx:postIndex]...)
			if m.Val == nil {
				m.Val = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Children", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpression
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExpression
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Children = append(m.Children, &Expr{})
			if err := m.Children[len(m.Children)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExpression(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExpression
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ByItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpression
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ByItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ByItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expr", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpression
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExpression
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expr == nil {
				m.Expr = &Expr{}
			}
			if err := m.Expr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpression
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Desc = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipExpression(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExpression
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipExpression(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExpression
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExpression
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExpression
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthExpression
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowExpression
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipExpression(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthExpression = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExpression   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("expression.proto", fileDescriptorExpression) }

var fileDescriptorExpression = []byte{
	// 1111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x55, 0x49, 0x70, 0x5b, 0x45,
	0x10, 0x8d, 0x6c, 0xc5, 0x96, 0xc7, 0x8e, 0xdd, 0x1e, 0x48, 0x45, 0x52, 0x54, 0x56, 0x2a, 0x45,
	0x51, 0x81, 0x83, 0xa9, 0x0a, 0xa9, 0xdc, 0x2d, 0x5b, 0x49, 0x44, 0x79, 0x8b, 0xe4, 0x24, 0x70,
	0x1c, 0x7f, 0x8d, 0xa4, 0x49, 0xbe, 0x66, 0x7e, 0xe6, 0xcf, 0x77, 0xa4, 0x23, 0x4b, 0xd8, 0x97,
	0x62, 0xcd, 0x1e, 0x8a, 0xfd, 0x04, 0x5c, 0xd8, 0xc2, 0x7a, 0x4d, 0xe0, 0x02, 0x24, 0xc0, 0x09,
	0x8a, 0x32, 0x05, 0x21, 0x70, 0x62, 0x3b, 0x01, 0x81, 0xea, 0xfe, 0xf2, 0x0f, 0xdc, 0xde, 0xeb,
	0x79, 0xaf, 0xbb, 0xa7, 0xa7, 0xbf, 0xc4, 0x40, 0x76, 0x02, 0x2b, 0xc3, 0x50, 0x19, 0x3d, 0x19,
	0x58, 0xe3, 0x0c, 0x4f, 0x3b, 0x15, 0x2c, 0xe7, 0xaf, 0x6f, 0x9a, 0xa6, 0xa1, 0xc0, 0x2d, 0x88,
	0xe2, 0xb3, 0xad, 0x0d, 0x96, 0x2e, 0x77, 0x02, 0xcb, 0x6f, 0x60, 0x7d, 0x2e, 0xc8, 0xa6, 0xb6,
	0xa4, 0xb6, 0x8d, 0x6e, 0x1f, 0x9d, 0x44, 0xc3, 0x24, 0xc6, 0x97, 0xba, 0x81, 0x2c, 0xa5, 0xcf,
	0x7f, 0x53, 0x5c, 0x57, 0xed, 0x73, 0x01, 0x07, 0xd6, 0xbf, 0x22, 0xfc, 0x6c, 0xdf, 0x96, 0xd4,
	0xb6, 0x91, 0x2a, 0x42, 0x7e, 0x23, 0xcb, 0x78, 0x2d, 0xe5, 0xd7, 0xad, 0xd4, 0xd9, 0xfe, 0x2d,
	0xfd, 0xdb, 0x86, 0xb7, 0xb3, 0x6b, 0xee, 0x6a, 0x72, 0xb6, 0xb5, 0xc4, 0x06, 0x4a, 0xdd, 0x8a,
	0x93, 0x6d, 0x3e, 0xc1, 0xd2, 0xd8, 0x21, 0xd5, 0xfa, 0xbf, 0x9a, 0xe2, 0x3c, 0xcb, 0xd2, 0x75,
	0x19, 0x7a, 0x54, 0x24, 0xd3, 0xab, 0x4d, 0x91, 0x9b, 0xdf, 0x18, 0x66, 0x99, 0xb5, 0xa6, 0x78,
	0x86, 0xa5, 0xe7, 0x23, 0xdf, 0x87, 0x75, 0x7c, 0x88, 0xad, 0xaf, 0x68, 0xb7, 0x73, 0x07, 0xa4,
	0x38, 0x63, 0x03, 0xfb, 0x14, 0xe1, 0x3e, 0x3e, 0xcc, 0x06, 0x77, 0xf9, 0x46, 0xb8, 0x5b, 0xb7,
	0x43, 0x7f, 0x42, 0x76, 0xee, 0x80, 0x34, 0xaa, 0x6a, 0xce, 0x2a, 0xdd, 0x84, 0xf5, 0x68, 0x2e,
	0x75, 0x9d, 0x0c, 0x61, 0x80, 0x8f, 0xb0, 0xcc, 0x5c, 0x37, 0x3c, 0xec, 0x97, 0x94, 0x03, 0xc9,
	0x81, 0x8d, 0x10, 0x9b, 0x91, 0x9e, 0x6a, 0x0b, 0x1f, 0x1a, 0x7c, 0x9c, 0x6d, 0x88, 0x23, 0x91,
	0x15, 0x4e, 0x19, 0x0d, 0x4d, 0xbe, 0x81, 0x0d, 0x51, 0xa8, 0xac, 0xa3, 0x36, 0xb4, 0x92, 0x0c,
	0x7b, 0x64, 0x07, 0x54, 0xc2, 0x6a, 0xd2, 0xc1, 0xc1, 0x44, 0xba, 0xa4, 0xda, 0x12, 0x0e, 0xf1,
	0x51, 0x36, 0xb4, 0x5f, 0xf8, 0x91, 0x9c, 0x55, 0xa1, 0x83, 0x63, 0x29, 0xe4, 0xd3, 0xc6, 0x8f,
	0xda, 0xba, 0x2a, 0x1b, 0x70, 0x21, 0xc5, 0x33, 0xac, 0x7f, 0xde, 0x38, 0xb8, 0x32, 0x48, 0x48,
	0x36, 0xe1, 0xa7, 0x41, 0x3e, 0xcc, 0x06, 0x4a, 0xca, 0x21, 0xf9, 0x79, 0x90, 0x0f, 0xb2, 0xbe,
	0xd9, 0x25, 0xf8, 0x6c, 0x8c, 0x40, 0x19, 0x3e, 0x27, 0x50, 0xde, 0x0b, 0x17, 0x09, 0xcc, 0x97,
	0xe1, 0x12, 0x81, 0xdd, 0x65, 0xf8, 0x22, 0x06, 0x4b, 0xf0, 0xe5, 0x18, 0xa6, 0xc0, 0xa9, 0x95,
	0xf7, 0xc2, 0x57, 0x63, 0xbd, 0x7c, 0x53, 0xba, 0x0e, 0x6f, 0x02, 0x67, 0x6c, 0x7d, 0x49, 0xb9,
	0x05, 0x0b, 0x6f, 0x41, 0xef, 0xe0, 0x76, 0x63, 0xe1, 0x6d, 0xc0, 0xce, 0x66, 0x65, 0xc3, 0xd5,
	0x5a, 0xaa, 0xe1, 0xe0, 0x1d, 0xe2, 0x55, 0xd5, 0x6c, 0xc5, 0xfc, 0x1c, 0xf0, 0x21, 0x96, 0x5e,
	0xf4, 0xa3, 0x10, 0x4e, 0x8c, 0x63, 0x8e, 0x39, 0xa5, 0xa3, 0x10, 0x4e, 0x8e, 0x63, 0xdb, 0x73,
	0x91, 0x0f, 0xa7, 0x08, 0xcd, 0xa8, 0x15, 0x38, 0x3d, 0x8e, 0x79, 0x2b, 0xda, 0x21, 0x39, 0x13,
	0x0b, 0x4c, 0x1d, 0xce, 0x12, 0xc2, 0x26, 0xfe, 0x1e, 0xc7, 0x3e, 0x17, 0x2c, 0x5c, 0xa5, 0x10,
	0x96, 0xff, 0x87, 0x72, 0x4e, 0x9b, 0x48, 0x3b, 0x38, 0xb7, 0x09, 0xa3, 0xb5, 0xa8, 0x0d, 0xef,
	0x12, 0x9a, 0x5a, 0x69, 0xc2, 0x7b, 0x84, 0xe6, 0x94, 0x86, 0xf7, 0x63, 0x24, 0x3a, 0xf0, 0xc1,
	0x26, 0xf4, 0xec, 0x52, 0x36, 0x74, 0xf0, 0xe1, 0x26, 0x0e, 0x6c, 0x78, 0xb7, 0x35, 0x51, 0x30,
	0x6d, 0xb4, 0x27, 0x1c, 0x7c, 0x14, 0x7b, 0x97, 0x43, 0x38, 0x93, 0x45, 0xb4, 0x68, 0x8e, 0xc0,
	0xd9, 0x2c, 0x3a, 0xaa, 0x26, 0xd2, 0x75, 0x78, 0x36, 0x8b, 0x5d, 0xf6, 0xc4, 0x77, 0xe6, 0xf8,
	0x06, 0x96, 0x89, 0xc9, 0x81, 0x1a, 0xdc, 0x95, 0xc3, 0xcb, 0xe2, 0x30, 0xe0, 0xee, 0x1c, 0xca,
	0x66, 0xa5, 0x6e, 0xba, 0x16, 0xdc, 0x93, 0x43, 0xff, 0xac, 0x39, 0x22, 0x2d, 0x1c, 0xa5, 0x83,
	0xaa, 0x0c, 0xa4, 0x70, 0x70, 0x6f, 0x8e, 0x8f, 0xb0, 0xc1, 0xaa, 0x0c, 0x7c, 0xe1, 0x49, 0xb8,
	0x8f, 0x64, 0xfb, 0x82, 0x40, 0x5a, 0xb8, 0x9f, 0x64, 0x35, 0x67, 0xbd, 0x76, 0x00, 0x0f, 0x90,
	0x6c, 0xda, 0xe8, 0x15, 0x69, 0x1d, 0x3c, 0x48, 0x55, 0xa6, 0x45, 0xe8, 0xe0, 0xa1, 0x1c, 0x4e,
	0xbb, 0x16, 0x2d, 0x87, 0xf1, 0xba, 0x3e, 0x9c, 0xe3, 0xd7, 0xb1, 0xd1, 0x84, 0x57, 0x74, 0x5d,
	0x76, 0xe0, 0x91, 0xb8, 0x15, 0xe3, 0x09, 0x27, 0xe1, 0x51, 0x32, 0x2f, 0x59, 0xd5, 0x86, 0xc7,
	0x72, 0x38, 0xce, 0x4a, 0x03, 0xbe, 0xcf, 0xad, 0x3d, 0x7b, 0xa5, 0x01, 0x3f, 0x10, 0xa9, 0x34,
	0xe8, 0xdb, 0xb9, 0x4c, 0xea, 0x19, 0x34, 0x5e, 0xc8, 0x63, 0x0f, 0x08, 0xa7, 0xea, 0x75, 0xf8,
	0x38, 0x61, 0xb5, 0x68, 0x19, 0x3e, 0xc9, 0xa3, 0xec, 0x0e, 0x29, 0x2c, 0x5c, 0xcc, 0xe3, 0x44,
	0x10, 0x1e, 0x90, 0xf2, 0x10, 0x5c, 0xca, 0xd3, 0x9b, 0x1b, 0xed, 0x5a, 0xf0, 0x35, 0xa9, 0x28,
	0x7c, 0x99, 0xec, 0x08, 0xeb, 0xa2, 0x0b, 0x3f, 0xe6, 0xf9, 0x18, 0x63, 0xc8, 0x16, 0x1a, 0x94,
	0xe4, 0x4a, 0x9e, 0x76, 0x42, 0x74, 0xe1, 0x97, 0x5e, 0x9d, 0xee, 0xbc, 0x68, 0x4b, 0xf8, 0x35,
	0x8f, 0xd7, 0x9d, 0x11, 0xdd, 0x9e, 0xee, 0x37, 0x32, 0x12, 0x8f, 0x4b, 0xfc, 0x7e, 0x4d, 0x40,
	0x75, 0xfe, 0xa0, 0x92, 0x7b, 0x4c, 0x64, 0xe1, 0xcf, 0x3c, 0xde, 0x0b, 0xb7, 0xcf, 0x49, 0xf8,
	0x8b, 0x48, 0x4d, 0x7a, 0x06, 0xd7, 0x2a, 0x8f, 0x3b, 0x30, 0xa7, 0x3c, 0x6b, 0xc2, 0x38, 0x72,
	0x95, 0xaa, 0x96, 0x3b, 0xce, 0x0a, 0xcf, 0xc1, 0xd1, 0xcd, 0xf1, 0x23, 0x0b, 0x5f, 0x86, 0x9e,
	0x84, 0x57, 0x89, 0xee, 0xb6, 0x52, 0x38, 0x19, 0x3a, 0x78, 0x6d, 0x33, 0xbd, 0xad, 0xc4, 0xe7,
	0x78, 0x7d, 0x33, 0x66, 0xba, 0x2d, 0x34, 0x7a, 0xcd, 0xfb, 0x78, 0x01, 0xc5, 0x18, 0xc1, 0xdf,
	0x24, 0x78, 0xa2, 0x80, 0xfd, 0x21, 0x9d, 0xb2, 0x56, 0x74, 0xe1, 0xc9, 0x02, 0x5e, 0x00, 0xf9,
	0xc2, 0xf2, 0x41, 0xe9, 0x39, 0x78, 0x2a, 0x11, 0xcc, 0x49, 0xdb, 0x94, 0xf0, 0x74, 0xc2, 0xf7,
	0x0b, 0x5f, 0xd5, 0xe1, 0x99, 0x02, 0x76, 0x86, 0x1c, 0x7f, 0x34, 0x8e, 0x25, 0xf6, 0x8a, 0x0e,
	0x71, 0x35, 0x8e, 0x17, 0xd6, 0x1a, 0x58, 0xdb, 0xa9, 0x13, 0x89, 0xa4, 0x2a, 0xdb, 0x66, 0x45,
	0xc2, 0xc9, 0x02, 0x1f, 0x67, 0x23, 0x18, 0x98, 0x36, 0xda, 0x09, 0xa5, 0x43, 0x38, 0x95, 0xb8,
	0xf6, 0xe9, 0xc3, 0x91, 0x71, 0x12, 0x4e, 0x17, 0xf8, 0x46, 0x06, 0xff, 0x15, 0x2d, 0x0a, 0xd7,
	0x82, 0x33, 0x05, 0xda, 0x18, 0x0d, 0xcf, 0x15, 0xb1, 0x8d, 0x4a, 0xb8, 0x64, 0x23, 0xd7, 0x82,
	0xe7, 0x8b, 0xb4, 0x32, 0x21, 0xad, 0xcc, 0x0b, 0xc5, 0x78, 0x76, 0x81, 0xad, 0x9a, 0x23, 0xf0,
	0x62, 0x91, 0xbe, 0x08, 0x75, 0x48, 0xc2, 0x4b, 0x45, 0xfa, 0x88, 0x08, 0xbf, 0x5c, 0xec, 0xad,
	0xb0, 0x84, 0x57, 0x8a, 0xa5, 0x9b, 0xce, 0xaf, 0x4e, 0xa4, 0x3e, 0x5d, 0x9d, 0x48, 0x7d, 0xbb,
	0x3a, 0x91, 0x3a, 0xfe, 0xdd, 0xc4, 0x3a, 0xb6, 0xd1, 0x33, 0xed, 0xc9, 0x40, 0xe9, 0xa6, 0x27,
	0x82, 0x49, 0xa7, 0xea, 0xcb, 0xf4, 0x27, 0xb0, 0x98, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0xa2,
	0x34, 0xb6, 0xab, 0xbc, 0x06, 0x00, 0x00,
}