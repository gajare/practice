# Go Social Media Application

This project is a simple implementation of a social media application using Golang, Kafka, and PostgreSQL. The architecture consists of an API server that handles HTTP requests, a Kafka producer that sends messages to Kafka topics, and a Kafka consumer that processes messages and updates the PostgreSQL database.

## Prerequisites

- Docker and Docker Compose
- Go (Golang)

## Getting Started

### Step 1: Clone the Repository

If you have your project in a repository, clone it. If not, create a new directory and set up the files as mentioned.

```sh
git clone <repository-url>
cd go-social-media

docker-compose up -d    
#This will start Kafka, Zookeeper, and PostgreSQL in detached mode.
```

###	1.	Connect to the PostgreSQL container:
```sh 
    docker exec -it kafka_demo-postgres-1 psql -U user -d social_media
```

### 2.	Run the SQL commands from sql/schema.sql:
```
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50),
    content TEXT
);
\q
```

### Step 4: Create the Kafka Topic
#### 4.1.	Enter the Kafka container:

```sh
    docker exec -it kafka_demo-kafka-1 /bin/bash
```
#### 4.	2.	Create the social_media topic:


```sh
   kafka-topics --create --topic social_media --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
#This will create a Kafka topic called social_media.
\
```
#### 4.3 Verify the topic creation:

``` sh
    kafka-topics --list --bootstrap-server localhost:9092
```

## Step 5: Run the API Server

``` sh
    cd api
    go run main.go

    #This will start the API server on http://localhost:8080.
```

## Step 6: Run the Kafka Consumer
``` sh
    cd consumer
    go run main.go
```
### Step 7: Test the API

``` sh
curl -X POST http://localhost:8080/post -H "Content-Type: application/json" -d '{"user_id": "user1", "content": "Hello, World!"}'
```


### 1.	Connect to PostgreSQL:
    ``` sh
    docker exec -it kafka_demo-postgres-1 psql -U user -d social_media
    ```

### 2.	Query the posts table:
``` sh
SELECT * FROM posts;

```