# `GO WEB SERVER`
A fast web server that connects to a remote render postgreSQL database and uses the mux framework.

## `Running locally`
The server requires a valid postgreSQL database connection string, stored in your `.env` file.
I've setup a taskfile to automatically format and run the webserver using the command:
```bash
task serve
```
NOTE: you need to have Task installed locally.
