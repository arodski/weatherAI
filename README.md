# weatherAI
No AI, just marketing magic.

## How to Run the Program

### Prerequisites
- Go 1.23.1 or later
- An API key from OpenWeatherMap

### Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/arodski/weatherAI.git
    cd weatherAI
    ```

2. Create a .env file in the root directory with the following content:
    ```env
    OPENWEATHER_URL=https://api.openweathermap.org/data/2.5/weather
    OPENWEATHER_API_KEY=your_openweather_api_key
    ADDR=:8080
    ```

3. Install the dependencies:
    ```sh
    go mod tidy
    ```

### Running the Server

1. Navigate to the api directory:
    ```sh
    cd cmd/api
    ```

2. Run the server:
    ```sh
    go run *.go
    ```

3. The server will start on port `8080`. You can access the weather endpoint at:
    ```
    http://localhost:8080/v1/weather?lat={latitude}&lon={longitude}
    ```

### Example Request

To get the weather for San Francisco (latitude: 37.7749, longitude: -122.4194), you can use the following command in the terminal:
```
curl -X GET "http://localhost:8080/v1/weather?lat=37.7749&lon=-122.4194"
```