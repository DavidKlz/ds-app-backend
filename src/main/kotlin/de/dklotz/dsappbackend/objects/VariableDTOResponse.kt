package de.dklotz.dsappbackend.objects

import com.fasterxml.jackson.annotation.JsonProperty
import java.util.*

data class VariableDTOResponse (
        @JsonProperty("id")
        val id: UUID,
        @JsonProperty("isRequired")
        val isRequired:Boolean,
        @JsonProperty("isEditable")
        val isEditable:Boolean,
        @JsonProperty("name")
        val name: String,
        @JsonProperty("values")
        val values: List<String>,
        @JsonProperty("controlTypeId")
        val controlTypeId: Int,
        @JsonProperty("dataTypeId")
        val dataTypeId: Int
)