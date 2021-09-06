package database

import (
	"chat/models"
	"errors"
)

func GetLastMessageID() (int, error) {
	var id int
	err := postgres.Get(&id, "SELECT last_value FROM messages_id_seq")
	return id, err
}

func InsertMessageGetID(message string, senderID int) (int, error) {
	var messageID models.ID
	err := postgres.Get(&messageID, "INSERT INTO messages(message, sender_id) VALUES($1, $2) RETURNING id", message, senderID)
	return messageID.ID, err
}

func GetMessagesDownFrom(downFrom int, limit int) (models.MessageRes, error) {
	messages := models.MessageRes{}
	if limit > 100 {
		return messages, errors.New("max amount of retrieving messages is 100")
	}

	err := postgres.Select(&messages.Messages, "SELECT messages.id, messages.message, users.username, to_char(messages.sent_at, 'YYYY-MM-DD HH:MM:SS') AS sent_at, users.password IS NOT NULL AS is_registered FROM (SELECT * FROM messages WHERE NOT deleted AND id < $1 ORDER BY id DESC LIMIT $2) AS messages INNER JOIN users ON messages.sender_id = users.id ORDER BY messages.id ASC", downFrom, limit)
	return messages, err
}

func GetMessagesUpFrom(upFrom int, limit int) (models.MessageRes, error) {
	messages := models.MessageRes{}
	if limit > 100 {
		return messages, errors.New("max amount of retrieving messages is 100")
	}

	err := postgres.Select(&messages.Messages, "SELECT messages.id, messages.message, users.username, to_char(messages.sent_at, 'YYYY-MM-DD HH:MM:SS') AS sent_at, users.password IS NOT NULL AS is_registered FROM (SELECT * FROM messages WHERE NOT deleted AND id >= $1 ORDER BY id DESC LIMIT $2) AS messages INNER JOIN users ON messages.sender_id = users.id ORDER BY messages.id ASC", upFrom, limit)
	return messages, err
}

func GetLastMessages(limit int) (models.MessageRes, error) {
	messages := models.MessageRes{}
	if limit > 100 {
		return messages, errors.New("max amount of retrieving messages is 100")
	}

	err := postgres.Select(&messages.Messages, "SELECT messages.id, messages.message, users.username, to_char(messages.sent_at, 'YYYY-MM-DD HH:MM:SS') AS sent_at, users.password IS NOT NULL AS is_registered FROM (SELECT * FROM messages WHERE NOT deleted ORDER BY id DESC LIMIT $1) AS messages INNER JOIN users ON messages.sender_id = users.id ORDER by id ASC", limit)
	return messages, err
}

func DeleteMessage(messageID int, userID int) error {
	res, err := postgres.Exec("UPDATE messages SET deleted=true WHERE id=$1 AND sender_id=$2", messageID, userID)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows == 0 {
		return errors.New("did not affect rows")
	}
	return err
}
