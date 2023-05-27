import { useGetHealthCheckQuery } from "../services/feelme_api";

export const HealthCheck = () => {
  const { data, isLoading, isError, error, endpointName } =
    useGetHealthCheckQuery({});
  return (
    <div className="App">
      <h1 className="font-bold">Welcome to FeelMe</h1>
      <p>
        Status: {!isLoading && !isError ? data?.message : JSON.stringify(error)}
      </p>
    </div>
  );
};
