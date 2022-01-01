# Portless Deploy
A deployment service that will run on a local network that has all incoming ports blocked.

The service will check a mongo db in the cloud every 5 mins for any deployments. Once a deployment is found a scrpit will run which will
- Clone the repo
- Apply any custom commands
- Deploy the service
