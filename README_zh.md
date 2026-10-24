# OpenSirius Server

> 面向 **World Dai Star: Yume no Stellarium（ワールドダイスター 夢のステラリウム / ユメステ）** 的非官方服务端兼容实现。

OpenSirius Server 是一个使用 Go 编写的游戏服务端研究项目，主要用于协议分析、数据结构研究、服务端实现与相关后端技术学习。

项目通过实现客户端所使用的部分 HTTP API、MessagePack 数据结构、用户状态、Master Data、Live 等功能，尝试复现客户端运行所需的基础服务端行为。

> [!IMPORTANT]
> 本项目不是官方服务端，也不以完整复刻官方线上服务为目标。当前仍处于开发阶段，部分接口仅实现了满足客户端流程所需的最小逻辑，行为可能与官方服务端存在差异。

## 项目特点

* 使用 **Go + Gin** 实现 HTTP API
* 使用 **GORM** 管理持久化数据
* 支持 **SQLite / MySQL**
* 实现游戏使用的 **MessagePack** 请求与响应结构
* 提供 **LZ4 Block Array** 编解码支持
* 实现独立的游戏账号、Session 与用户状态存储
* 支持从捕获数据或模板生成初始 User Data
* 支持 Master Data 文件分发与 decoded JSON 导入
* 对部分 API 提供 Integration / Repository / Service 测试
* 内置请求限速、重复请求保护和请求体大小限制

## 当前实现状态

状态说明：

* ✅：已有实际业务逻辑，可用于当前兼容流程
* 🟡：已实现基础兼容逻辑，但仍包含 seed、固定结果或未完整校验的部分
* 🚧：尚未实现或仍在开发

| 功能                            | 主要接口                                   |  状态 | 备注                                       |
| ----------------------------- | -------------------------------------- | :-: | ---------------------------------------- |
| Environment                   | `/api/Environment`                     |  ✅  | 返回客户端版本、资源地址、Master Data 地址与 Realtime 地址 |
| Environment Ping              | `/api/Environment/Ping`                |  ✅  | 基础连通性检查                                  |
| 新账号注册                         | `/api/Account/Register`                |  ✅  | 创建游戏账号与 Login Token                      |
| 账号认证                          | `/api/Account/Authenticate`            |  ✅  | Login Token 换取游戏 Session                 |
| Confirmation Code             | `/api/Account/GetConfirmationCode`     |  ✅  | 生成并持久化确认码                                |
| 注销账号                          | `/api/Account/Delete`                  |  ✅  | 已接入账号服务逻辑                                |
| 游戏登录                          | `/api/Login`                           |  ✅  | 登录并返回客户端所需初始化结果                          |
| User Data                     | `/api/data/user`                       |  ✅  | 支持持久化用户状态、模板与生成模式                        |
| Master Data Manifest          | `/api/data/master`                     |  ✅  | 返回 Master Data URI、版本与发布时间               |
| Master Data 文件                | `/master-data/production/*`            |  ✅  | 本地 Master Data 文件分发                      |
| Asset 文件                      | `/production/*`                        |  ✅  | 本地静态资源文件分发                               |
| Episode                       | `/api/Episodes/...`                    |  ✅  | Episode Detail 与 Detail Asset            |
| Home                          | `/api/Home/CheckReceiveLoginBonus` 等   |  ✅  | Login Bonus、External Payment 基础逻辑        |
| Tutorial / Player             | `/api/Player/...`                      |  ✅  | Tutorial、Game Hint、Splash 状态更新           |
| Inbox                         | `/api/Inbox/...`、`/api/Inboxes/...`    |  ✅  | 查询、领取单个/全部 Package                       |
| Friends                       | `/api/Friends/...`                     |  ✅  | 好友列表、请求、接受、拒绝、删除                         |
| Circle                        | `/api/Circles/...`                     |  ✅  | Circle 创建、邀请、加入、离开及相关状态                  |
| Shop                          | `/api/Shops/ViewPage`                  |  ✅  | 商店页面访问状态与基础响应                            |
| Notification                  | `/api/Home/GetNotifications...`        |  ✅  | 通知列表、内容与已读时间                             |
| Friend Invitation             | `/api/FriendInvitation/Update`         |  🟡 | 当前主要为兼容性响应                               |
| Solo Live                     | `/api/Lives/Start`、`FinishAndValidate` |  🟡 | Start 流程已实现；Finish 仍有 seed/固定计算结果        |
| Another Notation Live         | `/api/Lives/FinishAnotherNotationLive` |  🟡 | 已接入基础结束流程，结算仍需完善                         |
| Triple Cast Live              | `/api/Lives/StartTripleCastLive`       |  🟡 | 已实现基础 Live Unit 生成                       |
| League                        | `/api/Leagues/TopMenuInformation`      |  🟡 | 当前 Top Menu 数据仍包含固定 seed                 |
| Triple Cast Top Menu          | `/api/TripleCast/TopMenuInformation`   |  🟡 | 当前主要用于客户端流程兼容                            |
| Realtime / Multi Live         | -                                      |  🚧 | DTO 已有部分结构，独立 Realtime 服务尚未完成            |
| Gacha / Mission / Event 等完整业务 | -                                      |  🚧 | 部分 DTO 已定义，API 与业务逻辑仍待实现                 |

