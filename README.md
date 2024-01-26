# PalWorld-Admin

`PalWorld-Admin` 是一个用于管理 `PalWorld` 游戏配置文件和进行存档备份的可视化工具。

## 功能特性

- **配置文件管理：** 可以轻松查看和编辑 `PalWorld` 游戏的配置文件。
- **存档备份（开发中）：** 此功能正在开发中，敬请期待。
- **前端界面：** 一个直观、用户友好的前端界面，提供更好的使用体验。
- **简单的管理员密码保护：** 使用管理员密码来保护工具的访问权限。

![界面图片](https://github.com/limitcool/palworld-admin/blob/main/images/screenshot.png?raw=true)

## 安装及运行

```bash
# linux amd64
wget https://github.com/limitcool/palworld-admin/releases/download/v0.1.1/palworld-admin-v0.1.1-linux-amd64.tar.gz
tar -xzvf palworld-admin-v0.1.1-linux-amd64.tar.gz
chmod u+x 
./palworld-admin
```

## 配置文件

在使用 `PalWorld-Admin` 之前，请根据您的需求修改配置文件。

### 配置项

- **adminpassword：** 管理员密码，用于保护工具的访问权限。请将 `initcool-https://blog.nmslwsnd.com` 替换为您自己的安全密码。
- **palworldconfigfilepath：** `PalWorld` 游戏配置文件的路径。默认为 `/root/palworld/data/Config/LinuxServer/PalWorldSettings.ini`，您可以根据实际情况修改。
- **port：** 工具的访问端口。默认为 `8080`，您可以根据需要更改。

### 修改配置文件方法

- **POSIX (Linux/BSD):** 配置文件路径为 `~/.palworld-Admin/config/config.yaml`。
- **Windows:** 配置文件路径为 `%LOCALAPPDATA%\Palworld-Admin/config/config.yaml`。

## 反馈与支持

我们欢迎您提供宝贵的反馈意见，帮助我们不断改进工具。您可以加入我们的 QQ 群：

![QQ群:699024161](https://github.com/limitcool/palworld-admin/blob/main/images/qqgroup.jpg?raw=true)

