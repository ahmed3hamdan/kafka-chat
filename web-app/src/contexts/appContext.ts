import ApiSdk from "@sdk/ApiSdk";
import { createContext, useContext } from "react";
import { API_BASE_URL } from "@config/index";

export interface AppContextAuth {
  userID: string;
  token: string;
}

interface AppContext {
  auth: AppContextAuth | null;
  setAuth: (value: AppContextAuth | null) => void;
  api: ApiSdk;
}

const appContext = createContext<AppContext>({
  auth: null,
  setAuth: () => null,
  api: new ApiSdk({ baseURL: API_BASE_URL }),
});

export const useAppContext = () => useContext(appContext);

export default appContext;
