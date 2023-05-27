import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import "./index.css";
import { setupStore } from "./app/store";
import { Provider } from "react-redux";

// if (process.env.NODE_ENV === "development") {
//   const { worker } = await import("./mocks/browser");
//   worker.start();
// }

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <Provider store={setupStore()}>
      <App />
    </Provider>
  </React.StrictMode>
);
