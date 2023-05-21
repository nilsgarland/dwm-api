# dwm-api

## Getting Started

### Prerequisites

You need to have Go (version 1.17 or later) installed on your machine. Check your Go installation by typing the following in your terminal:

    go version

You should see the Go version in the output.

Additionally, you need SQLite3 installed to create and manage the database.

### Installing

1. Clone this repository to your local machine:

    git clone https://github.com/nilsgarland/dwm-api.git

2. Navigate to the project directory:

    cd dwm-api

3. Download the necessary Go packages:

    go get -u github.com/gorilla/mux
    go get github.com/mattn/go-sqlite3

4. Create the SQLite database:

    sqlite3 ./database/database.db

Inside the SQLite shell, create the meetings table:
sql
    CREATE TABLE meetings (
        id INTEGER PRIMARY KEY,
        day TEXT NOT NULL,
        name TEXT NOT NULL,
        location TEXT NOT NULL,
        description TEXT NOT NULL,
        picture INTEGER NOT NULL,
        time TEXT NOT NULL
    );

Exit the SQLite shell with `.exit`.

### Running the Application

In the project directory, run the Go script:

    go run main.go

This will start the server on `localhost:8000`.

## API Usage

The API provides the following endpoints:

### `GET /meetings`

This endpoint retrieves all meetings from the database.

### `POST /meetings`

This endpoint creates a new meeting. The request body should be a JSON object with the following structure:

```json
{
    "day": "Monday",
    "name": "Meeting Name",
    "location": "Meeting Location",
    "description": "Meeting Description",
    "picture": 1,
    "time": "2023-05-27"
}
