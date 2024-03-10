package de.dklotz.dsappbackend.repositories

import de.dklotz.dsappbackend.entities.Variable
import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.stereotype.Repository
import java.util.UUID

@Repository
interface VariableRepository : JpaRepository<Variable, UUID> {
}