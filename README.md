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

- cover.prof: raw coverage data

- HTML report: opens automatically (uses go tool cover -html=coverage.out)

### Git Hooks

This project uses custom Git hooks stored in the .githooks directory to ensure code quality before commits or pushes.

To enable these hooks locally, run the following command once:

```bash
git config core.hooksPath .githooks
```

This tells Git to use the .githooks/ folder instead of the default .git/hooks/.

Now any tracked hooks (like pre-push) will automatically run during Git operations.
