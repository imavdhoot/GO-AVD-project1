# GO-AVD-project1
- Exercise project to get hands on GO-lang. 
- Manages merchant and members accounts.

### SETUP
- clone this repo
- cd to this repo
- run this command `go get` to install all dependencies
- install mysql on your local machine if not installed already
- source `./model/dbScripts/createDataStore.sql` file to create database & populate it with sample data.
- you can use POSTMAN for testing APIs. Refer [APIs section](#APIs) 





### APIs
- for errorful request, status will be 1
- for successful request status will be 0
- below are some sample response for error request

eg. 1:  
```
{
    "error": "Error 1062: Duplicate entry 'this merchant' for key 'merchants.idxName'",
    "status": 1
}
```
eg. 2:
```
{
    "error": "invalid character '\"' after object key:value pair",
    "status": 1
}
```



curl --location --request POST 'localhost:9091/merchant/add' \
--header 'Content-Type: application/json' \
--data-raw '{
"name": "this merchant",
"address": "130, orchard road 089775"
}'

{
"address": "130, orchard road 089775",
"merchantId": "311b046f-e1fd-4fc7-8310-586318c03dcb",
"message": "merchant added successfully",
"name": "this merchant",
"status": 0
}

curl --location --request PUT 'localhost:9091/merchant/4ec7bae8-4622-448e-b46f-e462acda0365' \
--header 'Content-Type: application/json' \
--data-raw '{
"address": "orchard 989777"
}'

{
"merchantId": "4ec7bae8-4622-448e-b46f-e462acda0365",
"message": "merchant updated successfully",
"status": 0
}




curl --location --request GET 'localhost:9091/merchant/merc1'
{
"address": "Bedok, 083664",
"merchantId": "merc1",
"message": "merchant fetched successfully",
"name": "merchant_one",
"status": 0
}






curl --location --request GET 'localhost:9091/members/list/merc1?page=2'


{
"count": 2,
"members": [
{
"id": 3,
"name": "avd3",
"email": "avd3@gmail.com",
"merchantId": "merc1"
},
{
"id": 4,
"name": "avd4",
"email": "avd4@gmail.com",
"merchantId": "merc1"
}
],
"message": "Member fetched successfully",
"page": 2,
"pageSize": 2,
"status": 0
}
