import React from "react";
import {
  selectPeriodSelected,
  periodAll,
  periodMonth,
  periodYear,
} from "../../features/period-selection/periodSelectionSlice";
import { useDispatch } from "react-redux";

const PeriodSelector = () => {
  const period = selectPeriodSelected;
  const dispatch = useDispatch();
  
  return (
    <div className="flex flex-col justify-end gap-6">
      <button className="bg-violet-300 px-1 py-2 w-20 rounded-md" onClick={() => dispatch(periodAll())}>All</button>
      <button className="bg-violet-300 px-1 py-2 w-20 rounded-md" onClick={() => dispatch(periodYear())}>year</button>
      <button className="bg-violet-300 px-1 py-2 w-20 rounded-md" onClick={() => dispatch(periodMonth())}>month</button>
    </div>
  );
};

export default PeriodSelector;
