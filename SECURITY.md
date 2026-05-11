# Security Policy

## Security Considerations

This bot gives **complete remote control** over your Windows PC. Treat your credentials with extreme care.

## Sensitive Information

### ⚠️ Critical Security Items

1. **BOT_TOKEN**
   - **NEVER commit to git** (`.env` is in `.gitignore`)
   - **NEVER share publicly**
   - Treat like a password
   - If exposed, revoke immediately via @BotFather

2. **MY_ID**
   - Your Telegram user ID
   - Limits bot access to only you
   - If exposed, someone can use your bot

3. **.env File**
   - Contains all credentials
   - Should **NEVER** be committed to git
   - Should **NEVER** be shared
   - Keep it private and secure

## Bot Security Features

- ✅ Authorization check: Only your Telegram ID can use the bot
- ✅ Command timeout: 5 minutes max per command
- ✅ Process isolation: Commands run in hidden windows
- ✅ VPN/Remote friendly: Works over Telegram's secure connection

## Recommendations

### Network Security

- Bot communicates with Telegram servers only
- All communication is encrypted (HTTPS)
- No open ports required
- Safe behind firewalls/routers

### File Access

- Bot can access any file on your PC
- Downloaded files go to `Downloads` folder by default
- Screenshots saved in project directory (consider moving them)
- Temporary audio files in `data/temp/` (auto-downloaded)

### System Access

- Bot can execute system commands
- Bot can modify system settings (shutdown, restart, lock)
- Bot can read clipboard, type text
- Bot can launch any application

### What Can Be Compromised

⚠️ If someone gains your BOT_TOKEN:
- They can control your PC remotely
- They can read all your files
- They can execute any command
- They can monitor your screen/webcam
- They can turn off your PC

## Protecting Yourself

### Essential Steps

1. **Revoke tokens if exposed**
   ```
   Talk to @BotFather on Telegram
   Select your bot
   → /mybots → Select Bot → API Token → Revoke current token
   ```

2. **Update .env immediately** with new token

3. **Review bot logs** for suspicious activity
   ```powershell
   Get-EventLog -LogName Application -Source "*bot*" -Newest 50
   ```

4. **Monitor Telegram** for unauthorized access attempts

### Additional Security

- Use Telegram's two-factor authentication
- Enable device authentication in Telegram settings
- Consider using a VPN
- Disable bot when not in use

### Auditing Commands

Review what commands are being executed:
- Monitor network activity
- Review file access patterns

## Reporting Security Issues

⚠️ **Do NOT open public issues for security vulnerabilities**

If you discover a security issue:

1. **DO NOT** create a public GitHub issue
2. **DO NOT** share details publicly
3. **Email security details privately** or contact maintainer directly
4. Include:
   - Description of the vulnerability
   - Steps to reproduce
   - Potential impact
   - Suggested fix (if you have one)

We appreciate responsible disclosure and will:
- Investigate promptly
- Notify you of fixes
- Credit you if desired (with your permission)

## Code Review

All code is reviewed before merging:
- Security implications
- Input validation
- Privilege escalation risks
- Information disclosure risks
- Resource consumption issues

## Dependencies

### Go Dependencies
- `telebot.v3` - Telegram bot framework (maintained)
- `godotenv` - Environment loading (maintained)

### Python Dependencies
- Keep dependencies updated
- Review requirements.txt for security patches
- Use `pip install --upgrade -r requirements.txt` regularly

## Legal Disclaimer

This software is provided "AS IS" without warranty. Users are responsible for:
- Complying with local laws
- Using the bot ethically
- Securing their own systems
- Taking necessary precautions

## Best Practices Checklist

- [ ] `.env` file is in `.gitignore`
- [ ] BOT_TOKEN is kept secret
- [ ] Repository is private or non-sensitive
- [ ] Regularly update Go and Python
- [ ] Review command execution logs
- [ ] Keep Windows updated
- [ ] Run antivirus/malware protection
- [ ] Monitor for suspicious activity
- [ ] Disable bot when traveling
- [ ] Have a token revocation plan ready

## Questions?

For security-related questions (non-vulnerability):
- Check this file first
- Review README.md
- Check GitHub discussions

---

**Security is everyone's responsibility. Please report issues responsibly.**
