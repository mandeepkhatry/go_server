go_server
------------------

A simple server made in Go as a part of Brankas Skill Test.

**How to Run?**
```
go run cmd/main.go
```
*API server will run at: `127.0.0.1:5000`*


**API Info**

| Endpoint | Body Fields | Method| Remarks |
|----------|------|--------|--------|
| `/` | | GET| Index page|
| `/upload`| `auth`, `upload` | POST | Upload file with auth token|