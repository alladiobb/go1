## Rodando o MYSQL no Container 

"-d" é para liberar o terminal 

```
docker compose up -d
```

```
docker compose exec mysql bash
```

```
mysql -uroot -p golang
```

```
go mod init
go mod tidy
```

### SQL 
```
 create table products (id varchar(255), name varchar(255), price decimal(10,2), primary key(id));
```