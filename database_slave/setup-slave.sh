#!/bin/bash
set -e

PGPASSWORD=replica_password pg_basebackup -h postgres-master -D /var/lib/postgresql/data -U replicator -v -P --wal-method=stream
