#/bin/bash
set -ex

rm -rf classtest
go generate
nose2 -v

rm -rf classtest
