#!/bin/bash

set -eu

EULER_PROBLEM="$1"
export EULER_PROBLEM

go test -count=1 -race -timeout 1m -run TestSolutions -v github.com/mathyourlife/project-euler-go/pkg/problems
