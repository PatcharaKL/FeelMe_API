import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import type { RootState } from "../../app/store";

// Define a type for the slice state
interface UserState {
  id: number | null;
  name: string | null;
  position: string | null;
}

// Define the initial state using that type
const initialState: UserState = {
  id: null,
  name: null,
  position: null,
};

export const userSlice = createSlice({
  name: "user",
  // `createSlice` will infer the state type from the `initialState` argument
  initialState,
  reducers: {
    setID: (state, action: PayloadAction<number>) => {
      state.id = action.payload;
      localStorage.setItem("userID", String(action.payload));
    },
    setUser: (state, action: PayloadAction<UserState>) => {
      state.name = action.payload.name;
      state.position = action.payload.position;
    },
    clearUser: (state) => {
      state.name = "";
      state.position = "";
    },
  },
});

export const { setUser, clearUser, setID } = userSlice.actions;

// Other code such as selectors can use the imported `RootState` type
export const selectSidebarSelectedItemID = (state: RootState) =>
  state.sidebarSelection.selectedItemID;

export default userSlice.reducer;
