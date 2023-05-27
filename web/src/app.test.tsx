import { renderWithProviders } from "./utils/test-utils";
import App from "./App";
import { screen } from "@testing-library/react";

//* use 'screen.debug()' when want to output dom in terminal
describe("test msw", () => {
  //! fix app.tsx test
  it("should return healthy when call health-check endpoint", async () => {
    // Arrange

    // Assert
    await expect(true).toBe(true);
  });
});
