import { combineReducers, configureStore } from "@reduxjs/toolkit";
import type { PreloadedState } from "@reduxjs/toolkit";

import { feelmeAPI } from "../services/feelme_api";

import sidebarSelectionReducer from "../features/sidebar-selection/sidebarSelectionSlice";
import authReducer from "./../features/auth/authSlice";
import userReducer from "./../features/auth/userSlice";
import { RefreshTokenMiddleware } from "./middleware";

// Create the root reducer separately so we can extract the RootState type
const rootReducer = combineReducers({
  sidebarSelection: sidebarSelectionReducer,
  auth: authReducer,
  user: userReducer,
  [feelmeAPI.reducerPath]: feelmeAPI.reducer,
});
export const setupStore = (preloadedState?: PreloadedState<RootState>) => {
  return configureStore({
    reducer: rootReducer,
    preloadedState,
    middleware: (getDefaultMiddleware) =>
      getDefaultMiddleware().concat(
        feelmeAPI.middleware
        // RefreshTokenMiddleware
      ),
    devTools: true,
  });
};

export type RootState = ReturnType<typeof rootReducer>;
export type AppStore = ReturnType<typeof setupStore>;
export type AppDispatch = AppStore["dispatch"];
