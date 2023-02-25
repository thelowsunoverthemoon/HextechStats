<h1 align="center"> Hextech Stats </h1>
<p align="center">profile management app for league player performance graphs</p>

## features

* profile management app to visualize/compare LoL player statistics with graphs
* includes Go REST API queried by the client that implements CRUD operations using a SQLite database
* queries the Riot REST API for data, and uses concurrency via goroutines to analyze multiple JSON data at once
* client that displays multiple graphs using Recharts with responsive UI, dropdowns, and buttons
* uses :
  - **gin** for REST API web framework
  - **react.js** for client
  - **recharts** for visualization
  - **axios** for http requests

## screen captures

<img src="https://i.imgur.com/qQxOCBw.gif">
<img src="https://i.imgur.com/cconGbT.gif">

## run

To run client side, use

```Batch
npm start
```

in the client folder. Make sure to change the "host" variable in Operations.js to the correct server path. To build and start REST API, use

```Batch
go build
htstats
```
in the root folder. Make sure to change the DATA_BASE path to the correct SQLite database path.
