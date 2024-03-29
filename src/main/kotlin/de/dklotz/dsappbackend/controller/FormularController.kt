package de.dklotz.dsappbackend.controller

import de.dklotz.dsappbackend.objects.FormularDTORequest
import de.dklotz.dsappbackend.objects.FormularDTOResponse
import de.dklotz.dsappbackend.services.FormularService
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.DeleteMapping
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.PutMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController
import org.springframework.web.server.ResponseStatusException
import java.util.UUID

@RestController
@RequestMapping("/formular")
class FormularController(var service: FormularService) {
    @GetMapping("/all")
    fun getFormulare() : List<FormularDTOResponse> {
        return service.getFormulare();
    }

    @PostMapping("/create")
    fun createFormular(@RequestBody newForm: FormularDTORequest) : FormularDTOResponse {
        return service.createFormular(newForm)
    }

    @GetMapping("/{id}")
    fun getFormular(@PathVariable id: UUID) : FormularDTOResponse {
        return service.getFormular(id) ?: throw ResponseStatusException(HttpStatus.NOT_FOUND, "Formular not found")
    }

    @DeleteMapping("/{id}")
    fun deleteFormular(@PathVariable id: UUID) {
        service.deleteFormular(id)
    }

    @PutMapping("/{id}")
    fun updateFormular(@PathVariable id: UUID, @RequestBody updatedForm: FormularDTORequest): FormularDTOResponse {
        return service.updateFormular(id, updatedForm)
    }
}