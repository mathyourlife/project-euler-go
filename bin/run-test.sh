#!/bin/bash

set -eu

EULER_PROBLEM="$1"
export EULER_PROBLEM

go test -race -timeout 1m -run TestSolutions -v github.com/mathyourlife/project-euler-go/pkg/problems
