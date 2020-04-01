# LSIF code intelligence

This project is an early adopter of Microsoft's [LSIF](https://code.visualstudio.com/blogs/2019/02/19/lsif) standard. LSIF (Language Server Index Format) is a format used to store the results of language server queries that are computed ahead-of-time. We uses this format to provide jump-to-definition, find-reference, and hover docstring functionality.

LSIF dumps are generated by running an LSIF indexer in a build or continuous integration environment. The dump is uploaded to a Sourcegraph instance via [Sourcegraph CLI](https://github.com/sourcegraph/src-cli). An LSIF server, proxied by the frontend for auth, answers relevant LSP queries to provide fast and precise code intelligence.

### Architecture

This project is split into two parts, both written in TypeScript. These parts are deployable independently, but for now they run in the same docker container. The HTTP [server](./src/server/server.ts) receives LSIF uploads and answers LSP queries by looking at relevant SQLite databases on-disk. The [worker](./src/worker/worker.ts) dequeues unconverted LSIF uploads from Postgres and converts them into SQLite databases that can be queried by the server.

### Entrypoint

The dockerfile in this directory builds the application so that the output is self-contained. The docker image which is run in a production environment is produced later by either [./cmd/lsif-server](.././cmd/lsif-server) or by [./cmd/server](.././cmd/server) with the rest of the instance.

### Documentation

- Usage documentation is provided on [Sourcegraph.com](https://docs.sourcegraph.com/user/code_intelligence/lsif).
- API endpoint documentation is provided in [api.md](./docs/api.md).
- Database configuration and migrations are described in [database.md](./docs/database.md).
- Data models are described in [datamodel.md](./docs/datamodel.md) and [datamodel.pg.md](./docs/datamodel.pg.md).