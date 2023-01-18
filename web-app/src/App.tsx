import { RouterProvider } from "react-router-dom";
import { CssBaseline } from "@mui/material";
import { Provider } from "react-redux";
import store, { persistor } from "@store/index";
import router from "./router";
import { useEffect } from "react";
import useAuth from "@hooks/useAuth";
import apiSdk from "@lib/apiSdk";
import { PersistGate } from "redux-persist/integration/react";

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
      <PersistGate loading={null} persistor={persistor}>
        <Connector />
        <CssBaseline />
        <RouterProvider router={router} />
      </PersistGate>
    </Provider>
  );
};

export default App;
