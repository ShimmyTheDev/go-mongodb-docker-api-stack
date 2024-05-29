#!/bin/bash
mc alias set myminio http://minio:9000 $MINIO_ACCESS_KEY $MINIO_SECRET_KEY
mc mb myminio/shopire-backups
