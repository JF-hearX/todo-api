# Issues I had while developing

- Connecting to database using sqlx to mysql
    - using mariadb image fix the issue.  using mysql image kept erroring out when connecting with D

- Change DB migration libraries
    - From golang-migrate/migrate to github.com/rubenv/sql-migrate
        - easier to implement and code.
        