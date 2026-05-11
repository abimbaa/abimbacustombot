# Changelog

All notable changes to Abimba Custom Bot will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-05-11

### Added
- Initial release of Abimba Custom Bot
- 40+ Telegram commands for PC control
- Screenshot and webcam capture functionality
- File upload/download operations
- System information and status commands
- Media playback control
- Process and application management
- Macro system for command sequences
- Notes and task planning features
- Speech-to-text audio transcription
- Clipboard operations
- Text input simulation
- Power management (shutdown, restart, lock, sleep)
- Shell command execution
- Volume and audio control
- Comprehensive help documentation
- Environment-based configuration
- Virtual environment setup for Python dependencies
- Setup and build automation scripts
- Windows startup integration scripts
- Comprehensive README with setup instructions
- Security policy and guidelines
- Contributing guidelines
- Project structure cleanup for production

### Features
- **Telegram Integration**: Full Telegram Bot API integration using telebot library
- **Command Registry**: Extensible command handler system for easy additions
- **Python Integration**: Native Python script execution for complex operations
- **Hidden Windows**: All processes run hidden by default
- **Command Timeout**: 5-minute timeout for command execution
- **Authorization**: Single-user access control via Telegram ID
- **File Handling**: Automatic downloads to Downloads folder, file uploads to Telegram
- **Audio Processing**: Real-time audio file transcription using Whisper
- **Macro System**: Save and execute sequences of commands
- **Data Persistence**: JSON-based storage for macros, notes, and plans

### Documentation
- Complete README with all commands documented
- Setup guide with prerequisites
- Windows startup integration guide with 3 methods
- Security policy and best practices
- Contributing guidelines
- Development documentation

### Infrastructure
- Go module configuration
- Python virtual environment setup
- Build script with error checking
- Setup automation script
- Multiple startup launcher options (PowerShell, Batch, VBScript)
- Comprehensive .gitignore

---

## Versioning

- **Major**: Breaking changes, major features
- **Minor**: New features, backward compatible
- **Patch**: Bug fixes, documentation updates

## Guidelines for Contributors

When adding new features:
1. Update CHANGELOG.md with your changes
2. Follow semantic versioning
3. Update documentation
4. Test thoroughly before committing

---

**Last Updated**: 2026-05-11
**Maintainer**: abimbaa
