import React, { useEffect, useState } from "react";
import { useGetOverallHappinessScoreQuery } from "../../services/feelme_api";
import { PieChart, Pie, Cell } from "recharts";
import PeriodSelector from "./periodSelector";
import { selectPeriodSelected } from "../../features/period-selection/periodSelectionSlice";
import { useAppSelector } from "../../app/hooks";

const OverallScore = ({ employeeID = "" }: any) => {
  const selectedPeriod = useAppSelector(state => state.period.period)
  const [period, setPeriod] = useState("");
  const {
    data: hpScore,
    isLoading,
    isSuccess,
  } = useGetOverallHappinessScoreQuery({ period: selectedPeriod, id: employeeID });
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
    <div className="flex w-full items-center justify-around gap-6">
      {!isLoading && isSuccess && (
        <div className="flex h-full w-full items-center justify-around overflow-hidden rounded-2xl border border-violet-50 py-7 shadow-md shadow-violet-100">
          <div className="relative flex space-y-4">
            <p className="absolute inset-0 text-center text-xl font-bold">
              OVER ALL
            </p>
            <p className="absolute inset-0 flex items-center justify-center gap-2 text-4xl font-bold">
              {getScore().overall}
              <p className="text-sm">pts.</p>
            </p>
            <MatrixScore
              height={260}
              width={260}
              label="Overall Score"
              hpScore={getScore().overall}
              barColor="#975fff"
            />
          </div>
          <div>=</div>
          <div className="relative flex space-y-4">
            <p className="absolute inset-0 text-center text-sm font-semibold">
              Workplace
            </p>
            <p className="absolute inset-0 flex items-center justify-center gap-2 text-2xl font-bold">
              {getScore().working}
              <p className="text-xs">pts.</p>
            </p>
            <MatrixScore
              height={160}
              width={160}
              label="Overall Score"
              hpScore={getScore().overall}
              />
          </div>
              <div>+</div>
          <div className="relative flex space-y-4">
            <p className="absolute inset-0 text-center text-sm font-semibold">
              Employee
            </p>
            <p className="absolute inset-0 flex items-center justify-center gap-2 text-2xl font-bold">
              {getScore().self}
              <p className="text-xs">pts.</p>
            </p>
            <MatrixScore
              height={160}
              width={160}
              label="Overall Score"
              hpScore={getScore().self}
              />
          </div>
              <div>+</div>
          <div className="relative flex space-y-4">
            <p className="absolute inset-0 text-center text-sm font-semibold">
              Social
            </p>
            <p className="absolute inset-0 flex items-center justify-center gap-2 text-2xl font-bold">
              {getScore().social}
              <p className="text-xs">pts.</p>
            </p>
            <MatrixScore
              height={160}
              width={160}
              label="Overall Score"
              hpScore={getScore().social}
            />
          </div>
        </div>
      )}
      <PeriodSelector />
    </div>
  );
};

const MatrixScore = ({ hpScore, label, width, height, barColor = "#84d8ab" }: any) => {
  return (
    <div className="h-fit w-fit">
      <PieChart width={width} height={height}>
        <Pie
          data={[
            { name: "l", value: hpScore },
            { name: "h", value: 100 - hpScore },
          ]}
          dataKey="value"
          innerRadius="60%"
          startAngle={90}
          endAngle={480}
        >
          <Cell fill={barColor}/>
          <Cell fill="#8884d85e" />
        </Pie>

        {/* <text
          textAnchor="middle"
          dominantBaseline="middle"
          className="text-gray-500"
        >
          {label}
        </text> */}
        {/* <text
          x={80}
          y={90}
          textAnchor="middle"
          dominantBaseline="middle"
          className="text-3xl font-bold text-violet-600"
          fill="#2b2b2b"
        >
          {hpScore}
        </text> */}
      </PieChart>
    </div>
  );
};
export default OverallScore;
