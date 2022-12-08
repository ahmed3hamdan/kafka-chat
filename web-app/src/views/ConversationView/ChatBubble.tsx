import { Message, User } from "@sdk/types";
import React, { HTMLAttributes } from "react";
import { Avatar, css } from "@mui/material";

const rootCss = css`
  display: flex;
  column-gap: 8px;
`;

const rootRightCss = css`
  flex-direction: row-reverse;
`;

const contentBoxCss = css`
  border-radius: 20px;
  padding: 8px 16px;
`;

const contentBoxRightCss = css`
  color: white;
  background-color: #3f51b5;
`;

const contentBoxLeftCss = css`
  background-color: #f5f5f5;
`;

interface MessageItemProps extends HTMLAttributes<HTMLDivElement> {
  message: Message;
}

const ChatBubble: React.FC<MessageItemProps> = ({ message, ...rest }) => {
  const self: User = {
    userID: "user-0",
    name: "Ahmad Hamdan",
    username: "ahmadhamdan",
  };
  const isSelf = message.userID === self.userID;
  return (
    <div css={[rootCss, isSelf && rootRightCss]} {...rest}>
      {!isSelf && <Avatar alt="Ahmad Hamdan">A</Avatar>}
      <div
        css={[contentBoxCss, isSelf ? contentBoxRightCss : contentBoxLeftCss]}
      >
        {message.content}
      </div>
    </div>
  );
};

export default ChatBubble;
