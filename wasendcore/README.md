# Wasend SDK

A powerful SDK for managing WhatsApp sessions across multiple programming languages. This SDK provides a simple and intuitive interface for creating, managing, and interacting with WhatsApp sessions.

## Features

* Create and manage WhatsApp sessions
* QR code authentication
* Session status management (create, start, stop, restart, delete, get QR code, get session info, get all sessions)
* Contact management (get contact, get all contacts, check if contact exists, get profile picture)
* Message sending (text, image, video, file, voice, link with custom preview, mark messages as seen)
* Group management (create, manage participants, settings, invites)
* Account protection features
* Message logging
* Webhook support
* Multi-language support (TypeScript/JavaScript, Python, Java, .NET, Go)

## Installation

### TypeScript/JavaScript (npm)

```bash
npm install @wasend/core
# or
yarn add @wasend/core
```

### Python (pip)

```bash
pip install wasend-dev
```

### Java (Maven)

```xml
<dependency>
    <groupId>com.wasend</groupId>
    <artifactId>wasend-core</artifactId>
    <version>1.0.0</version>
</dependency>
```

### .NET (NuGet)

```bash
dotnet add package Wasend.Core
```

### Go

```bash
go get github.com/wasenddev/wasend-sdk-go
```

## Quick Start

### TypeScript/JavaScript Example

```go
import { WasendClient, SessionCreateRequest } from '@wasend/core';

async function main() {
    // Initialize the client
    const client = new WasendClient({
        apiKey: 'your-api-key',
        baseUrl: 'https://api.wasend.dev'
    });

    // Create a new session
    const sessionParams: SessionCreateRequest = {
        sessionName: 'my-whatsapp-session',
        phoneNumber: '+919876543210', // Example phone number
        enableAccountProtection: true,
        enableMessageLogging: true,
        enableWebhook: false
    };
    const session = await client.createSession(sessionParams);

    // Get QR code for authentication
    const qrCode = await client.getQRCode(session.uniqueSessionId);
    console.log('Scan this QR code with WhatsApp:', qrCode.data);

    // Start the session
    await client.startSession(session.uniqueSessionId);

    // Get session information
    const sessionInfo = await client.getSessionInfo(session.uniqueSessionId);
    console.log('Session status:', sessionInfo.status);
}

main().catch(console.error);
```

### Python Example

```python
from wasend_dev import WasendClient

# Initialize the client
client = WasendClient(api_key='your-api-key', base_url='https://api.wasend.dev')

try:
    # Create a new session
    session = client.create_session(
        session_name='my-whatsapp-session-python',
        phone_number='+919876543210',  # Example phone number
        enable_account_protection=True,
        enable_message_logging=True,
        enable_webhook=False
    )
    # Assuming session object has unique_session_id or similar attribute/key
    # e.g., session.unique_session_id or session['uniqueSessionId']
    # For this example, let's assume it's session.unique_session_id based on TS example structure

    print(f"Session created with ID: {session.unique_session_id}")

    # Get QR code for authentication
    qr_code_response = client.get_qr_code(session_id=session.unique_session_id)
    print(f"Scan this QR code with WhatsApp: {qr_code_response.data}")

    # Start the session
    client.start_session(session_id=session.unique_session_id)
    print(f"Session {session.unique_session_id} starting...")

    # Get session information
    session_info = client.get_session_info(session_id=session.unique_session_id)
    print(f"Session status: {session_info.status}")

except Exception as e:
    print(f"An error occurred: {e}")
```

### Go Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/wasenddev/wasend-sdk-go/wasend"
)

