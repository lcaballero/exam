#!/bin/bash

GOROOT=/opt/go1.6


function pooled_buffer() {
	mockgen -source assert/t2.go \
		-destination assert/t2_mock_test.go -package assert
}

function all() {
	pooled_buffer
}


$1

