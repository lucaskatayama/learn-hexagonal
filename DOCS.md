# Hexagonal Architecture

![arch](https://herbertograca.files.wordpress.com/2018/11/080-explicit-architecture-svg.png)

Following a simpler diagram:

```mermaid
---
title: Hexagonal architecture
---
graph LR
    subgraph ui
        brokerui[broker]
        grpc
        rest
    end
    
    subgraph app
        
        service <==> entity
        service --> brokerp([BrokerPort])
        service --> emailp([EmailPort])
        service --> dbp([DBPort])
    end
    
    
    subgraph infra
        subgraph broker
            KafkaAdapter
            brokermock[MockAdapter]
        end
        subgraph db
            PostgresAdapter
            CassandraAdapter
            dbmock[MockAdapter]
        end
        subgraph email
            SESAdapter
            SendgridAdapter
            emailmock[MockAdapter]
        end
    end
    
    subgraph pkg 
        kafka
        ses
        pq
        gocql
        sendgrid
    end
    
    brokerui --> service
    grpc --> service
    rest --> service
    
    brokerp -.-> KafkaAdapter
    brokerp -.-> brokermock
    emailp -.-> SESAdapter
    emailp -.-> SendgridAdapter
    emailp -.-> emailmock
    dbp -.-> PostgresAdapter
    dbp -.-> CassandraAdapter
    dbp -.-> dbmock
    
    KafkaAdapter --- kafka
    SESAdapter --- ses
    SendgridAdapter --- sendgrid
    PostgresAdapter --- pq
    CassandraAdapter --- gocql
```

Achieved by the folder structure:

```
.
├── app
│   └── service
|       ├── service.x
|       ├── ports.x
│       └── entities.x
├── infra
│   ├── broker
|   |   ├── mock
|   |   |   └── adapter.x
|   |   └── kafka
|   |       └── adapter.x
│   ├── db
|   |   ├── mock
|   |   |   └── adapter.x
|   |   ├── postgres
|   |   |   └── adapter.x
|   |   └── cassandra
|   |       └── adapter.x
│   └── email
|       ├── mock
|       |   └── adapter.x
|       ├── ses
|       |   └── adapter.x
|       └── sendgrid
|           └── adapter.x
└── ui
    ├── cli
    ├── broker
    ├── grpc
    └── rest
```
 

