import React from "react";
import { icons } from "../../assets/icons";
import OverallScore from "./OverallScore";
import { HealthBar } from "../employees/HealthBar";
import { HappinessScoreTrend } from "./HappinessScoreTrend";
import { useGetEmployeeQuery } from "../../services/feelme_api";

const CloseIcon = icons.close;

const EmployeeDashboard = ({ employeeID, setDashboardVisible }: any) => {
  const {
    data: employee,
    isLoading,
    isSuccess,
  } = useGetEmployeeQuery(employeeID);
  return (
    <div className=" fixed inset-0 z-10 flex items-center justify-center bg-black bg-opacity-25 backdrop-blur-sm">
      {/* container */}
      <div className="flex h-5/6 w-4/5 flex-col overflow-hidden rounded-xl bg-white">
        {/* header */}
        <div className="flex items-center border-b-2 border-violet-100 bg-violet-600 px-14 py-3">
          <div className="text-2xl text-white">Dashboard</div>
          <button
            className="ml-auto"
            onClick={() => setDashboardVisible(false)}
          >
            <CloseIcon className="text-white" />
          </button>
        </div>
        {/* content */}
        <div className="flex h-full w-full flex-col gap-14 px-14 py-8">
          {!isLoading && isSuccess ? (
            <div className="flex justify-between gap-10">
              {/* Information */}
              <div className="flex flex-col gap-3">
                <div>
                  <HealthBar hp={employee?.hp} />
                </div>
                <div>
                  <div className="text-4xl font-semibold">
                    {employee?.name} {employee?.surname}
                  </div>
                  <div className="grid grid-cols-2 divide-x">
                    <div className="p-1">{employee?.department_name}</div>
                    <div className="p-1 px-2">{employee?.position_name}</div>
                  </div>
                </div>
              </div>
              <div className="h-full w-2/3">
                <OverallScore employeeID={employeeID}/>
              </div>
            </div>
          ) : (
            <div className="flex h-fit w-full animate-pulse items-center justify-center bg-gray-300">
              Loading...
            </div>
          )}
          {/* chart */}
          <div className=" h-96 w-full">
            <HappinessScoreTrend id={employeeID} />
          </div>
        </div>
      </div>
    </div>
  );
};

export default EmployeeDashboard;
