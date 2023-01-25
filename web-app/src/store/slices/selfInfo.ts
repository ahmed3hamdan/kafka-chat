import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import storage from "redux-persist/lib/storage";
import { persistReducer } from "redux-persist";
import apiSdk from "@lib/apiSdk";
import serializeError from "@store/utils/serializeError";

export const getSelfInfo = createAsyncThunk("selfInfo/get", (arg, { rejectWithValue }) => {
  return apiSdk.getSelfInfo().catch(error => rejectWithValue(serializeError(error)));
});

interface SelfInfoState {
  status: "idle" | "loading" | "error" | "success";
  lastFetchedAt: number;
  data: {
    userID: number;
    name: string;
    username: string;
  };
}

const initialState: SelfInfoState = {
  status: "idle",
  lastFetchedAt: 0,
  data: {
    userID: 0,
    name: "",
    username: "",
  },
};

const selfInfoSlice = createSlice({
  name: "selfInfo",
  initialState,
  reducers: {
    clearSelfInfo: () => initialState,
  },
  extraReducers: builder =>
    builder
      .addCase(getSelfInfo.pending, state => {
        state.status = "loading";
      })
      .addCase(getSelfInfo.rejected, state => {
        state.status = "error";
      })
      .addCase(getSelfInfo.fulfilled, (state, action) => {
        state.status = "success";
        state.lastFetchedAt = Date.now();
        state.data = action.payload;
      }),
});

export const { clearSelfInfo } = selfInfoSlice.actions;

const persistConfig = {
  key: "selfInfo",
  version: 1,
  blacklist: ["status"],
  storage,
};

export default persistReducer(persistConfig, selfInfoSlice.reducer);
