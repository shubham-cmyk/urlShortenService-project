# urlShortenService

Library Used : 
1. Golang fiber
2. Google UUID
3. Redis Client 


You can configure the Your redis databse from the .env File ( Deprecated for now would add soon )

User can Define it's own short URL he need ( It's optional the logical would still work)

You would get the reponse as :- 
1. Long URL 
2. Short URL 
3. Expiration Time of the key

Default expiration time of the key is set to be 20 minute

How to use it 
1. Connect it to the redis server use .env
2. Use postMan or Curl command to send POST request 

CURL - curl -H "Content-Type: application/json" -X POST -d '{"url":"www.google.com"}' http://localhost:3000/api/


How to Deploy it : 

1. You can use docker-compose up -d to deploy it using docker
2. You can use the local host but then you need to connect your redis database to the .env

Sample Request : 
1. You can use the POSTMAN to send a API requst or use CURL 
2. Once you have set the Path you can use your  localhost:3000/{shortURL}
3. If you want to get the long URL again use localhost:3000/get/{shortURL}


Example : 

Request : curl -H "Content-Type: application/json" -X POST -d '{"url":"www.google.com"}' http://localhost:3000/api/
Response : {"url":"www.google.com","short":"localhost:3000/0ffb2f","time":20}


Request : curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:3000/get/0ffb2f
Response : 

HTTP/1.1 200 OK
Date: Sun, 18 Sep 2022 20:31:56 GMT
Content-Type: application/json
Content-Length: 50

{"url":"www.google.com","short":"0ffb2f","time":0}


Request : curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:3000/0ffb2f
Response : 
HTTP/1.1 308 Permanent Redirect
Date: Sun, 18 Sep 2022 20:32:29 GMT
Content-Length: 0
Location: www.google.com

