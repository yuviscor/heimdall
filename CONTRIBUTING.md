# Contributing to Heimdall 🤝

Thank you for your interest in contributing to Heimdall! This document provides guidelines and information for contributors.

## 📜 Code of Conduct

By participating in this project, you agree to abide by our Code of Conduct. Please be respectful and inclusive in all interactions.

## 🚀 How Can I Contribute?

### Types of Contributions

- **🐛 Bug Reports**: Report bugs and issues
- **✨ Feature Requests**: Suggest new features
- **📝 Documentation**: Improve or add documentation
- **🔧 Code Contributions**: Submit code improvements
- **🧪 Testing**: Add tests or improve test coverage

### Before You Start

1. Check existing issues to avoid duplicates
2. Read the project documentation
3. Familiarize yourself with the codebase structure

## 🛠️ Development Setup

### Prerequisites

- Go 1.24.2 or higher
- Git
- Task (optional, for build automation)

### Local Development

1. **Fork and Clone**
   ```bash
   git clone https://github.com/YOUR_USERNAME/heimdall.git
   cd heimdall
   ```

2. **Add Upstream Remote**
   ```bash
   git remote add upstream https://github.com/MowlCoder/heimdall.git
   ```

3. **Create a Feature Branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

4. **Install Task (Optional)**
   ```bash
   # macOS
   brew install go-task/tap/go-task
   
   # Linux
   sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b ~/.local/bin
   
   # Windows
   scoop install task
   ```

## 📝 Code Style Guidelines

### Go Code Style

- Follow [Effective Go](https://golang.org/doc/effective_go.html) guidelines
- Use `gofmt` for code formatting
- Run `go vet` to check for common mistakes
- Keep functions small and focused
- Use meaningful variable and function names
- Add comments for exported functions and complex logic

### Testing

- Write tests for new functionality
- Maintain good test coverage
- Use descriptive test names
- Test both success and failure cases

## 🔄 Pull Request Process

### Before Submitting

1. **Update Documentation**
   - Update README.md if needed
   - Update configuration examples if applicable

2. **Test Your Changes**
   ```bash
   # Run all tests
   go test ./...
   
   # Test the built application
   task run
   ```

3. **Code Quality Checks**
   ```bash
   # Format code
   go fmt ./...
   
   # Run linter
   go vet ./...
   
   # Check for unused imports
   go mod tidy
   ```

### Pull Request Guidelines

1. **Create a Descriptive Title**
   - Use present tense ("Add feature" not "Added feature")
   - Be specific and concise

2. **Write a Detailed Description**
   - Explain what the PR does
   - Link to related issues
   - Include screenshots if applicable
   - List any breaking changes

3. **Example PR Description**
   ```markdown
   ## Description
   Adds support for custom HTTP headers in service health checks.
   
   ## Changes
   - Added `headers` field to service configuration
   - Updated health checker to include custom headers
   - Added configuration examples
   - Updated documentation
   
   ## Testing
   - Added unit tests for header functionality
   - Tested with various header configurations
   - Verified backward compatibility
   
   ## Related Issues
   Closes #123
   ```

### Review Process

1. **Automated Checks**
   - All tests must pass
   - Code must be properly formatted
   - No linting errors

2. **Code Review**
   - At least one maintainer must approve
   - Address all review comments
   - Update PR based on feedback

3. **Merge Requirements**
   - All checks pass
   - Approved by maintainers
   - No merge conflicts

## 🐛 Reporting Bugs

### Before Reporting

1. Check existing issues for duplicates
2. Try to reproduce the issue
3. Check the documentation

### Bug Report Template

```markdown
## Bug Description
Brief description of the issue

## Steps to Reproduce
1. Step 1
2. Step 2
3. Step 3

## Expected Behavior
What should happen

## Actual Behavior
What actually happens

## Environment
- OS: [e.g., Ubuntu 20.04]
- Go Version: [e.g., 1.24.2]
- Heimdall Version: [e.g., commit hash]

## Configuration
```json
{
  "services": [...],
  "telegram": {...}
}
```

## Logs
```
[ERROR]: Service is not healthy...
```

## 📄 License

By contributing to Heimdall, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to Heimdall! 🚀 