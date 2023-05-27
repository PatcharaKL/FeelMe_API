import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import type { RootState } from "../../app/store";

// Define a type for the slice state
interface AuthState {
  authenticated?: boolean;
  accessToken: string;
  refreshToken: string;
}

// Define the initial state using that type
const initialState: AuthState = {
  authenticated: false,
  accessToken: "",
  refreshToken: "",
};

export const authSlice = createSlice({
  name: "auth",
  // `createSlice` will infer the state type from the `initialState` argument
  initialState,
  reducers: {
    login: (state, action: PayloadAction<AuthState>) => {
      state.authenticated = true;
      state.accessToken = action.payload.accessToken;
      state.refreshToken = action.payload.refreshToken;
      localStorage.setItem("refreshToken", action.payload.refreshToken);
      localStorage.setItem("accessToken", action.payload.accessToken);
    },
    logout: (state) => {
      localStorage.removeItem("refreshToken");
      localStorage.removeItem("accessToken");
      localStorage.setItem("sideBarSelectedID", "1");
      state.authenticated = false;
      state.accessToken = "";
      state.refreshToken = "";
      localStorage.removeItem("userID");
    },
  },
});

export const { login, logout } = authSlice.actions;

// Other code such as selectors can use the imported `RootState` type
export const selectSidebarSelectedItemID = (state: RootState) =>
  state.sidebarSelection.selectedItemID;

export default authSlice.reducer;
