# vaas-final
Carleton College CS348: Final Project. Team Vaas

## Running the CLI
To run the CLI, run the following command from the root directory of the project:
```
./start.sh
```
Note: if you executing the CLI for for first time, you may need to run the following command to make the start script executable first:

```
chmod +x start.sh
```
If you would like to access the DB manually, and not be able to access the CLI, please run 
```
docker-compose up --build
``` 
Note: Aaron will need to add your public IP address to the list of allowed IP addresses on the Mongo cluster