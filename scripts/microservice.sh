#!/bin/bash

function start {
    if [ "$(sudo docker images -q gabrielys/go-calculator:latest)" == "" ]; 
    then
        sudo packer build .
        sudo docker run --rm --name=go-calculator --network=host -p 8000:8000 gabrielys/go-calculator
    elif ["$(sudo docker ps -q -f name=go-calculator)" == ""];
    then
        sudo docker run --rm --name=go-calculator --network=host -p 8000:8000 gabrielys/go-calculator
    else
        echo "Container is running"
    fi
}

function stop {
    if [ "$(sudo docker ps -q -f name=go-calculator)" == "" ]; 
    then
        echo "Can not stop because container is not running"
    else
        sudo docker stop go-calculator
    fi
}

function status {
    if [ "$(sudo docker ps -q -f name=go-calculator)" == "" ]; 
    then
        echo "NOT RUNNING"
    else
        echo "RUNNING"
    fi
}