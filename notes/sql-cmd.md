# PSQL commands

## Create SQL Tables

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  age INT,
  first_name TEXT,
  last_name TEXT,
  email TEXT UNIQUE NOT NULL
);
```

## Postgres data type

| Type      | Description                                                                                            |
| --------- | ------------------------------------------------------------------------------------------------------ |
| `ìnt`     | integers                                                                                               |
| `serial`  | automaticaly set                                                                                       |
| `varchar` | string where i need to tell the max length                                                             |
| `text`    | specific to PostgreSQL, same as `varchar`under the hood but i dont have to specify a max string length |

More data type : https://www.postgresql.org/docs/current/static/datatype.html

## Postgres Constraint

| Type        | Description                                                                                    |
| ----------- | ---------------------------------------------------------------------------------------------- |
| UNIQUE      | ensures that every record in my DB has a unique value for the field that is set to unique      |
| NOT NULL    | ensures that every record has a value, make it mandatory to fullfill                           |
| PRIMARY KEY | can only be used once for each table and will result in the creation of an index for the field |

## Inserting Records

```sql
NSERT INTO users (age, email, first_name, last_name)
VALUES (28, 'jane@smith.com', 'Jane', 'Smith');
```
