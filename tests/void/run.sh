#!/bin/sh
scriptdir=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
cd "$scriptdir"
docker build --no-cache -t d4:void . && docker run -e DISTRODETECT --rm --name d4_void d4:void
