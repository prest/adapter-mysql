# MySQL Adapter to pRest project https://postgres.rest/

- [ ] Not Implemented
- [X] Implemented

[-] Not Applicable

## Select - GET
Select operations over a TABLE
- [X] /DATABASE/SCHEMA/TABLE (show all rows, find by database and table)
- [X] /DATABASE/SCHEMA/TABLE?_select=column (select statement by columns)
- [X] /DATABASE/SCHEMA/TABLE?_select=* (select all from TABLE)
- [X] /DATABASE/SCHEMA/TABLE?_count=* (use count function)
- [X] /DATABASE/SCHEMA/TABLE?_count=column (use count function)
- [X] /DATABASE/SCHEMA/TABLE?_page=2&_page_size=10 (pagination, page_size 10 by default)
- [X] /DATABASE/SCHEMA/TABLE?FIELD=VALUE (filter)
- [X] /DATABASE/SCHEMA/TABLE?_renderer=xml (JSON by default)

[-] /DATABASE/SCHEMA/TABLE?_select=column[array id] (select statement by array colum)

Join
- [X] /DATABASE/SCHEMA/Table?_join=Type:Table2:Table.field:Operator:Table2.field

Select operations over a VIEW
- [X] /DATABASE/SCHEMA/TABLE (show all rows, find by database and VIEW)
- [X] /DATABASE/SCHEMA/VIEW?_select=column (select statement by columns in VIEW)
- [X] /DATABASE/SCHEMA/VIEW?_select=* (select all from VIEW)
- [X] /DATABASE/SCHEMA/VIEW?_count=* (use count function)
- [X] /DATABASE/SCHEMA/VIEW?_count=column (use count function)
- [X] /DATABASE/SCHEMA/VIEW?_page=2&_page_size=10 (pagination, page_size 10 by default)
- [X] /DATABASE/SCHEMA/VIEW?FIELD=VALUE (filter)
- [X] /DATABASE/SCHEMA/VIEW?_renderer=xml (JSON by default)

Database structure
- [ ] /databases (show all databases)
- [ ] /databases?_count=* (count all databases)
- [ ] /databases?_renderer=xml (JSON by default)
- [ ] /schemas (show all schemas)
- [ ] /schemas?_count=* (count all schemas)
- [ ] /schemas?_renderer=xml (JSON by default)
- [ ] /tables (show all tables)
- [ ] /tables?_renderer=xml (JSON by default)
- [ ] /DATABASE/SCHEMA (show all tables, find by schema)
- [ ] /DATABASE/SCHEMA?_renderer=xml (JSON by default)

## Insert - POST
- [X]/DATABASE/SCHEMA/TABLE
```json
{
    "FIELD1": "string value",
    "FIELD2": 1234567890
}
```

## Update - PATCH/PUT
- [X] /DATABASE/SCHEMA/TABLE?FIELD1=xyz
```json
{
    "FIELD1": "string value",
    "FIELD2": 1234567890
}
```

## Delete - DELETE
- [X] /DATABASE/SCHEMA/TABLE?FIELD1=xyz