import { Pagination } from "@mui/material";
import { useGetEmployeesQuery } from "../../services/feelme_api";
import SearchIcon from "@mui/icons-material/Search";
import { useEffect, useState } from "react";
import CustomPagination from "./CustomPagination";
import EmployeeDashboard from "../dashboard/EmployeeDashboard";
import { HealthBar } from "./HealthBar";
import EditProfile from "./EditProfile";
interface Employees {
  account_id: number;
  hp: number;
  name: string;
  surname: string;
  avatar_url: string;
  position_name: string;
  setDashboardVisible: any;
  editVisible: any;
  setEditVisible: any;
}

export const Employees = () => {
  const {
    data: employees,
    isLoading,
    isSuccess,
    isFetching,
    error,
  } = useGetEmployeesQuery({});
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 8;
  const indexOfLastItem = currentPage * itemsPerPage;
  const indexOfFirstItem = indexOfLastItem - itemsPerPage;
  const paginate = (pageNumber: any) => setCurrentPage(pageNumber);
  const [dashboardVisible, setDashboardVisible] = useState({
    status: false,
    selectedID: null,
  });
  const [editVisible, setEditVisible] = useState({
    boardShow: false,
    status: false,
    selectedID: null,
  });

  return (
    <div className="flex h-full flex-col gap-5">
      {editVisible.boardShow && (
        <EditProfile setEditVisible={setEditVisible} editVisible={editVisible}/>
      )}

      {dashboardVisible.status && (
        <EmployeeDashboard
          employeeID={dashboardVisible.selectedID}
          setDashboardVisible={setDashboardVisible}
        />
      )}

      {!isLoading && isSuccess && !isFetching && (
        <>
          <Header setEditVisible={setEditVisible} editVisible={editVisible} />
          <CustomPagination
            itemsPerPage={itemsPerPage}
            totalItems={employees.length}
            paginate={paginate}
          />
        </>
      )}
      <div className="grid grid-cols-4 gap-5">
        {!isLoading &&
          isSuccess &&
          employees
            .slice(indexOfFirstItem, indexOfLastItem)
            .map((employee: Employees) => (
              <EmployeesCard
                key={employee.account_id}
                account_id={employee.account_id}
                name={employee.name}
                surname={employee.surname}
                avatar_url={employee.avatar_url}
                position_name={employee.position_name}
                hp={employee.hp}
                setDashboardVisible={setDashboardVisible}
                editVisible={editVisible}
                setEditVisible={setEditVisible}
              />
            ))}
      </div>
    </div>
  );
};

const Header = (props: any) => {
  return (
    <div className="flex h-fit w-full justify-between rounded-lg bg-violet-50 py-4 px-10">
      <h1 className="text-4xl font-bold text-violet-900">Employees</h1>
      <div className="m-auto flex w-full justify-end space-x-10">
        <label className="relative block w-[40%] self-end">
          <span className="sr-only">Search</span>
          <span className="absolute inset-y-0 left-0 flex items-center pl-2">
            <SearchIcon className="text-violet-500" />
          </span>
          <input
            type="text"
            className="w-full rounded-lg border border-violet-300 py-2 pl-9 pr-3 placeholder-violet-500 ring-violet-500 focus:outline-none focus:ring-1"
            placeholder="Search for employees.."
          ></input>
        </label>
        <button
          className=""
          onClick={() =>
            props.setEditVisible({
              selectedID: 0,
              status: !props.editVisible.status,
            })
          }
        >
          EDIT
        </button>
      </div>
    </div>
  );
};

const EmployeesCard = ({
  account_id,
  hp,
  name,
  surname: lastName,
  avatar_url,
  position_name,
  setDashboardVisible,
  editVisible,
  setEditVisible,
}: Employees) => {
  const EditOverlay = () => {
    return (
      <>
        <button
          onClick={() =>
            setEditVisible({
              selectedID: account_id,
              boardShow: true,
              status: true,
            })
          }
          className="absolute top-0 right-0 flex h-10 w-10 items-center justify-center rounded-xl border-2 border-white bg-orange-400 text-white"
        >
          Edit
        </button>
      </>
    );
  };

  const Overlay = () => {
    return (
      <div className="absolute inset-0 flex items-center justify-center rounded-lg bg-red-400/20 ring-2 ring-red-600/50"></div>
    );
  };

  return (
    <>
      <div
        className="relative"
        onClick={() =>
          setDashboardVisible({
            status: !editVisible.status,
            selectedID: account_id,
          })
        }
      >
        {editVisible.status && <EditOverlay />}
        {hp <= 0 && <Overlay />}
        <div className="flex h-fit w-64 flex-col items-center gap-3 overflow-hidden rounded-lg bg-violet-50 px-4 py-8 text-center shadow-lg shadow-violet-100">
          <HealthBar hp={hp} />
          <CardImage avatarURL={avatar_url}></CardImage>
          <div className="w-full">
            <p className="truncate text-xl font-bold">
              {name} {lastName}
            </p>
            <p className="truncate">{position_name}</p>
          </div>
        </div>
      </div>
    </>
  );
};

const CardImage = ({ avatarURL }: any) => {
  return (
    <>
      <img
        className="h-48 w-48 rounded-full object-scale-down ring-4 ring-emerald-300"
        src={avatarURL}
      ></img>
    </>
  );
};
