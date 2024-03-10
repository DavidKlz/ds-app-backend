package de.dklotz.dsappbackend.entities

import jakarta.persistence.*

import java.util.UUID

@Entity
data class Variable(
        @Id @GeneratedValue(strategy = GenerationType.UUID)
        val id: UUID,
        val isRequired: Boolean,
        val isEditable: Boolean,
        val name: String,
        @Enumerated(EnumType.ORDINAL)
        val controlType: ControlType,
        @Enumerated(EnumType.ORDINAL)
        val dataType: DataType)
