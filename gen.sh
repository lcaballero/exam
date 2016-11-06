#!/bin/bash

GOROOT=/opt/go1.6


function pooled_buffer() {
	mockgen -source assert/exam.go \
		-destination assert/exam_mock_test.go -package assert
}

function all() {
	pooled_buffer
}

function go_install() {
	go install
	if [ "$?" != "0" ]; then
		echo "repository doesn't compile."
		exit 1
	fi
}

function go_test() {
	go test --race ./... 
	if [ "$?" != "0" ]; then
		echo "all tests build and pass."
		exit 1
	fi
}

function go_fmt() {
	gofmt -w .	
}

function go_vet() {
	go vet
	if [ "$?" != "0" ]; then
		echo "making sure go code doesn't have any suspicious constructs"
		exit 1
	fi
}

function go_misspell() {
	misspell ./**/*
	if [ "$?" != "0" ]; then
		echo "code has misspellings in comments or strings, etc"
		exit 1
	fi
}

function go_assign() {
	ineffassign .
	if [ "$?" != "0" ]; then
		echo "passes checks for ineffectual assignments"
		exit 1
	fi
}

function go_cyclo() {
	gocyclo -over 15 .
	if [ "$?" != "0" ]; then
		echo "cyclo complexity is too hight (max 15)"
		exit 1
	fi
}

function go_lint() {
#	find $(pwd) \( ! -regex '.*\.git.*' -and ! -regex '.*\.idea.*' -and -regex '.*\.go$' -and ! -regex '.*_test\.go$' \) -type f -exec golint {} \;
	go_src=$( find $(pwd) \( ! -regex '.*\.git.*' -and ! -regex '.*\.idea.*' -and -regex '.*\.go$' -and ! -regex '.*_test\.go$' \) -type f )
	for g in $go_src ; do
		golint -min_confidence 0.3 "$g"
		if [ "$?" != "0" ]; then
			echo "golint found errors $g"
			exit 1
		fi		
	done
}	

function checks() {
	go_install
	go_test
	go_fmt
	go_vet
	go_misspell
	go_assign
	go_cyclo
	go_lint
}

$1

