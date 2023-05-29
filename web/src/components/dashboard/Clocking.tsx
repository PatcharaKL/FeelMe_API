import React from "react";
import { useGetClockingRecordQuery } from "../../services/feelme_api";

interface Clocking {
  "clock-in": string;
  "clock-out": string;
}

interface Record {
  date: string;
  clockIn: string;
  clockOut: string;
  workTime: number;
}

interface Response {
  Data: {
    name: string;
    record: Clocking[];
  };
}

const transformData = (data: Response): Record[] => {
  const transformResult: Record[] = [];
  console.log(Boolean(data.Data?.record));
  if (data.Data.record) {
    data.Data.record.forEach((rec) => {
      const date = rec["clock-in"].split(" ")[0];
      const clockIn = rec["clock-in"].split(" ")[1];
      const clockOut = rec["clock-out"].split(" ")[1];
      const clockInTime = new Date(rec["clock-in"]);
      const clockOutTime = new Date(rec["clock-out"]);

      // Convert the duration to hours
      const workTimeHours = clockOutTime.getHours() - clockInTime.getHours();

      const transformedRecord: Record = {
        date: date,
        clockIn: clockIn,
        clockOut: clockOut,
        workTime: workTimeHours,
      };

      transformResult.push(transformedRecord);
    });
  } else {
    transformResult.push({
      date: "-",
      clockIn: "-",
      clockOut: "-",
      workTime: 0,
    });
  }

  return transformResult.reverse();
};

const Clocking = (props: any) => {
  const { data, isLoading, isSuccess } = useGetClockingRecordQuery(props.id);
  return (
    <>
      <div className="text-xl font-semibold text-gray-800">
        Clocking History
      </div>
      <div className="text-md font-light text-gray-500">
        Clock-in and Clock-out each day
      </div>
      <div className="mt-6 h-full overflow-y-scroll">
        {!isLoading && (
          <div className="relative">
            <table className="w-full text-left text-sm text-gray-500">
              <thead className="sticky top-0 bg-violet-50 text-xs uppercase text-gray-700">
                <tr>
                  <th scope="col" className="px-6 py-3">
                    Date
                  </th>
                  <th scope="col" className="px-6 py-3">
                    Clock-in
                  </th>
                  <th scope="col" className="px-6 py-3">
                    Clock-out
                  </th>
                  <th scope="col" className="px-6 py-3">
                    Duration
                  </th>
                </tr>
              </thead>
              <tbody>
                {!isLoading &&
                  isSuccess &&
                  data.Data.record &&
                  transformData(data).map((data) => {
                    return (
                      <tr className="border-b bg-white">
                        <th
                          scope="row"
                          className="whitespace-nowrap px-6 py-4 font-medium text-gray-900"
                        >
                          {data.date}
                        </th>
                        <td className="px-6 py-4">{data.clockIn}</td>
                        <td className="px-6 py-4">{data.clockOut}</td>
                        <td className="px-6 py-4">
                          {!isNaN(data.workTime) ? data.workTime : "-"}
                        </td>
                      </tr>
                    );
                  })}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </>
  );
};

export default Clocking;
