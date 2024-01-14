# Hexagonal Architecture Code Sample project with Golang + Sequelize

A project to learn how to apply Hexagonal Architecture concepts in a real case project.

Based on: **[here](https://github.com/codeedu/fc2-arquitetura-hexagonal)**

## How to run?

The application is based on Go and Sequelize. However, to facilitate development, a docker-compose.yml file was created that raises a container with the Sqlite3 database.

To run the application, just run the command below:

```bash
docker compose up -d
```

After that, for you to enter the container and be able to execute Sequelize commands, run the command below:

```bash
docker exec -it app-product bash
```

If the message below appears:

```bash
> www-data@7cb0866f66c4:/go/src$
```

It means that you are inside the container and can execute Sequelize commands.

Now, to create a `go.mod` file, run the command below:

```bash
go mod init github.com/your-user-name/your-project-name
```

After that, to install the dependencies, run the command below:

```bash
go mod tidy
```

And finally, use the Visual Studio Code to update the dependencies in the `go.mod` file. How? Just open the `go.mod` file and click on the `Go: Install/Update Tools` option that will appear at the top of the file.

## Working with `gomock` and `mockgen`

To work with `mockgen` in this project, open now the terminal, use the docker command below to enter the container:

```bash
docker exec -it app-product bash
```

After that, run the command below to install the `mockgen` tool:

```bash
mockgen -destination=application/mocks/application.go -source=application/product.go application
```

After that, if you open the `application` folder, you will see that the `mocks` folder has been created with the `application.go` file inside it.

And to install the dependencies, run the command below:

```bash
go mod tidy
```



