-- Enable UUIDs
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table
CREATE TABLE "user"
(
    "userID"    UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "name"      TEXT             NOT NULL unique,
    "username"  TEXT             NOT NULL,
    "password"  BYTEA            NOT NULL,
    "createdAt" DATE             NOT NULL DEFAULT NOW()
);

-- Messages table
CREATE TABLE "message"
(
    "messageID"  UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "fromUserID" UUID             NOT NULL references "user" ("userID"),
    "toUserID"   UUID             NOT NULL references "user" ("userID"),
    "content"    TEXT             NOT NULL,
    "createdAt"  DATE             NOT NULL DEFAULT NOW()
);

-- Conversations table
CREATE TABLE "conversation"
(
    "userID"        UUID NOT NULL references "user" ("userID"),
    "withUserID"    UUID NOT NULL references "user" ("userID"),
    "latestMessage" UUID NOT NULL references "user" ("userID"),
    "createdAt"     DATE NOT NULL DEFAULT NOW(),
    "updatedAt"     DATE NOT NULL DEFAULT NOW()
);