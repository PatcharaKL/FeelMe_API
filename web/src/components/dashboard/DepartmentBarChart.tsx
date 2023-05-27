import React from "react";
import {
  BarChart,
  CartesianGrid,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
  Legend,
  Bar,
} from "recharts";
import { useGetDepartmentProportionQuery } from "../../services/feelme_api";

interface FuzzyData {
  Value: number;
  Department: string;
}
interface RawUserHappinessHistory {
  fuzzy_co_worker_points_average: FuzzyData[];
  fuzzy_self_points_average: FuzzyData[];
  fuzzy_work_points_average: FuzzyData[];
  value_over_all_average: FuzzyData[];
}

interface UserHappinessHistory {
  self_points: number;
  work_points: number;
  co_worker_points: number;
  overall_score: number;
  department: string;
}

const transformToChartData = (
  rawData: RawUserHappinessHistory
): UserHappinessHistory[] => {
  const transformedData: UserHappinessHistory[] = [];

  const {
    fuzzy_self_points_average,
    fuzzy_work_points_average,
    fuzzy_co_worker_points_average,
    value_over_all_average,
  } = rawData;

  for (let i = 0; i < fuzzy_self_points_average.length; i++) {
    const department = fuzzy_self_points_average[i].Department;
    const selfPoints = fuzzy_self_points_average[i].Value;
    const workPoints = fuzzy_work_points_average[i].Value;
    const coWorkerPoints = fuzzy_co_worker_points_average[i].Value;
    const overallScore = value_over_all_average[i].Value;

    const transformedEntry: UserHappinessHistory = {
      self_points: selfPoints,
      work_points: workPoints,
      co_worker_points: coWorkerPoints,
      overall_score: overallScore,
      department: department,
    };

    transformedData.push(transformedEntry);
  }
  console.log(transformedData);
  return transformedData;
};

const DepartmentBarChart = () => {
  const { data, isLoading, isSuccess } = useGetDepartmentProportionQuery({});
  return (
    <div className="h-full w-full">
      <div className="text-xl font-semibold text-gray-800">
        Happiness Score Proportion
      </div>
      <div className="text-md font-light text-gray-500">
        happiness score proportion of all departments
      </div>
      {!isLoading ? (
        <div className="w-full mt-7 h-full rounded-xl border border-violet-100 bg-gradient-to-tr from-indigo-50 py-5 pr-10">
          <ResponsiveContainer width="100%" height="100%">
            <BarChart
              width={600}
              height={300}
              data={transformToChartData(data)}
            >
              <CartesianGrid strokeDasharray="3 3" />
              <XAxis dataKey="department" />
              <YAxis domain={[0, 100]} />
              <Tooltip />
              <Legend verticalAlign="top" height={36} />
              <Bar name="Overall" dataKey="overall_score" fill="#8884d8" />
              <Bar name="Workplace" dataKey="work_points" fill="#5fd28bb1" />
              <Bar name="Self" dataKey="self_points" fill="#ff828c" />
              <Bar name="Social" dataKey="co_worker_points" fill="#caa782" />
            </BarChart>
          </ResponsiveContainer>
        </div>
      ) : (
        <div>Loading..</div>
      )}
    </div>
  );
};

export default DepartmentBarChart;
