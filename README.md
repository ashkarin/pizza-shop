# Pizza Sho
To run integration tests, please set `TEST` in `docker-compose.yaml` to `'true'`. To run the server, please set this property to `'false'`.

A cook can be implemented as a background process, which accesses orders using an atomic operation to ensure ownership. When the pizza is cooked, the cook will send an update to the database via API.

Alternatively, this system can be implemented with a message queue, which will serve as a waiter.