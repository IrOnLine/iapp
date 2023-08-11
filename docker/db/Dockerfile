FROM mysql:5.7

# Copy initialization sql script
COPY initdb.sql /docker-entrypoint-initdb.d/

# Persist data to volume
VOLUME /var/lib/mysql