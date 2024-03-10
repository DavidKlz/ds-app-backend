package de.dklotz.dsappbackend.repositories

import de.dklotz.dsappbackend.entities.Formular
import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.stereotype.Repository
import java.util.UUID

@Repository
interface FormularRepository : JpaRepository<Formular, UUID> {

}