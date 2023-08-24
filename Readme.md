
# Mcache


Mcache is a simple and efficient in-memory opensource cache server built using Go (Golang) 1.20. It provides a lightweight caching solution for your applications, helping you improve performance and reduce the load on your primary data stores.

Mcache, builds upon the principles of caching and draws inspiration from systems like Memcache to provide a lightweight, in-memory caching solution for Go applications. It leverages the efficiency of Go 1.20 to offer a performant and concurrent caching system that can help improve the speed and efficiency of your applications.

Mcache leverages the power of Representational State Transfer (REST) to ensure seamless compatibility with a wide range of systems. By adhering to REST principles, Mcache provides a simple and standardized way for clients to interact with the cache server over HTTP. Clients can easily perform common cache operations such as setting, getting, and deleting key-value pairs using standard HTTP methods like GET, POST, PUT, and DELETE. This RESTful design not only simplifies integration with various programming languages and platforms but also allows developers to harness the full potential of Mcache in a familiar and intuitive manner. Whether you're building a web application, a mobile app, or a microservices architecture, Mcache's RESTful API ensures that caching is both accessible and efficient across your entire system, promoting faster data retrieval and improved application performance.

Mcache is engineered with scalability and performance in mind. Its support for vertical scaling means that as your application's demands grow, you can efficiently allocate additional resources to Mcache to handle increased workloads. Moreover, Mcache's multithreaded architecture enables it to handle concurrent requests seamlessly, making it well-suited for high-throughput applications. Its memory-safe design ensures that data integrity and application stability are maintained even under heavy workloads, preventing common memory-related issues. With Mcache, you can confidently scale your caching layer as your application evolves, knowing that it will deliver both speed and reliability, while also keeping your data secure and readily accessible. Whether you're running a small-scale project or a large enterprise application, Mcache's combination of vertical scalability, multithreaded capabilities, and memory safety ensures optimal caching performance under all conditions.




## Tech Stack

**Server:** Golang



## Liscence

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)




## Features

- Light weight
- Scalable
- Containerized
- Portable
- REST Interface


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`PORT` - default: 4567

`HOST` - eg: localhost

`AUTH_ENABLED` - true/false | if enabeled basic auth should be provided in all request.

`AUTH_USER` - eg: Admin

`AUTH_PASSWORD` - eg: ********

`PERSIST` - true/false | if enabeled, mcache will persist the data.

## Installation

Install mcache native using repo

```bash
  git clone https://github.com/Midhun-Chandrasekhar/mcache.git
  cd mcache
  go mod download
  go build -o mcache
  ./mcache
```
    
Install mcache docker using repo

```bash
  git clone https://github.com/Midhun-Chandrasekhar/mcache.git
  cd mcache
  docker-compose up -d
```

Install mcache docker

```bash
  docker pull csekharjr/mcache:latest
  docker run --name docker-mcache -p 4567:4567 mcache
```
    
## Usage/Examples

### Get Keys
```javascript
Enpoint: /keys
Method: GET
Info: Retrive all keys available in cache

```

### Get Value
```javascript
Enpoint: /get?key={key}  
Method: GET
Info:  Retrieve value for given key
```

### Set Key
```javascript
Enpoint: /set?key={key}&value={value} 
Method: POST
Info: Add new key-value to cache
```

### Set Key
```javascript
Enpoint: /delete?key={key}
Method: DELETE
Info: Delete key from cache  
```

## Roadmap

- Alpha v1.0.0 - Released

- Alpha v2.0.0 (Distributed cache)


## Authors

- Midhun Chandrasekhar [@csekhar.jr](https://github.com/Midhun-Chandrasekhar)

## 🚀 About Me
I'm a full stack Technical architect experienced in building large scale distributed system in the domaince of fintech, AI, blockchain, healthcare, transportaion, hospitality and etc.


# Hi, I'm Midhun Chandrasekar! 👋



## 🔗 Links
[![portfolio](https://img.shields.io/badge/my_portfolio-000?style=for-the-badge&logo=ko-fi&logoColor=white)](https://csekhar.vercel.app/)
[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](www.linkedin.com/in/csekhar-jr)


## Other Common Github Profile Sections
👩‍💻 I'm currently working on Generative AI

🧠 I'm currently learning Quantm computing

👯‍♀️ I'm looking to collaborate for Mcache

🤔 I'm looking for help with the further development of Mcache

💬 Ask me about Technology

📫 csekhar.jr@gmail.com



## Contributing

Contributions are always welcome!

Mail me to get in touch|!

## Support

For support, email csekhar.jr@gmail.com.

