#!/bin/sh
cd "$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)"
docker build -t d4:void .
docker run -i -t d4:void /bin/bash
