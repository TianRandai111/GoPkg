# GoPkg
## 3.Web开发基础
### 3.1 数据库操作
```sql
#查询数据库
show databases 

#查看数据库创建语句
show create database [库名]

#创建数据库
create database [库名]

#选择数据库
use database [库名]

#修改数据库
alter database [库名] charset=[字符编码]

#删除数据库
drop database [库名]
```

创建表






### 3.3 数据操作

#### 3.3.2 修改数据

`update 表名 set 字段名=修改值 where 判断条件`

>样例
>
>```sql
>MariaDB [test]> select * from test;
>+----+--------------+--------------+-------+
>| id | name         | addr         | score |
>+----+--------------+--------------+-------+
>|  1 | 步荀仙       | 不知道       |   100 |
>|  2 | 蒙斯克       | 克哈         |    99 |
>|  3 | NULL         | 地址不详     |    89 |
>|  4 | 刀锋女王     | 查尔         |    78 |
>|  5 | 阿塔尼斯     | 艾尔         |   100 |
>|  6 | 霍纳         | 地址不详     |   100 |
>|  7 | 霍纳         | 地址不详     |   100 |
>|  8 | 凯瑞甘       | 地址不详     |   100 |
>|  9 | 诺娃         | 帝国         |    99 |
>+----+--------------+--------------+-------+
>```



- 修改1号学生的地址

```sql
MariaDB [test]> update test set addr='新世界' where id=1;
MariaDB [test]> select * from test;
+----+--------------+--------------+-------+
| id | name         | addr         | score |
+----+--------------+--------------+-------+
|  1 | 步荀仙       | 新世界       |   100 |
```



- 修改学生[霍纳]成绩

```sql
MariaDB [test]> update test set addr='艾尔' where name='霍纳';
MariaDB [test]> select * from test;；
+----+--------------+--------------+-------+
| id | name         | addr         | score |
+----+--------------+--------------+-------+
|  6 | 霍纳         | 艾尔         |   100 |
|  7 | 霍纳         | 艾尔         |   100 |
```



- 修改学生[刀锋女王]地址，并同时修改成绩

```sql
MariaDB [test]> update test set addr='新盖茨堡',score=1000 where name='刀锋女王';
MariaDB [test]> select * from test;
+----+--------------+--------------+-------+
| id | name         | addr         | score |
+----+--------------+--------------+-------+
|  4 | 刀锋女王     | 新盖茨堡     |  1000 |
```



- 修改地址为[艾尔]的同学的成绩

```sql
MariaDB [test]> update test set score=99 where addr='艾尔';
MariaDB [test]> select * from test;
+----+--------------+--------------+-------+
| id | name         | addr         | score |
+----+--------------+--------------+-------+
|  5 | 阿塔尼斯     | 艾尔         |    99 |
|  6 | 霍纳         | 艾尔         |    99 |
|  7 | 霍纳         | 艾尔         |    99 |
```



- 忽略条件会更改所有数据

```sql
MariaDB [test]> update test set addr='地球',score=60;
MariaDB [test]> select * from test;
+----+--------------+--------+-------+
| id | name         | addr   | score |
+----+--------------+--------+-------+
|  1 | 步荀仙       | 地球   |    60 |
|  2 | 蒙斯克       | 地球   |    60 |
|  3 | NULL         | 地球   |    60 |
|  4 | 刀锋女王     | 地球   |    60 |
|  5 | 阿塔尼斯     | 地球   |    60 |
|  6 | 霍纳         | 地球   |    60 |
|  7 | 霍纳         | 地球   |    60 |
|  8 | 凯瑞甘       | 地球   |    60 |
|  9 | 诺娃         | 地球   |    60 |
+----+--------------+--------+-------+
```



- 将id为2、3的学生成绩该为65

```sql
 MariaDB [test]> update test set score=65 where id=2 or id=3;
MariaDB [test]> select * from test;
+----+--------------+--------+-------+
| id | name         | addr   | score |
+----+--------------+--------+-------+
|  2 | 蒙斯克       | 地球   |    65 |
|  3 | NULL         | 地球   |    65 |

```




#### 3.3.3 删除数据

语法: `dalete from 表名 [where 条件]` 

> 样例
>
> ```sql
> MariaDB [test]> select * from test;
> +----+--------------+--------+-------+
> | id | name         | addr   | score |
> +----+--------------+--------+-------+
> |  1 | 步荀仙       | 地球   |    60 |
> |  2 | 蒙斯克       | 地球   |    65 |
> |  3 | NULL         | 地球   |    65 |
> |  4 | 刀锋女王     | 地球   |    60 |
> |  5 | 阿塔尼斯     | 地球   |    60 |
> |  6 | 霍纳         | 地球   |    60 |
> |  7 | 霍纳         | 地球   |    60 |
> |  8 | 凯瑞甘       | 地球   |    60 |
> |  9 | 诺娃         | 地球   |    60 |
> +----+--------------+--------+-------+
> ```
>
>

- 删除学号是1号的学生

```sql
MariaDB [test]> delete from test where id=1;
MariaDB [test]> select * from test;
+----+--------------+--------+-------+
| id | name         | addr   | score |
+----+--------------+--------+-------+
|  2 | 蒙斯克       | 地球   |    65 |
|  3 | NULL         | 地球   |    65 |
|  4 | 刀锋女王     | 地球   |    60 |
|  5 | 阿塔尼斯     | 地球   |    60 |
|  6 | 霍纳         | 地球   |    60 |
|  7 | 霍纳         | 地球   |    60 |
|  8 | 凯瑞甘       | 地球   |    60 |
|  9 | 诺娃         | 地球   |    60 |
+----+--------------+--------+-------+
```



- 删除成绩打于等于65的学生

```sql
MariaDB [test]> delete from test where score>=65;
Query OK, 2 rows affected (0.01 sec)

MariaDB [test]> select * from test;
+----+--------------+--------+-------+
| id | name         | addr   | score |
+----+--------------+--------+-------+
|  4 | 刀锋女王     | 地球   |    60 |
|  5 | 阿塔尼斯     | 地球   |    60 |
|  6 | 霍纳         | 地球   |    60 |
|  7 | 霍纳         | 地球   |    60 |
|  8 | 凯瑞甘       | 地球   |    60 |
|  9 | 诺娃         | 地球   |    60 |
+----+--------------+--------+-------+
```



- 删除表中所有记录

```sql
MariaDB [test]> delete from test;
Query OK, 6 rows affected (0.01 sec)

MariaDB [test]> select * from test;
Empty set (0.01 sec)

```



