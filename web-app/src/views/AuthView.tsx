import { Navigate, Outlet } from "react-router-dom";
import useAuth from "@hooks/useAuth";

const AuthView = () => {
  const { loggedIn } = useAuth();
  if (loggedIn) {
    return <Navigate to="/" replace />;
  }
  return <Outlet />;
};

export default AuthView;
