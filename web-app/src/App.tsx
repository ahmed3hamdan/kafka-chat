import { useMemo, useState } from "react";
import { RouterProvider } from "react-router-dom";
import { CssBaseline } from "@mui/material";
import router from "./router";
import ApiSdk from "@sdk/ApiSdk";
import { API_BASE_URL } from "@config/index";
import appContext, { AppContextAuth } from "@contexts/appContext";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const queryClient = new QueryClient();

const App = () => {
  const [auth, setAuth] = useState<AppContextAuth | null>(null);

  const api = useMemo(
    () =>
      new ApiSdk({
        baseURL: API_BASE_URL,
        authToken: auth === null ? undefined : auth.token,
      }),
    [auth]
  );

  return (
    <appContext.Provider value={{ auth, setAuth, api }}>
      <QueryClientProvider client={queryClient}>
        <CssBaseline />
        <RouterProvider router={router} />
      </QueryClientProvider>
    </appContext.Provider>
  );
};

export default App;
