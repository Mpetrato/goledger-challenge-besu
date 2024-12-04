# GoLedger Challenge - Besu Edition

## Run Project

### First set up the environment

To set up the environment, you need to clone this repository.
To set up the environment, you need to run the following commands:

```bash
cd besu
./startDev.sh
```

Get the deploy address on **"SimpleStorageModule"** after startDev finish, and set the value to **/api/docker-compose.yml** enviroment **"CONTRACT_ADDRESS"**.

Get the **privateKey** on **/besu/genesis/genesis.json** and set the value on **/api/docker-compose.yml** enviroment **"PRIVATE_KEY"**.

### To run the project use the following command:

```bash
cd api
docker-compose up
```

## Endpoints
1. **SET:**
	- method: "POST"
	- Endpoint: "localhost:3000/api/v1/contract"

2. **GET:**
	- method: "GET"
	- Endpoint: "localhost:3000/api/v1/contract"

3. **SYNC:**
	- method: "POST"
	- Endpoint: "localhost:3000/api/v1/contract/sync"

4. **CHECK:**
	- method: "GET"
	- Endpoint: "localhost:3000/api/v1/contract/check"

