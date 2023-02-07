# Web and Software Architecture 
## Project 2022-23, "WASA Photo"

This is a project done through the Web and Software Architecture Course, 
taught at the Università degli Studi di Roma "La Sapienza".

### How To Build the WebApp

Please refer to following original commands:

If you're not using the WebUI, or if you don't want to embed the WebUI into the final executable, then:

```shell
go build ./cmd/webapi/
```

If you're using the WebUI and you want to embed it into the final executable:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run build-embed
exit
# (outside the NPM container)
go build -tags webui ./cmd/webapi/
```

### How To Run the WebApp

You can launch the backend only using:

```shell
go run ./cmd/webapi/
```

If you want to launch the WebUI, open a new tab and launch:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run dev
```

### How to build container images

#### Backend

```sh
$ docker build -t wasa-photos-backend:latest -f Dockerfile.backend .
```

#### Frontend

```sh
$ docker build -t wasa-photos-frontend:latest -f Dockerfile.frontend .
```

### How to run container images

#### Backend

```sh
$ docker run -it --rm -p 3000:3000 wasa-photos-backend:latest
```

#### Frontend

```
$ docker run -it --rm -p 8081:80 wasa-photos-frontend:latest
```


### Logger of main changes:
V1. Enrolling so to obtain an SSH Key.\
V2. Main structure unfolded.\
V3. OpenAPI built, extension of the README.md file.\
V4.

## License

See [LICENSE](LICENSE).
