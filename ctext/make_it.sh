#!/usr/bin/env bash

gcc -c -o sum.o sum.c
ar -rsc libsum.a sum.o