功能状态以当前代码为准。DTO 中出现某一数据类型并不代表对应业务 API 已经完整实现。

## API 与数据格式

游戏 API 主要使用：

* HTTP
* `application/vnd.msgpack`
* MessagePack Array 形式的数据对象
* Result / Update / Delete / Effect 响应 Envelope
* LZ4 Block Array 压缩格式

协议相关实现主要位于：

```text
api/        MessagePack 数据结构、枚举和序列化辅助
apiwire/    API Envelope、请求解析、LZ4 编解码
dto/        API / Realtime DTO 与 Schema
```

如果你主要关注协议研究，建议优先阅读 `apiwire/`、`dto/` 与 `references/`。

## 项目结构

```text
.
├─ api/               # API MessagePack 编解码及公共数据结构
├─ apiwire/           # Wire format、Envelope、LZ4
├─ cmd/               # 辅助命令行工具
├─ configs/           # 配置文件示例
├─ core/              # 应用启动、配置、数据库、日志、鉴权
├─ dto/               # API / Realtime DTO
├─ handlers/          # HTTP Handler 与路由实现
├─ masterdata/        # Master Data 查询、同步与 JSON 导入
├─ middleware/        # Session、权限、安全限制、日志中间件
├─ references/        # 协议研究用参考数据（提交前请确保已脱敏）
├─ repositories/      # 数据持久化层
├─ router/            # Gin Router 注册
├─ services/          # 业务逻辑层
├─ userdata/          # User Data Builder、Template、Union
├─ utils/             # 通用工具
└─ main.go            # 程序入口
```

## 运行环境

建议环境：

* Go **1.26+**
* SQLite（开发环境默认）或 MySQL
* 使用 SQLite 时需要可用的 **CGO / C 编译器**，因为项目使用 `github.com/mattn/go-sqlite3`

## 快速开始

### 1. 获取依赖

```bash
go mod download
```

### 2. 准备配置

程序会在配置不存在时自动创建默认文件：

```text
data/config.yml
data/database.yml
data/security.yml
```

也可以先使用仓库内的示例：

```bash
mkdir -p data
cp configs/config.yml.example data/config.yml
cp configs/database.yml.example data/database.yml
cp configs/security.yml.example data/security.yml
```

Windows PowerShell：

```powershell
New-Item -ItemType Directory -Force data
Copy-Item configs/config.yml.example data/config.yml
Copy-Item configs/database.yml.example data/database.yml
Copy-Item configs/security.yml.example data/security.yml
```

然后根据自己的研究环境修改 `data/config.yml` 中的版本号、资源地址和 Master Data 路径。

### 3. 启动

```bash
go run .
```

默认监听：

```text
:8080
```

也可以编译后运行：

```bash
go build -o opensirius-server .
./opensirius-server
```

## 数据库

项目支持 SQLite 和 MySQL。

开发模式下默认使用 SQLite：

```yaml
sqlite:
  path: data/app.db
```

MySQL 配置位于 `data/database.yml`：

```yaml
mysql:
  host: 127.0.0.1
  port: 3306
  user: root
  password: change-me
  database: ymst_server
```

> [!NOTE]
> 当前数据库选择还会受到 Gin 运行模式影响：开发/测试模式使用 SQLite，release 模式使用 MySQL。将 `env` 设置为 `production` 前，请先准备好可用的 MySQL 配置。

启动时会自动执行 GORM Migration，并初始化项目需要的基础表。

## User Data

