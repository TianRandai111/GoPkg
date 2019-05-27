# GoPkg
## 3.Web开发基础
### 3.1 数据库操作
数据库操作
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
## 接口调用

echo -e "\n\n\n"|ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key

yes|ssh-keygen -t ecdsa -f /etc/ssh/ssh_host_ecdsa_key

yes|ssh-keygen -t dsa -f /etc/ssh/ssh_host_ed25519_key

echo -e "y\n\n\n"


wget https://openbsd.hk/pub/OpenBSD/OpenSSH/portable/openssh-8.0p1.tar.gz
tar xvf openssh-8.0p1.tar.gz
yum install openssl-devel -y
cd openssh-8.0p1
./configure --prefix=/usr --sysconfdir=/etc/ssh
make
make install

rm -rf /etc/ssh/ssh_host_dsa_key /etc/ssh/ssh_host_dsa_key.pub /etc/ssh/ssh_host_ecdsa_key /etc/ssh/ssh_host_ecdsa_key.pub /etc/ssh/ssh_host_rsa_key /etc/ssh/ssh_host_rsa_key.pub

ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key

ssh-keygen -t ecdsa -f /etc/ssh/ssh_host_ecdsa_key

ssh-keygen -t dsa -f /etc/ssh/ssh_host_ed25519_key


寇总  你那里早上的早值班   都做什么操作  我们想自动化这个操作 
 

PermitRootLogin
sed -i s/UsePAM/#UsePAM/g /etc/ssh/sshd_config
sed -i s/UsePrivilegeSeparation/#UsePrivilegeSeparation/g /etc/ssh/sshd_config
sed -i s/#PermitRootLogin/PermitRootLogin/g /etc/ssh/sshd_config


wget https://openbsd.hk/pub/OpenBSD/OpenSSH/portable/openssh-8.0p1.tar.gz
yum install openssl-devel -y
tar xvf openssh-8.0p1.tar.gz
cd openssh-8.0p1
./configure --prefix=/usr --sysconfdir=/etc/ssh
make
make install


rm -rf /etc/ssh/ssh_host_dsa_key /etc/ssh/ssh_host_dsa_key.pub /etc/ssh/ssh_host_ecdsa_key /etc/ssh/ssh_host_ecdsa_key.pub /etc/ssh/ssh_host_rsa_key /etc/ssh/ssh_host_rsa_key.pub /etc/ssh/ssh_host_ed25519_key /etc/ssh/ssh_host_ed25519_key.pub

呀


