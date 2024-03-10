package de.dklotz.dsappbackend.entities

enum class DataType(val index: Int) {
        TEXT(1), NUMBER(2), DATE(3), BOOL(4), SELECTION(5), MULTI_SELECTION(6);


        companion object {
                fun getDataTypeForIndex(index: Int) : DataType {
                        return entries.first { it.index == index }
                }
        }
}