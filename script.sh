#docker build namedockerimage foldername
docker build -t stevanisiska/docker-list-pets .

#docker build dengan menampilkan detail
docker build -t stevanisiska/docker-list-pets --progress=plain --no-cache

#create container docker
docker container create --name docker-list-pets --publish 8181:8181 stevanisiska/docker-list-pets