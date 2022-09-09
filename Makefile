BIN_DIR = bin

ifeq ($(OS), Windows_NT)
        SHELL := powershell.exe
        .SHELLFLAGS := -NoProfile -Command
        SHELL_VERSION = $(shell (Get-Host | Select-Object Version | Format-Table -HideTableHeaders | Out-String).Trim())
        OS = $(shell "{0} {1}" -f "windows", (Get-ComputerInfo -Property OsVersion, OsArchitecture | Format-Table -HideTableHeaders | Out-String).Trim())
        PACKAGE = $(shell (Get-Content go.mod -head 1).Split(" ")[1])
        CHECK_DIR_CMD = if (!(Test-Path $@)) { $$e = [char]27; Write-Error "$$e[31mDirectory $@ doesn't exist$${e}[0m" }
        RM_F_CMD = Remove-Item -erroraction silentlycontinue -Force
        RM_RF_CMD = ${RM_F_CMD} -Recurse
		EXT := .exe
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
        CHECK_DIR_CMD = test -d $@ || (echo "\033[31mDirectory $@ doesn't exist\033[0m" && false)
        RM_F_CMD = rm -f
        RM_RF_CMD = ${RM_F_CMD} -r
endif

.PHONY: api

all: terminal


terminal:
	go build -ldflags "-s -w" -o ${BIN_DIR}/omniterm${EXT} ./cmd/omniterm


omniterm:
	go run ./cmd/omniterm/main.go


test: all
	go test ./...

clean:
	${RM_F_CMD} ${BIN_DIR}/omniterm${EXT}

rebuild: clean all

about:
	@echo "OS: ${OS}"
	@echo "Shell: ${SHELL} ${SHELL_VERSION}"
	@echo "Protoc version: $(shell protoc --version)"
	@echo "Go version: $(shell go version)"
	@echo "Go package: ${PACKAGE}"
	@echo "OpenSSL version: $(shell openssl version)"
