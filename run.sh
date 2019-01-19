#!/bin/bash
sudo docker build -t emotech-sample -f Dockerfile .
sudo  docker run -p 12345:5000 emotech-sample