#!/bin/bash

# Check if .env.example exists
if [ -e .env.example ]; then
    # Replace '[' and ']' with ''
    sed 's/\[//g; s/\]//g' .env.example > .env

    echo "Copied .env.example to .env, replacing '[' and ']' characters."
else
    echo ".env.example not found. Please make sure the file exists in the current directory."
fi
