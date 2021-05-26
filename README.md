# Chefkoch code challenge

This application was created as the result of a coding challenge for chefkoch.de. Its parameters were as follows:

- We want you to create a REST-service using go (https://golang.org/)
- The service shall manage ToDo's
- A ToDo consists of an arbitrary list of Subtasks and is structured as follows:
```
{
  id [mandatory]
  name [mandatory]
  description
  tasks: [
      {
          id [mandatory]
          name [mandatory]
          description
      }
  ]
}
```  
- The service shall serve the following endpoints:
    - GET /todos → Returns a list of all Todos
    - POST /todos → Expects a Todo (without id) and returns a Todo with id
    - GET /todos/{id} → Returns a Todo
    - PUT /todos/{id} → Overwrites an existing Todo
    - DELETE /todos/{id} → Deletes a Todo
- All ToDo's have to be persisted, the means are up to the applicant.

# Running the application

## Running with the Makefile

To run the application, run the command `make run`. The application will then be available on http://localhost:8080

To run the application with this command, one needs to install some tools:

Name | Installation | website 
--- | --- | --- 
GoMock | `GO111MODULE=on go get github.com/golang/mock/mockgen@v1.5.0` | https://github.com/golang/mock
Wire | `go get github.com/google/wire/cmd/wire` | https://github.com/google/wire
GolangCI Lint | `curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh` &#124; `sh -s -- -b $(go env GOPATH)/bin v1.40.1` | https://github.com/golangci/golangci-lint
Nancy | Manual from their github website | https://github.com/sonatype-nexus-community/nancy
Cobra | `go get -u github.com/spf13/cobra` | https://github.com/spf13/cobra
Docker | https://docs.docker.com/engine/install/ | https://www.docker.com/
Docker-Compose | https://docs.docker.com/compose/install/ | https://docs.docker.com/compose/
Golang | https://golang.org/doc/install | https://golang.org

## Running without the Makefile

```shell
#  The application will then be available on http://localhost:8080
docker build . -f build/Dockerfile -t todo:latest
docker-compose -f docker-compose/docker-compose.yml up
```

# Tools

The following is a description of all tools used: 

**Gomock** - 
is a mocking framework for the Go programming language. It integrates well with Go's built-in testing package, 
but can be used in other contexts too.

**GolangCI Lint** -
is a fast Go linters runner. It runs linters in parallel, uses caching, supports yaml config, has 
integrations with all major IDE and has dozens of linters included.

**Nancy** - 
is a tool to check for vulnerabilities in your Golang dependencies, powered by Sonatype OSS Index.  It works with Nexus
IQ Server, allowing you a smooth experience as a Golang developer, using the best tools in the market!

**Wire** -
is a code generation tool that automates connecting components using dependency injection. Dependencies between 
components are represented in Wire as function parameters, encouraging explicit initialization instead of global 
variables. Because Wire operates without runtime state or reflection, code written to be used with Wire is useful even 
for hand-written initialization.

**Cobra** -
is a library providing a simple interface to create powerful modern CLI interfaces similar to git & go tools.
Cobra is also an application that will generate your application scaffolding to rapidly develop a Cobra-based 
application.

**Viper** -
is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within 
an application, and can handle all types of configuration needs and formats. It supports:

**Gorm** -
is a developer friendly ORM library for golang.

**Gin** -
Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times 
faster thanks to http router. If you need performance and good productivity, you will love Gin.