func main() {
	// Initialize the client
	// Ensure WasendClientOptions or similar struct is used if constructor expects it
	client, err := wasend.NewClient("your-api-key", "https://api.wasend.dev") // Or wasend.NewClient(&wasend.Config{...})
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	// Create a new session
	sessionParams := wasend.SessionCreateRequest{
		SessionName:             "my-whatsapp-session-go",
		PhoneNumber:             "+919876543210", // Example phone number
		EnableAccountProtection: true,
		EnableMessageLogging:    true,
		EnableWebHook:           false,
	}
	session, err := client.CreateSession(sessionParams)
	if err != nil {
		log.Fatalf("Error creating session: %v", err)
	}
	fmt.Println("Session created with ID:", session.UniqueSessionId)

	// Get QR code for authentication
	qrCode, err := client.GetQRCode(session.UniqueSessionId)
	if err != nil {
		log.Fatalf("Error getting QR code: %v", err)
	}
	fmt.Println("Scan this QR code with WhatsApp:", qrCode.Data)

	// Start the session
	_, err = client.StartSession(session.UniqueSessionId)
	if err != nil {
		log.Fatalf("Error starting session: %v", err)
	}
	fmt.Println("Session starting...")

	// Get session information
	sessionInfo, err := client.GetSessionInfo(session.UniqueSessionId)
	if err != nil {
		log.Fatalf("Error getting session info: %v", err)
	}
	fmt.Println("Session status:", sessionInfo.Status)
}
```

### .NET Example

```csharp
using System;
using System.Threading.Tasks;
using Wasend.Core;

public class WasendExample
{
    public static async Task Main(string[] args)
    {
        // Initialize the client
        var client = new WasendClient(new WasendClientOptions
        {
            ApiKey = "your-api-key",
            BaseUrl = "https://api.wasend.dev"
        });

        try
        {
            // Create a new session
            var sessionParams = new SessionCreateRequest
            {
                SessionName = "my-whatsapp-session-dotnet",
                PhoneNumber = "+919876543210",
                EnableAccountProtection = true,
                EnableMessageLogging = true,
                EnableWebHook = false
            };
            var session = await client.CreateSessionAsync(sessionParams);
            Console.WriteLine($"Session created with ID: {session.UniqueSessionId}");

            // Get QR code for authentication
            var qrCode = await client.GetQRCodeAsync(session.UniqueSessionId);
            Console.WriteLine($"Scan this QR code with WhatsApp: {qrCode.Data}");

            // Start the session
            await client.StartSessionAsync(session.UniqueSessionId);
            Console.WriteLine($"Session {session.UniqueSessionId} starting...");

            // Get session information
            var sessionInfo = await client.GetSessionInfoAsync(session.UniqueSessionId);
            Console.WriteLine($"Session status: {sessionInfo.Status}");
        }
        catch (Exception ex)
        {
            Console.WriteLine($"An error occurred: {ex.Message}");
        }
    }
}
```

## Session Management

### Creating a Session

```go
const session = await client.sessions.createSession({
    sessionName: 'my-whatsapp-session',
    phoneNumber: '+919876543210',
    enableAccountProtection: true,
    enableMessageLogging: true,
    enableWebhook: true,
    webhookUrl: 'https://your-webhook-url.com/callback'
});
```

### Session Configuration Options

* `sessionName`: A unique name for your session
* `phoneNumber`: The WhatsApp phone number to use (format: +[country code][number])
* `enableAccountProtection`: Enable additional security features
* `enableMessageLogging`: Enable message history logging
* `enableWebhook`: Enable webhook notifications
* `webhookUrl`: URL for receiving webhook notifications (required if enableWebhook is true)

### Managing Sessions

```go
// Get all sessions
const allSessions = await client.sessions.getAllSessions();

// Get specific session info
const sessionInfo = await client.sessions.getSessionInfo(sessionId);

// Start a session
await client.sessions.startSession(sessionId);

// Stop a session
await client.sessions.stopSession(sessionId);

// Restart a session
await client.sessions.restartSession(sessionId);

// Delete a session
await client.sessions.deleteSession(sessionId);
```

## Group Management

The SDK provides comprehensive group management features for WhatsApp groups.

### Creating a Group

```go
const group = await client.createGroup(sessionId, {
    name: "My Group",
    participants: [{ id: "+919876543210" }],
    description: "Group description"
});
```

### Group Configuration Options

* `name`: The name of the group
* `participants`: Array of participants with their phone numbers
* `description`: Optional group description
* `pictureUrl`: Optional URL for group picture
* `tags`: Optional array of tags for the group

### Managing Groups

```go
// Get all groups
const groups = await client.getGroups(sessionId);

