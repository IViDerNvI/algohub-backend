# algohub

A Golang-based user information management service.

## Usage

### Using docker

To run the service using Docker, ensure you have Docker installed on your system. Then, build the Docker image by navigating to the project directory and executing the following command:

```shell
docker pull ividernvi/algohub:latest
```

Once the image is built, you can run the service using:

```shell
docker run -d -p 8080:8080 -p 8443:8443 -e ALGOHUB_MYSQL_HOST=<mysql_hostname> -e ALGOHUB_MINIO_ENDPOINT=<minio_endpoint> --name algohub algohub-app
```

This will start the service and map port 8080 of the container to port 8080 on your host machine. You can verify that the container is running by using:

```shell
docker ps
```

To stop and remove the container, use:

```shell
docker stop algohub-container && docker rm algohub-container
```

### Using docker compose

To run the service using Docker Compose, ensure you have Docker and Docker Compose installed on your system. Then, navigate to the project directory and execute the following command:

```shell
docker-compose up -d
```

This will start the service in detached mode. You can verify that the containers are running by using:

```shell
docker ps
```

To stop the service, use:

```shell
docker-compose down
```

## Contribute
We welcome contributions to the algohub project! To contribute, follow these steps:

1. **Fork the repository**: Click the "Fork" button on the top right of this repository's GitHub page.

2. **Clone your fork**: Clone your forked repository to your local machine using:
    ```shell
    git clone https://github.com/<your-username>/algohub.git
    ```

3. **Create a branch**: Create a new branch for your feature or bug fix:
    ```shell
    git checkout -b feature-or-bugfix-name
    ```

4. **Make your changes**: Implement your changes and commit them with clear and concise commit messages:
    ```shell
    git commit -m "Description of your changes"
    ```

5. **Push your changes**: Push your branch to your forked repository:
    ```shell
    git push origin feature-or-bugfix-name
    ```

6. **Submit a pull request**: Open a pull request from your branch to the `main` branch of this repository. Provide a detailed description of your changes and the problem they solve.

### Guidelines
- Ensure your code adheres to the project's coding standards.
- Write clear and concise documentation for any new features.
- Include tests for your changes, if applicable.

Thank you for contributing!