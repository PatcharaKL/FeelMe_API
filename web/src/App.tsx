import React from "react";
import FeelMe from "./pages/FeelMe";
import { LoginPage } from "./pages/Login";
import { useAppDispatch, useAppSelector } from "./app/hooks";
import { login } from "./features/auth/authSlice";
import { setSelectedItem } from "./features/sidebar-selection/sidebarSelectionSlice";
import { setID } from "./features/auth/userSlice";

function App() {
  const auth = useAppSelector((state) => state.auth.authenticated);
  const dispatch = useAppDispatch();
  const accessToken = localStorage.getItem("accessToken");
  const refreshToken = localStorage.getItem("refreshToken");
  if (accessToken && refreshToken) {
    dispatch(login({ accessToken, refreshToken }));
    dispatch(setSelectedItem(Number(localStorage.getItem("sideBarSelectedID"))));
    dispatch(setID(Number(localStorage.getItem('userID'))))
  }
  const Display = () => {
    if (auth) {
      return <FeelMe />;
    } else {
      return <LoginPage />;
    }
  };
  return (
    <React.Fragment>
      <Display />
    </React.Fragment>
  );
}

export default App;
