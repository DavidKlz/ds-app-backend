package de.dklotz.dsappbackend.entities

enum class ControlType(val index: Int) {
        TEXT_FIELD(1), TEXT_AREA(2), CALENDAR(3), DROPDOWN(4), CHECKBOX(5), RADIO(6);

        companion object {
                fun getControlTypeForIndex(index: Int) : ControlType {
                        return entries.first { it.index == index }
                }
        }
}