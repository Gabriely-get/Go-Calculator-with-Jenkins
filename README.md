# Calculator with Packer

## Required
- Install and configuring Docker
- Install and configuring Packer
- Download this project

## Running
Open the project folder on terminal

### Running with microservice.sh file
- For start: type: *. ./scripts/microservice.sh && start*
- For stop: type: *. ./scripts/microservice.sh && stop*
- For verify status, type: *. ./scripts/microservice.sh && status*

### Running without microservice.sh file
- For create the image, type: *packer build .*
- For run the container, type: *sudo docker run --rm --name=go-calculator --network=host -p 8000:8000 gabrielys/go-calculator* 
    - For stop, type: *Ctrl+c* or *docker stop go-calculator* 

## Testing

- Type: *./scripts/smoke-test-calculator.sh* to execute the test

## Endpoints
Endpoints and examples of appropriate response

- Addition: http://localhost:8000/calc/sum/{value1}/{value2}
   ``` json
     {
        "result": "10"
     }
     ``` 
- Subtraction: http://localhost:8000/calc/sub/{value1}/{value2}
  ``` json
     {
        "result": "5"
     }
     ``` 
- Division: http://localhost:8000/calc/div/{value1}/{value2}
  ``` json
     {
        "result": "7"
     }
     ``` 
- Multiplication: http://localhost:8000/calc/mult/{value1}/{value2}
  ``` json
     {
        "result": "30"
     }
     ``` 
- History: http://localhost:8000/calc/history
  ``` json
    [
        {
            "time": "2022-08-19T02:13:52.853500749-03:00",
            "result": "6.00 - 6.00 = 0.00",
            "operation": "SUBTRACTION"
        },
        {
            "time": "2022-08-19T02:13:52.859714718-03:00",
            "result": "6.00 + 6.00 = 12.00",
            "operation": "ADDITION"
        }
    ]
     ``` 
