import { shallowEqual, useDispatch, useSelector } from "react-redux";
import { AppDispatch, RootState } from "@store/index";
import { useCallback } from "react";
import { login as loginAction, register as registerAction, logout as logoutAction } from "@store/slices/auth";
import { LoginParams, RegisterParams } from "@sdk/ApiSdk";

const useAuth = () => {
  const dispatch: AppDispatch = useDispatch();
  const authState = useSelector((state: RootState) => state.auth, shallowEqual);
  const login = useCallback((params: LoginParams) => dispatch(loginAction(params)).unwrap(), [dispatch]);
  const register = useCallback((params: RegisterParams) => dispatch(registerAction(params)).unwrap(), [dispatch]);
  const logout = useCallback(() => dispatch(logoutAction()), [dispatch]);
  return { ...authState, login, register, logout };
};

export default useAuth;
