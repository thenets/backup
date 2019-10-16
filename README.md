# thenets backup
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fthenets%2Fbackup.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fthenets%2Fbackup?ref=badge_shield)


## Requirements

- mysql-client
- postgres-client
- openssh-client
- rsync
- awk

## Config files

### Directory (SSH)

The current user must have access to the target server from SSH. Add the `key.pub` to the `authorized_keys` to allow the access.

`my-app-content.dir.ini`

```properties
# Server origin dir
SERVER_DIR=backup@1.2.3.4:/opt/app/uploads/

# Target dir to put all backup files
TARGET_DIR=/mnt/volume_backup

# Delete files older than X days
DELETE_OLDER_THAN_X_DAYS=5

# Compress subdirs
# 0 False; 1 True
COMPRESS_SUBDIR=1

# Set compression level
# 1 - low (fast)
# 9 - high (slow)
GZIP=-1
```

### MySQL

`my-database.mysql.ini`

```properties
# MySQL Params
MYSQL_HOST=mysql.example.com
MYSQL_PORT=3306
MYSQL_USER=backup
MYSQL_PASS=+voD/QvMzv821s9uJsBs/PCtdflura4Q2C4gayfAHiA=

# Target dir to put all backup files
TARGET_DIR=/mnt/volume_backup

# Delete files older than X days
DELETE_OLDER_THAN_X_DAYS=20

# Compress subdirs
# 0 False; 1 True
COMPRESS_SUBDIR=1

# Set compression level
# 1 - low (fast)
# 9 - high (slow)
GZIP=-1
```

### PostgreSQL

`my-database.postgres.ini`

```properties
# PostgreSQL Params
POSTGRES_HOST=postgres.example.com
POSTGRES_PORT=25060
POSTGRES_USER=backup
POSTGRES_PASS=LFDrtmkCwUYuLTCMwVjFmGIjKqBwOivxNyJkmRarihg=

# Target dir to put all backup files
TARGET_DIR=/mnt/volume_backup

# Delete files older than X days
DELETE_OLDER_THAN_X_DAYS=20

# Compress subdirs
# 0 False; 1 True
COMPRESS_SUBDIR=1

# Set compression level
# 1 - low (fast)
# 9 - high (slow)
GZIP=-1
```





## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fthenets%2Fbackup.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fthenets%2Fbackup?ref=badge_large)