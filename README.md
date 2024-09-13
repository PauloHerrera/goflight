# Go Flights

This service helps individuals who have recently lost a relative find the next available flights from their local airport at special discounted rates.

### How it works:

**Input**: Simply enter the departure airport and desired travel dates.

**Search**: Our system will search for the most suitable flights based on search criteria.

**Discounts**: We'll provide a list of flights that include exclusive discounts.

**Choose**: Users can submit up to three outbound and three return flights (if needed), ranking them according to their preference.

## Installation

#### Prerequisites:

- [Go 1.23](https://go.dev/doc/install#releases);
- [Docker](https://docs.docker.com/engine/install/) (optional but advisable);

#### Add your env vars:

1. Get your access key from [Serp API](https://serpapi.com/). This project consumes their [Google Flights integration](https://serpapi.com/google-flights-api).
2. Create your `.env` file following the `.env.example`.
   - Add your Google Flight integration key;
   - Your local server address (with the port number). Ex: `localhost:3000`;
   - Your database uri following the pattern `mongodb://user:password@host:port/`;
   - The MongoDB credentials from your preference. You can repeat the user and password above for local tests.

You don't need to configure a MongoDB instance. The application is ready to delivery it for you through `docker-compose.yml`. It will generate a MongoDB container and configure the database, including authentication using the information from your `.env` file.

## Running the project

The easiest way to run the application:

```
make run
```

It will build and run the MongoDB and application server containers.

If you don't want to use containers or already have your own MongoDB instance, you can just run the go application. You must configure your MongoDB instance and add your own string connection uri at `.env`. To run only the go API:

```
make server
```

## Application routes

### GET `\fligths`:

Returns the a list of up to five outbound and five return flights with the calculated discount.

```
https://goflight.wow/?departure_airport=CGB&departure_date=2024-01-13&return_airport=GRU&return_date=2024-01-21&flight_type=1
```

**Params**:

**departure_airport**: Origin airport code, with 3 characters. Ex: `AAA`

**departure_date**: The outbound flight date with format `yyyy-mm-dd`.

**flight_type**: 1 - round trip; 2 - one way flight.

**return_airport**: The destination airport code, with 3 characters. Also the airports for the return trip.

**return_date**: Return date with format `yyyy-mm-dd`.

Important:

- It allows multiple airport from departure and return. Submit them with a comma separator. Ex: `AAA,BBB,CCC`.
- It accepts only airports based on Brazil. There is a json resource containing the valid airports at `public_br_airports.json`. You can also check them [here](https://data.opendatasoft.com/explore/dataset/airports-code%40public/table/?refine.country_name=Brazil).

### PUT `\fligths`:

Submit the user flight preferences

**Params**:

**user_id**: Identify the user, passed as query string. Should be replaced by authentication in the future. Ex: `\fligths?user_id:123`.

**order**: Ranks user preferences. Valid values: 1, 2, 3.

**flight info**: Repeats the flight info returned on get request. You can check an exemple at `post_flights_sample.json`.

Important:

- It saves as much information as possible to help the operator to schedule the right flight, because the trip itself does not have an ID;
- Also, the decision of saving nothing (for now) on the search call envolves 2 key points:
  - Return data as fast as possible;
  - Be as cheap as possible saving resources such as async comns or a db cache layer such as Redis;

## Future contributions

- Create an API to list brazilian airports and allow a search by city. Could be done in another microservice if it makes sense;
- Consider flight transfers in the flight details information;
- Increase the unit tests coverage;
- Persists a user session through authentication;
