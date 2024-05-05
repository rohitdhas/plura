# 💡 Plura

Plura is a powerful HTTP benchmarking tool built with Go and Cobra, designed to benchmark API endpoints.

## 📦 Installation

Ensure you have Go installed on your machine. Then, you can install the HTTP Benchmarking Tool by running:

```bash
go install github.com/rohitdhas/plura@latest
```

## 👨🏻‍💻 Usage

### Benchmarking with CLI Parameters

You can run benchmarks directly from the command line by providing the URL of the endpoint, the number of requests, and the concurrency level.

```bash
plura exec --url <URL> --requests <numRequests> --concurrency <concurrencyLevel>
```

Example:

```bash
plura exec --url http://example.com/api --requests 1000 --concurrency 10
```

## Benchmarking with YAML File (WIP)

Alternatively, you can specify the benchmark parameters in a YAML file and pass it to the tool.

```bash
plura exec --file <path-to-yaml-file>
```

The YAML file should have the following structure:

```bash
url: http://example.com/api
requests: 1000
concurrency: 10
```

Example:

```bash
plura exec --file benchmark.yaml
```

## Command-line Options

**`exec` Command**

- `--url or -u`: URL of the API endpoint to benchmark.
- `--requests or -r`: Number of requests to make.
- `--concurrency or -c`: Concurrency level (number of concurrent requests).
- `--file or -f`: Path to the YAML file containing benchmark parameters.

## Contributing

Contributions are welcome! Please fork the repository, make your changes, and submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
