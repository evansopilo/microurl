# Micro-URL
A web service for shortening URLs. Takes a long URL the returns a shorten version to the client, when the short URL is clicked the client application is directed to the original URL.

## Functional requirements.
1. User can submit and receive a shorter version of a URL.
2. Created URL can be customized.
3. URL can be created with expiration time.
4. Advanced analytics on URL interactions ie. clicks.
5. Generate QR code for the shorter URL.

## Non-functional requirements.
1. Security, the web service shoud be secured.
2. Scalable, the web service should at scale.
3. Latency, low latency when serving client request.
4. Peformance, the service should offer high peformance to users.

## requirements
1. MongoDB 
2. Redis Cache
3. Go
4. Docker 
