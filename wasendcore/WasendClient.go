package wasendcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/wasenddev/wasend-sdk-go/wasendcore/jsii"
)

// Main Wasend SDK Client.
//
// This class provides access to the Wasend API for sending messages,
// managing contacts, and other messaging operations.
//
// Example:
//   const client = new WasendClient({
//     apiKey: 'your-api-key'
//   });
//
//   const result = await client.sendMessage({
//     to: '+1234567890',
//     content: 'Hello, World!'
//   });
//
type WasendClient interface {
	// Add participants to group.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/{sessionId}/groups/{groupId}/participants/add' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "participants": ["+1234567890", "+0987654321"],
	//       "notify": true
	//     }'
	//
	AddGroupParticipants(sessionId *string, groupId *string, request *ParticipantsRequest)
	// Check if a phone number is registered in WhatsApp.
	//
	// Returns: Promise resolving to the check result.
	//
	// Example:
	//   curl -X GET 'http://localhost:3001/contacts/check-exists?session={sessionId}&phone=1213213213' \
	//     -H 'Authorization: Bearer {apiKey}'
	//
	CheckContactExists(sessionId *string, params *CheckContactExistsQueryParams) *WANumberExistResult
	// Create a new group.
	//
	// Returns: Promise resolving to the created group.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/{sessionId}/groups' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "name": "My Group",
	//       "participants": [{"id": "+1234567890"}],
	//       "description": "Group description"
	//     }'
	//
	CreateGroup(sessionId *string, request *CreateGroupRequest) *Group
	CreateSession(request *SessionCreateRequest) *Session
	// Delete group picture.
	//
	// Example:
	//   curl -X DELETE 'http://localhost:3001/{sessionId}/groups/{groupId}/picture' \
	//     -H 'Authorization: Bearer {apiKey}'
	//
	DeleteGroupPicture(sessionId *string, groupId *string)
	DeleteSession(sessionId *string)
	// Demote participants from admin.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/{sessionId}/groups/{groupId}/admin/demote' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "participants": ["+1234567890", "+0987654321"],
	//       "notify": true
	//     }'
	//
	DemoteGroupParticipants(sessionId *string, groupId *string, request *ParticipantsRequest)
	// Retrieves all chats for a given session.
	//
	// Returns: A promise that resolves to an SdkResponse object. If successful, `data` contains an array of Chat objects. If failed, `error` contains the error message.
	//
	// Example:
	//   const client = new WasendClient({ apiKey: 'your-api-key' });
	//   const result = await client.getAllChats('my-session', { limit: 10, sortBy: 'timestamp', sortOrder: 'desc' });
	//   if (result.success && result.data) {
	//     console.log('Chats:', result.data);
	//   } else {
	//     console.error('Failed to get chats:', result.error);
	//   }
	//
	GetAllChats(session *string, options *GetChatsOptions) *SdkResponse
	// Retrieves the profile picture of a chat (user or group).
	//
	// Returns: A promise that resolves to an SdkResponse object. If successful, `data` contains a `ChatPictureResponse` object. If failed, `error` contains the error message.
	//
	// Example:
	//   const client = new WasendClient({ apiKey: 'your-api-key' });
	//   const session = 'my-session';
	//   const chatId = '1234567890@c.us';
	//
	//   const result = await client.getChatPicture(session, chatId);
	//   if (result.success && result.data) {
	//     if (result.data.url) {
	//       console.log('Chat picture URL:', result.data.url);
	//     } else {
	//       console.log('Chat does not have a picture or it could not be retrieved.');
	//     }
	//
	//     // Force refresh
	//     const refreshedResult = await client.getChatPicture(session, chatId, { refresh: true });
	//     if (refreshedResult.success && refreshedResult.data) {
	//        console.log('Refreshed picture URL:', refreshedResult.data.url);
	//     } else {
	//        console.error('Failed to refresh chat picture:', refreshedResult.error);
	//     }
	//   } else {
	//     console.error('Failed to get chat picture:', result.error);
	//   }
	//
	GetChatPicture(session *string, chatId *string, options *GetChatPictureOptions) *SdkResponse
	// Retrieves an overview of chats for a given session, allowing pagination and filtering by chat IDs.
	//
	// Returns: A promise that resolves to an SdkResponse object. If successful, `data` contains an array of ChatOverview objects. If failed, `error` contains the error message.
	//
	// Example:
	//   const client = new WasendClient({ apiKey: 'your-api-key' });
	//   const result = await client.getChatsOverview('my-session', { limit: 10, ids: ['1234567890@c.us'] });
	//   if (result.success && result.data) {
	//     console.log(result.data);
	//   } else {
	//     console.error('Failed to get chats overview:', result.error);
	//   }
	//
	GetChatsOverview(session *string, options *GetChatsOverviewOptions) *SdkResponse
	// Get contact basic info.
	//
	// Returns: Promise resolving to the contact information.
	//
	// Example:
	//   curl -X GET 'http://localhost:3001/contacts?session={sessionId}&contactId=+11111111111' \
	//     -H 'Authorization: Bearer {apiKey}'
	//
	GetContact(sessionId *string, params *GetContactQueryParams) *Contact
	// Get contact's profile picture URL.
	//
	// Returns: Promise resolving to the profile picture response.
	//
	// Example:
	//   curl -X GET 'http://localhost:3001/contacts/profile-picture?session={sessionId}&contactId=+11111111111&refresh=false' \
	//     -H 'Authorization: Bearer {apiKey}'
	//
	GetContactProfilePicture(sessionId *string, params *GetProfilePictureQueryParams) *ProfilePictureResponse
	// Get all contacts.
	//
	// Returns: Promise resolving to the list of contacts.
	//
	// Example:
	//   curl -X GET 'http://localhost:3001/contacts/all?session={sessionId}&sortBy=name&sortOrder=asc&limit=10&offset=0' \
	//     -H 'Authorization: Bearer {apiKey}'
	//
	GetContacts(sessionId *string, params *GetContactsQueryParams) *[]*Contact
	// Get a specific group.
	//
	// Returns: Promise resolving to the group.
	//
	// Example:
	//   curl -X GET 'http://localhost:3001/{sessionId}/groups/{groupId}' \
	//     -H 'Authorization: Bearer {apiKey}'
	//
	GetGroup(sessionId *string, groupId *string) *Group
	// Get group info admin only setting.
	//
	// Returns: Promise resolving to the settings security change info.
	//
	// Example:
	//   curl -X GET 'http://localhost:3001/{sessionId}/groups/{groupId}/settings/security/info-admin-only' \
	//     -H 'Authorization: Bearer {apiKey}'
	//
	GetGroupInfoAdminOnly(sessionId *string, groupId *string) *SettingsSecurityChangeInfo
	// Get group invite code.
	//
	// Returns: Promise resolving to the invite code.
	//
	// Example:
	//   curl -X GET 'http://localhost:3001/{sessionId}/groups/{groupId}/invite-code' \
	//     -H 'Authorization: Bearer {apiKey}'
	//
	GetGroupInviteCode(sessionId *string, groupId *string) *string
	// Get group join info.
	//
	// Returns: Promise resolving to the group info.
	//
	// Example:
	//   curl -X GET 'http://localhost:3001/{sessionId}/groups/join-info?code=https://chat.whatsapp.com/1234567890abcdef' \
	//     -H 'Authorization: Bearer {apiKey}'
	//
	GetGroupJoinInfo(sessionId *string, code *string) *Group
	// Get group messages admin only setting.
	//
	// Returns: Promise resolving to the settings security change info.
	//
	// Example:
	//   curl -X GET 'http://localhost:3001/{sessionId}/groups/{groupId}/settings/security/messages-admin-only' \
	//     -H 'Authorization: Bearer {apiKey}'
	//
	GetGroupMessagesAdminOnly(sessionId *string, groupId *string) *SettingsSecurityChangeInfo
	// Get group participants.
	//
	// Returns: Promise resolving to the list of participants.
	//
	// Example:
	//   curl -X GET 'http://localhost:3001/{sessionId}/groups/{groupId}/participants' \
	//     -H 'Authorization: Bearer {apiKey}'
	//
	GetGroupParticipants(sessionId *string, groupId *string) *[]*GroupParticipant
	// Get group picture.
	//
	// Returns: Promise resolving to the group picture response.
	//
	// Example:
	//   curl -X GET 'http://localhost:3001/{sessionId}/groups/{groupId}/picture?refresh=true' \
	//     -H 'Authorization: Bearer {apiKey}'
	//
	GetGroupPicture(sessionId *string, groupId *string, refresh *bool) *GroupPictureResponse
	// Get all groups.
	//
	// Returns: Promise resolving to the list of groups.
	//
	// Example:
	//   curl -X GET 'http://localhost:3001/{sessionId}/groups?sortBy=creation&sortOrder=desc&limit=50&offset=0' \
	//     -H 'Authorization: Bearer {apiKey}'
	//
	GetGroups(sessionId *string, params *GetGroupsQueryParams) *[]*Group
	// Get groups count.
	//
	// Returns: Promise resolving to the count response.
	//
	// Example:
	//   curl -X GET 'http://localhost:3001/{sessionId}/groups/count' \
	//     -H 'Authorization: Bearer {apiKey}'
	//
	GetGroupsCount(sessionId *string) *CountResponse
	// Get message by ID.
	//
	// Returns: Promise resolving to the API response.
	GetMessage(messageId *string) *SdkResponse
	// Retrieves a specific message by its ID from a given chat.
	//
	// Returns: A promise that resolves to an SdkResponse object. If successful, `data` contains the `WAMessage` object. If failed, `error` contains the error message.
	//
	// Example:
	//   const client = new WasendClient({ apiKey: 'your-api-key' });
	//   const session = 'my-session';
	//   const chatId = '1234567890@c.us';
	//   const messageId = 'messageId123';
	//
	//   const result = await client.getMessageById(session, chatId, messageId);
	//   if (result.success && result.data) {
	//     console.log('Retrieved message:', result.data);
	//
	//     // Retrieve message and attempt to include media
	//     const mediaResult = await client.getMessageById(session, chatId, messageId, { downloadMedia: true });
	//     if (mediaResult.success && mediaResult.data) {
	//        console.log('Message with media (if any):', mediaResult.data);
	//     } else {
	//        console.error('Failed to get message with media:', mediaResult.error);
	//     }
	//   } else {
	//     console.error('Failed to get message by ID:', result.error);
	//   }
	//
	GetMessageById(session *string, chatId *string, messageId *string, options *GetMessageByIdOptions) *SdkResponse
	// Get messages from a chat.
	//
	// This method retrieves a list of messages from a specified chat, allowing for pagination,
	// media download preferences, and filtering based on various criteria.
	//
	// Returns: A promise that resolves to an SdkResponse object. If successful, `data` contains an array of WAMessage objects. If failed, `error` contains the error message.
	//
	// Example:
	//   const client = new WasendClient({ apiKey: 'your-api-key' });
	//   const session = 'my-session';
	//   const chatId = '1234567890@c.us';
	//
	//   // Get the last 10 messages
	//   const result = await client.getMessages(session, chatId, { limit: 10 });
	//   if (result.success && result.data) {
	//     console.log('Last 10 messages:', result.data);
	//
	//     // Get messages from the last 24 hours, with media, sent by me
	//     const twentyFourHoursAgo = Math.floor((Date.now() - 24 * 60 * 60 * 1000) / 1000);
	//     const myMessagesResult = await client.getMessages(session, chatId, {
	//       limit: 50, // Max 50
	//       downloadMedia: true,
	//       filter: {
	//         timestampGte: twentyFourHoursAgo,
	//         fromMe: true,
	//       },
	//     });
	//     if (myMessagesResult.success && myMessagesResult.data) {
	//        console.log('My messages from last 24h with media:', myMessagesResult.data);
	//     } else {
	//        console.error('Failed to get my messages:', myMessagesResult.error);
	//     }
	//   } else {
	//    console.error('Failed to get messages:', result.error);
	//   }
	//
	GetMessages(session *string, chatId *string, options *GetMessagesOptions) *SdkResponse
	// Join a group.
	//
	// Returns: Promise resolving to the join group response.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/{sessionId}/groups/join' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "code": "https://chat.whatsapp.com/1234567890abcdef"
	//     }'
	//
	JoinGroup(sessionId *string, request *JoinGroupRequest) *JoinGroupResponse
	// Leave a group.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/{sessionId}/groups/{groupId}/leave' \
	//     -H 'Authorization: Bearer {apiKey}'
	//
	LeaveGroup(sessionId *string, groupId *string)
	// Utility method to process a message with proper timing and typing indicators This follows WhatsApp's guidelines to avoid being flagged as spam.
	//
	// Returns: Promise resolving to the sent message.
	//
	// Example:
	//   // Send a text message with proper processing
	//   const message = await client.processMessage({
	//     session: "sessionId",
	//     to: "+1234567890",
	//     text: "Hello, World!"
	//   });
	//
	//   // Send an image with caption and proper processing
	//   const imageMessage = await client.processMessage({
	//     session: "sessionId",
	//     to: "+1234567890",
	//     imageUrl: "https://example.com/image.jpg",
	//     text: "Check out this image!"
	//   });
	//
	ProcessMessage(request *SendRequest, options *map[string]*string) *WAMessage
	// Promote participants to admin.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/{sessionId}/groups/{groupId}/admin/promote' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "participants": ["+1234567890", "+0987654321"],
	//       "notify": true
	//     }'
	//
	PromoteGroupParticipants(sessionId *string, groupId *string, request *ParticipantsRequest)
	// Marks messages in a specific chat as read.
	//
	// This can be done by specifying the number of
	// latest messages to mark as read, or by specifying a number of days to mark messages from.
	//
	// If the API request fails due to network issues or server-side errors, this method will
	// return a `ReadChatMessagesResponse` object with `success: false` and an error message.
	//
	// Returns: A promise that resolves to a `ReadChatMessagesResponse` object indicating the success or failure
	// of the operation, along with an optional descriptive message.
	//
	// Example:
	//   const client = new WasendClient({ apiKey: 'your-api-key' });
	//   const session = 'my-session';
	//   const chatId = '1234567890@c.us';
	//
	//   // Mark the last 5 messages as read
	//   const response1 = await client.readMessages(session, chatId, { messages: 5 });
	//   if (response1.success) {
	//     console.log(response1.message || 'Successfully marked 5 messages as read.');
	//   } else {
	//     console.error('Failed to mark messages as read:', response1.message);
	//   }
	//
	//   // Mark messages from the last 2 days as read
	//   const response2 = await client.readMessages(session, chatId, { days: 2 });
	//   if (response2.success) {
	//     console.log(response2.message || 'Successfully marked messages from last 2 days as read.');
	//   } else {
	//     console.error('Failed to mark messages as read:', response2.message);
	//   }
	//
	ReadMessages(session *string, chatId *string, options *ReadMessagesOptions) *ReadChatMessagesResponse
	// Remove participants from group.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/{sessionId}/groups/{groupId}/participants/remove' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "participants": ["+1234567890", "+0987654321"],
	//       "notify": true
	//     }'
	//
	RemoveGroupParticipants(sessionId *string, groupId *string, request *ParticipantsRequest)
	RestartSession(sessionId *string) *Session
	// Get account information.
	//
	// Returns: Promise resolving to the API response.
	RetrieveAccount() *SdkResponse
	RetrieveAllSessions() *GetAllSessionsResponse
	// Get the current configuration.
	//
	// Returns: The configuration object (without sensitive data).
	RetrieveConfig() *WasendConfigInfo
	RetrieveQRCode(sessionId *string) *QRCodeResponse
	RetrieveSessionInfo(sessionId *string) *Session
	// Revoke group invite code.
	//
	// Returns: Promise resolving to the new invite code.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/{sessionId}/groups/{groupId}/invite-code/revoke' \
	//     -H 'Authorization: Bearer {apiKey}'
	//
	RevokeGroupInviteCode(sessionId *string, groupId *string) *string
	// Send a message using the unified send endpoint.
	//
	// Returns: Promise resolving to the sent message.
	//
	// Example:
	//   // Send a text message
	//   const textMessage = await client.send({
	//     session: "sessionId",
	//     to: "+1234567890",
	//     text: "Hello from WASend!"
	//   });
	//
	//   // Send an image
	//   const imageMessage = await client.send({
	//     session: "sessionId",
	//     to: "+1234567890",
	//     imageUrl: "https://example.com/image.jpg",
	//     text: "Check out this image!",
	//     mimetype: "image/jpeg"
	//   });
	//
	//   // Send a video
	//   const videoMessage = await client.send({
	//     session: "sessionId",
	//     to: "+1234567890",
	//     videoUrl: "https://example.com/video.mp4",
	//     text: "Check out this video!",
	//     mimetype: "video/mp4"
	//   });
	//
	//   // Send a document
	//   const documentMessage = await client.send({
	//     session: "sessionId",
	//     to: "+1234567890",
	//     documentUrl: "https://example.com/document.pdf",
	//     filename: "document.pdf",
	//     mimetype: "application/pdf"
	//   });
	//
	//   // Send an audio message
	//   const audioMessage = await client.send({
	//     session: "sessionId",
	//     to: "+1234567890",
	//     audioUrl: "https://example.com/audio.mp3",
	//     mimetype: "audio/mpeg"
	//   });
	//
	Send(request *SendRequest) *WAMessage
	// Send a file message.
	//
	// Returns: Promise resolving to the sent message.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/sendFile' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "session": "sessionId",
	//       "to": "+11111111111",
	//       "url": "https://example.com/file.pdf",
	//       "fileName": "document.pdf",
	//       "mimeType": "application/pdf"
	//     }'
	//
	SendFile(request *MessageFileRequest) *WAMessage
	// Send an image message.
	//
	// Returns: Promise resolving to the sent message.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/sendImage' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "session": "sessionId",
	//       "to": "+11111111111",
	//       "url": "https://example.com/image.jpg",
	//       "caption": "Check out this image!"
	//     }'
	//
	SendImage(request *MessageImageRequest) *WAMessage
	// Send a text message with custom link preview.
	//
	// Returns: Promise resolving to the sent message.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/send/link-custom-preview' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "session": "sessionId",
	//       "to": "+11111111111",
	//       "text": "Check out this link!",
	//       "preview": {
	//         "title": "Example Website",
	//         "description": "A great website",
	//         "thumbnailUrl": "https://example.com/thumbnail.jpg"
	//       }
	//     }'
	//
	SendLinkCustomPreview(request *MessageLinkCustomPreviewRequest) *WAMessage
	// Send a message to a recipient.
	//
	// Returns: Promise resolving to the API response.
	SendMessage(request *MessageRequest) *SdkResponse
	// Mark a message as seen.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/sendSeen' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "session": "sessionId",
	//       "to": "+11111111111",
	//       "messageId": "messageId"
	//     }'
	//
	SendSeen(request *SendSeenRequest)
	// Send a text message.
	//
	// Returns: Promise resolving to the sent message.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/sendText' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "session": "sessionId",
	//       "to": "+11111111111",
	//       "text": "Hello, World!"
	//     }'
	//
	SendText(request *MessageTextRequest) *WAMessage
	// Send a video message.
	//
	// Returns: Promise resolving to the sent message.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/sendVideo' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "session": "sessionId",
	//       "to": "+11111111111",
	//       "url": "https://example.com/video.mp4",
	//       "caption": "Check out this video!"
	//     }'
	//
	SendVideo(request *MessageVideoRequest) *WAMessage
	// Send a voice message.
	//
	// Returns: Promise resolving to the sent message.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/sendVoice' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "session": "sessionId",
	//       "to": "+11111111111",
	//       "url": "https://example.com/voice.mp3",
	//       "duration": 30
	//     }'
	//
	SendVoice(request *MessageVoiceRequest) *WAMessage
	// Set group description.
	//
	// Example:
	//   curl -X PUT 'http://localhost:3001/{sessionId}/groups/{groupId}/description' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "description": "New group description"
	//     }'
	//
	SetGroupDescription(sessionId *string, groupId *string, request *DescriptionRequest)
	// Set group info admin only setting.
	//
	// Example:
	//   curl -X PUT 'http://localhost:3001/{sessionId}/groups/{groupId}/settings/security/info-admin-only' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "enabled": true
	//     }'
	//
	SetGroupInfoAdminOnly(sessionId *string, groupId *string, request *SettingsSecurityChangeInfo)
	// Set group messages admin only setting.
	//
	// Example:
	//   curl -X PUT 'http://localhost:3001/{sessionId}/groups/{groupId}/settings/security/messages-admin-only' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "enabled": true
	//     }'
	//
	SetGroupMessagesAdminOnly(sessionId *string, groupId *string, request *SettingsSecurityChangeInfo)
	// Set group picture.
	//
	// Example:
	//   curl -X PUT 'http://localhost:3001/{sessionId}/groups/{groupId}/picture' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "url": "https://example.com/picture.jpg",
	//       "format": "jpeg",
	//       "cropToSquare": true
	//     }'
	//
	SetGroupPicture(sessionId *string, groupId *string, request *ProfilePictureRequest)
	// Set group subject.
	//
	// Example:
	//   curl -X PUT 'http://localhost:3001/{sessionId}/groups/{groupId}/subject' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "subject": "New group name"
	//     }'
	//
	SetGroupSubject(sessionId *string, groupId *string, request *SubjectRequest)
	StartSession(sessionId *string) *Session
	// Start typing indicator.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/startTyping' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "session": "sessionId",
	//       "to": "+11111111111"
	//     }'
	//
	StartTyping(request *ChatRequest)
	StopSession(sessionId *string) *Session
	// Stop typing indicator.
	//
	// Example:
	//   curl -X POST 'http://localhost:3001/stopTyping' \
	//     -H 'Authorization: Bearer {apiKey}' \
	//     -H 'Content-Type: application/json' \
	//     -d '{
	//       "session": "sessionId",
	//       "to": "+11111111111"
	//     }'
	//
	StopTyping(request *ChatRequest)
}

