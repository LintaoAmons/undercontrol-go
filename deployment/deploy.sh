#!/bin/bash

cd "$(dirname "$0")"

directory="$HOME/.local/share/utlgo"
if [ ! -d "$directory" ]; then
  mkdir -p "$directory"
fi


kubectl apply -f deployment.yaml
