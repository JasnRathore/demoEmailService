
# Email Tracking Demo

Ever wondered if your emails are actually being read? This Go application demonstrates how to send emails with invisible tracking pixels that notify you the moment someone opens your message. Perfect for understanding email engagement, marketing campaigns, or just satisfying your curiosity about email delivery.

## Project Requirements

- **Go 1.24.4** or higher
- **SMTP server access** (Gmail, Outlook, or any SMTP provider)
- **Network connectivity** for sending emails and receiving tracking events
- **Port availability** (default: 8080) for the tracking server

## Dependencies

This project leverages two custom Go packages for email functionality:

- `github.com/jasnrathore/goemail` - Handles email composition and sending
- `github.com/jasnrathore/trackingmail` - Manages email tracking with invisible pixels
- `gopkg.in/gomail.v2` - Core email sending library
- `gopkg.in/alexcesaro/quotedprintable.v3` - Email encoding support

All dependencies are managed through Go modules and will be automatically downloaded when you build the project.

## Getting Started

### Configuration Setup

Before running the application, you'll need to configure your email credentials and targets.

**Step 1: Update Email Profile**

Modify the email profile in `main.go` with your actual SMTP credentials:

```go
prof := goemail.NewProfile(
    "Your Name",
    "your-email@gmail.com",
    "Your Name <your-email@gmail.com>",
    "smtp.gmail.com:587",
    "your-app-password",
)
```

**Step 2: Configure Target Recipients**

Update the `mails` slice with your intended recipients:

```go
mails := []Target{
    {
        Mail: "recipient1@example.com",
        Id:   1,
    },
    {
        Mail: "recipient2@example.com",
        Id:   2,
    },
}
```

**Step 3: Adjust Tracking Configuration**

The tracking server configuration can be customized based on your needs:

```go
tracker := emailtracker.NewTracker(
    emailtracker.Config{
        Port:   8080,           // Local server port
        Domain: "localhost:8080", // Your domain (use your public domain for production)
        Path:   "/pixel",       // Tracking endpoint path
    },
    // Callback function handles tracking events
)
```

## How to Run the Application

### Install Dependencies

```bash
go mod tidy
```

### Build and Run

```bash
go build -o email-tracker
./email-tracker
```

Or run directly:

```bash
go run main.go
```

### What Happens Next

1. **Tracking Server Starts**: A local HTTP server launches on port 8080
2. **Emails Are Sent**: Each target receives an email with an embedded tracking pixel
3. **Real-time Tracking**: When recipients open emails, you'll see console output like:

```
recipient1@example.com
Email opened: {ID:1 Timestamp:2025-06-27T10:30:45Z UserAgent:Mozilla/5.0...}
```

4. **Continuous Monitoring**: The application runs indefinitely, tracking all email opens

## Relevant Examples

### Basic Email Tracking

The core functionality revolves around embedding invisible tracking pixels in HTML emails:

```go
trackURL := tracker.GenerateLink(strconv.Itoa(item.Id))
body := `
<html>
    <body>
        Hello, this is a tracked email.<br>
    </body>
</html>
`
err := prof.SendMailWithTracking(item.Mail, "Subject Line", body, nil, trackURL)
```

### Custom Email Templates

You can enhance the email body with more sophisticated HTML:

```go
body := fmt.Sprintf(`
<html>
    <head>
        <style>
            body { font-family: Arial, sans-serif; }
            .header { color: #2c3e50; }
        </style>
    </head>
    <body>
        <h1 class="header">Welcome to Our Newsletter</h1>
        <p>Thanks for subscribing! Here's your personalized content.</p>
        <p>Best regards,<br>The Team</p>
    </body>
</html>
`)
```

### Advanced Event Handling

The tracking callback can be extended to perform additional actions:

```go
func(evt emailtracker.OpenEvent) {
    // Find recipient by ID
    var recipient string
    for _, item := range mails {
        if strconv.Itoa(item.Id) == evt.ID {
            recipient = item.Mail
            break
        }
    }
    
    // Log to file, database, or external service
    log.Printf("Email tracking event: %s opened email at %v", 
               recipient, evt.Timestamp)
    
    // Send webhook notification
    // notifyWebhook(recipient, evt)
}
```

### Production Considerations

For production deployment, consider these modifications:

```go
tracker := emailtracker.NewTracker(
    emailtracker.Config{
        Port:   443,
        Domain: "your-domain.com",  // Use your actual domain
        Path:   "/track",
    },
    handleTrackingEvent,
)
```

## Understanding Email Tracking

This application demonstrates the common practice of email tracking used by marketing platforms, newsletters, and CRM systems. When an email is opened, the recipient's email client requests the tracking pixel image from your server, triggering the tracking event.

**Privacy Note**: Always ensure compliance with privacy regulations (GDPR, CAN-SPAM, etc.) and consider informing recipients about tracking in your email communications.

## Troubleshooting

**SMTP Authentication Issues**: Ensure you're using app-specific passwords for Gmail or proper authentication for your email provider.

**Port Conflicts**: If port 8080 is occupied, modify the `Port` field in the tracker configuration.

**Firewall Restrictions**: Ensure your tracking server port is accessible if testing across networks.

---

Ready to track your email engagement? This demo provides a solid foundation for understanding email tracking mechanics and can be extended for more sophisticated email marketing solutions. Dive into the code, experiment with different configurations, and see how email tracking works under the hood!