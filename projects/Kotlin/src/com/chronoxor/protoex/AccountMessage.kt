// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.3.0.0

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.protoex

@Suppress("MemberVisibilityCanBePrivate", "RemoveRedundantCallsOfConversionMethods")
open class AccountMessage : Comparable<Any?>
{
    var body: Account

    @Transient open var fbeType: Long = 13

    constructor(body: Account = Account())
    {
        this.body = body
    }

    @Suppress("UNUSED_PARAMETER")
    constructor(other: AccountMessage)
    {
        this.body = other.body
    }

    open fun clone(): AccountMessage
    {
        // Serialize the struct to the FBE stream
        val writer = com.chronoxor.protoex.fbe.AccountMessageModel()
        writer.serialize(this)

        // Deserialize the struct from the FBE stream
        val reader = com.chronoxor.protoex.fbe.AccountMessageModel()
        reader.attach(writer.buffer)
        return reader.deserialize()
    }

    override fun compareTo(other: Any?): Int
    {
        if (other == null)
            return -1

        if (!AccountMessage::class.java.isAssignableFrom(other.javaClass))
            return -1

        @Suppress("UNUSED_VARIABLE")
        val obj = other as AccountMessage? ?: return -1

        @Suppress("VARIABLE_WITH_REDUNDANT_INITIALIZER", "CanBeVal", "UnnecessaryVariable")
        var result = 0
        return result
    }

    override fun equals(other: Any?): Boolean
    {
        if (other == null)
            return false

        if (!AccountMessage::class.java.isAssignableFrom(other.javaClass))
            return false

        @Suppress("UNUSED_VARIABLE")
        val obj = other as AccountMessage? ?: return false

        return true
    }

    override fun hashCode(): Int
    {
        @Suppress("CanBeVal", "UnnecessaryVariable")
        var hash = 17
        return hash
    }

    override fun toString(): String
    {
        val sb = StringBuilder()
        sb.append("AccountMessage(")
        sb.append("body="); sb.append(body)
        sb.append(")")
        return sb.toString()
    }

    open fun toJson(): String = com.chronoxor.protoex.fbe.Json.engine.toJson(this)

    companion object
    {
        const val fbeTypeConst: Long = 13
        fun fromJson(json: String): AccountMessage = com.chronoxor.protoex.fbe.Json.engine.fromJson(json, AccountMessage::class.java)
    }
}
