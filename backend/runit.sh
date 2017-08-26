#!/usr/bin/env bash

#docker run --rm -v $(pwd)/data:/data -w /data contd/say "Hello there"
docker run -d -p 8080:8080 contd/say