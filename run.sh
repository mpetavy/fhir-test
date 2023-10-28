#!/bin/sh

# https://smilecdr.com/docs/fhir_standard/fhir_crud_operations.html

curl -v -k -X POST -H 'Content-Type: application/json; charset=utf-8' --data @patient.json https://hapi.fhir.org/baseR4/Patient