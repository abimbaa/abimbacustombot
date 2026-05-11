# Contributing to Abimba Custom Bot

Thank you for your interest in contributing! This guide will help you get started.

## Code of Conduct

Please be respectful and constructive in all interactions.

## How to Contribute

### Reporting Bugs

1. Check if the bug has already been reported in Issues
2. Create a new issue with:
   - Clear title describing the bug
   - Detailed description of the problem
   - Steps to reproduce
   - Expected vs actual behavior
   - Your environment (Windows version, Go version, Python version)
   - Error messages or logs

### Suggesting Features

1. Check existing issues/discussions first
2. Create an issue describing:
   - The feature you want
   - Why it would be useful
   - Example use cases
   - Possible implementation approach (optional)

### Code Contributions

1. **Fork the repository**
   ```powershell
   git clone https://github.com/YOUR_USERNAME/abimbaCustomBot.git
   cd abimbaCustomBot
   ```

2. **Create a feature branch**
   ```powershell
   git checkout -b feature/your-feature-name
   ```

3. **Make your changes**
   - Follow the existing code style
   - Add comments for complex logic
   - For Go: use `gofmt` for formatting
   - For Python: use `black` for formatting (install: `pip install black`)

4. **Test your changes**
   ```powershell
   .\build.ps1
   .\dispatcher.exe
   ```
   Test the bot thoroughly with your changes

5. **Commit with clear messages**
   ```powershell
   git add .
   git commit -m "feat: description of your changes"
   ```
   
   Use conventional commits:
   - `feat:` for new features
   - `fix:` for bug fixes
   - `docs:` for documentation
   - `refactor:` for code improvements
   - `test:` for tests
   - `chore:` for maintenance

6. **Push to your fork**
   ```powershell
   git push origin feature/your-feature-name
   ```

7. **Create a Pull Request**
   - Title: Clear description of changes
   - Description: Explain what you changed and why
   - Reference related issues with `#123`
   - Include any breaking changes

## Development Guidelines

### Adding New Commands

To add a new command to the bot:

1. **Edit `cmd/executor/main.go`:**
   ```go
   var registry map[string]CommandFunc
   func init() {
       registry = map[string]CommandFunc{
           "/mycommand": handleMyCommand,
           // ...
       }
   }
   
   func handleMyCommand(args string) {
       // Your implementation
       fmt.Println("Output goes here")
   }
   ```

2. **Update help text in `handleHelp()`:**
   ```go
   "/mycommand": {"Short description", "Long description with usage"},
   ```

3. **For Python-based commands:**
   - Create a new script in `scripts/`
   - Add to `requirements.txt` if needed
   - Call from Go using `runProcess(VENV_PYTHON, "./scripts/myscript.py", args)`

4. **Build and test:**
   ```powershell
   .\build.ps1
   .\dispatcher.exe
   ```

### Performance Considerations

- Commands have a 5-minute timeout
- Long operations should show progress
- Avoid blocking operations in the main loop
- Use goroutines for concurrent operations if needed

### Security

- Never log sensitive information (tokens, IDs, file contents)
- Validate user input before executing
- Sanitize file paths
- Use subprocess isolation
- Review Python code for security issues

## Building & Testing

### Clean build with testing:

```powershell
.\build.ps1 -Clean -Rebuild

# Test the bot
.\dispatcher.exe
```

### Check for Python syntax errors:

```powershell
python -m py_compile scripts/myscript.py
```

### Run linters (optional):

```powershell
# Go
go vet ./...

# Python
pip install pylint
pylint scripts/
```

## Documentation

- Update README.md for user-facing changes
- Add inline comments for complex code
- Document new features with examples
- Update the command list in `/help`

## Pull Request Process

1. Ensure your code builds: `.\build.ps1`
2. Test your changes thoroughly
3. Update documentation if needed
4. Keep commits organized and meaningful
5. Respond to review feedback
6. Your PR will be merged once approved

## Questions?

- Check existing issues and discussions
- Create a new discussion for questions
- Ask in PR reviews

## License

By contributing, you agree that your contributions will be licensed under the same license as the project.

---

**Thank you for contributing to make Abimba Custom Bot better! 🎉**
