import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import apiSdk from "@lib/apiSdk";
import { LoginParams, RegisterParams } from "@sdk/ApiSdk";
import serializeError from "@store/utils/serializeError";

export const login = createAsyncThunk("auth/login", (params: LoginParams, { rejectWithValue }) => {
  return apiSdk.login(params).catch(error => rejectWithValue(serializeError(error)));
});

export const register = createAsyncThunk("auth/register", (params: RegisterParams, { rejectWithValue }) => {
  return apiSdk.register(params).catch(error => rejectWithValue(serializeError(error)));
});

type AuthState = { loggedIn: false; userID: null; token: null } | { loggedIn: true; userID: number; token: string };

const initialState: AuthState = {
  loggedIn: false,
  userID: null,
  token: null,
};

const authSlice = createSlice<AuthState, { logout: () => Extract<AuthState, { loggedIn: false }> }, "auth">({
  name: "auth",
  initialState,
  reducers: {
    logout: () => initialState,
  },
  extraReducers: builder =>
    builder
      .addCase(login.fulfilled, (state, action) => {
        const { userID, token } = action.payload;
        return {
          loggedIn: true,
          userID,
          token,
        };
      })
      .addCase(register.fulfilled, (state, action) => {
        const { userID, token } = action.payload;
        return {
          loggedIn: true,
          userID,
          token,
        };
      }),
});

export const { logout } = authSlice.actions;

export default authSlice.reducer;
