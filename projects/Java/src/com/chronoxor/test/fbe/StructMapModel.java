// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.3.0.0

package com.chronoxor.test.fbe;

import java.io.*;
import java.lang.*;
import java.lang.reflect.*;
import java.math.*;
import java.nio.ByteBuffer;
import java.nio.charset.*;
import java.time.*;
import java.util.*;

import com.chronoxor.fbe.*;
import com.chronoxor.test.*;

// Fast Binary Encoding StructMap model
public final class StructMapModel extends Model
{
    public final FieldModelStructMap model;

    public StructMapModel() { model = new FieldModelStructMap(getBuffer(), 4); }
    public StructMapModel(Buffer buffer) { super(buffer); model = new FieldModelStructMap(getBuffer(), 4); }

    // Get the model size
    public long fbeSize() { return model.fbeSize() + model.fbeExtra(); }
    // Get the model type
    public static final long fbeTypeConst = FieldModelStructMap.fbeTypeConst;
    public long fbeType() { return fbeTypeConst; }

    // Check if the struct value is valid
    public boolean verify()
    {
        if ((getBuffer().getOffset() + model.fbeOffset() - 4) > getBuffer().getSize())
            return false;

        int fbeFullSize = readInt32(model.fbeOffset() - 4);
        if (fbeFullSize < model.fbeSize())
            return false;

        return model.verify();
    }

    // Create a new model (begin phase)
    public long createBegin()
    {
        long fbeBegin = getBuffer().allocate(4 + model.fbeSize());
        return fbeBegin;
    }

    // Create a new model (end phase)
    public long createEnd(long fbeBegin)
    {
        long fbeEnd = getBuffer().getSize();
        int fbeFullSize = (int)(fbeEnd - fbeBegin);
        write(model.fbeOffset() - 4, fbeFullSize);
        return fbeFullSize;
    }

    // Serialize the struct value
    public long serialize(StructMap value)
    {
        long fbeBegin = createBegin();
        model.set(value);
        long fbeFullSize = createEnd(fbeBegin);
        return fbeFullSize;
    }

    // Deserialize the struct value
    public StructMap deserialize() { var value = new StructMap(); deserialize(value); return value; }
    public long deserialize(StructMap value)
    {
        if ((getBuffer().getOffset() + model.fbeOffset() - 4) > getBuffer().getSize())
        {
            value = new StructMap();
            return 0;
        }

        int fbeFullSize = readInt32(model.fbeOffset() - 4);
        assert (fbeFullSize >= model.fbeSize()) : "Model is broken!";
        if (fbeFullSize < model.fbeSize())
        {
            value = new StructMap();
            return 0;
        }

        value = model.get(value);
        return fbeFullSize;
    }

    // Move to the next struct value
    public void next(long prev)
    {
        model.fbeShift(prev);
    }
}