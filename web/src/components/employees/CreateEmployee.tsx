import { useEffect, useState } from "react";
import { icons } from "../../assets/icons";
import {
  useCreateUserMutation,
  useGetEmployeeQuery,
  useUpdateUserDataMutation,
  useUploadImageMutation,
} from "../../services/feelme_api";
import Dropdown from "react-dropdown";
import "react-dropdown/style.css";

const CloseIcon = icons.close;
const EditProfile = ({ setAddVisible }: any) => {
  const [updateUserData, { isLoading: updating }] = useUpdateUserDataMutation();
  const [uploadFile, { isLoading: uploading, isError, error }] =
    useUploadImageMutation();

  const [createUser, { isLoading, isSuccess }] = useCreateUserMutation();
  const [userBody, setUserBody] = useState({
    email: "",
    name: "",
    surname: "",
    password: "",
    conPW: "",
    department_id: 0,
    position_id: 0,
  });

  const createHandler = (e: any) => {
    e.preventDefault();
    if (
      userBody.email &&
      userBody.password &&
      userBody.name &&
      userBody.position_id &&
      userBody.department_id &&
      userBody.surname
    ) {
      if (userBody.password == userBody.conPW) {
        createUser({
          email: userBody.email,
          password: userBody.password,
          name: userBody.name,
          surname: userBody.surname,
          position_id: Number(userBody.position_id),
          department_id: Number(userBody.department_id),
          company_id: 1,
        })
          .unwrap()
          .then((res) => console.log(res))
          .catch((e) => console.log(e));
        setAddVisible(false);
      }
    } else {
      alert("Password don't match");
      setUserBody({ ...userBody, password: "", conPW: "" });
    }
  };

  const departmentOptions = [
    { value: "1", label: "CE" },
    { value: "2", label: "MLE" },
    { value: "3", label: "IT" },
  ];
  const defaultDepartmentOption = departmentOptions[userBody.department_id - 1];
  const positionOptions = [
    { value: "1", label: "Fullstack Developer" },
    { value: "2", label: "Frontend Developer" },
    { value: "3", label: "Backend Developer" },
    { value: "4", label: "Human Resource" },
  ];
  const defaultPositionOption = positionOptions[userBody.position_id - 1];

  return (
    <>
      <div className=" fixed inset-0 z-10 flex items-center justify-center bg-black bg-opacity-25 backdrop-blur-sm">
        <div className="flex h-2/3 w-1/4 flex-col rounded-xl bg-white">
          {/* Header */}
          <div className="flex items-center border-b-2 border-violet-100 bg-violet-600 px-14 py-3">
            <div className="text-2xl text-white">Register</div>
            <button className="ml-auto" onClick={() => setAddVisible(false)}>
              <CloseIcon className="text-white" />
            </button>
          </div>
          {/* Content */}
          {!isLoading && (
            <form
              onSubmit={(e) => createHandler(e)}
              className="grid h-full w-full grid-cols-2 justify-center space-x-1 space-y-4 bg-slate-50 p-6"
            >
              <input
                className="col-span-2 h-10 rounded-md border border-gray-300 px-2"
                value={userBody.email}
                placeholder="Email"
                onChange={(e) =>
                  setUserBody((prev) => ({ ...prev, email: e.target.value }))
                }
                required
              />
              <input
                type="password"
                className="col-span-2 h-10 rounded-md border border-gray-300 px-2"
                value={userBody.password}
                placeholder="Password"
                onChange={(e) =>
                  setUserBody((prev) => ({ ...prev, password: e.target.value }))
                }
                required
              />
              <input
                type="password"
                className="col-span-2 h-10 rounded-md border border-gray-300 px-2"
                value={userBody.conPW}
                placeholder="Confirm Password"
                onChange={(e) =>
                  setUserBody((prev) => ({ ...prev, conPW: e.target.value }))
                }
                required
              />
              <input
                className="h-10 rounded-md border border-gray-300 px-2"
                value={userBody.name}
                placeholder="First name"
                onChange={(e) =>
                  setUserBody((prev) => ({ ...prev, name: e.target.value }))
                }
                required
              />
              <input
                type="text"
                value={userBody.surname}
                placeholder="Last name"
                className="h-10 rounded-md border  border-gray-300 px-2"
                onChange={(e) =>
                  setUserBody((prev) => ({ ...prev, surname: e.target.value }))
                }
                required
              />
              <Dropdown
                className="col-span-2"
                options={departmentOptions}
                value={defaultDepartmentOption}
                onChange={(id) =>
                  setUserBody((prev) => ({ ...prev, department_id: id.value }))
                }
                required
              />

              <Dropdown
                className="col-span-2"
                options={positionOptions}
                value={defaultPositionOption}
                onChange={(id) =>
                  setUserBody((prev) => ({ ...prev, position_id: id.value }))
                }
                required
              />
              <div className="col-span-2 flex">
                <button
                  type="submit"
                  className="m-auto h-10 w-1/3 rounded-lg bg-violet-600 text-white hover:bg-violet-700"
                >
                  Register
                </button>
              </div>
            </form>
          )}
        </div>
      </div>
    </>
  );
};

export default EditProfile;
