# Chalenges, trade-offs and key design decisions I had while developing

- Connecting to database using sqlx to mysql
    - using mariadb image fix the issue.  using mysql image kept erroring out when connecting with DSN
        - this took longer to debug than would have liked.
    - changed back to mysql after long trouble shooting.  Port Issues

- Change DB migration libraries
    - From golang-migrate/migrate to github.com/rubenv/sql-migrate
        - easier to implement and code.
        

- Write error for PATCH when not all objects in the array has been update to make it clear
    - This was confusing at some times as I did not get the error to make it clear what I was doing wrong.

- MySQL Auto_increment increments on error
    - Can be solved with stored procedures but migration for stored procedures is not running with mariadb and the migration package
    - This causes slices to be wrong when there is errors when creating new records

- .env when using docker was not parsing through to container
    - switched to config.yaml file and is easier to use.

- Started Later than would have liked.
    - did not have time to implement unit tests
    - did not have time to implement api auth tokens
        - Would have loved to implement Authentik (Keyclock alternative) Running Locally at home to test how implementation wourks using OIDC

