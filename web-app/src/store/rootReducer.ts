import { combineReducers } from "@reduxjs/toolkit";
import authReducer from "./slices/auth";
import selfInfoReducer from "./slices/selfInfo";

const rootReducer = combineReducers({
  auth: authReducer,
  selfInfo: selfInfoReducer,
});

export default rootReducer;
