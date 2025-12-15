package repository

import (
	"AgenticDikti/internal/model"
	"context"
)

const (
	selectChatBySessionidQuery = `SELECT s.role, s.chatinput FROM chat_logs s WHERE s.sessionid = $1 ORDER BY timestamp DESC LIMIT 10`
	insertChatQuery            = `INSERT INTO chat_logs (sessionid, chatid, chatinput, timestamp, role, emergency, universityid) VALUES ($1, $2, $3, $4, $5, $6, $7)`
)

func (q *Queries) SelectChatBySessionid(ctx context.Context, sessionId string) (res []model.ChatHistory, err error) {
	rows, err := q.db.QueryContext(ctx, selectChatBySessionidQuery, sessionId)

	// Scan(&res.Role, &res.ChatInput)
	if err != nil {
		return []model.ChatHistory{}, err
	}
	defer rows.Close()

	var chats []model.ChatHistory

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var chat model.ChatHistory
		if err := rows.Scan(&chat.Role, &chat.ChatInput); err != nil {
			return chats, err
		}
		chats = append(chats, chat)
	}
	if err = rows.Err(); err != nil {
		return chats, err
	}
	return chats, nil
}

func (q *Queries) InsertChat(ctx context.Context, userLog model.ChatLogs, aiLog model.ChatLogs) (err error) {
	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Rollback on panic or error
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Insert user chat
	_, err = tx.ExecContext(
		ctx,
		insertChatQuery,
		userLog.SessionID,
		userLog.ChatID,
		userLog.ChatInput,
		userLog.Timestamp,
		userLog.Role,
		userLog.Emergency,
		userLog.UniversityID,
	)
	if err != nil {
		return err
	}

	// Insert AI chat
	_, err = tx.ExecContext(
		ctx,
		insertChatQuery,
		aiLog.SessionID,
		aiLog.ChatID,
		aiLog.ChatInput,
		aiLog.Timestamp,
		aiLog.Role,
		aiLog.Emergency,
		aiLog.UniversityID,
	)
	if err != nil {
		return err
	}

	return nil
}
