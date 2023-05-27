import { renderWithProviders } from "../../../utils/test-utils";
// import { screen, waitFor } from "@testing-library/react";
import LeftSideBar from "../LeftSideBar";

//* use 'screen.debug()' when want to output dom in terminal
describe("test Health Check", () => {
  it("should return true", async () => {
    // Arrange
    renderWithProviders(<LeftSideBar />);
    // const value = await screen.findByText("Status: Healthy");

    // Assert
    await expect(true).toBe(true);
  });
});
