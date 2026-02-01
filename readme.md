# Full-stack Application with dockercompose.yaml

### create a .env file in a isnside the root project folder

```sh
POSTGRES_USER=appuser
POSTGRES_PASSWORD=changeme
POSTGRES_DB=appdb
``` 
you can give the value as per your likings also.

### create a *.gitignore* file in your project's root folder and add .env to it

```sh
echo .env > .gitignore
```

### Start the docker-compose and build the stack

```sh
docker compose up -d --build
```
### visit the localhost to see the application running