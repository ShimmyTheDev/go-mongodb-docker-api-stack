# Go docker API stack example

An example of a a dockerized Go stack for self-hosted purpose.

#### Services

- Simple **Go** service that tests logs
- **Grafana**/**loki**/**promtail** Log observability tools
- **MongoDB** - NoSQL Database
- **MinIO** - S3 & Kubernetes Native Object Storage

---

### Instructions

##### Quickstart

```
git clone https://github.com/ShimmyTheDev/go-mongodb-docker-api-stack \
&& cd go-mongodb-docker-api-stack \
&& touch .env \
&& echo -e "MONGODB_DATABASE=database
MONGODB_USER=db_user
MONGODB_PASSWORD=da_password
MONGO_INITDB_ROOT_USERNAME=username
MONGO_INITDB_ROOT_PASSWORD=password
MONGODB_USER=db_user
MONGODB_PASSWORD=db_password
MONGO_INITDB_ROOT_USERNAME=db_init_user
MONGO_INITDB_ROOT_PASSSWORD=db_init_password
MINIO_ACCESS_KEY=minio_access_key
MINIO_SECRET_KEY=minio_secret_key" >> .env \
&& docker compose up -d
```

---

##### Ports

- **MongoDB** - 27017
- **MinIO** - 9000
- **Grafana** - 3000
- **loki** - 3100
- **API** - 8000
