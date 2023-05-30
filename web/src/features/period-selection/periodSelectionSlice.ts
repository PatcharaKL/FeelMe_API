import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import type { RootState } from "../../app/store";

// Define a type for the slice state
interface sidebarSelectionState {
  period: string;
}

// Define the initial state using that type
const initialState: sidebarSelectionState = {
  period: "month",
};

export const periodSelectionSlice = createSlice({
  name: "periodSelection",
  // `createSlice` will infer the state type from the `initialState` argument
  initialState,
  reducers: {
    periodAll: (state) => {
      state.period = "";
    },
    periodYear: (state) => {
      state.period = "year";
    },
    periodMonth: (state) => {
      state.period = "month";
    },
  },
});

export const { periodAll, periodYear, periodMonth } =
  periodSelectionSlice.actions;

// Other code such as selectors can use the imported `RootState` type
export const selectPeriodSelected = (state: RootState) => state.period.period;

export default periodSelectionSlice.reducer;
