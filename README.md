# OpenSirius Server

[Chinese Version](./README_zh.md)

> An unofficial server implementation for **World Dai Star: Yume no Stellarium (ワールドダイスター 夢のステラリウム / Yumeste)**.

OpenSirius Server is a Go-based game server research project intended for protocol analysis, data structure research, backend implementation, and related technical study.

The project implements part of the HTTP API, MessagePack data structures, user state, Master Data, Live-related features, and other server behaviors required by the client.

> [!IMPORTANT]
> This project is not an official server and does not aim to fully reproduce the official online service. It is still under active development. Some endpoints only implement the minimum logic required for client compatibility, and behavior may differ from the official service.

## Features

- HTTP API implemented with **Go + Gin**
- Persistent storage through **GORM**
- **SQLite / MySQL** support
- MessagePack request and response structures used by the game
- **LZ4 Block Array** encoding and decoding
- Account, session, and persistent user-state handling
- User Data generation from captured data or templates
- Master Data file distribution and decoded JSON import
- Integration, repository, and service tests for selected features
- Built-in rate limiting, duplicate request protection, and request body size limits

## Implementation Status

Status legend:

- ✅ Implemented with actual business logic
- 🟡 Partially implemented; may still contain seed data, fixed responses, or incomplete validation
- 🚧 Not implemented or still under development

| Feature | Main Endpoint | Status | Notes |
| --- | --- | :---: | --- |
| Environment | `/api/Environment` | ✅ | Returns client version, asset URL, Master Data URL, and Realtime URL |
| Environment Ping | `/api/Environment/Ping` | ✅ | Basic connectivity check |
| Account Registration | `/api/Account/Register` | ✅ | Creates a game account and login token |
| Account Authentication | `/api/Account/Authenticate` | ✅ | Exchanges a login token for a game session |
| Confirmation Code | `/api/Account/GetConfirmationCode` | ✅ | Generates and persists confirmation codes |
| Account Deletion | `/api/Account/Delete` | ✅ | Connected to account service logic |
| Game Login | `/api/Login` | ✅ | Performs login and returns initialization data required by the client |
| User Data | `/api/data/user` | ✅ | Supports persistent user state, templates, and generated data |
| Master Data Manifest | `/api/data/master` | ✅ | Returns Master Data URI, version, and publish time |
| Master Data Files | `/master-data/production/*` | ✅ | Serves local Master Data files |
| Asset Files | `/production/*` | ✅ | Serves local static assets |
| Episode | `/api/Episodes/...` | ✅ | Episode Detail and Detail Asset |
| Home | `/api/Home/CheckReceiveLoginBonus` etc. | ✅ | Login Bonus and External Payment base logic |
| Tutorial / Player | `/api/Player/...` | ✅ | Tutorial, Game Hint, and Splash state updates |
| Inbox | `/api/Inbox/...`, `/api/Inboxes/...` | ✅ | Query and claim single/all packages |
| Friends | `/api/Friends/...` | ✅ | Friend list, requests, accept, reject, and delete |
| Circle | `/api/Circles/...` | ✅ | Circle creation, invitations, joining, leaving, and related state |
| Shop | `/api/Shops/ViewPage` | ✅ | Shop page state and base response |
| Notification | `/api/Home/GetNotifications...` | ✅ | Notification list, content, and read time |
| Friend Invitation | `/api/FriendInvitation/Update` | 🟡 | Primarily a compatibility response at the moment |
| Solo Live | `/api/Lives/Start`, `/api/Lives/FinishAndValidate` | 🟡 | Start flow implemented; Finish still contains seeded/fixed calculations |
| Another Notation Live | `/api/Lives/FinishAnotherNotationLive` | 🟡 | Basic finish flow implemented; result logic still needs work |
| Triple Cast Live | `/api/Lives/StartTripleCastLive` | 🟡 | Basic Live Unit generation implemented |
| League | `/api/Leagues/TopMenuInformation` | 🟡 | Top Menu data still contains fixed seed data |
| Triple Cast Top Menu | `/api/TripleCast/TopMenuInformation` | 🟡 | Mainly implemented for client-flow compatibility |
| Realtime / Multi Live | - | 🚧 | Some DTO structures exist, but the standalone Realtime service is not complete |
| Gacha / Mission / Event and other full business flows | - | 🚧 | Some DTOs exist, but APIs and business logic are incomplete |

Implementation status is based on the current codebase. The existence of a DTO does not necessarily mean that the related business API is fully implemented.

## API and Data Format

The game API mainly uses:

- HTTP
- `application/vnd.msgpack`
- MessagePack array-based objects
- Result / Update / Delete / Effect response envelopes
- LZ4 Block Array compression

Protocol-related code is mainly located in:

```text
api/        MessagePack structures, enums, and serialization helpers
apiwire/    API envelopes, request parsing, and LZ4 handling
dto/        API / Realtime DTOs and schemas
```

If you are primarily interested in protocol research, start with `apiwire/`, `dto/`, and `references/`.

## Project Structure

```text
.
├─ api/               # API MessagePack codecs and common data structures
├─ apiwire/           # Wire format, envelopes, and LZ4
├─ cmd/               # Helper command-line tools
├─ configs/           # Example configuration files
├─ core/              # Application startup, config, database, logging, authentication
├─ dto/               # API / Realtime DTOs
├─ handlers/          # HTTP handlers and route implementations
├─ masterdata/        # Master Data query, sync, and JSON import
├─ middleware/        # Session, permissions, security limits, request logging
├─ references/        # Protocol research reference data; sanitize before publishing
├─ repositories/      # Persistence layer
├─ router/            # Gin router registration
├─ services/          # Business logic layer
├─ userdata/          # User Data builders, templates, and union logic
├─ utils/             # General utilities
└─ main.go            # Application entry point
```

