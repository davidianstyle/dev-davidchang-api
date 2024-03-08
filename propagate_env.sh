#!/bin/bash

# Check if .env file exists
if [ -e .env ]; then
    # Copy .env file to all subdirectories
    for dir in */; do
        # Ensure the directory is not the current directory
        if [ "$dir" != "./" ]; then
            # Copy .env to subdirectory
            cp .env "$dir"
            echo "Copied .env to $dir"
        fi
    done

    # Source .env in the root directory
    source .env
    echo "Sourced .env in the root directory."
else
    echo ".env file not found. Please make sure the file exists in the current directory."
fi
