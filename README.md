# PalWorld-Admin

`PalWorld-Admin` 是一个用于管理 `PalWorld` 游戏配置文件和进行存档备份的跨平台可视化工具。

## 功能特性

- **配置文件管理：** 可以轻松查看和编辑 `PalWorld` 游戏的配置文件。(可在配置文件里面填写重启命令后自动重启Palworld服务端)
- **存档备份：** 定时备份存档,默认单位为秒级。
- **存档恢复（开发中）：** 通过前端恢复指定存档。
- **前端界面：** 一个直观、用户友好的前端界面，提供更好的使用体验。
- **简单的管理员密码保护：** 使用管理员密码来保护工具的访问权限。

![界面图片](https://github.com/limitcool/palworld-admin/blob/main/images/screenshot.png?raw=true)

## 安装及运行

```bash
# linux amd64
wget https://github.com/limitcool/palworld-admin/releases/download/v0.1.4/palworld-admin-v0.1.4-linux-amd64.tar.gz
tar -xzvf palworld-admin-v0.1.4-linux-amd64.tar.gz
chmod u+x palworld-admin
./palworld-admin
```

## 配置文件

在使用 `PalWorld-Admin` 之前，请根据您的需求修改配置文件。

```yaml
PalSavedPath: "C:\\Users\\Andorid\\AppData\\Local\\Pal\\Saved"  # palworld游戏目录
AdminPassword: "initcool-https://blog.nmslwsnd.com"    # 管理员面板密码
Port: 8080              # http服务监听端口
SaveConfig:
    BackupInterval: 60            # 每60秒备份一次存档
    MaxRetentionDays: 7           # 存档最大保留时间,默认为7天
    BackupDirectory: backups/         # 存档保存目录
RestartCommand: "docker restart palworld" # 自动重启的命令
#RestartCommand: "taskkill /f /im PalServer-Win64-Test-Cmd.exe && start D:\\Pal-Steam-Cmd\\steamapps\\common\\PalServer\\Pal\\Binaries\\Win64\\PalServer-Win64-Test-Cmd.exe"    # (Windows) 需要修改路径为实际的的PalServer-Win64-Test-Cmd.exe路径
```

### 修改配置文件方法

- **POSIX (Linux/BSD):** 配置文件路径为 `~/.palworld-admin/config.yaml`。
- **Windows:** 配置文件路径为 `%LOCALAPPDATA%/palworld-admin/config.yaml`。

## 路线图

- **存档恢复功能（预计下个版本）：** 计划在下一个版本中添加存档恢复功能。

## 反馈与支持

我们欢迎您提供宝贵的反馈意见，帮助我们不断改进工具。您可以加入我们的 `QQ 群`：`699024161`

![QQ群:699024161](https://github.com/limitcool/palworld-admin/blob/main/images/qqgroup.jpg?raw=true)
