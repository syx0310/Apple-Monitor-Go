# Apple-Monitor-Go

`Apple-Monitor-Go` is a monitoring tool designed to track Apple device availability in stores. It allows users to configure specific devices and locations to monitor, and sends notifications when a device becomes available using Bark and WeChat Work (WeCom).

## Features

- Monitor availability of Apple devices in specific stores.
- Filter monitored stores by keywords (whitelist feature).
- Supports multiple notification methods:
  - **Bark** (mobile notifications).
  - **WeCom** (corporate notifications).
- Configurable query parameters and cron jobs for periodic monitoring.

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Notification Setup](#notification-setup)
- [Development](#development)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. **Download the binary**  
   The binary is automatically released through GoReleaser and can be downloaded from the [Releases](https://github.com/syx0310/Apple-Monitor-Go/releases) section.

2. **Build from source**  
   Alternatively, you can build the binary from source:

   ```bash
   git clone https://github.com/your-repo/Apple-Monitor-Go.git
   cd Apple-Monitor-Go
   go build -o apple-monitor-go .
   ```

## Configuration

1. **Create the configuration file**  
   Copy the example configuration file `config.yaml.example` to `config.yaml`.

   ```bash
   cp config.yaml.example config.yaml
   ```

2. **Modify the configuration**  
   Customize the `config.yaml` file with your desired Apple devices, stores, and notification settings.

   Example configuration can be found in `config.yaml.example`

## Usage

Run the application by executing the binary:

```bash
./apple-monitor-go run
```

The application will start monitoring the configured devices and send notifications when stock is available in the specified stores.

## Notification Setup

### Bark Setup

1. Download the [Bark App](https://apps.apple.com/cn/app/bark/id1403753865).
2. Configure the Bark notification settings in `config.yaml` by adding your `bark_key` and `bark_api_url`.

### WeChat Work (WeCom) Setup

1. Set up a bot in your WeChat Work (WeCom) application.
2. Add the generated webhook URL to the `wecom_url` field in your `config.yaml`.

## Development

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-repo/Apple-Monitor-Go.git
   ```

2. **Run the application**:

   ```bash
   go run main.go
   ```

3. **Test changes**:

   The application includes cron job scheduling for periodic monitoring and checks for device availability.

## Contributing

Contributions are welcome! If you find a bug or want to add a feature, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
