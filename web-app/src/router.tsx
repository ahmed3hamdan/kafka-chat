import { createBrowserRouter } from "react-router-dom";
import LoginView from "@views/LoginView";
import MainView from "@views/MainView";
import ConversationView from "@views/ConversationView";

const router = createBrowserRouter([
  {
    path: "/login",
    element: <LoginView />,
  },
  {
    path: "/",
    element: <MainView />,
    children: [
      {
        path: ":userID",
        element: <ConversationView />,
      },
    ],
  },
]);

export default router;
