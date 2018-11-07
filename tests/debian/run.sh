#!/bin/sh
cd "$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)"
docker build --no-cache -t d4:debian . && docker run -e DISTRODETECT --rm --name d4_debian d4:debian
