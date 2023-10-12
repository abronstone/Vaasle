# vaas-final
Carleton College CS348: Final Project. Team Vaas

## Running the CLI
To start the containers and run the CLI, run the following command from the root directory of the project (Unix):
```
./start.sh
```
Or on Windows:
```
bash ./start.sh
```
If you are running this script for the first time, you may need to give it permission to execute. To do this, run the following command:
```
chmod u+x start.sh
```
If you would like to access the DB manually, and not be able to access the CLI, please run 
```
docker-compose up --build
``` 
Note: Aaron will need to add your public IP address to the list of allowed IP addresses on the Mongo cluster