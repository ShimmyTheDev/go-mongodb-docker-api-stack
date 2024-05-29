#!/bin/bash
TIMESTAMP=$(date +%F-%H%M)
BACKUP_DIR=/backups/$TIMESTAMP
mkdir -p $BACKUP_DIR
mongodump --host mongo1 --out $BACKUP_DIR
