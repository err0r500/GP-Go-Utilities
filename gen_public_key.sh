#!/usr/bin/env bash

openssl req -new -config cert/local.conf -x509 -sha256 -days 1825 -key cert/private.key -out cert/public.crt

openssl x509 -in cert/public.crt -text -noout
