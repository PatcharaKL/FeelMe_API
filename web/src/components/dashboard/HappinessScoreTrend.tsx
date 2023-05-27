import {
  LineChart,
  CartesianGrid,
  Legend,
  Tooltip,
  XAxis,
  YAxis,
  ResponsiveContainer,
  Line,
  Brush,
} from "recharts";
import { useGetHappinessPointsQuery } from "../../services/feelme_api";

// interface RawUserHappinessHistory {
//   id: number;
//   period: string;
//   record: {
//     happiness_points: {
//       self_points: number;
//       work_points: number;
//       co_worker_points: number;
//     };
//     date: string;
//   }[];
// }

interface FuzzyData {
  Value: number;
  DateTime: string;
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
  date: string;
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
    const date = fuzzy_self_points_average[i].DateTime;
    const selfPoints = fuzzy_self_points_average[i].Value;
    const workPoints = fuzzy_work_points_average[i].Value;
    const coWorkerPoints = fuzzy_co_worker_points_average[i].Value;
    const overallScore = value_over_all_average[i].Value;

    const transformedEntry: UserHappinessHistory = {
      self_points: selfPoints,
      work_points: workPoints,
      co_worker_points: coWorkerPoints,
      overall_score: overallScore,
      date: date,
    };

    transformedData.push(transformedEntry);
  }
  console.log(transformedData);
  return transformedData;
};

export const HappinessScoreTrend = ({ id = "" }: any) => {
  const { data, isLoading, isSuccess, error } = useGetHappinessPointsQuery(id);

  return (
    <div className="h-full w-full">
      <div className="text-xl font-semibold text-gray-800">
        Happiness Score Trend
      </div>
      <div className="text-md font-light text-gray-500">
        happiness score trend from collected questionnaire every work day
      </div>
      {isLoading ? (
        <div className="h-full w-full animate-pulse bg-gray-100">Loading..</div>
      ) : (
        <ResponsiveContainer width="100%" height="100%">
          <LineChart
            data={transformToChartData(data)}
            margin={{
              top: 20,
              bottom: 50,
            }}
          >
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="date" />
            <YAxis
              dataKey={"Happiness points"}
              domain={[0, 100]}
              label={{
                value: "Happiness Score",
                angle: -90,
                position: "insideLeft",
                textAnchor: "middle",
              }}
            />
            <Tooltip />
            <Legend verticalAlign="top" height={36} />
            <Line
              type="monotone"
              name="Self Score"
              dataKey="self_points"
              stroke="#ff828c"
            />
            <Line
              type="monotone"
              name="Workplace Score"
              dataKey="work_points"
              stroke="#5fd28b"
            />
            <Line
              type="monotone"
              name="Co-Worker Score"
              dataKey="co_worker_points"
              stroke="#caa782"
            />
            <Line
              type="monotone"
              name="Overall Score"
              dataKey="overall_score"
              strokeWidth={3}
              stroke="#8884d8"
            />
            <Brush
              dataKey={"date"}
              height={20}
              stroke="#8884d8"
              startIndex={transformToChartData(data).length - 7}
            />
          </LineChart>
        </ResponsiveContainer>
      )}
    </div>
  );
};
