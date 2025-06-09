# comida.app - backend

### Coverage

To generate and view a full HTML coverage report:

1. Install make

2. Run the coverage command
   In the project root:

```bash
make coverage
```

This will:

- Run tests across all packages

- Generate a coverage.out file

- Open an HTML report in your browser showing which lines are covered

#### 📁 Output

- coverage.out: raw coverage data

- HTML report: opens automatically (uses go tool cover -html=coverage.out)
