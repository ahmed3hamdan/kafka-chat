export interface Message {
  userID: string;
  content: string;
}

export interface User {
  userID: string;
  name: string;
  username: string;
}

export interface SingleConversation extends User {
  lastMessage: Message;
}
