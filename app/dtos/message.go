package dtos

type MessageDTO struct {
    MessageContent string `json:"message_content"`
    RecipientPhone string `json:"recipient_phone_number"`
    SendingStatus  string `json:"sending_status"`
}
