import { useEffect, useState } from "react";
import { icons } from "../../assets/icons";
import {
  useGetEmployeeQuery,
  useUpdateUserDataMutation,
  useUploadImageMutation,
} from "../../services/feelme_api";
import Dropdown from "react-dropdown";
import "react-dropdown/style.css";

const CloseIcon = icons.close;
const EditProfile = ({ editVisible, setEditVisible }: any) => {
  const [updateUserData, { isLoading: updating }] = useUpdateUserDataMutation();
  const [uploadFile, { isLoading: uploading, isError, error }] =
    useUploadImageMutation();
  const handleUpload = (event: any) => {
    if (event.target.files[0]) {
      uploadFile({ file: event.target.files[0], id: editVisible.selectedID })
        .unwrap()
        .then((response) => {
          console.log(response);
        })
        .catch((error) => {
          alert("Failed to upload image");
          console.error(error);
        });
    }
  };

  const {
    data: user,
    isLoading,
    isSuccess,
  } = useGetEmployeeQuery(editVisible.selectedID);
  const [userBody, setUserBody] = useState({
    account_id: 0,
    name: "",
    surname: "",
    department_id: 0,
    position_id: 0,
  });
  useEffect(() => {
    if (isSuccess) {
      setUserBody({
        account_id: user.account_id,
        name: user.name,
        surname: user.surname,
        department_id: Number(
          departmentOptions.find(
            (option) => option.label === user.department_name
          )?.value
        ),
        position_id: Number(
          positionOptions.find((option) => option.label === user.position_name)
            ?.value
        ),
      });
    }
  }, [user]);

  const updateHandler = () => {
    if (
      userBody.account_id &&
      userBody.name &&
      userBody.position_id &&
      userBody.department_id &&
      userBody.surname
    )
      updateUserData({
        account_id: editVisible.selectedID,
        name: userBody.name,
        surname: userBody.surname,
        position_id: Number(userBody.position_id),
        department_id: Number(userBody.department_id),
      })
        .unwrap()
        .then((res) => console.log(res))
        .catch((e) => console.log(e));
    setEditVisible({
      boardShow: false,
      selectedID: 0,
      status: true,
    })
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
            <div className="text-2xl text-white">Edit User</div>
            <button
              className="ml-auto"
              onClick={() =>
                setEditVisible({
                  boardShow: false,
                  selectedID: null,
                  status: true,
                })
              }
            >
              <CloseIcon className="text-white" />
            </button>
          </div>
          {/* Content */}
          {!isLoading && (
            <div className="grid h-full w-full grid-cols-2 justify-center space-x-1 space-y-4 bg-slate-50 p-6">
              <div className="col-span-2">
                <img
                  src={user.avatar_url}
                  alt="user profile picture"
                  className="m-auto w-56 rounded-full ring-8 ring-gray-300"
                />
              </div>
              <input
                accept="image/*"
                type="file"
                alt="user profile picture"
                className="col-span-2 self-center rounded-sm bg-slate-200 hover:bg-slate-300"
                onChange={handleUpload}
              />
              <input
                className="h-10 rounded-md border border-gray-300 px-2"
                value={userBody.name}
                placeholder={user.name}
                onChange={(e) =>
                  setUserBody((prev) => ({ ...prev, name: e.target.value }))
                }
              />
              <input
                type="text"
                value={userBody.surname}
                placeholder={user.surname}
                className="h-10 rounded-md border  border-gray-300 px-2"
                onChange={(e) =>
                  setUserBody((prev) => ({ ...prev, surname: e.target.value }))
                }
              />
              <Dropdown
                className="col-span-2"
                options={departmentOptions}
                value={defaultDepartmentOption}
                onChange={(id) =>
                  setUserBody((prev) => ({ ...prev, department_id: id.value }))
                }
              />

              <Dropdown
                className="col-span-2"
                options={positionOptions}
                value={defaultPositionOption}
              />
              <div className="col-span-2 flex">
                <button
                  onClick={updateHandler}
                  className="m-auto h-10 w-1/3 rounded-lg bg-violet-600 text-white hover:bg-violet-700"
                >
                  Done
                </button>
              </div>
            </div>
          )}
        </div>
      </div>
    </>
  );
};

export default EditProfile;
