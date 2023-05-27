import { renderWithProviders } from "../../utils/test-utils";
import { FeelMeLogo } from "../FeelMeLogo";
import { screen } from "@testing-library/react";

//* use 'screen.debug()' when want to output dom in terminal
describe("test Health Check", () => {
  it("should return true", async () => {
    // Arrange
    renderWithProviders(<FeelMeLogo />);
    const feelText = await screen.findByText("Feel");
    const meText = await screen.findByText("them");

    // Assert
    await expect(feelText).toBeDefined;
    await expect(meText).toBeDefined;
  });
});