// Get specific group info
const groupInfo = await client.getGroup(sessionId, groupId);

// Get group participants
const participants = await client.getGroupParticipants(sessionId, groupId);

// Add participants to group
await client.addGroupParticipants(sessionId, groupId, {
    participants: ["+919876543210"],
    notify: true
});

// Remove participants from group
await client.removeGroupParticipants(sessionId, groupId, {
    participants: ["+919876543210"],
    notify: true
});

// Promote participants to admin
await client.promoteGroupParticipants(sessionId, groupId, {
    participants: ["+919876543210"],
    notify: true
});

// Demote participants from admin
await client.demoteGroupParticipants(sessionId, groupId, {
    participants: ["+919876543210"],
    notify: true
});
```

### Group Settings

```go
// Set group description
await client.setGroupDescription(sessionId, groupId, {
    description: "New group description"
});

// Set group subject/name
await client.setGroupSubject(sessionId, groupId, {
    subject: "New group name"
});

// Set group picture
await client.setGroupPicture(sessionId, groupId, {
    url: "https://example.com/picture.jpg",
    format: "jpeg",
    cropToSquare: true
});

// Delete group picture
await client.deleteGroupPicture(sessionId, groupId);
```

### Group Security Settings

```go
// Set info admin only
await client.setGroupInfoAdminOnly(sessionId, groupId, {
    enabled: true
});

// Set messages admin only
await client.setGroupMessagesAdminOnly(sessionId, groupId, {
    enabled: true
});

// Get current security settings
const infoAdminOnly = await client.getGroupInfoAdminOnly(sessionId, groupId);
const messagesAdminOnly = await client.getGroupMessagesAdminOnly(sessionId, groupId);
```

### Group Invite Management

```go
// Get group invite code
const inviteCode = await client.getGroupInviteCode(sessionId, groupId);

// Revoke group invite code
const newInviteCode = await client.revokeGroupInviteCode(sessionId, groupId);

// Join a group using invite code
const joinInfo = await client.getGroupJoinInfo(sessionId, "https://chat.whatsapp.com/1234567890abcdef");
await client.joinGroup(sessionId, {
    code: "https://chat.whatsapp.com/1234567890abcdef"
});

