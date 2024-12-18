# Fetch Assessment: Receipt Processor

This assessment is for the company Fetch. This assessment required me to create a webservice that fulfils the specified documented API listed:
- POST: Submit a receipt to process and return the ID assigned to the receipt
- GET: Return the points awarded for the specified receipt

The rules to calculate the points for each receipt is listed below:
- One point for every alphanumeric character in the retailer name.
- 50 points if the total is a round dollar amount with no cents.
- 25 points if the total is a multiple of 0.25.
- 5 points for every two items on the receipt. 
- If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
- 6 points if the day in the purchase date is odd.
- 10 points if the time of purchase is after 2:00pm and before 4:00pm.


## How to Run with Docker

To run this webservice, we need to deploy Docker. Run the command `docker-compose up`. Once this is called, the program should be running. 

## How to Run without Docker

With the project directory in the terminal, type in `go run main.go`. This would run the code on the main file. Once this is running you can start using the routes.

## Example Test Case to Input

To start, you would need a program that can send request. I used Postman to send them. 
To send the request you would add a request, and label it as a POST route. Make sure that in the headers the `Content-Type` is of application/json. 
Then, for the route itself, you would use `localhost:8080/receipts/process`
As for the body, you can refer to the test case below:

Input example test case:
```json
{
    "retailer": "Costco",
    "purchaseDate": "2024-01-01",
    "purchaseTime": "15:24",
    "total": "39.47",
    "items": [
        {
            "shortDescription": "Northwest Alaskan Salmon Fillets 10 LB",
            "price": "20.99"
        },{
            "shortDescription": "Downy Soft Liquid Fabric Softener 150 FL OZ",
            "price": "16.99"
        },{
            "shortDescription": "Reese's Peanut Butter Cups",
            "price": "1.49"
        }
    ]
}
```
After sending the input, the response should return an ID in this JSON format:
```json
{
    "id": "8f1ab475-1139-40be-a044-85ed1f3fc421"
}
```

This would not be the exact ID, but it will be of a similar structure.

You would take this ID and place it into new GET route for the points. To do so, you add another request and label it as a GET route. 
Then you would input the route as so: `receipts/8f1ab475-1139-40be-a044-85ed1f3fc421/points`

After sending the GET request, you should get a response like this:
```json
{
    "points": 27
}
```
