#!/bin/bash

gox -os="linux darwin windows freebsd" -arch="386 amd64" -output="{{.Dir}}-{{.OS}}-{{.Arch}}/{{.Dir}}"

for os in linux darwin windows freebsd; do
    for arch in 386 amd64; do
        # cp README.md LICENSE gojira-$os-$arch
        zip -r gojira-$os-$arch.zip gojira-$os-$arch
        rm -rf gojira-$os-$arch
    done
done