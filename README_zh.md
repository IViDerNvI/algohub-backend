# algohub-backend

Algohub 程序

## 使用方法

### 使用 Docker

要使用 Docker 运行服务，请确保您的系统已安装 Docker。然后，进入项目目录并执行以下命令以构建 Docker 镜像：

```shell
docker pull ividernvi/algohub:latest
```

镜像构建完成后，可以使用以下命令运行服务：

```shell
docker run -d -p 8080:8080 -p 8443:8443 -e ALGOHUB_MYSQL_HOST=<mysql_hostname> -e ALGOHUB_MINIO_ENDPOINT=<minio_endpoint> --name algohub algohub
```

此命令将启动服务，并将容器的 8080 端口映射到主机的 8080 端口。您可以通过以下命令验证容器是否正在运行：

要停止并移除容器，请使用以下命令：

```shell
docker stop algohub-container && docker rm algohub-container
```

### 使用 Docker Compose

要使用 Docker Compose 运行服务，请确保您的系统已安装 Docker 和 Docker Compose。然后，进入项目目录并执行以下命令：

```shell
docker-compose up -d
```

此命令将以分离模式启动服务。您可以通过以下命令验证容器是否正在运行：

```shell
docker ps
```

要停止服务，请使用以下命令：

```shell
docker-compose down
```

## 贡献

我们欢迎对 algohub 项目的贡献！要参与贡献，请按照以下步骤操作：

Fork 仓库：点击此仓库 GitHub 页面右上角的 "Fork" 按钮。

克隆您的 Fork：将您 Fork 的仓库克隆到本地计算机：

创建分支：为您的功能或 Bug 修复创建一个新分支：

进行更改：实现您的更改，并使用清晰简洁的提交信息提交更改：

推送更改：将您的分支推送到您 Fork 的仓库：

提交 Pull Request：从您的分支向此仓库的 main 分支提交一个 Pull Request。请提供详细的更改描述以及解决的问题。

### 贡献指南
确保您的代码符合项目的编码标准。
为任何新功能编写清晰简洁的文档。
如果适用，请为您的更改添加测试。
感谢您的贡献！