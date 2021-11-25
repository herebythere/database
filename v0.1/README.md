# Database

Container to deploy a postgres database

Utility functions to interface with a postgres database

## Types

```
DatabaseDetails {
	Name     string
	Host     string
	Password string
	Port     int64
	Username string
}

type SQLStatement struct {
	SQL    string
	Values []unknown
}
```

## Interfaces

```
DatabaseInterface {
    pool    <postgres pool>
}

DatabaseInterface::Query(statement SQLStatement)->[][]unknown
```

## Containers

### Requirements

```
dnf install python3 golang podman podman-compose
```

### Configuration

```
./config/database.json

{
    "container_port": 5432,
    "host_port": 3015,
    "database_name": "db",
    "username": "user",
    "password": "password"
}
```

### Scripts

```
python3 build.py

--config        config filepath
--templates     templates directory
--dest          destination directory
```

```
python3 run.py

--file          podman-compose filepath
```

```
python3 down.py

--file          podman-compose filepath
```

## License

BSD-3-Clause License
