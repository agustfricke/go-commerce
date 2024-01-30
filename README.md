```bash
sudo docker run --name go-commerce-mysql -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -d mysql:8.0-debian
sudo docker exec -it go-commerce-mysql bash
mysql -u root -p
create database gocommerce;
use gocommerce;
```

```sql
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id            INT AUTO_INCREMENT NOT NULL,
    first_name    VARCHAR(128) NOT NULL,
    last_name     VARCHAR(128) NOT NULL,
    email         VARCHAR(128) NOT NULL,
    password      VARCHAR(128) NOT NULL,
    social_id     VARCHAR(128),
    avatar        VARCHAR(128) NOT NULL,
    otp_enabled   BOOLEAN DEFAULT FALSE,
    otp_verified  BOOLEAN DEFAULT FALSE,
	otp_secret    VARCHAR(128),
	otp_auth_url  VARCHAR(128),
  PRIMARY KEY (`id`)
);
```
