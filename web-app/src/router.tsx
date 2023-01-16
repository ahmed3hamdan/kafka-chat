import { createBrowserRouter } from "react-router-dom";
import LoginView from "@views/LoginView";
import RegisterView from "@views/RegisterView";
import MainView from "@views/MainView";
import ConversationView from "@views/ConversationView";
import AuthView from "@views/ConversationView/AuthView";

const router = createBrowserRouter([
  {
    element: <AuthView />,
    children: [
      {
        path: "/login",
        element: <LoginView />,
      },
      {
        path: "/register",
        element: <RegisterView />,
      },
    ],
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
