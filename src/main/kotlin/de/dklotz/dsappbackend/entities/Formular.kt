package de.dklotz.dsappbackend.entities

import jakarta.persistence.CascadeType
import jakarta.persistence.Entity
import jakarta.persistence.GeneratedValue
import jakarta.persistence.GenerationType
import jakarta.persistence.Id
import jakarta.persistence.OneToMany
import java.util.UUID

@Entity
data class Formular(
    @Id @GeneratedValue(strategy = GenerationType.UUID)
    val id: UUID,
    val name: String,
    @OneToMany(cascade = [CascadeType.ALL])
    val variables: Set<Variable>)