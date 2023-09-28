# Template project of Clean-Architecture
This is a simple CRUD application build on a Clean-Architecture written in Golang.

## :triangular_flag_on_post: DEMO
### 1. use Docker to set up an environment to run the application
```bash
> make up 
```
### 2. Call API by curl command
#### create
```bash
> curl -i -XPOST -d '{"title": "test", "body": "this is a test article"}' localhost:8080/articles
```

#### read
```bash
> curl -i -XGET localhost:8080/articles
> curl -i -XGET localhost:8080/articles/1
```

#### update
```bash
> curl -i -XPUT -d '{"title": "updated title", "body": "this article is updated"}' localhost:8080/articles/1
```

#### delete
```bash
> curl -i -XDELETE localhost:8080/articles/1
```

## :book: Explanation


| layer | directory |
| ---- | ---- |
| Enterprise Business Rules | entities |
| Application Business Rules | usecase |
| Interface Adapters | adapter |
| Frameworks & Drivers | driver |

### Enterprise Business Rules
Core business rules are implemented in this layer.

### Application Business Rules
By describing the business logic API, this layer expresses what this software will achieve.

### Interface Adapters 
This layer describes the input(controller), output(presenter) and data presistence process(gateway).

### Frameworks & Drivers
This layer describes establishing DB connections, configuring routing and using the framework.