## Go and Python Test

Write a rest webservice in Go that allows to operate over text files as resources. The minimum requirements for the service is:

    Create a text file with some contents stored in a given path.

    Retrieve the contents of a text file under the given path.

    Replace the contents of a text file.

    Delete the resource that is stored under a given path.


We would also need to get some statistics per folder basis and retrieve them through another entry point. These statistics are:

    Total number of files in that folder.

    Average number of alphanumeric characters per text file (and standard deviation) in that folder.

    Average word length (and standard deviation) in that folder. 

    Total number of bytes stored in that folder.

    Note: All these computations must be calculated recursively from the provided path to the entry point.


Use all necessary libraries, including third-party libraries but make sure that they are easily fetched. Keep in mind that json is the transport format to be used.
# Tech used

- Python 3 (script)
- Golang 1.11 (api)
- Docker (container)

# Testing

For your convenience I have added:
 - unit tests in go
 - postman tests on each endpoint
 - dockerfile
 - the executable compiled for linux os ready to be tested
 
 # Run
 
 1 command line bash script have been added for your convenience
 that bind the port 12345 on the host 
 
 ` bash run.sh`
