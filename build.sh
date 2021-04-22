#!/bin/sh

echo "Building svelte"
cd svelte
npm ci
npm run build
cd ..

echo "Building Golang"
go mod download
go build .

echo "Build successful"