项目支持多种 User Data 初始化方式，并提供模板化生成能力。

相关配置：

```yaml
game:
  user_data_seed_json: references/captured/example.api.json
  user_data_seed_mode: captured_json
  generated_user_profile: template_minimal

  generated_user:
    player_rank: 1
    current_stamina: 100
    coin: 100000
    free_jewel: 3000
    paid_jewel: 0
    tutorial_status: 99
```

User Data 的生成、Union 与模板实现位于 `userdata/`。

## Master Data

项目支持三种与 Master Data 相关的用途：

1. `/api/data/master` 返回客户端使用的 Master Data Manifest
2. `/master-data/production/*` 分发本地 Master Data 文件
3. 启动时从 decoded JSON 目录导入 Master Data 到数据库，供业务逻辑查询

主要配置：

```yaml
game:
  master_data_root: data/master-data/production
  master_decoded_json_root: ../ymst_masterdata/data/master/decoded
  master_data_uri: 2026-05-09/mastermemory_xxx.db
  master_data_version: xxx
```

仓库不应直接提交来源不明或无权分发的官方资源。请自行准备合法取得的测试数据。

## Debug API

当前提供以下 User State 调试接口：

| API                                       | 作用                    |
| ----------------------------------------- | --------------------- |
| `GET /debug/user-state/stats`             | 查看当前 User State 统计信息  |
| `GET /debug/user-state/objects`           | 查看存储的数据对象             |
| `POST /debug/user-state/rebuild`          | 重建 User State         |
| `GET /debug/user-state/generated`         | 查看生成后的 User Data      |
| `GET /debug/user-state/generated-msgpack` | 获取生成后的 MessagePack 数据 |
| `GET /api/v1/health`                      | 服务与数据库健康检查            |

> [!WARNING]
> `/debug/user-state/*` 当前不属于游戏 Session 鉴权路径。它们只应在受控的本地/开发环境使用，不要直接暴露到公网。

## 安全配置

`data/security.yml` 提供基础请求保护：

* 最小请求间隔
* 单 IP 时间窗口限流
* 重复请求抑制
* 最大请求体限制
* Login 请求频率限制

示例：

```yaml
enabled: true
min_request_interval: 120ms
window_size: 1m
max_requests_per_ip: 120
duplicate_ttl: 2s
max_body_bytes: 65536
```

进行批量协议回放或自动化测试时，如果请求被限制，可以在受控环境中适当调整这些参数。

## 开发说明

项目目前仍以兼容客户端流程和协议研究为优先目标，因此代码中可能同时存在：

* 已持久化的真实业务逻辑
* 根据协议样本构造的兼容结果
* seed / fallback 数据
* 尚未连接完整业务系统的 DTO

提交新功能时，建议尽量遵循当前分层：

```text
handler -> service -> repository -> database
```

协议编解码逻辑尽量放入 `api/` 或 `apiwire/`，避免将复杂 MessagePack 处理直接堆积在 Handler 中。

如果新增接口，建议同时补充对应的单元测试或 Integration Test。

## 测试

```bash
go test ./...
```

现有测试覆盖部分：

* API Wire / Envelope
* LZ4
* Account
* Session
* User Repository
* User Data Builder
* Episode
* Service 层业务逻辑

## 贡献

Issue、协议结构补充、测试样本、兼容性修复和代码重构均可作为贡献方向。

提交协议样本或抓包数据前，请务必：

1. 删除账号凭据、Session Token、JWT、Cookie 等认证信息；
2. 删除能够识别真实用户身份的数据；
3. 不提交官方私钥、证书私钥、未公开密钥或其他敏感材料；
4. 不提交无权重新分发的完整客户端或官方资源文件。

## 免责声明

本项目仅用于软件开发、协议研究、兼容性测试和技术学习。

本项目与游戏官方、开发商、发行商及其他相关权利方不存在隶属、授权或合作关系。项目名称及相关游戏名称、商标、素材等权利归其各自权利人所有。

本仓库不应包含官方客户端、商业资源、私钥、Token 或其他无权公开传播的内容。使用者应自行确保其获取和使用测试数据、客户端及资源的方式符合所在地区法律法规和相关服务条款。

因使用本项目产生的账号、数据、服务或其他风险由使用者自行承担。

## License

本项目采用 **BSD 2-Clause License** 开源。

详见 [LICENSE](./LICENSE)。
