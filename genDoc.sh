#!/usr/bin/env bash

rm -rf ./doc
goctl api doc --dir ./myapi/ --o ./doc/