// Leave a group
await client.leaveGroup(sessionId, groupId);
```

### Group Query Options

When retrieving groups, you can use various query parameters:

```go
const groups = await client.getGroups(sessionId, {
    sortBy: "creation",
    sortOrder: "desc",
    limit: 50,
    offset: 0,
    exclude: ["participants"],
    status: "ACTIVE",
    search: "group name",
    tags: ["tag1", "tag2"]
});
```

### Group Status

A group can have the following statuses:

* `ACTIVE`: Group is active and can be used
* `ARCHIVED`: Group has been archived
* `DELETED`: Group has been deleted

## Authentication

### QR Code Authentication

1. Create a session
2. Get the QR code
3. Scan the QR code with WhatsApp on your phone
4. The session will automatically connect once scanned

```go
const qrCode = await client.getQRCode(sessionId);
console.log('QR Code data:', qrCode.data);
```

## Session Status

A session can have the following statuses:

* `CREATED`: Session has been created but not started
* `STARTING`: Session is in the process of starting
* `CONNECTED`: Session is connected and ready to use
* `STOPPED`: Session has been stopped
* `ERROR`: Session encountered an error

## Webhook Integration

To receive notifications about session events:

1. Enable webhooks when creating a session
2. Provide a valid webhook URL
3. Handle incoming webhook notifications on your server

```go
const session = await client.sessions.createSession({
    sessionName: 'webhook-enabled-session',
    phoneNumber: '+919876543210',
    enableWebhook: true,
    webhookUrl: 'https://your-server.com/webhook'
});
```

## Error Handling

```go
try {
    const session = await client.sessions.createSession({
        sessionName: 'test-session',
        phoneNumber: '+919876543210'
    });
} catch (error) {
    console.error('Error creating session:', error);
}
```

## Contact Management

The SDK allows you to manage contacts associated with a session.

### Get Contact Details

```go
const contactId = "contact_jid@c.us";
const contact = await client.getContact(sessionId, contactId);
console.log('Contact details:', contact);
```

### Get All Contacts

```go
const params = { limit: 10, offset: 0, sortBy: "name", sortOrder: "asc" };
const contacts = await client.getContacts(sessionId, params);
console.log('Contacts list:', contacts);
```

### Check if Contact Exists

```go
const phoneNumber = "1234567890";
const exists = await client.checkContactExists(sessionId, { phone: phoneNumber });
console.log(`Contact with phone ${phoneNumber} exists:`, exists);
```

### Get Profile Picture URL

```go
const contactIdForPic = "contact_jid@c.us";
const picInfo = await client.getProfilePictureUrl(sessionId, { contactId: contactIdForPic, refresh: false });
console.log('Profile picture URL:', picInfo.url);
```

## Message Sending

The SDK provides methods to send various types of messages. All `send...` methods typically require the `sessionId` and a request object specific to the message type.

### Send Text Message

```go
const textMessage = {
    to: "recipient_jid@c.us",
    text: "Hello from Wasend SDK!"
};
const sentMessageInfo = await client.sendTextMessage(sessionId, textMessage);
console.log('Text message sent:', sentMessageInfo);
```

### Send Image Message

```go
const imageMessage = {
    to: "recipient_jid@c.us",
    url: "https://example.com/image.jpg",
    caption: "Check out this image!"
};
const sentImageInfo = await client.sendImageMessage(sessionId, imageMessage);
console.log('Image message sent:', sentImageInfo);
```

### Send Video Message

```go
const videoMessage = {
    to: "recipient_jid@c.us",
    url: "https://example.com/video.mp4",
    caption: "Watch this cool video!"
};
const sentVideoInfo = await client.sendVideoMessage(sessionId, videoMessage);
console.log('Video message sent:', sentVideoInfo);
```

### Send File/Document Message

```go
const fileMessage = {
    to: "recipient_jid@c.us",
    url: "https://example.com/document.pdf",
    fileName: "document.pdf",
    mimeType: "application/pdf"
};
const sentFileInfo = await client.sendFileMessage(sessionId, fileMessage);
console.log('File message sent:', sentFileInfo);
```

### Send Voice Message

```go
const voiceMessage = {
    to: "recipient_jid@c.us",
    url: "https://example.com/audio.ogg",
};
const sentVoiceInfo = await client.sendVoiceMessage(sessionId, voiceMessage);
console.log('Voice message sent:', sentVoiceInfo);
```

### Send Link with Custom Preview

```go
const linkPreview = {
    title: "Wasend SDK",
    description: "Powerful WhatsApp SDK",
    thumbnailUrl: "https://example.com/logo.png"
};

const linkMessage = {
    to: "recipient_jid@c.us",
    text: "Check out this SDK: https://wasend.dev",
    preview: linkPreview
};
const sentLinkInfo = await client.sendLinkWithCustomPreview(sessionId, linkMessage);
console.log('Link message with custom preview sent:', sentLinkInfo);
```

### Mark Message as Seen

```go
const seenRequest = {
    to: "sender_jid@c.us",
    messageId: "message_id_to_mark_seen"
};
await client.sendSeen(sessionId, seenRequest);
console.log('Message marked as seen.');
```

## Best Practices

1. Always store the `sessionId` after creating a session
2. Implement proper error handling
3. Use environment variables for API keys
4. Enable account protection for production sessions
5. Implement webhook handling for real-time updates
6. Regularly check session status
7. Clean up unused sessions

## Support

For support, please contact:

* Email: admin@wasend.dev
* Website: https://wasend.dev
* Documentation: https://docs.wasend.dev

## License

This SDK is licensed under the Apache-2.0 License. See the [LICENSE](LICENSE) file for details.

For more detailed API information, please refer to the [API.md](API.md) file.
