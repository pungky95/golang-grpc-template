PROTO_DIR = proto
DATABASE_DEV = cockroachdb://pungky:EyliO7ppTtcMvmlNNTWiNw@be-eazy-2242.8nk.cockroachlabs.cloud:26257/user_db?sslmode=verify-full

ifeq ($(OS), Windows_NT)
	SHELL := powershell.exe
	.SHELLFLAGS := -NoProfile -Command
	SHELL_VERSION = $(shell (Get-Host | Select-Object Version | Format-Table -HideTableHeaders | Out-String).Trim({}))
	OS = $(shell "{0} {1}" -f "windows", (Get-ComputerInfo -Property OsVersion, OsArchitecture | Format-Table -HideTableHeaders | Out-String).Trim({}))
	PACKAGE = $(shell (Get-Content go.mod -head 1).Split(" ")[1])
	HELP_CMD = Select-String "^[a-zA-Z_-]+:.*?\#\# .*$$" "./Makefile" | Foreach-Object { $$_data = $$_.matches -split ":.*?\#\# "; $$obj = New-Object PSCustomObject; Add-Member -InputObject $$obj -NotePropertyName ('Command') -NotePropertyValue $$_data[0]; Add-Member -InputObject $$obj -NotePropertyName ('Description') -NotePropertyValue $$_data[1]; $$obj } | Format-Table -HideTableHeaders @{Expression={ $$e = [char]27; "$$e[36m$$($$_.Command)$${e}[0m" }}, Description
	RM_F_CMD = Remove-Item -erroraction silentlycontinue -Force
	RM_RF_CMD = ${RM_F_CMD} -Recurse
else
	SHELL := bash
	SHELL_VERSION = $(shell echo $$BASH_VERSION)
	UNAME := $(shell uname -s)
	VERSION_AND_ARCH = $(shell uname -rm)
	ifeq ($(UNAME),Darwin)
		OS = macos ${VERSION_AND_ARCH}
	else ifeq ($(UNAME),Linux)
		OS = linux ${VERSION_AND_ARCH}
	else
    $(error OS not supported by this Makefile)
	endif
	PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
	HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	RM_F_CMD = rm -f
	RM_RF_CMD = ${RM_F_CMD} -r
endif

.DEFAULT_GOAL := help

generate-graphql: ## Generate GraphQL example: generate-graphql
	go run github.com/99designs/gqlgen generate

init-graphql: ## Init GraphQL example: init-graphql
	go run github.com/99designs/gqlgen init

about: ## Display info related to the build
	@echo "OS: ${OS}"
	@echo "Shell: ${SHELL} ${SHELL_VERSION}"
	@echo "Protoc version: $(shell protoc --version)"
	@echo "Go version: $(shell go version)"
	@echo "Go package: ${PACKAGE}"
	@echo "Openssl version: $(shell openssl version)"

help: ## Show this help
	@${HELP_CMD}

generate-migration: ## Generate migration file example: generate-migration name=migration_name
	@[ "${name}" ] || ( echo "migration name not set"; exit 1 )
	migrate create -ext sql -dir migrations $(name)

migrate-up-dev: ## Migrate Up example: migrate-up database="postgres://username:password@host:port/db_name?sslmode=disable"
	migrate -path migrations -database "${DATABASE_DEV}" -verbose up $(N)

migrate-down-dev: ## Migrate Down example: migrate-down database="postgres://username:password@host:port/db_name?sslmode=disable"
	migrate -path migrations -database "${DATABASE_DEV}" -verbose down $(N)

db-reset-dev: migrate-down-dev migrate-up-dev
	

.PHONY: migrate-up-dev migrate-down-up db-reset-dev
