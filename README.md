# Krini Project Management

## Overview
Basic implementation for a project management service for the KRINI WebUI. This is not production ready code and mainly
created for the ELIXIR Biohackathon 2019.
It currently uses SQLite as a database backend. This will be changed in the future for production deployments. 

## Concept
This service keeps tracks of users of the platform and their associated projects. Running workflows are associated with individual projects to keep track of them. The tracking component is currently WIP. For details please see the Krini design document.

## Build and Run
### Requirements
- go version 1.13+ (others not tested)
  
### Build process
Clone the repository and go into the directory. Use `go build` to build the binary and execute it.

### Docker
Use the provided dockerfile to create an image.
`docker build --rm -f "Dockerfile" -t projecthandler:latest .`

### Run
#### Parameters:
- -p --port: Port the service should run on
- -d --dbPath: Path of the sqlite database file
  
#### Run Docker
To run the docker image after building it just run: `docker run -p 8080:8080 projecthandler`.
Please note that at the moment there are repository based images available.