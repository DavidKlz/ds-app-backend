package de.dklotz.dsappbackend.objects

import com.fasterxml.jackson.annotation.JsonProperty
import java.util.*

data class FormularDTOResponse(
        @JsonProperty("id")
        var id: UUID,
        @JsonProperty("name")
        var name: String,
        @JsonProperty("variables")
        val variables: Set<VariableDTOResponse>,
)