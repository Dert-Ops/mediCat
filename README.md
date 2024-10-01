# catmedi

Full stack microservice social media app project, like yor dreams...

# Usefully Shortcuts

all volumes delete : docker volume rm $(docker volume ls -q)

database quick connection : docker exec -it postgres-master psql -U admin1 -d medicat_db -p 5432

database quick connection2 : psql -h localhost -U admin1 -d medicat_db -p 5432

down, build and compose up single make command : make re 


# Node.js Installation for MacOS

1-) Install NVM (Node Version Manager)

curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash

2-) Download and install Node.js

nvm install 20

## restart your terminal or "source ~/.bashrc" or "source ~/.zshrc"

3-) verifies the right Node.js version

node -v

4-) verifies the right npm version

npm -v

# Vite Installation for MacOS with NPM

npm install -D vite

# Run with developer mode

npm run dev