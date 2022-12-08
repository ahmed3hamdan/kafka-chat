import React from "react";
import {
  AppBar,
  css,
  IconButton,
  InputAdornment,
  OutlinedInput,
  Toolbar
} from "@mui/material";
import { Send as SendIcon } from "@mui/icons-material";
import { Virtuoso } from "react-virtuoso";
import { Message } from "@sdk/types";
import { Field, FieldProps, Form, Formik } from "formik";
import ChatBubble from "./ChatBubble";

const rootCss = css`
  height: 100%;
  display: flex;
  flex-direction: column;
`;

const messagesVirtuosoCss = css`
  flex-grow: 1;
`;

const messagesVirtuosoHeaderCss = css`
  height: 24px;
`;

const messageFormCss = css`
  padding: 24px;
`;

const messagesVirtuosoItemCss = css`
  padding: 2px 24px;
`;

const messages: Message[] = [
  {
    userID: "user-0",
    content: "Hello 1"
  },
  {
    userID: "user-1",
    content: "Hello 2"
  },
  {
    userID: "user-0",
    content: "Hello 3"
  },
  {
    userID: "user-0",
    content: "Hello 1"
  },
  {
    userID: "user-1",
    content: "Hello 1"
  },
  {
    userID: "user-1",
    content: "Hello 1"
  }
];

interface MessageFormValues {
  message: string;
}

const initialValues: MessageFormValues = {
  message: ""
};

const MessagesVirtuosoHeader: React.FC = () => <div css={messagesVirtuosoHeaderCss} />;

const MessagesVirtuosoItem: React.FC<React.HTMLAttributes<HTMLDivElement>> = props => <div css={messagesVirtuosoItemCss} {...props} />;

const Index = () => {
  const handleSubmit = async ({ message }: MessageFormValues) => {
    console.log(message);
  };
  return (
    <div css={rootCss}>
      <AppBar position="static">
        <Toolbar>Ahmad Hamdan</Toolbar>
      </AppBar>
      <Virtuoso
        css={messagesVirtuosoCss}
        data={messages}
        alignToBottom
        components={{
          Header: MessagesVirtuosoHeader,
          Item: MessagesVirtuosoItem,
        }}
        itemContent={(index, data) => <ChatBubble message={data} />}
      />
      <Formik<MessageFormValues>
        initialValues={initialValues}
        onSubmit={handleSubmit}
      >
        <Form autoComplete="off" noValidate css={messageFormCss}>
          <Field name="message">
            {({ field }: FieldProps<string>) => (
              <OutlinedInput
                {...field}
                placeholder="Type a message..."
                multiline
                maxRows={4}
                fullWidth
                autoFocus
                endAdornment={
                  field.value.trim().length > 0 && (
                    <InputAdornment position="end">
                      <IconButton type="submit" edge="end">
                        <SendIcon />
                      </IconButton>
                    </InputAdornment>
                  )
                }
              />
            )}
          </Field>
        </Form>
      </Formik>
    </div>
  );
};

export default Index;
