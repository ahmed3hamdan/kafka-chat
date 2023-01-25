import React, { useCallback, useEffect } from "react";
import { RouterProvider } from "react-router-dom";
import { Box, CircularProgress, CssBaseline } from "@mui/material";
import { Provider } from "react-redux";
import { PersistGate } from "redux-persist/integration/react";
import store, { persistor } from "@store/index";
import apiSdk from "@lib/apiSdk";
import useAuth from "@hooks/useAuth";
import useSelfInfo from "@hooks/useSelfInfo";
import router from "./router";

const Connector: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const { token } = useAuth();
  const selfInfoSlice = useSelfInfo();

  useEffect(() => {
    apiSdk.setAuthorization(token);
    if (token === null) {
      selfInfoSlice.clear();
      return;
    }
    if (selfInfoSlice.status === "idle" && selfInfoSlice.lastFetchedAt === 0) {
      selfInfoSlice.get();
    }
  }, [token, selfInfoSlice]);

  const handleTryAgainClick = useCallback(() => {
    selfInfoSlice.status === "error" && selfInfoSlice.lastFetchedAt === 0 && selfInfoSlice.get();
  }, [selfInfoSlice]);

  if (selfInfoSlice.status === "loading" && selfInfoSlice.lastFetchedAt === 0) {
    return (
      <Box sx={{ display: "flex", justifyContent: "center", alignItems: "center", height: "100%" }}>
        <CircularProgress />
      </Box>
    );
  }

  if (selfInfoSlice.status === "error") {
    return (
      <>
        <div>An error occurred while fetching initial data</div>
        <button onClick={handleTryAgainClick}>Try Again</button>
      </>
    );
  }

  return <>{children}</>;
};

const App = () => {
  return (
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistor}>
        <CssBaseline />
        <Connector>
          <RouterProvider router={router} />
        </Connector>
      </PersistGate>
    </Provider>
  );
};

export default App;
