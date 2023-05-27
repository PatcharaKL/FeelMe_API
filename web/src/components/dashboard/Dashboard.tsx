import React from "react";
import { HappinessScoreTrend } from "./HappinessScoreTrend";
import OverallScore from "./OverallScore";
import DepartmentBarChart from "./DepartmentBarChart";

const Dashboard = () => {
  return (
    <div className="grid h-full w-full grid-cols-6 gap-10 p-8">
      {/* Overall Score */}
      <div className="col-span-6 divide-y h-fit rounded-xl bg-gradient-to-tr from-violet-500 to-violet-50 p-12 text-white">
        <div className="text-5xl font-bold">Hello, Welcome Back!</div>
        <div className="text-lg font-light text-gray-100">
          We glad you are here with us today! <br />
          let's take a look on how's is your company doing today.
        </div>
      </div>
      <div className="col-span-6 flex w-full flex-col gap-3 rounded-xl">
        <div>
          <div className="text-xl font-semibold text-gray-800">
            Happiness Score
          </div>
          <div className="text-md font-light text-gray-500">
            your company all-time happiness score
          </div>
        </div>
        <div className=" self-center">
          <OverallScore />
        </div>
      </div>
      <div className=" col-span-3 h-[25rem]">
        <DepartmentBarChart />
      </div>
      {/* happiness score trend */}
      <div className="col-span-3 h-[30rem]">
        <HappinessScoreTrend />
      </div>
      
    </div>
  );
};

export default Dashboard;
