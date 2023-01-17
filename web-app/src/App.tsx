import { RouterProvider } from "react-router-dom";
import { CssBaseline } from "@mui/material";
import { Provider } from "react-redux";
import store from "@store/index";
import router from "./router";
import { useEffect } from "react";
import useAuth from "@hooks/useAuth";
import apiSdk from "@lib/apiSdk";

const Connector = () => {
  const { token } = useAuth();
  useEffect(() => {
    apiSdk.setAuthorization(token);
  }, [token]);
  return null;
};

const App = () => {
  return (
    <Provider store={store}>
      <Connector />
      <CssBaseline />
      <RouterProvider router={router} />
    </Provider>
  );
};

export default App;
