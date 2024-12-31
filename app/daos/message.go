package daos

type MessageDAO struct {
    ID             string `bson:"_id,omitempty"`
    MessageContent string `bson:"message_content"`
    RecipientPhone string `bson:"recipient_phone_number"`
    SendingStatus  string `bson:"sending_status"`
}
