#!/bin/sh
goose -dir migrations postgres "postgres://wheels:wheels@localhost:5432/wheels" up
