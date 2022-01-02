# Portless Deploy
A deployment service that will run on a local network that has all incoming ports blocked.

The service will check a mongo db in the cloud every X mins for any deployments. Once a deployment is found a scrpit will run which will
- Clone the repo
- Apply any custom commands
- Deploy the service


## Arguments

Arg 1: DEPLOY_MONGO_CONNECTION_STRING: Connection String to Database

Arg 2: DEPLOY_MONGO_DATABASE: Name of Database

Arg 3: DEPLOY_MONGO_COLLECTION: Name of collection

## Running
 `$ ./deploy "DEPLOY_MONGO_CONNECTION_STRING" "DEPLOY_MONGO_DATABASE" "DEPLOY_MONGO_COLLECTION" "5" &`

`$ ps`

`$ disown PROCESSID`

## Stopping
`$ ps -A | grep "deploy"`

`$ kill -9 PROCESSID`

## Example of deployment scripts
https://github.com/horvatic/zracni-udar-service/tree/main/deploy/scripts

https://github.com/horvatic/zracni-udar-ui/tree/main/deploy/scripts
