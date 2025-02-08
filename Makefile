# -------------------------------------------------------------------------------
# Makefile for go project
# -------------------------------------------------------------------------------

# -------------------------------------------------------------------------------
# variables
# -------------------------------------------------------------------------------
MAIN_FILE := main.go

BIN_DIR := ./bin
APP := github520cli # the name of the application

# the name of the docker image
# DOCKER_IMAGE_NAME := docker-image-name-example

# args
STATIC_ARGS := -ldflags='-extldflags "-static"'
RACE_ARGS := -race

# this dir is for the test output when u run `test` cmd
TEST_OUTPUT_DIR := ./test_output

# color code
GREEN_COLOR_CODE_HEAD := \033[32m
GREEN_COLOR_CODE_END := \033[0m
RED_COLOR_CODE_HEAD := \033[31m
RED_COLOR_CODE_END := \033[0m

# -------------------------------------------------------------------------------
# commands
# -------------------------------------------------------------------------------

# build the application
.PHONY: build
build:
	go build -o $(BIN_DIR)/$(APP) $(MAIN_FILE)

# build the application statically
# you can use this if you want to run the application on a system without a C compiler
# if you use glibc, you can also use the static version of glibc(like musl)
.PHONY: build-static
build-static:
	go build $(STATIC_ARGS) -o $(BIN_DIR)/$(APP) $(MAIN_FILE)

# build the application with docker
# u need to set the DOCKER_IMAGE_NAME variable
.PHONY: build-docker
build-docker:
	docker build -t $(DOCKER_IMAGE_NAME) .
	@$(show_info)"Docker image built successfully: $(GREEN_COLOR_CODE_HEAD)$(DOCKER_IMAGE_NAME)$(GREEN_COLOR_CODE_END)";

# run the application with race detector
# use this cmd when u develop
.PHONY: run
run:
	go run $(RACE_ARGS) $(MAIN_FILE)

# test application
# it will create a coverage.out file(for test coverage)
.PHONY: test
test:
	mkdir -p $(TEST_OUTPUT_DIR)
	go test -v -covermode=set -coverprofile=$(TEST_OUTPUT_DIR)/coverage.out ./...

# show the test coverage
# it will create a coverage.html file from coverage.out
# so, u need to run `test` cmd first
.PHONY: test-coverage
test-coverage: test
	go tool cover -html=$(TEST_OUTPUT_DIR)/coverage.out -o $(TEST_OUTPUT_DIR)/coverage.html
	@$(show_info)"Please open $(GREEN_COLOR_CODE_HEAD)$(TEST_OUTPUT_DIR)/coverage.html$(GREEN_COLOR_CODE_END) to see the test coverage"

# clean the application
# it will remove the bin directory and the test output directory
.PHONY: clean
clean:
	go clean
	rm -rf $(BIN_DIR) $(TEST_OUTPUT_DIR)

# -------------------------------------------------------------------------------
# functions
# -------------------------------------------------------------------------------

# echo time
TIMESTAMP_FORMAT := %Y-%m-%d %H:%M:%S
define timestamp
	$(shell date "+$(TIMESTAMP_FORMAT)")
endef

# show info
INFO_PREFIX := *INFO
define show_info
	@echo -e "$(GREEN_COLOR_CODE_HEAD)$(INFO_PREFIX):$(GREEN_COLOR_CODE_END)"
	@echo -e $(1)
endef

# show error
# it will exit the program
ERROR_PREFIX := *ERROR
define show_error
	@echo -e "$(RED_COLOR_CODE_HEAD)$(ERROR_PREFIX):$(RED_COLOR_CODE_END)"
	@echo -e $(1)
	exit 1
endef
