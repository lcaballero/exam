#!/bin/bash

GOROOT=/opt/go1.6


function pooled_buffer() {
	mockgen -source assert/exam.go \
		-destination assert/exam_mock_test.go -package assert
}

function all() {
	pooled_buffer
}

$1

