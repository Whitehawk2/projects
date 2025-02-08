#!/bin/env bash

make install
docker build -t projects/seniority:scratchrouter .
