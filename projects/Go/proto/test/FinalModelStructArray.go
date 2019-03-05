// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.2.0.0

package test

import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// Fast Binary Encoding StructArray final model
type FinalModelStructArray struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset

    F1 *FinalModelArrayByte
    F2 *FinalModelArrayOptionalByte
    F3 *FinalModelArrayBytes
    F4 *FinalModelArrayOptionalBytes
    F5 *FinalModelArrayEnumSimple
    F6 *FinalModelArrayOptionalEnumSimple
    F7 *FinalModelArrayFlagsSimple
    F8 *FinalModelArrayOptionalFlagsSimple
    F9 *FinalModelArrayStructSimple
    F10 *FinalModelArrayOptionalStructSimple
}

// Create a new StructArray final model
func NewFinalModelStructArray(buffer *fbe.Buffer, offset int) *FinalModelStructArray {
    fbeResult := FinalModelStructArray{buffer: buffer, offset: offset}
    fbeResult.F1 = NewFinalModelArrayByte(buffer, 0, 2)
    fbeResult.F2 = NewFinalModelArrayOptionalByte(buffer, 0, 2)
    fbeResult.F3 = NewFinalModelArrayBytes(buffer, 0, 2)
    fbeResult.F4 = NewFinalModelArrayOptionalBytes(buffer, 0, 2)
    fbeResult.F5 = NewFinalModelArrayEnumSimple(buffer, 0, 2)
    fbeResult.F6 = NewFinalModelArrayOptionalEnumSimple(buffer, 0, 2)
    fbeResult.F7 = NewFinalModelArrayFlagsSimple(buffer, 0, 2)
    fbeResult.F8 = NewFinalModelArrayOptionalFlagsSimple(buffer, 0, 2)
    fbeResult.F9 = NewFinalModelArrayStructSimple(buffer, 0, 2)
    fbeResult.F10 = NewFinalModelArrayOptionalStructSimple(buffer, 0, 2)
    return &fbeResult
}

// Get the allocation size
func (fm *FinalModelStructArray) FBEAllocationSize(fbeValue *StructArray) int {
    fbeResult := 0 +
        fm.F1.FBEAllocationSize(fbeValue.F1[:]) +
        fm.F2.FBEAllocationSize(fbeValue.F2[:]) +
        fm.F3.FBEAllocationSize(fbeValue.F3[:]) +
        fm.F4.FBEAllocationSize(fbeValue.F4[:]) +
        fm.F5.FBEAllocationSize(fbeValue.F5[:]) +
        fm.F6.FBEAllocationSize(fbeValue.F6[:]) +
        fm.F7.FBEAllocationSize(fbeValue.F7[:]) +
        fm.F8.FBEAllocationSize(fbeValue.F8[:]) +
        fm.F9.FBEAllocationSize(fbeValue.F9[:]) +
        fm.F10.FBEAllocationSize(fbeValue.F10[:]) +
        0
    return fbeResult
}

// Get the final size
func (fm *FinalModelStructArray) FBESize() int { return 0 }

// Get the final extra size
func (fm *FinalModelStructArray) FBEExtra() int { return 0 }

// Get the final type
func (fm *FinalModelStructArray) FBEType() int { return 125 }

// Get the final offset
func (fm *FinalModelStructArray) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelStructArray) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelStructArray) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelStructArray) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FinalModelStructArray) Verify() int {
    fm.buffer.Shift(fm.FBEOffset())
    fbeResult := fm.VerifyFields()
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeResult
}

// Check if the struct fields are valid
func (fm *FinalModelStructArray) VerifyFields() int {
    fbeCurrentOffset := 0
    fbeFieldSize := 0


    fm.F1.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F1.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F2.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F2.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F3.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F3.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F4.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F4.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F5.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F5.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F6.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F6.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F7.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F7.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F8.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F8.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F9.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F9.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F10.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F10.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    return fbeCurrentOffset
}

// Get the struct value
func (fm *FinalModelStructArray) Get() (*StructArray, int, error) {
    fbeResult := NewStructArray()
    fbeSize, err := fm.GetValue(fbeResult)
    return fbeResult, fbeSize, err
}

