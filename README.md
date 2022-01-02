# Portless Deploy
A deployment service that will run on a local network that has all incoming ports blocked.

The service will check a mongo db in the cloud every 5 mins for any deployments. Once a deployment is found a scrpit will run which will
- Clone the repo
- Apply any custom commands
- Deploy the service


## Setup
This is built to be a background process run every 5 mins


## Environment Variables

Arg 1: MONGO_CONNECTION_STRING: Connection String to Database

Arg 2: MONGO_DATABASE: Name of Database

Arg 3: MONGO_COLLECTION: Name of collection

## Example of deployment scripts
https://github.com/horvatic/zracni-udar-service/tree/main/deploy/scripts
