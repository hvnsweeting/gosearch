#!/bin/bash

set -x
go get -u github.com/avelino/awesome-go
AWESOME_DIR="$(go list -f '{{ .Dir }}' github.com/avelino/awesome-go)"
mkdir -p data
cp "$AWESOME_DIR/README.md" data/
if ! ( command -v go-bindata ); then
    go get github.com/jteeuwen/go-bindata
fi

go-bindata data/

if ( git diff --shortstat bindata.go | grep '1 file changed, 1 insertion(+), 1 deletion(-)' ); then
    echo "ONLY TIMESTAMP CHANGED, reverting the change"
    git checkout bindata.go
fi
