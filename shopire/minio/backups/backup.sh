#!/bin/bash
TIMESTAMP=$(date +%F-%H%M)
mc cp /data myminio/shopire-backups/$TIMESTAMP
