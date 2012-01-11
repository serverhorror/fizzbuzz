# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include ${GOROOT}/src/Make.inc

TARG=fizzbuzz
GOFILES=\
	fizzbuzz.go\

include ${GOROOT}/src/Make.pkg

## Convenience Targets

run: main
	./main

help: main
	./main -h

main: main.6
	6l -o main main.6

main.6: main.go fizzbuzz.6
	6g -o main.6 main.go

fizzbuzz.6: ${GOFILES}
	6g -o fizzbuzz.6 ${GOFILES}

