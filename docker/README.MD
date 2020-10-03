# docker 101

# DOCKERFILE references: https://docs.docker.com
# Docker hub: https://hub.docker.com/

# BASH aliases
alias ds='docker start $(docker ps --all -q -f name=$1)'		
alias db='docker build --rm'		
alias drmc='docker rm $(docker ps -q -f status=exited)'		
alias drmi='docker rmi $(docker images -q -f dangling=true)'		
alias dps="docker ps --all"		
alias d="docker"		
alias dc="docker-compose"

# Docker build commans
docker image build -t myapp .
docker container run -p 8888:80 myapp

# Docker hub login
docker login --username=yourhubusername --email=youremail@company.com

# Uploading a container to the hub