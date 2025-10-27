mkdir -p data
podman run \
    --cgroup-manager=cgroupfs \
    -d \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=postgres \
    --privileged \
    --mount type=bind,source=./data,target=/var/lib/postgresql/data \
    -p 5432:5432 \
    --name postgresGator \
    postgres:12
