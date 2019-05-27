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

### 3.2 表操作

>旧表表数据
>
>```sql
>MariaDB [test]> select * from test;
>+----+--------+------+-------+
>| id | name   | addr | score |
>+----+--------+------+-------+
>|  1 | 测试   | NULL |   100 |
>|  2 | 测试   | NULL |   100 |
>|  3 | 测试   | NULL |   100 |
>|  4 | 测试   | NULL |   100 |
>|  5 | 测试   | NULL |   100 |
>+----+--------+------+-------+
>```
>
>旧表表结构
>
>```sql
>+-------+-------------+------+-----+--------------+----------------+
>| Field | Type        | Null | Key | Default      | Extra          |
>+-------+-------------+------+-----+--------------+----------------+
>| id    | int(11)     | NO   | PRI | NULL         | auto_increment |
>| name  | varchar(20) | YES  |     | NULL         |                |
>| addr  | varchar(20) | YES  |     | 地址不详     |                |
>| score | int(11)     | YES  |     | NULL         |                |
>+-------+-------------+------+-----+--------------+----------------+
>```
>
>

复制表 

语法：`carete table 新表名 select 字段名 from 旧表名`（复制数据，不能复制旧表的主键）

试例：

```sql
MariaDB [test]> create table test1 select * from test;
Query OK, 5 rows affected (0.33 sec)
Records: 5  Duplicates: 0  Warnings: 0

MariaDB [test]> select * from test1;
+----+--------+------+-------+
| id | name   | addr | score |
+----+--------+------+-------+
|  1 | 测试   | NULL |   100 |
|  2 | 测试   | NULL |   100 |
|  3 | 测试   | NULL |   100 |
|  4 | 测试   | NULL |   100 |
|  5 | 测试   | NULL |   100 |
+----+--------+------+-------+
5 rows in set (0.00 sec)

MariaDB [test]> desc test1;  #该表中没有主键
+-------+-------------+------+-----+--------------+-------+
| Field | Type        | Null | Key | Default      | Extra |
+-------+-------------+------+-----+--------------+-------+
| id    | int(11)     | NO   |     | 0            |       |
| name  | varchar(20) | YES  |     | NULL         |       |
| addr  | varchar(20) | YES  |     | 地址不详     |       |
| score | int(11)     | YES  |     | NULL         |       |
+-------+-------------+------+-----+--------------+-------+
4 rows in set (0.00 sec)
```

语法：`create table 新表名 like 旧表名`（复制所有的表结构，但是不复制数据）

试例：

```sql
MariaDB [test]> create table test2 like test;
Query OK, 0 rows affected (0.06 sec)

MariaDB [test]> select * from test2;
Empty set (0.00 sec)

MariaDB [test]> desc test2; #表结构全部服务制过来，但是没有数据
+-------+-------------+------+-----+--------------+----------------+
| Field | Type        | Null | Key | Default      | Extra          |
+-------+-------------+------+-----+--------------+----------------+
| id    | int(11)     | NO   | PRI | NULL         | auto_increment |
| name  | varchar(20) | YES  |     | NULL         |                |
| addr  | varchar(20) | YES  |     | 地址不详     |                |
| score | int(11)     | YES  |     | NULL         |                |
+-------+-------------+------+-----+--------------+----------------+
4 rows in set (0.01 sec)

```




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



#### 3.3.4 清空表

> ```sql
> +----+--------+------+-------+
> | id | name   | addr | score |
> +----+--------+------+-------+
> |  1 | 高手   |      |   100 |
> | 10 | 测试   | NULL |   100 |
> +----+--------+------+-------+
> ```

语法：`truncate table 表名`

试例：

```sql
MariaDB [test]> truncate table test;
Query OK, 0 rows affected (0.20 sec)
```

```
脚下留心：delete from 表和truncate table 表区别？
delete from 表：遍历表记录，一条一条的删除
truncate table：将原表销毁，再创建一个同结构的新表。就清空表而言，这种方法效率高。
```

#### 3.3.5 查询表

> ```sql
> +----+--------+------+-------+
> | id | name   | addr | score |
> +----+--------+------+-------+
> |  1 | 测试   | NULL |   100 |
> |  2 | 测试   | NULL |   100 |
> |  3 | 测试   | NULL |   100 |
> |  4 | 测试   | NULL |   100 |
> |  5 | 测试   | NULL |   100 |
> +----+--------+------+-------+
> ```

语法：`select 字段名 from 表名`

试例1：

```sql
MariaDB [test]> select name,addr from test;
+--------+------+
| name   | addr |
+--------+------+
| 测试   | NULL |
| 测试   | NULL |
| 测试   | NULL |
| 测试   | NULL |
| 测试   | NULL |
+--------+------+
```

试例2：

```sql
MariaDB [test]> select id,name,addr,score from test;
+----+--------+------+-------+
| id | name   | addr | score |
+----+--------+------+-------+
|  1 | 测试   | NULL |   100 |
|  2 | 测试   | NULL |   100 |
|  3 | 测试   | NULL |   100 |
|  4 | 测试   | NULL |   100 |
|  5 | 测试   | NULL |   100 |
+----+--------+------+-------+

```

试例3：(*代表全部字段)

```sql
MariaDB [test]> select * from test;
+----+--------+------+-------+
| id | name   | addr | score |
+----+--------+------+-------+
|  1 | 测试   | NULL |   100 |
|  2 | 测试   | NULL |   100 |
|  3 | 测试   | NULL |   100 |
|  4 | 测试   | NULL |   100 |
|  5 | 测试   | NULL |   100 |
+----+--------+------+-------+
5 rows in set (0.00 sec)

```

### 3.4 SQL分类

DDL（data definition language）数据库定义语言CREATE、ALTER、DROP、SHOW

DML（data manipulation language）数据操纵语言SELECT、UPDATE、INSERT、DELETE

DCL（Data Control Language）数据库控制语言,是用来设置或更改数据库用户或角色权限的语句

 ### 3.5 数据表的文件介绍

一个数据库对应一个文件夹

一个表对应一个或多个文件

引擎是myisam，一个表对应三个文件

 ![image](./image/1536654269605.png)

引擎是innodb,一个表对应一个表结构文件

 ![1536654519700](./image/1536654519700.png)

所有的innodb引擎的数据统一的存放在data\ibdata1文件中。如果数据量很大，MySQL会自动的创建ibdata2，ibdata3，…，目的就是为了便于管理。

 引擎是memory，数据存储在内存中，重启服务数据丢失，但是读取速度非常快。

### 3.6 字符集

