import { Navigate, Outlet } from "react-router-dom";
import { useAppContext } from "@contexts/appContext";

const AuthView = () => {
  const { auth } = useAppContext();
  if (auth !== null) {
    return <Navigate to="/" replace />;
  }
  return <Outlet />;
};

export default AuthView;