// Get the struct value by the given pointer
func (fm *FinalModelStructArray) GetValue(fbeValue *StructArray) (int, error) {
    fm.buffer.Shift(fm.FBEOffset())
    fbeSize, err := fm.GetFields(fbeValue)
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeSize, err
}

// Get the struct fields values
func (fm *FinalModelStructArray) GetFields(fbeValue *StructArray) (int, error) {
    var err error = nil
    fbeCurrentOffset := 0
    fbeCurrentSize := 0
    fbeFieldSize := 0

    fm.F1.SetFBEOffset(fbeCurrentOffset)
    if slice, size, err := fm.F1.Get(); err != nil {
        return fbeCurrentSize, err
    } else {
        copy(fbeValue.F1[:], slice)
        fbeFieldSize = size
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F2.SetFBEOffset(fbeCurrentOffset)
    if slice, size, err := fm.F2.Get(); err != nil {
        return fbeCurrentSize, err
    } else {
        copy(fbeValue.F2[:], slice)
        fbeFieldSize = size
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F3.SetFBEOffset(fbeCurrentOffset)
    if slice, size, err := fm.F3.Get(); err != nil {
        return fbeCurrentSize, err
    } else {
        copy(fbeValue.F3[:], slice)
        fbeFieldSize = size
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F4.SetFBEOffset(fbeCurrentOffset)
    if slice, size, err := fm.F4.Get(); err != nil {
        return fbeCurrentSize, err
    } else {
        copy(fbeValue.F4[:], slice)
        fbeFieldSize = size
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F5.SetFBEOffset(fbeCurrentOffset)
    if slice, size, err := fm.F5.Get(); err != nil {
        return fbeCurrentSize, err
    } else {
        copy(fbeValue.F5[:], slice)
        fbeFieldSize = size
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F6.SetFBEOffset(fbeCurrentOffset)
    if slice, size, err := fm.F6.Get(); err != nil {
        return fbeCurrentSize, err
    } else {
        copy(fbeValue.F6[:], slice)
        fbeFieldSize = size
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F7.SetFBEOffset(fbeCurrentOffset)
    if slice, size, err := fm.F7.Get(); err != nil {
        return fbeCurrentSize, err
    } else {
        copy(fbeValue.F7[:], slice)
        fbeFieldSize = size
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F8.SetFBEOffset(fbeCurrentOffset)
    if slice, size, err := fm.F8.Get(); err != nil {
        return fbeCurrentSize, err
    } else {
        copy(fbeValue.F8[:], slice)
        fbeFieldSize = size
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F9.SetFBEOffset(fbeCurrentOffset)
    if slice, size, err := fm.F9.Get(); err != nil {
        return fbeCurrentSize, err
    } else {
        copy(fbeValue.F9[:], slice)
        fbeFieldSize = size
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F10.SetFBEOffset(fbeCurrentOffset)
    if slice, size, err := fm.F10.Get(); err != nil {
        return fbeCurrentSize, err
    } else {
        copy(fbeValue.F10[:], slice)
        fbeFieldSize = size
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    return fbeCurrentSize, err
}

// Set the struct value
func (fm *FinalModelStructArray) Set(fbeValue *StructArray) (int, error) {
    fm.buffer.Shift(fm.FBEOffset())
    fbeResult, err := fm.SetFields(fbeValue)
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeResult, err
}

// Set the struct fields values
func (fm *FinalModelStructArray) SetFields(fbeValue *StructArray) (int, error) {
    var err error = nil
    fbeCurrentOffset := 0
    fbeCurrentSize := 0
    fbeFieldSize := 0

    fm.F1.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1.Set(fbeValue.F1[:]); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F2.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F2.Set(fbeValue.F2[:]); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F3.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F3.Set(fbeValue.F3[:]); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F4.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F4.Set(fbeValue.F4[:]); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F5.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F5.Set(fbeValue.F5[:]); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F6.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F6.Set(fbeValue.F6[:]); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F7.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F7.Set(fbeValue.F7[:]); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F8.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F8.Set(fbeValue.F8[:]); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F9.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F9.Set(fbeValue.F9[:]); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F10.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F10.Set(fbeValue.F10[:]); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    return fbeCurrentSize, err
}