#!/bin/sh

echo "Building the server..."
(cd server && go build -o ./bin/server .)

echo "Building the frontend..."
(cd buldan-front && npm install && npm run build)

echo "Starting the server..."
./server/bin/server ./buldan-front/dist/ $1
