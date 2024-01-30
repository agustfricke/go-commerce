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
    name    VARCHAR(128) NOT NULL,
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

```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "name": "Agustttt",
  "email": "agust@agust.com",
  "password": "root"
}' http://localhost:6969/signup
```

```bash
{"data":{"ID":1,"Name":"Agustttt","Email":"agust@agust.com","Password":"$2a$10$T7CdMQDgXdiBR6HLtPMvoOOlkFCxssU.26qmgx3bbZKbowAylpUma","Verified":false,"SocialID":"","Avatar":"","Otp_enabled":false,"Otp_verified":false,"Otp_secret":"","Otp_auth_url":""},"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDY2NzQ4MDQsImlhdCI6MTcwNjU4ODQwNCwibmJmIjoxNzA2NTg4NDA0LCJzdWIiOjF9.Jn6l4D48At7tSZFtLipV1AeMQKRpBVDles5zdVGDlNA"}%
```

