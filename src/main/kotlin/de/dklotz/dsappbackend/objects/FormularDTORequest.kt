package de.dklotz.dsappbackend.objects

import com.fasterxml.jackson.annotation.JsonProperty

data class FormularDTORequest(
        @JsonProperty("name")
        var name: String,
        @JsonProperty("variables")
        val variables: Set<VariableDTORequest>,
)