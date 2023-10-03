#!/usr/bin/env bash

set -e

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)

commands=(
	"go run main.go"
	"go run main.go --help"
	"go run main.go version"
	"go run main.go version --help"
)

function execute_command() {
	for command in "${commands[@]}"; do
		echo -e "\e[32m"
		echo "**************************************************"
		printf "* %-46s *\n" "$command"
		echo "**************************************************"
		echo -e "\e[0m"

		eval "$command"
	done
}

function main() {
	cd "${SCRIPT_DIR}"
	execute_command
}

main
