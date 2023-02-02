-- Users table
CREATE TABLE "user"
(
    "userID"    SERIAL PRIMARY KEY,
    "name"      TEXT      NOT NULL,
    "username"  TEXT      NOT NULL unique,
    "password"  BYTEA     NOT NULL,
    "createdAt" TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Messages table
CREATE TABLE "message"
(
    "messageID"   SERIAL PRIMARY KEY,
    "ownerUserID" INT       NOT NULL REFERENCES "user" ("userID"),
    "withUserID"  INT       NOT NULL REFERENCES "user" ("userID"),
    "fromUserID"  INT       NOT NULL REFERENCES "user" ("userID"),
    "toUserID"    INT       NOT NULL REFERENCES "user" ("userID"),
    "key"         CHAR(21)  NOT NULL,
    "content"     TEXT      NOT NULL,
    "createdAt"   TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE ("ownerUserID", "key")
);

-- Conversations table
CREATE TABLE "conversation"
(
    "conversationID"        SERIAL PRIMARY KEY,
    "ownerUserID"           INT       NOT NULL REFERENCES "user" ("userID"),
    "withUserID"            INT       NOT NULL REFERENCES "user" ("userID"),
    "key"                   CHAR(21)  NOT NULL,
    "lastMessageFromUserID" INT REFERENCES "user" ("userID"),
    "lastMessageContent"    TEXT,
    "createdAt"             TIMESTAMP NOT NULL DEFAULT NOW(),
    "updatedAt"             TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE ("ownerUserID", "key")
);

ALTER TABLE "conversation"
    ADD CONSTRAINT "conversation_ownerUserID_withUserID_key" UNIQUE ("ownerUserID", "withUserID");