## Requirements

Recommended environment:

- Go **1.26+**
- SQLite for development, or MySQL
- A working **CGO / C compiler** when using SQLite, because the project uses `github.com/mattn/go-sqlite3`

## Quick Start

### 1. Install Dependencies

```bash
go mod download
```

### 2. Prepare Configuration

If the configuration files do not exist, the application can generate default files under:

```text
data/config.yml
data/database.yml
data/security.yml
```

You may also start from the examples included in the repository:

```bash
mkdir -p data
cp configs/config.yml.example data/config.yml
cp configs/database.yml.example data/database.yml
cp configs/security.yml.example data/security.yml
```

Windows PowerShell:

```powershell
New-Item -ItemType Directory -Force data
Copy-Item configs/config.yml.example data/config.yml
Copy-Item configs/database.yml.example data/database.yml
Copy-Item configs/security.yml.example data/security.yml
```

Then edit `data/config.yml` according to your research environment, especially the game version, asset URL, and Master Data paths.

### 3. Start the Server

```bash
go run .
```

Default listen address:

```text
:8080
```

Or build it first:

```bash
go build -o opensirius-server .
./opensirius-server
```

## Database

The project supports SQLite and MySQL.

Development mode uses SQLite by default:

```yaml
sqlite:
  path: data/app.db
```

MySQL configuration is stored in `data/database.yml`:

```yaml
mysql:
  host: 127.0.0.1
  port: 3306
  user: root
  password: change-me
  database: ymst_server
```

> [!NOTE]
> Database selection is currently also affected by the Gin runtime mode. Development/test mode uses SQLite, while release mode uses MySQL. Before setting the environment to `production`, make sure a working MySQL configuration is available.

GORM migrations run automatically during startup and initialize the tables required by the project.

## User Data

The project supports multiple User Data initialization modes and template-based generation.

Related configuration:

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

User Data generation, union, and templates are implemented under `userdata/`.

## Master Data

Master Data is used in three main ways:

1. `/api/data/master` returns the Master Data manifest required by the client.
2. `/master-data/production/*` serves local Master Data files.
3. Decoded Master Data JSON can be imported into the database on startup for business-logic queries.

Main configuration:

```yaml
game:
  master_data_root: data/master-data/production
  master_decoded_json_root: ../ymst_masterdata/data/master/decoded
  master_data_uri: 2026-05-09/mastermemory_xxx.db
  master_data_version: xxx
```

Do not commit official resources or other copyrighted assets that you are not authorized to redistribute.

## Debug API

The following User State debug endpoints are currently available:

| API | Purpose |
| --- | --- |
| `GET /debug/user-state/stats` | View current User State statistics |
| `GET /debug/user-state/objects` | Inspect stored data objects |
| `POST /debug/user-state/rebuild` | Rebuild User State |
| `GET /debug/user-state/generated` | View generated User Data |
| `GET /debug/user-state/generated-msgpack` | Retrieve generated MessagePack data |
| `GET /api/v1/health` | Service and database health check |

> [!WARNING]
> `/debug/user-state/*` is currently outside the normal game-session authentication path. These endpoints should only be used in controlled local or development environments and should not be exposed directly to the public Internet.

## Security Configuration

`data/security.yml` provides basic request protection, including:

- Minimum request interval
- Per-IP request window limits
- Duplicate request suppression
- Maximum request body size
- Login request rate limiting

Example:

```yaml
enabled: true
min_request_interval: 120ms
window_size: 1m
max_requests_per_ip: 120
duplicate_ttl: 2s
max_body_bytes: 65536
```

When performing controlled protocol replay or automated testing, these values may be adjusted as needed.

## Development Notes

The current priority is client compatibility and protocol research. As a result, the codebase may contain a mixture of:

- Fully persisted business logic
- Compatibility responses constructed from protocol samples
- Seed or fallback data
- DTOs that are not yet connected to complete business systems

When adding new features, prefer the current layering model:

```text
handler -> service -> repository -> database
```

Protocol encoding and decoding should generally stay in `api/` or `apiwire/` rather than being implemented directly inside HTTP handlers.

When adding new endpoints, adding corresponding unit or integration tests is strongly recommended.

## Testing

```bash
go test ./...
```

Existing tests cover parts of:

- API Wire / Envelope
- LZ4
- Account
- Session
- User Repository
- User Data Builder
- Episode
- Service-layer business logic

## Contributing

Issues, protocol structure improvements, compatibility fixes, tests, and refactoring are welcome.
Before submitting packet captures or protocol samples, make sure to:

1. Remove account credentials, session tokens, JWTs, cookies, and other authentication data.
2. Remove information that can identify real users.
3. Do not submit official private keys, certificate private keys, unreleased secrets, or other sensitive material.
4. Do not submit complete game clients or official resources that you are not authorized to redistribute.

## Disclaimer

This project is intended for software development, protocol research, compatibility testing, and technical study.
This project is not affiliated with, endorsed by, authorized by, or otherwise associated with the game's developers, publishers, or other rights holders. Game names, trademarks, artwork, and related materials remain the property of their respective owners.
The repository should not contain official game clients, commercial assets, private keys, tokens, or other material that cannot legally be redistributed. Users are responsible for ensuring that the way they obtain and use clients, test data, and resources complies with applicable laws and service terms.
Users assume responsibility for any account, data, service, or other risks resulting from use of this project.

## License

This project is licensed under the **BSD 2-Clause License**.

See [LICENSE](./LICENSE) for details.
