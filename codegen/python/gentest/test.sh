#/bin/bash

rm -rf classtest
go generate
nose2 -v

rm -rf classtest
