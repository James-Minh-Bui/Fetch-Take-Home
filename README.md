# FETCH-TAKE-HOME

This project is a simple web service built with **Go** and **Gin** that allows users to process receipts, generate a unique ID for each receipt, and calculate points based on various rules. The service includes both a web interface and an API that can be accessed via **Postman** or any HTTP client.

## Features

- **Process Receipt**: Submit a receipt in JSON format and get a unique ID for the receipt.
- **Get Points**: Fetch the points for a specific receipt using its ID.
- **Web Interface**: A simple front-end where users can test the web service by submitting JSON data and receiving responses.

## API Endpoints

### `POST /receipts/process`
- **Description**: Submits a receipt to the system and generates a unique ID for it.
- **Request body**:
    ```json
    {
      "retailer": "Target",
      "purchaseDate": "2022-01-01",
      "purchaseTime": "13:01",
      "items": [
        {
          "shortDescription": "Mountain Dew 12PK",
          "price": "6.49"
        },
        {
          "shortDescription": "Emils Cheese Pizza",
          "price": "12.25"
        }
      ],
      "total": "35.35"
    }
    ```
- **Response**:
    ```json
    {
      "id": "7fb1377b-b223-49d9-a31a-5a02701dd310"
    }
    ```

### `GET /receipts/:id/points`
- **Description**: Retrieves the points for a specific receipt identified by its ID.
- **Example Request**: `GET /receipts/7fb1377b-b223-49d9-a31a-5a02701dd310/points`
- **Response**:
    ```json
    {
      "points": 28
    }
    ```

## Rules for Points Calculation

1. **Retailer Name**: Points are awarded based on the number of alphanumeric characters in the retailer's name.
2. **Total Price**: 
    - Bonus points if the total ends in `.00`.
    - Additional points if the total is a multiple of `0.25`.
3. **Items**:
    - Points for every two items on the receipt (5 points for each pair).
    - For items whose `shortDescription` length is a multiple of 3, points are awarded based on the price.
4. **Purchase Date**: If the purchase day is odd, 6 points are awarded.
5. **Purchase Time**: If the purchase time is between 2 PM and 4 PM, 10 points are awarded.

## Running the Application

1. Clone the repository:
    ```bash
    git clone git@github.com:<username>/FETCH-TAKE-HOME.git
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Run the application:
    ```bash
    go run main.go
    ```

4. Access the web interface at `http://localhost:8080/` or use Postman to interact with the API.

## Testing with Postman

You can test the following endpoints using Postman:

### POST Request to `/receipts/process`
- **Body**: Provide the receipt JSON as the body of the request.
- **Response**: Get a generated `id` for the receipt.

### GET Request to `/receipts/:id/points`
- **URL Params**: Provide the receipt `id` to retrieve the points.
- **Response**: Get the calculated points for the receipt.

## Front-End

The front-end is a simple HTML form that allows users to submit JSON data and view the response. You can use the interface to:
1. Submit a receipt and get an ID.
2. Retrieve points for a receipt using the generated ID.
