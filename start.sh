docker-compose up -d --build
docker-compose exec -it cli /bin/sh -c "go run cli.go"
echo "Closing down containers"
docker-compose down