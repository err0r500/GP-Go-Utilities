#!/usr/bin/env bash

brew list openssl || brew install openssl

openssl genrsa -out cert/private.key 4096