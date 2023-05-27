// Need to use the React-specific entry point to import createApi
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { login, logout } from "../features/auth/authSlice";
import type {
  BaseQueryFn,
  FetchArgs,
  FetchBaseQueryError,
} from "@reduxjs/toolkit/query";

const baseQuery = fetchBaseQuery({
  baseUrl: import.meta.env.VITE_BASE_URL,
  prepareHeaders: (headers, { getState }: any) => {
    const accessToken = getState().auth.accessToken;
    headers.set("Authorization", `Bearer ${accessToken}`);
    return headers;
  },
});

const baseQueryWithReAuth: BaseQueryFn<
  string | FetchArgs,
  unknown,
  FetchBaseQueryError
> = async (args, api, extraOptions) => {
  let result = await baseQuery(args, api, extraOptions);
  if (result.error && result.error.status === 401) {
    // try to get a new token
    const refreshToken = localStorage.getItem("refreshToken");
    const refreshResult: any = await baseQuery(
      {
        url: "/newtoken",
        body: { refreshToken: refreshToken },
        method: "POST",
      },
      api,
      extraOptions
    );
    if (refreshResult.data) {
      // store the new token
      api.dispatch(
        login({
          accessToken: refreshResult.data.accessToken,
          refreshToken: refreshResult.data.refreshToken,
        })
      );
      // store token in local storage
      // retry the initial query
      result = await baseQuery(args, api, extraOptions);
    } else {
      api.dispatch(logout());
    }
  }
  return result;
};

// Define a service using a base URL and expected endpoints
export const feelmeAPI = createApi({
  reducerPath: "feelmeAPI",
  baseQuery: baseQueryWithReAuth,
  endpoints: (builder) => ({
    getHealthCheck: builder.query({
      query: () => `/health-check`,
    }),
    getHappinessPoints: builder.query({
      query: (id: number) => ({
        url: `happiness-score-all-time?account-id=${id}`,
        method: "GET",
      }),
    }),
    getEmployees: builder.query({
      query: () => `/users/employees/`,
    }),
    login: builder.mutation({
      query: (credential) => ({
        url: "/login",
        method: "POST",
        body: credential,
      }),
    }),
    logout: builder.mutation({
      query: () => ({
        url: "/logout",
        method: "POST",
      }),
    }),
    getEmployee: builder.query({
      query: (id) => `/users/employees/?accountId=${id}`,
    }),
    refreshToken: builder.mutation({
      query: (refreshToken) => ({
        url: "/newtoken",
        method: "POST",
        body: { refreshToken: refreshToken },
      }),
    }),
    getOverallHappinessScore: builder.query({
      query: ({ period, id }) => ({
        url: `happiness-score?period=${period}&account-id=${id}`,
        method: "GET",
      }),
    }),
    getDepartmentProportion: builder.query({
      query: () => ({
        url: `/happiness-score-department`,
        method: "GET",
      }),
    }),
  }),
});

// Export hooks for usage in functional components, which are
// auto-generated based on the defined endpoints
export const {
  useGetHealthCheckQuery,
  useGetHappinessPointsQuery,
  useGetEmployeesQuery,
  useLoginMutation,
  useLogoutMutation,
  useGetEmployeeQuery,
  useGetOverallHappinessScoreQuery,
  useGetDepartmentProportionQuery,
} = feelmeAPI;
