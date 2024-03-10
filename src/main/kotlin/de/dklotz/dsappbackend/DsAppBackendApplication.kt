package de.dklotz.dsappbackend

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class DsAppBackendApplication

fun main(args: Array<String>) {
    runApplication<DsAppBackendApplication>(*args)
}
