# Developer Setup — go_shop (Windows cmd)

This file contains copy-paste commands and short explanations to set up a development environment for this project (Echo + gRPC + MongoDB + Protobuf + JWT + bcrypt).

> Shell: Windows cmd.exe (commands are written for cmd); if you prefer PowerShell, use that instead.

---

## 0) Quick verification (after installing tools)

Copy & run these to check basic tools:

```cmd
go version
protoc --version
where go
where protoc
%USERPROFILE%\go\bin\protoc-gen-go.exe --version 2>nul || echo "protoc-gen-go (check by running protoc with plugin)"
```

---

## 1) System prerequisites

1. Install Go (use the version in `go.mod` or Go 1.20+):
   - Download the Windows MSI: https://go.dev/dl/
   - After installation, ensure `%GOPATH%\bin` (or `%USERPROFILE%\go\bin`) is in `PATH`.

2. Install protoc (Protocol Buffers compiler):
   - Download the prebuilt `protoc` for Windows from https://github.com/protocolbuffers/protobuf/releases
   - Unzip and put `protoc.exe` somewhere in your PATH (e.g., `C:\bin\protoc\protoc.exe`) or add its folder to PATH.

3. Install Docker (optional but recommended) for local databases (MongoDB) or run MongoDB locally:
   - https://www.docker.com/get-started

---

## 2) Add Go bin to PATH (one-time)

If `%USERPROFILE%\go\bin` is not already in your PATH, run (cmd):

```cmd
setx PATH "%PATH%;%USERPROFILE%\go\bin"
```

Close and reopen your terminal to pick up the change.

---

## 3) Install global Go developer tools (copy/paste)

Run these from cmd (they install binaries into `%GOPATH%\bin`). Replace `@latest` with a specific version if you want pinned versions.

```cmd
REM Language server for editor (gopls)
go install golang.org/x/tools/gopls@latest

REM Linter aggregator
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

REM Debugger
go install github.com/go-delve/delve/cmd/dlv@latest

REM Protobuf generators (Go + gRPC)
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

REM Optional: buf (protos workflow)
go install github.com/bufbuild/buf/cmd/buf@latest

REM Optional: mock generators for testing
go install github.com/golang/mock/mockgen@latest
go install github.com/vektra/mockery/v2@latest
```

After these, ensure `%USERPROFILE%\go\bin` contains `protoc-gen-go.exe` and `protoc-gen-go-grpc.exe`.

---

## 4) Project module dependencies (inside project root)

Open a cmd in the project root (`c:\inet_training\go_shop`) and run:

```cmd
REM Download and tidy dependencies
go mod tidy
```

If you need to add a library, use `go get` or edit `go.mod` and run `go mod tidy`.

Recommended libraries used in this project (they should already be in go.mod):
- github.com/labstack/echo/v4 (HTTP server)
- go.mongodb.org/mongo-driver/v2 (MongoDB driver v2)
- google.golang.org/grpc and google.golang.org/protobuf (gRPC + protobuf)
- github.com/golang-jwt/jwt/v5 (JWT)
- golang.org/x/crypto/bcrypt (bcrypt)
- github.com/go-playground/validator/v10 (validation)

If any are missing you can add them, for example:

```cmd
go get github.com/labstack/echo/v4
go get go.mongodb.org/mongo-driver/v2
go get google.golang.org/grpc
go get google.golang.org/protobuf
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get github.com/go-playground/validator/v10
```

Then run `go mod tidy` again.

---

## 5) Generating protobuf files (example)

This project has `.proto` files under module folders (e.g., `modules/*/*Pb.proto`). Example command (run from project root). Adjust paths to your proto files:

```cmd
REM Example: generate Go and gRPC stubs for a proto file
protoc --proto_path=./modules/player/playerPb --go_out=./modules/player/playerPb --go_opt=paths=source_relative --go-grpc_out=./modules/player/playerPb --go-grpc_opt=paths=source_relative modules/player/playerPb/playerPb.proto
```

If you have many proto files, either repeat or use a script/buf to generate all at once.

---

## 6) Run the project (player service)

Start the player service with environment file (example):

```cmd
go run main.go ./env/dev/.env.player
```

Expected output:
- HTTP server on `:1325` (see `server.go`)
- gRPC server on configured port (e.g., `:1623`)

---

## 7) Useful troubleshooting commands

```cmd
REM Fix module cache corruption
go clean -modcache

REM Re-download tidy dependencies
go mod tidy

REM Show which process is using a port (Windows)
netstat -ano | findstr :1325

REM Kill process by PID (if needed)
taskkill /F /PID <PID>
```

If you see errors about `missing go.sum entry`, run the suggested `go get` command that the error shows, then `go mod tidy`.

---

## 8) Final tips

- Keep the entire codebase using the same mongo-driver major version (this project uses v2).
- Pin important tool versions in docs or a `tools.go` file for reproducible dev environments.
- Consider adding a small `dev.ps1` or `dev.bat` script to automate common steps (generate protos, `go mod tidy`, run server).

---

If you want, I can also:
- Add a `dev.bat` script to the repo that installs tools and runs `go mod tidy` (Windows cmd style).
- Generate example `protoc` commands for every `.proto` file in this repo.


---

Created for: go_shop
