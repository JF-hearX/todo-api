# Issues I had while developing

- Connecting to database using sqlx to mysql
    - using mariadb image fix the issue.  using mysql image kept erroring out when connecting with D

- Change DB migration libraries
    - From golang-migrate/migrate to github.com/rubenv/sql-migrate
        - easier to implement and code.
        

- Write error for PATCH when not all objects in the array has been update to make it clear
    - This was confusing at some times as I did not get the error to make it clear what I was doing wrong.