import { login, logout } from "../features/auth/authSlice";
import { feelmeAPI } from "../services/feelme_api";
import type { AppStore, RootState, AppDispatch } from "./store";

export const RefreshTokenMiddleware: any =
  (store: AppStore) => (next: any) => async (action: any) => {
    if (action.payload && action.payload.status === 401) {
      const refreshToken = store.getState().auth.refreshToken;
      console.log(refreshToken);
      try {
        const token = await store.dispatch(feelmeAPI.endpoints.refreshToken.initiate(refreshToken)).unwrap()
        await store.dispatch(
          login({
            accessToken: token?.accessToken,
            refreshToken: token?.refreshToken,
          })
        );
        console.log("from middleware");
        // Retry the original action
        return next(action);
      } catch (error) {
        console.log(error);
        await store.dispatch(feelmeAPI.endpoints.logout.initiate(null));
        await store.dispatch(logout());
      }
    }

    return next(action);
  };
