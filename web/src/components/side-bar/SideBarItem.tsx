import { useAppDispatch, useAppSelector } from "../../app/hooks";
import { icons } from "../../assets/icons";
import { setSelectedItem } from "../../features/sidebar-selection/sidebarSelectionSlice";
import { SideBarItemType } from "./type";

export const SideBarItem = ({
  id,
  icon: Icon = icons.default,
  name,
}: SideBarItemType) => {
  const dispatch = useAppDispatch();
  const selectedID = useAppSelector(
    (state) => state.sidebarSelection.selectedItemID
  );
  return (
    <button
      className={`group/unselected rounded-lg text-center font-medium transition duration-75 ease-in-out hover:scale-105 active:scale-100 ${
        selectedID == id
          ? "bg-violet-100 text-violet-900  hover:bg-violet-100 hover:text-violet-900"
          : "text-gray-500 hover:text-violet-800"
      }`}
      onClick={() => {
        dispatch(setSelectedItem(id));
        localStorage.setItem("sideBarSelectedID", String(id));
      }}
    >
      <span className="flex select-none gap-4 px-4 py-2">
        <Icon
          className={`${
            selectedID == id ? "text-violet-700" : "text-violet-500"
          }`}
        />
        {name}
      </span>
    </button>
  );
};
