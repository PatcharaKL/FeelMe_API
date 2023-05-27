import { renderWithProviders } from "../../utils/test-utils";
import { screen, waitFor } from "@testing-library/react";
import { HealthCheck } from "../HealthCheck";

//* use 'screen.debug()' when want to output dom in terminal
describe("test Health Check", () => {
  it("should return healthy when call health-check endpoint", async () => {
    // Arrange
    renderWithProviders(<HealthCheck />);
    const value = await screen.findByText("Status: Healthy");

    // Assert
    await expect(value).toBeDefined;
  });
});
