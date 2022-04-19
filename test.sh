#!/bin/bash

set -e

docker build -t affluences-atelier-tech-2022 .
docker run --rm -v $(pwd):/app affluences-atelier-tech-2022
