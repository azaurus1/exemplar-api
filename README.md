# exemplar-api

An exemplar snapified API service.

## API Service:
Uses:
- `postgres` as a database.
- `sqlc` for generating database code from `sql`
- `dbmate` for migrations
- `cobra` for commands
- `viper` for config + live-reloads
- Migrations
    - Database migrations performed immediately on application start.

## Snap:
Packages the application code into a Snap that can be deployed via `snapd`

## Features:
- Live-reload
    - use `snap set port=[PORT]` to change the port which the `http` service is running on, while the application is running.

## Installation:
1. `git clone` this repo
2. `cd` into repo
3. run `snapcraft`, this will create a `.snap` files
4. run `snap install exemplar-api.snap --dangerous`

## Data:
`migrations` and `config` are stored in `$SNAP_DATA` (e.g. `/snap/exemplar-api/[version]/`) this directory is writable only via `snapd`, e.g. via `snap set` or the `refresh` hook 
