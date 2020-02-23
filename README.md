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

## Install

```sh
make build
```

## Usage

```sh
make deploy
```

## Run tests

```sh
make test
```

## Creating new function

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