![Logo](https://cdn-images-1.medium.com/max/1600/1*OezhU9lHTNCk6O6FCUL5fQ.png)

<h1 align="center">Serverless Architecture Boilerplate (GoLang) 👋</h1>
<p>
  <img src="https://img.shields.io/badge/version-v0-blue.svg?cacheSeconds=2592000" />
  <a href="https://github.com/msfidelis/serverless-architecture-boilerplate-go">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" target="_blank" />
  </a>
  <a href="https://twitter.com/fidelissauro">
    <img alt="Twitter: fidelissauro" src="https://img.shields.io/twitter/follow/fidelissauro.svg?style=social" target="_blank" />
  </a>
</p>

> Boilerplate to organize and deploy big projects using AWS API Gateway and AWS Lambda with Serverless Framework

### 🏠 [Homepage](https://github.com/msfidelis/serverless-architecture-boilerplate-go)

## Structure

```
.
├── Dockerfile
├── Gopkg.lock
├── Gopkg.toml
├── README.md
├── bin (output for go binaries)
├── configs (configs folders for environment)
│   └── dev.yml 
├── go.mod
├── go.sum
├── modules (modules folder)
│   └── books (module / context)
│       ├── endpoints (api endpoints)
│       │   ├── create.go
│       │   ├── delete.go
│       │   ├── detail.go
│       │   ├── read.go
│       │   └── update.go
│       └── functions (workers / background functions)
│           └── worker.go
├── package-lock.json 
├── pkg (shared components)
│   ├── libs (libraries)
│   │   ├── dynamoclient
│   │   │   └── dynamoclient.go
│   │   └── sqsclient
│   │       └── sqsclient.go
│   └── models (models)
│       └── book
│           └── book.go
├── magefile.go
├── mage.go
└── serverless.yml
```

## Install

```sh
go run mage.go build
```

## Usage

```sh
go run mage.go deploy
```

## Run tests

```sh
go run mage.go test
```

# Testing boilerplate

### Listing books

```bash
curl -i https://5pas07ibi2.execute-api.us-east-1.amazonaws.com/dev/books
```

```json
HTTP/2 200 
content-type: application/json
content-length: 122213
date: Sun, 23 Feb 2020 19:42:16 GMT
x-amzn-requestid: 09b93530-dda7-46b0-a916-fcb4ae59c0bc
x-amz-apigw-id: IXZP3EspoAMFyFQ=
x-amzn-trace-id: Root=1-5e52d598-b446d86c871e237cf838922c;Sampled=0
x-cache: Miss from cloudfront
via: 1.1 37a135c363e9512a2b27aa63bc837339.cloudfront.net (CloudFront)
x-amz-cf-pop: GRU50-C1
x-amz-cf-id: vhEGO9RiFXPQw5q4GWRxd-w9l5Klg8NqI4u9-hbC231q9okFLI4Pbw==

[{"hashkey":"ccd31b13-4945-11ea-a024-da7fa31ac789","title":"D(jwItrDnA","author":"prSa(j)kSD","price":10.1,"updated"
:"","created":"","processed":true},{"hashkey":"cc867cb1-4945-11ea-8754-a2fe55138192","title":"pDy@eghdyS","author":"
m*ikSdohth","price":10.1,"updated":"","created":"","processed":true},{"hashkey":"c2ff73e9-4945-11ea-a024-da7fa31ac78
9","title":"DFse*wsa)(","author":"DpFoqHISd*","price":10.1,"updated":"","created":"","processed":true},{"hashkey":"8
e139c9a-4f9c-11ea-9f24-72f20a774bc0","title":"","author":"","price":0,"updated":"","created":"","processed":true},{"
hashkey":"6e3f7b60-4f9c-11ea-9f24-72f20a774bc0","title":"","author":"","price":0,"updated":"","created":"","processe
d":true},{"hashkey":"708117bf-4f9c-11ea-b2af-6e3ae06015d9","title":"","author":"","price":0,"updated":"","created":"
","processed":true}]
```

### Creating new Book 

```bash
curl -X POST  https://5pas07ibi2.execute-api.us-east-1.amazonaws.com/dev/books -H 'Content-type:application/json' -d '{"title": "American Gods", "price": 10.50, "author": "Neil Gaiman"}'
```


```json
curl -i -X POST  https://5pas07i
bi2.execute-api.us-east-1.amazonaws.com/dev/books -H 'Content-type:application/json' -d '{"title": "American Gods", 
"price": 10.50, "author": "Neil Gaiman"}' 
HTTP/2 201 
content-type: application/json
content-length: 265
date: Sun, 23 Feb 2020 19:45:41 GMT
x-amzn-requestid: 03f52228-3df8-413e-a22f-12877e41ebb3
x-amz-apigw-id: IXZv2FOXIAMFgDw=
x-amzn-trace-id: Root=1-5e52d665-f99746e82b4797d024a9d196;Sampled=0
x-cache: Miss from cloudfront
via: 1.1 aa3fc654df34a675869d8ecab4dd6bab.cloudfront.net (CloudFront)
x-amz-cf-pop: GRU50
x-amz-cf-id: NasPt3E-aQBmufDdZ3fVjCSaYYVHTnicDcyib1ZoAgqdkHCcwySxGA==

{"hashkey":"12bfd98c-5675-11ea-94ea-5ec3dff6689d","title":"American Gods","author":"Neil Gaiman","price":10.5,"updat
ed":"2020-02-23 19:45:41.394682371 +0000 UTC m=+329.476777607","created":"2020-02-23 19:45:41.39467731 +0000 UTC m=+
329.476772546","processed":false}
```

### Detail a book

```bash
curl -i https://5pas07ibi2.execu
te-api.us-east-1.amazonaws.com/dev/books/12bfd98c-5675-11ea-94ea-5ec3dff6689d
```

```json
HTTP/2 200 
content-type: application/json
content-length: 265
date: Sun, 23 Feb 2020 19:47:24 GMT
x-amzn-requestid: 7de185ab-a088-4bcd-8da8-5b8e249adfb7
x-amz-apigw-id: IXZ_0G4jIAMFfzA=
x-amzn-trace-id: Root=1-5e52d6cb-096a4604266412f06b182418;Sampled=0
x-cache: Miss from cloudfront
via: 1.1 356de6a1dc9aa6df67447cdc3e65d45e.cloudfront.net (CloudFront)
x-amz-cf-pop: GRU1-C1
x-amz-cf-id: cH8MwkQU9WLfuEkpNZCXYYS8QQGU0uQ5CsdmIfTEVn8lVrR9wMtoRg==

{"hashkey":"12bfd98c-5675-11ea-94ea-5ec3dff6689d","title":"American Gods","author":"Neil Gaiman","price":10.5,"updat
ed":"2020-02-23 19:45:41.526980165 +0000 UTC m=+328.914267950","created":"2020-02-23 19:45:41.39467731 +0000 UTC m=+
329.476772546","processed":false}%
```


### Updating a book 

```bash
curl -i -X PUT  https://5pas07
ibi2.execute-api.us-east-1.amazonaws.com/dev/books/12bfd98c-5675-11ea-94ea-5ec3dff6689d -H 'Content-type:application
/json' -d '{"title": "American Gods - Updated", "price": 20.00, "author": "Neil Gaiman"}'
```

```json
HTTP/2 200 
content-type: application/json
content-length: 179
date: Sun, 23 Feb 2020 19:50:51 GMT
x-amzn-requestid: 6720a008-07b5-4f0c-81de-3ee3bc2ed93b
x-amz-apigw-id: IXagXGkmoAMFwvA=
x-amzn-trace-id: Root=1-5e52d79b-b689cd30ffe1e9dc5ef25256;Sampled=0
x-cache: Miss from cloudfront
via: 1.1 a43a4e3a015929f71d6fc6fa15418703.cloudfront.net (CloudFront)
x-amz-cf-pop: GRU1-C1
x-amz-cf-id: 2W2jS7yoAbPsICQ7C3EMMKQl827HdbFPkb_WdsztzUwmstH7U4nIMw==

{"hashkey":"","title":"American Gods - Updated","author":"Neil Gaiman","price":20,"updated":"2020-02-23 19:50:51.896
7559 +0000 UTC m=+48.464494766","created":"","processed":false}
```


### Deleting a book

```bash
curl -i -X DELETE https://5pas
07ibi2.execute-api.us-east-1.amazonaws.com/dev/books/12bfd98c-5675-11ea-94ea-5ec3dff6689d
```

```json
HTTP/2 200 
content-type: application/json
content-length: 69
date: Sun, 23 Feb 2020 19:52:48 GMT
x-amzn-requestid: 19d8bf9a-9a9f-41f9-9558-481baae891d9
x-amz-apigw-id: IXaybEqAoAMFuAg=
x-amzn-trace-id: Root=1-5e52d80f-40cde16afd20a6ce870e7b0e;Sampled=0
x-cache: Miss from cloudfront
via: 1.1 32063733c6b1049f7b777e1f8ac028ad.cloudfront.net (CloudFront)
x-amz-cf-pop: GRU50
x-amz-cf-id: lNTlUeTq4RGRD1Us5cHgZ1spPqhzE91asXdSyci8k0IuISUuMOR_4w==

{"hashkey":"12bfd98c-5675-11ea-94ea-5ec3dff6689d","status":"deleted"}ls

```

# Creating a new function

### 1) create new function inside `modules` path

```bash
touch modules/mymodule/endpoints/myfunction.go
```

### 2) Add build instructions on Makefile

```bash
vim Makefile
```

```Makefile
build:
  dep ensure
  // ...
  env GOOS=linux go build -ldflags="-s -w" -o bin/mymodule/endpoints/myfunction modules/mymodule/endpoints/myfunction
```

### 3) Add function mapping on `serverless.yml` file

```bash
vim serverless.yml
```

```yml
# ...
functions:
  create:
    handler: bin/mymodule/endpoints/myfunction
    events:
      - http:
          path: /services/mypath
          method: get
    tags:
      TAGFUNCTION: Tag Value
```

### 4) Delete stack 

```bash
go run mage.go remove 
```

## Author

👤 **Matheus Fidelis**

* Twitter: [@fidelissauro](https://twitter.com/fidelissauro)
* Github: [@msfidelis](https://github.com/msfidelis)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/msfidelis/serverless-architecture-boilerplate-go/issues).

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
