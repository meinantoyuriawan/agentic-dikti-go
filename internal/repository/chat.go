package repository

import (
	"AgenticDikti/internal/model"
	"context"
)

const (
	selectChatBySessionidQuery = `SELECT s.role, s.chatinput FROM chat_logs s WHERE s.sessionid = $1 ORDER BY timestamp DESC LIMIT 10`
	insertChatQuery            = `INSERT INTO chat_logs (sessionid, chatid, chatinput, timestamp, role, emergency, universityid) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	insertBookingData          = `INSERT INTO counseling_bookings (nama, nim, schedule, universityid) VALUES ($1, $2, $3, $4)`
)

func (q *Queries) SelectChatBySessionid(ctx context.Context, sessionId string) (res model.ChatHistory, err error) {
	err = q.db.QueryRowContext(ctx, selectChatBySessionidQuery, sessionId).Scan(&res.Role, &res.ChatInput)
	if err != nil {
		return model.ChatHistory{}, err
	}
	return res, nil
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

func (q *Queries) InsertBooking(ctx context.Context, userBookData model.BookingData) (err error) {
	err = q.db.QueryRowContext(ctx, insertBookingData,
		userBookData.Nama,
		userBookData.Nim,
		userBookData.Schedule,
		userBookData.UniversityID,
	).Scan()

	if err != nil {
		return err
	}
	return nil
}
