## The Question
### Simple Database Query 
Terdapat sebuah table "USER" yang memiliki 3 kolom: `ID`, `username`, `parent` 

Dimana:
- Kolom ID adalah Primary Key.
- Kolom username adalah Nama User.
- Kolom Parent adalah ID dari User yang menjadi Creator untuk User tertentu.

Example:

| id | username | parent |
|----|----------|--------|
| 1  | Ali      |    2   |
| 2  | Budi     |    0   |
| 3  | Cecep    |    1   |

Tuliskan SQL Query untuk mendapatkan data berisi:

| id | username | parent_username |
|----|----------|-----------------|
| 1  | Ali      |    Budi         |
| 2  | Budi     |    NULL         |
| 3  | Cecep    |    Ali          |

* Kolom Parent Username adalah Username berdasarkan value parent 

## Answer
### Solution
### Table users schema
```
Create table users(
	id serial PRIMARY key,
  	username VARCHAR(50) UNIQUE not null,
  	parent_id int not null
)
```

#### Seeding of users table 
```
BEGIN;
insert INTO users(id, username, parent_id) VALUES(1, 'Ali', 2);
insert INTO users(id, username, parent_id) VALUES(2, 'Budi', 0);
insert INTO users(id, username, parent_id) VALUES(3, 'Cecep', 1);
Commit;
```

#### Get All data of users
```
select * from users
```

Result Query:
``` 
id	username	parent_id
1	Ali	2
2	Budi	0
3	Cecep	1
```

### Get Parent Username 
``` 
select 
	u.id,
    u.username, 
    ( select 
     	u1.username 
      	from users u1 
     	where u1.id = u.parent_id 
    ) as parent_user_name
from users u
```

Result:
``` 
id	username	parent_user_name
1	Ali	Budi
2	Budi	null
3	Cecep	Ali

```
### SQL Online Compiler
I create schema and query in SQL online compiler, You can visit to sql compiler:

[db-fiddle.com](https://www.db-fiddle.com/f/gEByiJnuy9HxK3ckYJTPAt/3)

![image](https://user-images.githubusercontent.com/16787866/149661854-b181fa62-57c7-4969-96d4-5f6ece0d0029.png)
