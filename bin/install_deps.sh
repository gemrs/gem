#!/bin/bash
pip install --user pytest -q
pip install --user codecov -q
go get github.com/wadey/gocovmerge
go get github.com/go-playground/overalls
