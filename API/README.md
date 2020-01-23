# API configuration
To configure your software you need to edit `.env` file. `.env` is the configuration file
written in plain text format. 
###### Configuration:
- `USER_LOGIN=xxx ` replace word `xxx` with your own login (username)
- `USER_PASSWORD=pass ` replace word `pass` with your own password
- `HOST_URL=http://localhost:3000/` replace `http://localhost:3000/` with your server ip (or domain) + port
- `DB_URL=postgres://whbmdhoo:HbKTACofj3_0a9nCbyThdTSNpm34QWwv@rajje.db.elephantsql.com:5432/whbmdhoo` replace `postgres://whbmdhoo:HbKTACofj3_0a9nCbyThdTSNpm34QWwv@rajje.db.elephantsql.com:5432/whbmdhoo:`
with your own postgres DB url
- `MANAGER_PASSWORD=mpassword` replace `mpassword` with your own password for manager account

# How to start API
You have to download `go` compiler and after this write 3 commands:
- `export GO111MODULE=on`
- `go build -o ./bin/API`
- `sudo ./bin/API`