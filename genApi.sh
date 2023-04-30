#!/usr/bin/env bash

goctl api go -api ./myapi/user/desc/*.api -dir ./myapi/user/ -style=goZero
goctl api go -api ./myapi/article/desc/*.api -dir ./myapi/article/ -style=goZero

