# MICRO URL

A web service for shortening URLs. Takes a long URL the returns a shorten version to the client, when the short URL is clicked the client application is directed to the original URL.

## Functional requirements.

1. User can submit and receive a shorter version of a URL.
2. Created URL can be customized.
3. URL should be short and readable.
4. URL can be created with expiration time.
5. Advanced analytics on URL interactions ie. clicks.
6. Generate QR code for the shorter URL.

## Non-functional requirements.

1. Security, the web service shoud be secured.
2. Scalable, the web service should at scale.
3. Latency, low latency when serving client request.
4. Peformance, the service should offer high peformance to users.

## Tools and Requirements

1. [MongoDB]("https://www.mongodb.com/) - MongoDB is a source-available cross-platform document-oriented database program. Classified as a NoSQL database program, MongoDB uses JSON-like documents with optional schemas. MongoDB is developed by MongoDB Inc.
2. [Redis Cache]("https://redis.io/") - Redis is an in-memory data structure store, used as a distributed, in-memory key–value database, cache and message broker, with optional durability.
3. [Go]("https://go.dev/") - Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.
4. [Docker]("https://www.docker.com/") - Docker is a set of platform as a service products that use OS-level virtualization to deliver software in packages called containers. Containers are isolated from one another and bundle their own software, libraries and configuration files; they can communicate with each other through well-defined channels.

## How to install

### Using Git (recommended)

1.  Clone the project from github. 
```bash
git clone https://github.com/evansopilo/MICRO-URL.git 
```

### Using manual download ZIP

1.  Download repository
2.  Uncompress to your desired directory

### Setting up environments (development or production)

1.  In the root this repository you will find a file named `.env.example`
2.  Create a new file by copying and pasting the file and then renaming it to just `.env`
3.  The file `.env` is already ignored, so you never commit your credentials.
4.  Change the values of the file to your environment (development or production)
5.  Upload the `.env` to your environment server(development or production)
6.  If you use the postman collection to try the endpoints, change value of the variable `server` on your environment to the url of your server, for development mode use <http://localhost:8080>

## Bugs or improvements

Feel free to report any bugs or improvements. Pull requests are always welcome.

## License

This project is open-sourced software licensed under the MIT License. See the LICENSE file for more information.