// The jsii proxy struct for WasendClient
type jsiiProxy_WasendClient struct {
	_ byte // padding
}

// Creates a new Wasend client instance.
func NewWasendClient(config *WasendConfig) WasendClient {
	_init_.Initialize()

	if err := validateNewWasendClientParameters(config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WasendClient{}

	_jsii_.Create(
		"@wasend/core.WasendClient",
		[]interface{}{config},
		&j,
	)

	return &j
}

// Creates a new Wasend client instance.
func NewWasendClient_Override(w WasendClient, config *WasendConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@wasend/core.WasendClient",
		[]interface{}{config},
		w,
	)
}

func (w *jsiiProxy_WasendClient) AddGroupParticipants(sessionId *string, groupId *string, request *ParticipantsRequest) {
	if err := w.validateAddGroupParticipantsParameters(sessionId, groupId, request); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addGroupParticipants",
		[]interface{}{sessionId, groupId, request},
	)
}

func (w *jsiiProxy_WasendClient) CheckContactExists(sessionId *string, params *CheckContactExistsQueryParams) *WANumberExistResult {
	if err := w.validateCheckContactExistsParameters(sessionId, params); err != nil {
		panic(err)
	}
	var returns *WANumberExistResult

	_jsii_.Invoke(
		w,
		"checkContactExists",
		[]interface{}{sessionId, params},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) CreateGroup(sessionId *string, request *CreateGroupRequest) *Group {
	if err := w.validateCreateGroupParameters(sessionId, request); err != nil {
		panic(err)
	}
	var returns *Group

	_jsii_.Invoke(
		w,
		"createGroup",
		[]interface{}{sessionId, request},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) CreateSession(request *SessionCreateRequest) *Session {
	if err := w.validateCreateSessionParameters(request); err != nil {
		panic(err)
	}
	var returns *Session

	_jsii_.Invoke(
		w,
		"createSession",
		[]interface{}{request},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) DeleteGroupPicture(sessionId *string, groupId *string) {
	if err := w.validateDeleteGroupPictureParameters(sessionId, groupId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"deleteGroupPicture",
		[]interface{}{sessionId, groupId},
	)
}

func (w *jsiiProxy_WasendClient) DeleteSession(sessionId *string) {
	if err := w.validateDeleteSessionParameters(sessionId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"deleteSession",
		[]interface{}{sessionId},
	)
}

func (w *jsiiProxy_WasendClient) DemoteGroupParticipants(sessionId *string, groupId *string, request *ParticipantsRequest) {
	if err := w.validateDemoteGroupParticipantsParameters(sessionId, groupId, request); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"demoteGroupParticipants",
		[]interface{}{sessionId, groupId, request},
	)
}

func (w *jsiiProxy_WasendClient) GetAllChats(session *string, options *GetChatsOptions) *SdkResponse {
	if err := w.validateGetAllChatsParameters(session, options); err != nil {
		panic(err)
	}
	var returns *SdkResponse

	_jsii_.Invoke(
		w,
		"getAllChats",
		[]interface{}{session, options},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetChatPicture(session *string, chatId *string, options *GetChatPictureOptions) *SdkResponse {
	if err := w.validateGetChatPictureParameters(session, chatId, options); err != nil {
		panic(err)
	}
	var returns *SdkResponse

	_jsii_.Invoke(
		w,
		"getChatPicture",
		[]interface{}{session, chatId, options},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetChatsOverview(session *string, options *GetChatsOverviewOptions) *SdkResponse {
	if err := w.validateGetChatsOverviewParameters(session, options); err != nil {
		panic(err)
	}
	var returns *SdkResponse

	_jsii_.Invoke(
		w,
		"getChatsOverview",
		[]interface{}{session, options},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetContact(sessionId *string, params *GetContactQueryParams) *Contact {
	if err := w.validateGetContactParameters(sessionId, params); err != nil {
		panic(err)
	}
	var returns *Contact

	_jsii_.Invoke(
		w,
		"getContact",
		[]interface{}{sessionId, params},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetContactProfilePicture(sessionId *string, params *GetProfilePictureQueryParams) *ProfilePictureResponse {
	if err := w.validateGetContactProfilePictureParameters(sessionId, params); err != nil {
		panic(err)
	}
	var returns *ProfilePictureResponse

	_jsii_.Invoke(
		w,
		"getContactProfilePicture",
		[]interface{}{sessionId, params},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetContacts(sessionId *string, params *GetContactsQueryParams) *[]*Contact {
	if err := w.validateGetContactsParameters(sessionId, params); err != nil {
		panic(err)
	}
	var returns *[]*Contact

	_jsii_.Invoke(
		w,
		"getContacts",
		[]interface{}{sessionId, params},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetGroup(sessionId *string, groupId *string) *Group {
	if err := w.validateGetGroupParameters(sessionId, groupId); err != nil {
		panic(err)
	}
	var returns *Group

	_jsii_.Invoke(
		w,
		"getGroup",
		[]interface{}{sessionId, groupId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetGroupInfoAdminOnly(sessionId *string, groupId *string) *SettingsSecurityChangeInfo {
	if err := w.validateGetGroupInfoAdminOnlyParameters(sessionId, groupId); err != nil {
		panic(err)
	}
	var returns *SettingsSecurityChangeInfo

	_jsii_.Invoke(
		w,
		"getGroupInfoAdminOnly",
		[]interface{}{sessionId, groupId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetGroupInviteCode(sessionId *string, groupId *string) *string {
	if err := w.validateGetGroupInviteCodeParameters(sessionId, groupId); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getGroupInviteCode",
		[]interface{}{sessionId, groupId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetGroupJoinInfo(sessionId *string, code *string) *Group {
	if err := w.validateGetGroupJoinInfoParameters(sessionId, code); err != nil {
		panic(err)
	}
	var returns *Group

	_jsii_.Invoke(
		w,
		"getGroupJoinInfo",
		[]interface{}{sessionId, code},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetGroupMessagesAdminOnly(sessionId *string, groupId *string) *SettingsSecurityChangeInfo {
	if err := w.validateGetGroupMessagesAdminOnlyParameters(sessionId, groupId); err != nil {
		panic(err)
	}
	var returns *SettingsSecurityChangeInfo

	_jsii_.Invoke(
		w,
		"getGroupMessagesAdminOnly",
		[]interface{}{sessionId, groupId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetGroupParticipants(sessionId *string, groupId *string) *[]*GroupParticipant {
	if err := w.validateGetGroupParticipantsParameters(sessionId, groupId); err != nil {
		panic(err)
	}
	var returns *[]*GroupParticipant

	_jsii_.Invoke(
		w,
		"getGroupParticipants",
		[]interface{}{sessionId, groupId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetGroupPicture(sessionId *string, groupId *string, refresh *bool) *GroupPictureResponse {
	if err := w.validateGetGroupPictureParameters(sessionId, groupId); err != nil {
		panic(err)
	}
	var returns *GroupPictureResponse

	_jsii_.Invoke(
		w,
		"getGroupPicture",
		[]interface{}{sessionId, groupId, refresh},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetGroups(sessionId *string, params *GetGroupsQueryParams) *[]*Group {
	if err := w.validateGetGroupsParameters(sessionId, params); err != nil {
		panic(err)
	}
	var returns *[]*Group

	_jsii_.Invoke(
		w,
		"getGroups",
		[]interface{}{sessionId, params},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetGroupsCount(sessionId *string) *CountResponse {
	if err := w.validateGetGroupsCountParameters(sessionId); err != nil {
		panic(err)
	}
	var returns *CountResponse

	_jsii_.Invoke(
		w,
		"getGroupsCount",
		[]interface{}{sessionId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetMessage(messageId *string) *SdkResponse {
	if err := w.validateGetMessageParameters(messageId); err != nil {
		panic(err)
	}
	var returns *SdkResponse

	_jsii_.Invoke(
		w,
		"getMessage",
		[]interface{}{messageId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetMessageById(session *string, chatId *string, messageId *string, options *GetMessageByIdOptions) *SdkResponse {
	if err := w.validateGetMessageByIdParameters(session, chatId, messageId, options); err != nil {
		panic(err)
	}
	var returns *SdkResponse

	_jsii_.Invoke(
		w,
		"getMessageById",
		[]interface{}{session, chatId, messageId, options},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) GetMessages(session *string, chatId *string, options *GetMessagesOptions) *SdkResponse {
	if err := w.validateGetMessagesParameters(session, chatId, options); err != nil {
		panic(err)
	}
	var returns *SdkResponse

	_jsii_.Invoke(
		w,
		"getMessages",
		[]interface{}{session, chatId, options},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) JoinGroup(sessionId *string, request *JoinGroupRequest) *JoinGroupResponse {
	if err := w.validateJoinGroupParameters(sessionId, request); err != nil {
		panic(err)
	}
	var returns *JoinGroupResponse

	_jsii_.Invoke(
		w,
		"joinGroup",
		[]interface{}{sessionId, request},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) LeaveGroup(sessionId *string, groupId *string) {
	if err := w.validateLeaveGroupParameters(sessionId, groupId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"leaveGroup",
		[]interface{}{sessionId, groupId},
	)
}

func (w *jsiiProxy_WasendClient) ProcessMessage(request *SendRequest, options *map[string]*string) *WAMessage {
	if err := w.validateProcessMessageParameters(request); err != nil {
		panic(err)
	}
	var returns *WAMessage

	_jsii_.Invoke(
		w,
		"processMessage",
		[]interface{}{request, options},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) PromoteGroupParticipants(sessionId *string, groupId *string, request *ParticipantsRequest) {
	if err := w.validatePromoteGroupParticipantsParameters(sessionId, groupId, request); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"promoteGroupParticipants",
		[]interface{}{sessionId, groupId, request},
	)
}

func (w *jsiiProxy_WasendClient) ReadMessages(session *string, chatId *string, options *ReadMessagesOptions) *ReadChatMessagesResponse {
	if err := w.validateReadMessagesParameters(session, chatId, options); err != nil {
		panic(err)
	}
	var returns *ReadChatMessagesResponse

	_jsii_.Invoke(
		w,
		"readMessages",
		[]interface{}{session, chatId, options},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) RemoveGroupParticipants(sessionId *string, groupId *string, request *ParticipantsRequest) {
	if err := w.validateRemoveGroupParticipantsParameters(sessionId, groupId, request); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"removeGroupParticipants",
		[]interface{}{sessionId, groupId, request},
	)
}

func (w *jsiiProxy_WasendClient) RestartSession(sessionId *string) *Session {
	if err := w.validateRestartSessionParameters(sessionId); err != nil {
		panic(err)
	}
	var returns *Session

	_jsii_.Invoke(
		w,
		"restartSession",
		[]interface{}{sessionId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) RetrieveAccount() *SdkResponse {
	var returns *SdkResponse

	_jsii_.Invoke(
		w,
		"retrieveAccount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) RetrieveAllSessions() *GetAllSessionsResponse {
	var returns *GetAllSessionsResponse

	_jsii_.Invoke(
		w,
		"retrieveAllSessions",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) RetrieveConfig() *WasendConfigInfo {
	var returns *WasendConfigInfo

	_jsii_.Invoke(
		w,
		"retrieveConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) RetrieveQRCode(sessionId *string) *QRCodeResponse {
	if err := w.validateRetrieveQRCodeParameters(sessionId); err != nil {
		panic(err)
	}
	var returns *QRCodeResponse

	_jsii_.Invoke(
		w,
		"retrieveQRCode",
		[]interface{}{sessionId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) RetrieveSessionInfo(sessionId *string) *Session {
	if err := w.validateRetrieveSessionInfoParameters(sessionId); err != nil {
		panic(err)
	}
	var returns *Session

	_jsii_.Invoke(
		w,
		"retrieveSessionInfo",
		[]interface{}{sessionId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) RevokeGroupInviteCode(sessionId *string, groupId *string) *string {
	if err := w.validateRevokeGroupInviteCodeParameters(sessionId, groupId); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"revokeGroupInviteCode",
		[]interface{}{sessionId, groupId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) Send(request *SendRequest) *WAMessage {
	if err := w.validateSendParameters(request); err != nil {
		panic(err)
	}
	var returns *WAMessage

	_jsii_.Invoke(
		w,
		"send",
		[]interface{}{request},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) SendFile(request *MessageFileRequest) *WAMessage {
	if err := w.validateSendFileParameters(request); err != nil {
		panic(err)
	}
	var returns *WAMessage

	_jsii_.Invoke(
		w,
		"sendFile",
		[]interface{}{request},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) SendImage(request *MessageImageRequest) *WAMessage {
	if err := w.validateSendImageParameters(request); err != nil {
		panic(err)
	}
	var returns *WAMessage

	_jsii_.Invoke(
		w,
		"sendImage",
		[]interface{}{request},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) SendLinkCustomPreview(request *MessageLinkCustomPreviewRequest) *WAMessage {
	if err := w.validateSendLinkCustomPreviewParameters(request); err != nil {
		panic(err)
	}
	var returns *WAMessage

	_jsii_.Invoke(
		w,
		"sendLinkCustomPreview",
		[]interface{}{request},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) SendMessage(request *MessageRequest) *SdkResponse {
	if err := w.validateSendMessageParameters(request); err != nil {
		panic(err)
	}
	var returns *SdkResponse

	_jsii_.Invoke(
		w,
		"sendMessage",
		[]interface{}{request},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) SendSeen(request *SendSeenRequest) {
	if err := w.validateSendSeenParameters(request); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"sendSeen",
		[]interface{}{request},
	)
}

func (w *jsiiProxy_WasendClient) SendText(request *MessageTextRequest) *WAMessage {
	if err := w.validateSendTextParameters(request); err != nil {
		panic(err)
	}
	var returns *WAMessage

	_jsii_.Invoke(
		w,
		"sendText",
		[]interface{}{request},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) SendVideo(request *MessageVideoRequest) *WAMessage {
	if err := w.validateSendVideoParameters(request); err != nil {
		panic(err)
	}
	var returns *WAMessage

	_jsii_.Invoke(
		w,
		"sendVideo",
		[]interface{}{request},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) SendVoice(request *MessageVoiceRequest) *WAMessage {
	if err := w.validateSendVoiceParameters(request); err != nil {
		panic(err)
	}
	var returns *WAMessage

	_jsii_.Invoke(
		w,
		"sendVoice",
		[]interface{}{request},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) SetGroupDescription(sessionId *string, groupId *string, request *DescriptionRequest) {
	if err := w.validateSetGroupDescriptionParameters(sessionId, groupId, request); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"setGroupDescription",
		[]interface{}{sessionId, groupId, request},
	)
}

func (w *jsiiProxy_WasendClient) SetGroupInfoAdminOnly(sessionId *string, groupId *string, request *SettingsSecurityChangeInfo) {
	if err := w.validateSetGroupInfoAdminOnlyParameters(sessionId, groupId, request); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"setGroupInfoAdminOnly",
		[]interface{}{sessionId, groupId, request},
	)
}

func (w *jsiiProxy_WasendClient) SetGroupMessagesAdminOnly(sessionId *string, groupId *string, request *SettingsSecurityChangeInfo) {
	if err := w.validateSetGroupMessagesAdminOnlyParameters(sessionId, groupId, request); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"setGroupMessagesAdminOnly",
		[]interface{}{sessionId, groupId, request},
	)
}

func (w *jsiiProxy_WasendClient) SetGroupPicture(sessionId *string, groupId *string, request *ProfilePictureRequest) {
	if err := w.validateSetGroupPictureParameters(sessionId, groupId, request); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"setGroupPicture",
		[]interface{}{sessionId, groupId, request},
	)
}

func (w *jsiiProxy_WasendClient) SetGroupSubject(sessionId *string, groupId *string, request *SubjectRequest) {
	if err := w.validateSetGroupSubjectParameters(sessionId, groupId, request); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"setGroupSubject",
		[]interface{}{sessionId, groupId, request},
	)
}

func (w *jsiiProxy_WasendClient) StartSession(sessionId *string) *Session {
	if err := w.validateStartSessionParameters(sessionId); err != nil {
		panic(err)
	}
	var returns *Session

	_jsii_.Invoke(
		w,
		"startSession",
		[]interface{}{sessionId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) StartTyping(request *ChatRequest) {
	if err := w.validateStartTypingParameters(request); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"startTyping",
		[]interface{}{request},
	)
}

func (w *jsiiProxy_WasendClient) StopSession(sessionId *string) *Session {
	if err := w.validateStopSessionParameters(sessionId); err != nil {
		panic(err)
	}
	var returns *Session

	_jsii_.Invoke(
		w,
		"stopSession",
		[]interface{}{sessionId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) StopTyping(request *ChatRequest) {
	if err := w.validateStopTypingParameters(request); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"stopTyping",
		[]interface{}{request},
	)
}

