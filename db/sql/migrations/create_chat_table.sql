CREATE TABLE chat_logs (
    chatid        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sessionid      UUID NOT NULL,
    role            VARCHAR(10) NOT NULL,        -- e.g., 'user', 'assistant', etc.
    chatinput      TEXT NOT NULL,
    timestamp      TIMESTAMP NOT NULL,
    emergency    BOOLEAN NOT NULL DEFAULT FALSE,
    universityid   INTEGER NOT NULL
);

-- Index 1: Fast retrieval by session (most common query: get chat history for a session)
CREATE INDEX idx_chat_messages_session
    ON chat_logs (sessionid);
