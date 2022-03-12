## Transactions test
Simple web application whcih provides API for transactions.

### Run locally
App can be run locally via docker, simply run:
```sh
docker-compose up -d
```
NOTE: please make sure you have import sample database right after app has been started for the first time!
```sh
make import-db
```
Then you can check swagger docs and do testing
```
http://localhost:8080/swagger/index.html
```

### Database diagram
![App Screenshot](https://i.imgur.com/ng7bUtg.png)

In diagram above we can see only 2 tables:
 - transactions - table where all transactions are saved.
 - accounts - table where user accounts details are saved.

Transactions table has `parent_id` which is used to determines if transaction is in the top level or it is related to another transaction,
whereas `order` field is responsible for order to display transactions in the list.
