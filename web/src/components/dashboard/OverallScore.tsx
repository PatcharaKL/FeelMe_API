import React, { useEffect, useState } from "react";
import { useGetOverallHappinessScoreQuery } from "../../services/feelme_api";
import { PieChart, Pie, Cell } from "recharts";

const OverallScore = ({ employeeID = "" }: any) => {
  const [period, setPeriod] = useState("");
  const {
    data: hpScore,
    isLoading,
    isSuccess,
  } = useGetOverallHappinessScoreQuery({ period: "", id: employeeID });

  const getScore = () => {
    return {
      overall:
        hpScore && !isLoading && isSuccess
          ? hpScore?.value_over_all_average
          : "__",
      working:
        hpScore && !isLoading && isSuccess
          ? hpScore?.fuzzy_work_points_average
          : "__",
      self:
        hpScore && !isLoading && isSuccess
          ? hpScore?.fuzzy_self_points_average
          : "__",
      social:
        hpScore && !isLoading && isSuccess
          ? hpScore?.fuzzy_co_worker_points_average
          : "__",
    };
  };
  return (
    <>
      {!isLoading && isSuccess && (
        <div className="flex justify-around gap-10">
          <div className="flex h-fit w-fit items-center justify-evenly  overflow-hidden rounded-2xl border border-violet-50 py-7 shadow-md shadow-violet-100">
            <MatrixScore label="Overall Score" hpScore={getScore().overall} />
          </div>
          <div className="flex h-fit w-fit items-center justify-evenly  overflow-hidden rounded-2xl border border-violet-50 py-7 shadow-md shadow-violet-100">
            <MatrixScore label="Workplace Score" hpScore={getScore().working} />
          </div>
          <div className="flex h-fit w-fit items-center justify-evenly  overflow-hidden rounded-2xl border border-violet-50 py-7 shadow-md shadow-violet-100">
            <MatrixScore label="Self Score" hpScore={getScore().self} />
          </div>
          <div className="flex h-fit w-fit items-center justify-evenly  overflow-hidden rounded-2xl border border-violet-50 py-7 shadow-md shadow-violet-100">
            <MatrixScore label="Social Score" hpScore={getScore().social} />
          </div>
        </div>
      )}
    </>
  );
};

const MatrixScore = ({ hpScore, label }: any) => {
  return (
    <div className="">
      <PieChart width={160} height={160}>
        <Pie
          data={[
            { name: "l", value: hpScore },
            { name: "h", value: 100 - hpScore },
          ]}
          dataKey="value"
          innerRadius="60%"
          startAngle={90}
          endAngle={480}
          cy={85}
        >
          <Cell fill="#84d8ab" />
          <Cell fill="#8884d85e" />
        </Pie>

        <text
          x={80}
          y={8}
          textAnchor="middle"
          dominantBaseline="middle"
          className="text-gray-500"
        >
          {label}
        </text>
        <text
          x={80}
          y={90}
          textAnchor="middle"
          dominantBaseline="middle"
          className="text-3xl font-bold text-violet-600"
          fill="#2b2b2b"
        >
          {hpScore}
        </text>
      </PieChart>
    </div>
  );
};
export default OverallScore;
