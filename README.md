# ON AIR!
fiber-air-docker development environment boilerplate

## TODO
- on air 세션 준비
- fiber 유저 준비
- gorm 외래키 준비 

![](static/ONAIR.png)

아키텍쳐 참고  
https://blog.puppyloper.com/menus/Golang/articles/Golang%EA%B3%BC%20Clean%20Architecture

- Handler
  - web framework 를 통해 request 가 직접 도달하는 layer
- UseCase
  - Handler 에서 실행되는 비즈니스 로직 layer
- Repository
  - UseCase 에서 실행되는 layer (DB, 등 기타 외부와의 연결)


## Run Project Development
```bash
vi .env/dev.env

PORT=:9000
POSTGRES_HOST=onair-pg
POSTGRES_DB=onair
POSTGRES_USER=onair
POSTGRES_PASSWORD=onair
POSTGRES_PORT=5432
POSTGRES_TZ=Asia/Seoul
PGDATA=/var/lib/postgresql/data/db-files/
SECRET_KEY=a-o{<%MyF.apGy'<LB-SZ^jK!3N{ni
MODE=debug
```

```bash
docker-compose up
```

## Connect To Postgresql 
```
docker exec -it <postgresql_container> /bin/bash
...
root@49d68bd0cacd:/# psql -U <user> <database>
psql (13.1 (Debian 13.1-1.pgdg100+1))
Type "help" for help.
...
```

```
# \dt
          List of relations
 Schema |   Name    | Type  |  Owner  
--------+-----------+-------+---------
 public | books | table | neulhan
(3 rows)
...

# SELECT * FROM books LIMIT 1;
 id |          created_at           |          updated_at           | deleted_at |  name | email | pass | logged_in 
----+-------------------------------+-------------------------------+------------+-------+-------+------+-----------
  1 | 2021-01-19 00:14:24.580872+09 | 2021-01-19 00:14:24.580872+09 |            |       |       |      | t
(1 row)
...
```
