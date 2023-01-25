-- Users table
CREATE TABLE "user"
(
    "userID"    SERIAL PRIMARY KEY,
    "name"      TEXT  NOT NULL,
    "username"  TEXT  NOT NULL unique,
    "password"  BYTEA NOT NULL,
    "createdAt" DATE  NOT NULL DEFAULT NOW()
);

-- Messages table
CREATE TABLE "message"
(
    "messageID"   SERIAL PRIMARY KEY,
    "ownerUserID" INT  NOT NULL references "user" ("userID"),
    "fromUserID"  INT  NOT NULL references "user" ("userID"),
    "toUserID"    INT  NOT NULL references "user" ("userID"),
    "key"         TEXT NOT NULL unique,
    "content"     TEXT NOT NULL,
    "createdAt"   DATE NOT NULL DEFAULT NOW()
);

-- Conversations table
CREATE TABLE "conversation"
(
    "conversationID" SERIAL PRIMARY KEY,
    "ownerUserID"    INT  NOT NULL references "user" ("userID"),
    "withUserID"     INT  NOT NULL references "user" ("userID"),
    "latestMessage"  INT  NOT NULL references "message" ("messageID"),
    "createdAt"      DATE NOT NULL DEFAULT NOW(),
    "updatedAt"      DATE NOT NULL DEFAULT NOW(),
    UNIQUE ("ownerUserID", "withUserID")
);