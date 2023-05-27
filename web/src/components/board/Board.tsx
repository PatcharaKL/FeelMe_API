import { useCallback } from "react";
import { useAppSelector } from "../../app/hooks";
import { HappinessScoreTrend } from "../dashboard/HappinessScoreTrend";
import { Employees } from "../employees/Employees";
import Dashboard from "../dashboard/Dashboard";

export const Board = () => {
  const selectedID = useAppSelector(
    (state) => state.sidebarSelection.selectedItemID
  );

  const Display = useCallback(() => {
    if (selectedID === 1) {
      return <Dashboard />;
    }
    if (selectedID === 2) {
      return <Employees />;
    }
    return <div>None..</div>;
  }, [selectedID]);

  return (
    <div
      id="board"
      className="flex h-full w-full flex-1 flex-col items-center justify-center 
      overflow-auto rounded-lg bg-white p-6 shadow-xl shadow-violet-200"
    >
      {Display()}
    </div>
  );
};
