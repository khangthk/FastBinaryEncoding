// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding

@file:Suppress("UnusedImport", "unused")

package protoex

import java.io.*
import java.lang.*
import java.lang.reflect.*
import java.math.*
import java.nio.charset.*
import java.time.*
import java.util.*
import fbe.*
import proto.*

@Suppress("EnumEntryName", "MemberVisibilityCanBePrivate", "RemoveRedundantCallsOfConversionMethods")
enum class StateExEnum
{
    unknown(0x00 + 0)
    , invalid(0x01 + 0)
    , initialized(0x02 + 0)
    , calculated(0x04 + 0)
    , broken(0x08 + 0)
    , happy(0x10 + 0)
    , sad(0x20 + 0)
    , good(initialized.raw.toInt() or calculated.raw.toInt())
    , bad(unknown.raw.toInt() or invalid.raw.toInt() or broken.raw.toInt())
    ;

    var raw: Byte = 0
        private set

    constructor(value: Byte) { this.raw = value.toByte() }
    constructor(value: Short) { this.raw = value.toByte() }
    constructor(value: Int) { this.raw = value.toByte() }
    constructor(value: Long) { this.raw = value.toByte() }
    constructor(value: StateExEnum) { this.raw = value.raw }

    fun hasFlags(flags: Byte): Boolean = ((raw.toInt() and flags.toInt()) != 0) && ((raw.toInt() and flags.toInt()) == flags.toInt())
    fun hasFlags(flags: StateExEnum): Boolean = hasFlags(flags.raw)

    val allSet: EnumSet<StateExEnum> get() = EnumSet.allOf(StateExEnum::class.java)
    val noneSet: EnumSet<StateExEnum> get() = EnumSet.noneOf(StateExEnum::class.java)
    val currentSet: EnumSet<StateExEnum> get()
    {
        val result = EnumSet.noneOf(StateExEnum::class.java)
        if ((raw.toInt() and unknown.raw.toInt()) != 0)
        {
            result.add(unknown)
        }
        if ((raw.toInt() and invalid.raw.toInt()) != 0)
        {
            result.add(invalid)
        }
        if ((raw.toInt() and initialized.raw.toInt()) != 0)
        {
            result.add(initialized)
        }
        if ((raw.toInt() and calculated.raw.toInt()) != 0)
        {
            result.add(calculated)
        }
        if ((raw.toInt() and broken.raw.toInt()) != 0)
        {
            result.add(broken)
        }
        if ((raw.toInt() and happy.raw.toInt()) != 0)
        {
            result.add(happy)
        }
        if ((raw.toInt() and sad.raw.toInt()) != 0)
        {
            result.add(sad)
        }
        if ((raw.toInt() and good.raw.toInt()) != 0)
        {
            result.add(good)
        }
        if ((raw.toInt() and bad.raw.toInt()) != 0)
        {
            result.add(bad)
        }
        return result
    }

    override fun toString(): String
    {
        val sb = StringBuilder()
        var first = true
        if (hasFlags(unknown))
        {
            sb.append(if (first) "" else "|").append("unknown")
            @Suppress("UNUSED_VALUE")
            first = false
        }
        if (hasFlags(invalid))
        {
            sb.append(if (first) "" else "|").append("invalid")
            @Suppress("UNUSED_VALUE")
            first = false
        }
        if (hasFlags(initialized))
        {
            sb.append(if (first) "" else "|").append("initialized")
            @Suppress("UNUSED_VALUE")
            first = false
        }
        if (hasFlags(calculated))
        {
            sb.append(if (first) "" else "|").append("calculated")
            @Suppress("UNUSED_VALUE")
            first = false
        }
        if (hasFlags(broken))
        {
            sb.append(if (first) "" else "|").append("broken")
            @Suppress("UNUSED_VALUE")
            first = false
        }
        if (hasFlags(happy))
        {
            sb.append(if (first) "" else "|").append("happy")
            @Suppress("UNUSED_VALUE")
            first = false
        }
        if (hasFlags(sad))
        {
            sb.append(if (first) "" else "|").append("sad")
            @Suppress("UNUSED_VALUE")
            first = false
        }
        if (hasFlags(good))
        {
            sb.append(if (first) "" else "|").append("good")
            @Suppress("UNUSED_VALUE")
            first = false
        }
        if (hasFlags(bad))
        {
            sb.append(if (first) "" else "|").append("bad")
            @Suppress("UNUSED_VALUE")
            first = false
        }
        return sb.toString()
    }

    companion object
    {
        private val mapping = HashMap<Byte, StateExEnum>()

        init
        {
            for (value in StateExEnum.values())
                mapping[value.raw] = value
        }

        fun mapValue(value: Byte): StateExEnum? { return mapping[value] }
    }
}