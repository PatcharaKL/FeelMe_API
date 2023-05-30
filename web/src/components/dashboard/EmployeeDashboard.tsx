import React from "react";
import { icons } from "../../assets/icons";
import OverallScore from "./OverallScore";
import { HealthBar } from "../employees/HealthBar";
import { HappinessScoreTrend } from "./HappinessScoreTrend";
import { useGetEmployeeQuery } from "../../services/feelme_api";
import Clocking from "./Clocking";

const CloseIcon = icons.close;

const EmployeeDashboard = ({ employeeID, setDashboardVisible }: any) => {
  const {
    data: employee,
    isLoading,
    isSuccess,
  } = useGetEmployeeQuery(employeeID);
  return (
    <div className="fixed inset-0 z-10 flex items-center justify-center bg-black bg-opacity-25 backdrop-blur-sm">
      {/* container */}
      <div className="relative flex h-5/6 w-4/5 flex-col overflow-y-scroll rounded-xl bg-white">
        {/* header */}
        <div className="flex sticky top-0 items-center border-b-2 border-violet-100 bg-violet-600 px-14 py-3">
          <div className="text-2xl text-white">Dashboard</div>
          <button
            className="ml-auto"
            onClick={() => setDashboardVisible(false)}
          >
            <CloseIcon className="text-white" />
          </button>
        </div>
        {/* content */}
        <div className="grid h-full w-full grid-cols-2 gap-14 px-14 py-8">
          {!isLoading && isSuccess ? (
            <>
              {/* Information */}
              <div className="top-0 grid grid-rows-3 grid-cols-2 divide-x grid-flow-col items-center">
                <img className="w-48 h-48 ring-4 ring-violet-400 object-cover row-span-3 rounded-full" src={employee?.avatar_url}/>
                <div className="col-span-3">
                  <HealthBar hp={employee?.hp} />
                </div>
                <div className="text-4xl font-semibold col-span-2">
                  {employee?.name} {employee?.surname}
                </div>
                <div className="p-1">{employee?.department_name}</div>
                <div className="p-1 px-2">{employee?.position_name}</div>
              </div>
              <div className="h-full w-full col-span-2">
                <OverallScore employeeID={employeeID} />
              </div>
            </>
          ) : (
            <div className="flex h-fit w-full animate-pulse items-center justify-center bg-gray-300">
              Loading...
            </div>
          )}
          {/* chart */}
          <div className="h-96 w-full">
            <HappinessScoreTrend id={employeeID} />
          </div>
          <div className=" w-ful h-3/4 overflow-hidden">
            <Clocking id={employeeID} />
          </div>
        </div>
      </div>
    </div>
  );
};

export default EmployeeDashboard;
