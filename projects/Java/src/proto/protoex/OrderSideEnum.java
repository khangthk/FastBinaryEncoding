// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding

package protoex;

import java.io.*;
import java.lang.*;
import java.lang.reflect.*;
import java.math.*;
import java.nio.charset.*;
import java.time.*;
import java.util.*;
import fbe.*;
import proto.*;

public enum OrderSideEnum
{
    buy((byte)0 + 0)
    , sell((byte)0 + 1)
    , tell((byte)0 + 2)
    ;

    private byte value;

    OrderSideEnum(byte value) { this.value = value; }
    OrderSideEnum(int value) { this.value = (byte)value; }
    OrderSideEnum(OrderSideEnum value) { this.value = value.value; }

    public byte getRaw() { return value; }

    public static OrderSideEnum mapValue(byte value) { return mapping.get(value); }

    @Override
    public String toString()
    {
        if (this == buy) return "buy";
        if (this == sell) return "sell";
        if (this == tell) return "tell";
        return "<unknown>";
    }

    private static final Map<Byte, OrderSideEnum> mapping = new HashMap<>();
    static
    {
        for (var value : OrderSideEnum.values())
            mapping.put(value.value, value);
    }
}