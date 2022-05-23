# GO-AVD-project1
- Exercise project to get hands on GO-lang. 
- Manages merchant and members accounts.

### SETUP
- clone this repo
- cd to this repo
- For installing all dependencies, use following command in terminal
```bash
go install
```
- install mysql on your local machine if not installed already
- source file or execute queries from file `./model/dbScripts/createDataStore.sql` to create database & populate it with sample data.
- you can use POSTMAN for testing APIs. Refer [APIs section](#APIs) 
- For starting the server, use following command in terminal
```bash
go run main.go
```
- For running the test cases, use following command
```bash
go test ./...
```


<br/>

### APIs
- for errorful request, in response body status field will be 1
- for successful request, in response body status field will be 0
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
<br/>

#### Merchant APIs (Request and Response): 
 - Create New Merchant API:
   - creates a new merchant with given details. Auto assign alphanumeric UNIQUE merchanId
   - merchant Name has to be unique. Min 3 and max 64 characters. can be alphanumeric and can have space and _
   - address min 6 and max 256 characters. can be alphanumeric and can have space and underscore(_) and comma (,)
   - Method: POST
```
Request -
Endpoint: localhost:9091/merchant/add
Body: 
{
    "name": "this merchant",
    "address": "130, orchard road 089775"
}
```

```
Response - 200 status OK
{
    "address": "130, orchard road 089775",
    "merchantId": "311b046f-e1fd-4fc7-8310-586318c03dcb",
    "message": "merchant added successfully",
    "name": "this merchant",
    "status": 0
}
```
<br/>

- Get Merchant API:
    - retrieves the details of given merchant
    - Method: GET
```
Request -
Endpoint: localhost:9091/merchant/:id
Example: localhost:9091/merchant/merc1   // where merc1 is merchant ID
```

```
Response - 200 status OK
{
    "address": "Bedok, 083664",
    "merchantId": "merc1",
    "message": "merchant fetched successfully",
    "name": "merchant_one",
    "status": 0
}
```
<br/>

- Update Merchant API:
    - updates the details of specified merchant
    - only address field can be updated in merchant record
    - Method: PUT
```
Request -
Endpoint: localhost:9091/merchant/:id
Example: localhost:9091/merchant/4ec7bae8-4622-448e-b46f-e462acda0365
here 4ec7bae8-4622-448e-b46f-e462acda0365 is merchant ID
Body: 
{
    "address": "orchard 989777"
}
```
```
Response - 200 status OK
{
    "merchantId": "4ec7bae8-4622-448e-b46f-e462acda0365",
    "message": "merchant updated successfully",
    "status": 0
}
```
<br/>



- Delete Merchant API:
    - Deletes the specified merchant from system
    - Method: DELETE
```
Request -
Endpoint: localhost:9091/merchant/:id
Example: localhost:9091/merchant/merc1   // where merc1 is merchant ID
```

```
Response - 200 status OK
{
    "merchantId": "merc3",
    "message": "merchant deleted successfully",
    "status": 0
}

```
<br/>


#### Member APIs (Request and Response):
- Create New Member API:
    - creates a new member in specified merchant account(team) with given details. Auto assign UNIQUE memberId
    - Name min 3 and max 64 characters. can be alphanumeric and can have space and underscore(_)
    - valid email address
    - merchantId of existing merchant
    - Method: POST
```
Request -
Endpoint: localhost:9091/member/add
Body: 
{
    "name": "this avd2",
    "email": "adsdas99@ade.com",
    "merchantId": "merc2"
}
```

```
Response - 200 status OK
{
    "email": "adsdas99@ade.com",
    "memberId": 12,
    "merchantId": "merc2",
    "message": "member added successfully",
    "name": "this avd2",
    "status": 0
}
```
<br/>

- Get Member API:
    - retrieves the details of given member
    - Method: GET
```
Request -
Endpoint: localhost:9091/member/:id
Example: localhost:9091/member/6   // where 6 is member ID
```

```
Response - 200 status OK
{
    "email": "adsdas99@ade.com",
    "memberId": 6,
    "merchantId": "merc1",
    "message": "Member fetched successfully",
    "name": "this avd",
    "status": 0
}
```
<br/>

- Update Member API:
    - updates the details of specified member
    - only email field can be updated in member record
    - Method: PUT
```
Request -
Endpoint: localhost:9091/member/:id
Example: localhost:9091/member/12       //here 12 is member ID
Body: 
{
    "email": "avd12@game.com"
}
```
```
Response - 200 status OK
{
    "memberId": 12,
    "message": "member updated successfully",
    "status": 0
}
```
<br/>

- Delete Member API:
    - Deletes the specified member from system
    - Method: DELETE
```
Request -
Endpoint: localhost:9091/member/:id
Example: localhost:9091/member/12   //here 12 is member ID
```

```
Response - 200 status OK
{
    "memberId": 12,
    "message": "Member deleted successfully",
    "status": 0
}

```
<br/>

- Get Member List By Merchant API:
    - Method: GET
    - Query params: 
      - page : number of page which needs to be fetched
    - pagination support. Page size is 4 records per page.
```
Request -
Endpoint:  localhost:9091/members/list/:merchantId?page=<page Number>

Example: localhost:9091/members/list/merc1?page=2
here merc1 is merchantId where as 2nd page is being fetched
```
```
Response - 200 status OK
{
    "count": 4,
    "members": [
        {
            "id": 5,
            "name": "avd5",
            "email": "avd5@gmail.com",
            "merchantId": "merc1"
        },
        {
            "id": 6,
            "name": "avd6",
            "email": "avd6@gmail.com",
            "merchantId": "merc1"
        },
        {
            "id": 7,
            "name": "avd7",
            "email": "avd7@gmail.com",
            "merchantId": "merc1"
        },
        {
            "id": 8,
            "name": "avd8",
            "email": "avd8@gmail.com",
            "merchantId": "merc1"
        }
    ],
    "message": "Members fetched successfully",
    "page": 2,
    "pageSize": 4,
    "status": 0
}
```
