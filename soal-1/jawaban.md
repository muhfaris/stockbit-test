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

You can visit to sql compiler:
[db-fiddle.com](https://www.db-fiddle.com/f/gEByiJnuy9HxK3ckYJTPAt/3)
![image](https://user-images.githubusercontent.com/16787866/149661854-b181fa62-57c7-4969-96d4-5f6ece0d0029.png)
