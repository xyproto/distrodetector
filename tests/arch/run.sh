#!/bin/sh
cd "$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)"
docker build --no-cache -t d4:arch . && docker run -e DISTRODETECT --rm --name d4_arch d4:arch
