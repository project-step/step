# Step

Step is a stream & batch computing library to provide fast analytical data computation.

Test:
From server side:
```
./bin/step session --mysql.pass 123456 --mysql.db pairportal --pass abcde
```

From client side: 
```
mysql -h 127.0.0.1 -P 3309 -u root -p
```
Then:
```
mysql> select name, userid, gender from users;
```