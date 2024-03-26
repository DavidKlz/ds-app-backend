package de.dklotz.dsappbackend.services

import de.dklotz.dsappbackend.entities.Formular
import de.dklotz.dsappbackend.objects.FormularDTORequest
import de.dklotz.dsappbackend.objects.FormularDTOResponse
import de.dklotz.dsappbackend.objects.VariableDTOResponse
import de.dklotz.dsappbackend.repositories.FormularRepository
import org.springframework.stereotype.Service
import java.util.*
import kotlin.jvm.optionals.getOrNull

@Service
class FormularService(var repository: FormularRepository, var variableService: VariableService) {

    fun createFormular(formular: FormularDTORequest) : FormularDTOResponse {
        val vars = formular.variables.map {
            variableService.saveVariable(it)
        }.toSet()

        val savedForm = repository.save(
            Formular(
                id = UUID.randomUUID(),
                name = formular.name,
                variables = vars,
            )
        )

        return FormularDTOResponse(
            id = savedForm.id,
            name = savedForm.name,
            variables = savedForm.variables.map {
                VariableDTOResponse(
                    id = it.id,
                    name = it.name,
                    isRequired = it.isRequired,
                    isEditable = it.isEditable,
                    controlTypeId = it.controlType.index,
                    dataTypeId = it.dataType.index,
                    values = it.values,
                )
            }.toSet()
        )
    }

    fun deleteFormular(id: UUID) {
        repository.findById(id).ifPresent {
            it.variables.forEach { variable ->
                variableService.deleteVariable(variable.id)
            }
        }

        repository.deleteById(id)
    }

    fun getFormulare() : List<FormularDTOResponse> {
        return repository.findAll().map { formular ->
            FormularDTOResponse(
                id = formular.id,
                name = formular.name,
                variables = formular.variables.map {variable ->
                    VariableDTOResponse(
                        id = variable.id,
                        isEditable = variable.isEditable,
                        controlTypeId = variable.controlType.index,
                        dataTypeId = variable.dataType.index,
                        name = variable.name,
                        isRequired = variable.isRequired,
                        values = variable.values,
                    )
                }.toSet()
            )
        }
    }

    fun getFormular(id: UUID) : FormularDTOResponse? {
        return repository.findById(id).map { formular ->
            FormularDTOResponse(
                id = formular.id,
                name = formular.name,
                variables = formular.variables.map {variable ->
                    VariableDTOResponse(
                        id = variable.id,
                        isEditable = variable.isEditable,
                        controlTypeId = variable.controlType.index,
                        dataTypeId = variable.dataType.index,
                        name = variable.name,
                        isRequired = variable.isRequired,
                        values = variable.values,
                    )
                }.toSet()
            )
        }.getOrNull()
    }

    fun updateFormular(id: UUID, updatedFormular: FormularDTORequest) : FormularDTOResponse {
        val updatedVars = updatedFormular.variables.map {
            variableService.saveVariable(it)
        }.toSet()

        val updatedForm = repository.save(Formular(
            id = id,
            name = updatedFormular.name,
            variables = updatedVars,
        ))

        return FormularDTOResponse(
            id = updatedForm.id,
            name = updatedForm.name,
            variables = updatedForm.variables.map {
                VariableDTOResponse(
                    id = it.id,
                    name = it.name,
                    isRequired = it.isRequired,
                    isEditable = it.isEditable,
                    controlTypeId = it.controlType.index,
                    dataTypeId = it.dataType.index,
                    values = it.values,
                )
            }.toSet()
        )
    }

}