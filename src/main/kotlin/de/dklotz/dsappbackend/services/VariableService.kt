package de.dklotz.dsappbackend.services

import de.dklotz.dsappbackend.entities.ControlType
import de.dklotz.dsappbackend.entities.DataType
import de.dklotz.dsappbackend.entities.Variable
import de.dklotz.dsappbackend.objects.VariableDTORequest
import de.dklotz.dsappbackend.repositories.VariableRepository
import org.springframework.stereotype.Service
import java.util.*

@Service
class VariableService(var repository: VariableRepository) {
    fun saveVariable(newVar: VariableDTORequest): Variable {
        return repository.save(
                Variable(
                    id = newVar.id ?: UUID.randomUUID(),
                    name = newVar.name,
                    controlType = ControlType.getControlTypeForIndex(newVar.controlTypeId),
                    dataType = DataType.getDataTypeForIndex(newVar.dataTypeId),
                    isEditable = newVar.isEditable,
                    isRequired = newVar.isRequired,
                    values = newVar.values,
                )
            )
    }

    fun deleteVariable(id: UUID) {
        repository.deleteById(id)
    }
}