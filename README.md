# Calculator with Jenkins

## Required
- Install and configure Docker
- Create a Docker Hub account: https://hub.docker.com/signup

## Configuring

### Jenkins

1. Pull Jenkins docker image:
   1.     docker pull jenkins/jenkins:lts
2. Create a Jenkins volume:
   1.     docker volume create jenkins_home
3. Run the Jenkins container:
   1.     docker run --name jenkins -p 8282:8080 -p 50000:50000 -v jenkins_home:/var/jenkins_home -v /var/run/docker.sock:/var/run/docker.sock -v $(which docker):$(which docker) jenkins/jenkins:lts
4. Give permission so that the container terminal can run docker commands
   1. Access the container terminal
      1.     docker exec -u root -it jenkins bash
   2. Give permission
      1.     chmod 666 /var/run/docker.sock 
5. Access Jenkins: http://localhost:8282/
   ###### The password will be shown in the terminal
6. Install suggested plugins -> Skip and continue as admin -> Verify the URL, finish and start
7. Click on *admin* in the top of the page -> *Configure* -> Change the password -> Go back to *Dashboard*

### Jenkins Credentials and Plugins
1. Create credentials: Go to `Manage Jenkins` -> `Manage Credentials` -> `Stores scoped to Jenkins` -> Jenkins -> Global credentials -> *Add Credentials*
   1. Select kind *Username with password* and fill with you Docker Hub credentials (username and password)
   2. Click *Add Credentials* again. Select kind *Secret text* and write your Docker Hub repository name
2. Install Packer plugin: Go to `Manage Jenkins` -> `Manage Plugins` -> `Available` -> Search and Install *Packer*
3. Configuring **Packer**:
      1. Go to `Manage Jenkins` -> `Global Tool Configuration` -> Packer -> `Add Packer`
      2. Name: *Packer_Go-Calculator_linux_amd64* -> *Install automatically* -> Choose a version compatible with your S.O. _E.g.: linux (amd64)_ -> Save 


## Running

### Creating jobs
1. On *Dashboard* click in `New Item`
2. Set the name, choose `Pipeline` and click `OK`
   >Choose appropriated names: *BAKE* for job 01 and *LAUNCH* for job 02.
3. Go to `Pipeline` -> `Definition` -> *Pipeline script from SCM* -> `SCM` -> Git -> Paste https://github.com/Gabriely-get/Go-Calculator.git
4. Change the branch **/master* to **/main*
5. Choose the Jenkinsfile path. The steps to create the 02 jobs are the same, the only thing that will change is the path:
   1. ./job1/Jenkinsfile
   2. ./job2/Jenkinsfile
6. Run both jobs
   - ps: Launch job only stops if you abort it manually

## Testing

- Open the project folder on terminal and type: *./scripts/smoke-test-calculator.sh* to execute the test

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
