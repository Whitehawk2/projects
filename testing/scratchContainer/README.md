# Simple Go Router test / demo Application

This is a simple Go application designed to listen on a port, and say hello. Meant to act as an auxilary to other things,
when needing a simple listener, that's already configured to run in a "FROM scratch" container.
Written originally to act as a demo for a debugging workshop.

## Features

- Lightweight, super simple, and fast
- Easy to configure and extend
- Containerized (via attached `Dockerfile` & `install.sh`) in a scratch docker container

## Prerequisites

- [Docker](https://www.docker.com/get-started) installed on your system

## Running the Application

To run the Simple Go Router Application, follow these steps:

1. **Build the Docker Image**

   First, you need to build the Docker image for the application. Replace `<Dockerfile-path>` with the path to your Dockerfile:

   ```bash
   docker build -t simple-go-router <Dockerfile-path>
   ```

2. **Run the Docker Container**

   Next, run the Docker container. Replace `<port>` with the desired port number:

   ```bash
   docker run -p <port>:<port> simple-go-router
   ```

3. **Access the Application**

   Once the container is running, access the application via your web browser or a tool like `curl` at `http://localhost:<port>`. Replace `<port>` with the port number you used in the previous step.

## Configuration

Placeholders for configuration details will be described here. You can specify routes, handlers, and other settings in the configuration file.

## Extending the Application

This application can be extended by adding new routes and handlers. Placeholders for code snippets will be provided where you can implement custom logic.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contact

For questions or support, please contact [Your Contact Information].
