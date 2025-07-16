# SignalSpectre: Real-time Cryptocurrency Anomaly Detection

SignalSpectre is a Go-based system designed for the real-time detection of anomalous price fluctuations in cryptocurrency markets. It leverages spectral analysis techniques, coupled with Kalman filter-powered outlier prediction, to identify unusual price movements that deviate significantly from expected trends. When anomalous behavior is detected, SignalSpectre triggers real-time alerts delivered via WebSockets, allowing for immediate response to potentially profitable or detrimental market shifts.

The core purpose of SignalSpectre is to provide a robust and reliable early warning system for cryptocurrency traders and analysts. By employing advanced signal processing methods, the system filters out normal market volatility and pinpoints genuine anomalies that warrant attention. This reduces the noise associated with typical price fluctuations and provides a clearer picture of potentially significant market events, such as flash crashes, pump-and-dump schemes, or unexpected regulatory announcements. The real-time WebSocket-based alerts enable rapid reaction, facilitating informed decision-making and minimizing potential losses while maximizing potential gains.

SignalSpectre offers a significant advantage over traditional price monitoring systems by incorporating predictive analytics. The Kalman filter component learns the typical price behavior of a given cryptocurrency over time and then uses this knowledge to predict future price movements. By comparing actual prices to these predictions, the system can identify outliers that are statistically unlikely to have occurred under normal market conditions. This proactive approach allows users to anticipate and react to market anomalies before they become widely apparent, providing a critical edge in the fast-paced cryptocurrency market.

## Key Features

*   **Real-time Spectral Analysis:** Implements a Fast Fourier Transform (FFT) to decompose cryptocurrency price data into its constituent frequencies. Anomalous events are identified by detecting significant deviations in frequency amplitudes compared to historical averages.
*   **Kalman Filter-Powered Outlier Prediction:** Utilizes a Kalman filter to estimate the expected price based on historical data. The system identifies outliers by comparing the predicted price with the actual price, flagging discrepancies exceeding a predefined threshold.
*   **Configurable Anomaly Sensitivity:** Allows users to adjust the sensitivity of anomaly detection through configurable parameters such as the standard deviation multiplier for outlier detection and the frequency amplitude threshold for spectral analysis.
*   **WebSocket-Based Alerting:** Provides real-time alerts of detected anomalies via a WebSocket connection. Alert messages include details about the cryptocurrency pair, timestamp of the anomaly, and the severity of the anomaly. The alert format is JSON.
*   **Modular Architecture:** Designed with a modular architecture that facilitates easy integration with different cryptocurrency exchanges and data sources. New data sources can be implemented by creating a new interface implementation.
*   **Exchange Agnostic:** The system is built to be modular in order to support any exchange that provides historical candlestick data via API. Current implementation utilizes the Binance API.
*   **Persistence Layer:** Configurable persistence layer allows the storing of anomaly alerts for later investigation, using either an in-memory store or database.

## Technology Stack

*   **Go:** The primary programming language for the core logic of the application, chosen for its concurrency features and performance.
*   **Binance API:** Used for fetching real-time cryptocurrency price data. The API is utilized to retrieve historical candlestick data.
*   **Gorilla WebSocket:** Used for establishing and managing WebSocket connections for real-time alert delivery.
*   **Gonum:** Used for numerical computation, specifically for the Fast Fourier Transform (FFT) and Kalman filter implementations.
*   **JSON:** Used for formatting alert messages transmitted via WebSockets.
*   **viper:** Used for configuration management.

## Installation

1.  **Prerequisites:** Ensure you have Go (version 1.18 or higher) installed and configured on your system. Verify by running `go version`.
2.  **Clone the Repository:** Clone the SignalSpectre repository from GitHub using the command: `git clone https://github.com/uhsr/SignalSpectre.git`
3.  **Navigate to the Project Directory:** Change your current directory to the SignalSpectre directory using the command: `cd SignalSpectre`
4.  **Install Dependencies:** Download and install the required dependencies using the command: `go mod download`
5.  **Build the Application:** Build the executable using the command: `go build -o signalspectre`

## Configuration

SignalSpectre uses environment variables and a `config.yaml` file for configuration. The `config.yaml` should be placed in the same directory as the executable.

**Environment Variables:**

*   `BINANCE_API_KEY`: Your Binance API key.
*   `BINANCE_API_SECRET`: Your Binance API secret.
*   `WS_PORT`: Port for the WebSocket server. Default is 8080.

**config.yaml:**



The `fft_threshold` and `kalman_std_dev` values allow you to fine tune the sensitivity of the anomaly detection.

## Usage

1.  **Run the Application:** Execute the built binary using the command: `./signalspectre`
2.  **Connect to the WebSocket Endpoint:** Open a WebSocket connection to `ws://localhost:8080/ws` (or the address and port specified in your configuration).
3.  **Receive Real-time Alerts:** The system will send JSON-formatted alert messages over the WebSocket connection whenever an anomaly is detected.

Example Alert Message:

`{"timestamp":"2023-10-27T10:00:00Z","symbol":"BTCUSDT","anomaly_type":"FFT","severity":"High","message":"Significant frequency deviation detected."}`

## Contributing

We welcome contributions to SignalSpectre! Please follow these guidelines:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix: `git checkout -b feature/your-feature-name`
3.  Implement your changes and write appropriate tests.
4.  Ensure all tests pass: `go test ./...`
5.  Commit your changes with clear and concise commit messages.
6.  Push your branch to your forked repository: `git push origin feature/your-feature-name`
7.  Submit a pull request to the main repository.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/uhsr/SignalSpectre/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the developers of the Gonum library for providing the numerical computation tools used in this project.