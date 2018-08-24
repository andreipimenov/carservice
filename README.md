# Test task

Architechture pattern explanation:
- There are 3 layers used:
    1. Model layer contains business specific entities with interfaces for storing and interacting with them.
    2. Interact and Storage layers are representing Interface layer. They depends on model layer. Here implemented business logic rules.
    3. Infrastructure layer represented by driver package. Here implemented functions for interacting with real DB. In my case - MongoDB.
- Layers could being tested separately and independently from DB or other frameworks and drivers.
- Each layer depends on top layers only (dependency rule).

Done:
- Car model with storage and interactor. It is possible to easily change DB without chaning business-logic code base.
- Server with API:
    1. Configuration by config.json with option to specify config file name by -config flag.
    1. /api/ping - for health checking.
    2. GET /api/cars?serialNumber={uint64} - for getting car by its serial number from the storage.
    3. For all unsupported API endpoints - not found error with user-friendly json-response.

Example of using:
0. Check that you have mongo db instance is running on addr, which defined in config.json
1. Build executable file and run with default config.json.
```
go build -o server
./server
```
2. Get car by serial number 1234567890
```
curl -X GET 127.0.0.1:8080/api/cars?serialNumber=1234567890
```


TODO: 
- Write tests.
- Add Dockerfile and docker-compose.