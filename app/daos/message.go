package daos

import "github.com/myKemal/insiderGo/app/dtos"

type MessageDAO struct {
	ID             string `bson:"_id,omitempty"`
	MessageContent string `bson:"message_content"`
	RecipientPhone string `bson:"recipient_phone_number"`
	SendingStatus  string `bson:"sending_status"`
}

func ConvertDAOToDTO(dao MessageDAO) dtos.MessageDTO {
	return dtos.MessageDTO{
		MessageContent: dao.MessageContent,
		RecipientPhone: dao.RecipientPhone,
		SendingStatus:  dao.SendingStatus,
	}
}

func ConvertDAOsToDTOs(daos []MessageDAO) []dtos.MessageDTO {
	dtos := make([]dtos.MessageDTO, len(daos))
	for i, dao := range daos {
		dtos[i] = ConvertDAOToDTO(dao)
	}
	return dtos
}
