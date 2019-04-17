// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.3.0.0

package test

import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// EnumTyped enum key
type EnumTypedKey uint8

// Convert EnumTyped enum key to string
func (k EnumTypedKey) String() string {
    return EnumTyped(k).String()
}

// EnumTyped enum
type EnumTyped uint8

// EnumTyped enum values
//noinspection GoSnakeCaseUsage
const (
    EnumTyped_ENUM_VALUE_0 = EnumTyped(0 + 0)
    EnumTyped_ENUM_VALUE_1 = EnumTyped(EnumTyped('1') + 0)
    EnumTyped_ENUM_VALUE_2 = EnumTyped(EnumTyped('1') + 1)
    EnumTyped_ENUM_VALUE_3 = EnumTyped(EnumTyped('3') + 0)
    EnumTyped_ENUM_VALUE_4 = EnumTyped(EnumTyped('3') + 1)
    EnumTyped_ENUM_VALUE_5 = EnumTyped(EnumTyped_ENUM_VALUE_3)
)

// Create a new EnumTyped enum
func NewEnumTyped() *EnumTyped {
    return new(EnumTyped)
}

// Create a new EnumTyped enum from the given value
func NewEnumTypedFromValue(value uint8) *EnumTyped {
    result := EnumTyped(value)
    return &result
}

// Get the enum key
func (e EnumTyped) Key() EnumTypedKey {
    return EnumTypedKey(e)
}

// Convert enum to optional
func (e *EnumTyped) Optional() *EnumTyped {
    return e
}

// Convert enum to string
func (e EnumTyped) String() string {
    if e == EnumTyped_ENUM_VALUE_0 {
        return "ENUM_VALUE_0"
    }
    if e == EnumTyped_ENUM_VALUE_1 {
        return "ENUM_VALUE_1"
    }
    if e == EnumTyped_ENUM_VALUE_2 {
        return "ENUM_VALUE_2"
    }
    if e == EnumTyped_ENUM_VALUE_3 {
        return "ENUM_VALUE_3"
    }
    if e == EnumTyped_ENUM_VALUE_4 {
        return "ENUM_VALUE_4"
    }
    if e == EnumTyped_ENUM_VALUE_5 {
        return "ENUM_VALUE_5"
    }
    return "<unknown>"
}

// Convert enum to JSON
func (e EnumTyped) MarshalJSON() ([]byte, error) {
    value := uint8(e)
    return fbe.Json.Marshal(&value)
}

// Convert JSON to enum
func (e *EnumTyped) UnmarshalJSON(buffer []byte) error {
    var result uint8
    err := fbe.Json.Unmarshal(buffer, &result)
    if err != nil {
        return err
    }
    *e = EnumTyped(result)
    return nil
}