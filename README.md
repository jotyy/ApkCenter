# 和行约车Apk管理中心

## 使用说明

### 目录结构

`/backend`：后端项目，使用Golang编写，打包详情见`Dockerfile`。

`/frontend`: 前端项目，使用Vue编写，打包详情见`Dockerfile`

`docker-compose.yml`：docker编排文件，包含前后端项目、postgres数据库及nginx作为文件代理服务器。

### 运行

执行以下命令可启动项目

```shell
docker-compose up
```

### 注意

如需迁移至其它服务器，前端项目内`/src/plugins/axios.js`的baseurl地址需修改。（此处可考虑使用env配置IP地址，没时间写）

