Here are detailed instructions to set up and run the application:
(created with claude.ai)

Setting up
Clone the repository:
    git clone https://github.com/example/go-docker-api

2. Build the Docker images:

    cd go-docker-api
    docker-compose build

3.  Configure environment variables:

- In api/.env, set DB_HOST, DB_USER, DB_PASS, etc
- In db/.env, set the MYSQL_ROOT_PASSWORD

4. Generate SSL certificates for HTTPS (optional):
    openssl req -newkey rsa:4096 -nodes -sha256 -keyout certs/domain.key -x509 -days 365 -out certs/domain.crt

Running
    docker-compose up -d

This will start the containers in the background.

Accessing the API
The API will be available on http://localhost or https://localhost if you configured SSL.

Test endpoints using curl or a REST client like Postman.

Monitoring
- Check container logs for errors:
    docker-compose logs -f

- Monitor application metrics with Prometheus by exposing metrics on /metrics.
- Trace requests with Jaeger by integrating OpenTracing.
- Set up healthchecks in Docker Compose.
- Use watchtower to automatically update containers.

Performance
- Profile and benchmark API endpoints to identify bottlenecks.
- Enable pprof in the Go application to profile CPU and memory usage.
- Optimize slow database queries.
- Set memory limits on containers to prevent OOM kills.
- Configure nginx worker processes and connections for load balancing.
Let me know if you have any other questions!
