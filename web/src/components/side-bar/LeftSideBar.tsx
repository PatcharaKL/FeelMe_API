// FeelThem

import { FeelMeLogo } from "../FeelMeLogo";
import { icons } from "../../assets/icons";

import { SideBarItemType } from "./type";
import { SideBarItem } from "./SideBarItem";
import {
  SideBarButton,
  SideBarContainer,
  SideBarItemsGroup,
  SideBarLogo,
} from "./SideBar";
import React from "react";
import { useAppDispatch } from "../../app/hooks";
import { logout } from "../../features/auth/authSlice";
import { useLogoutMutation } from "../../services/feelme_api";

const sidebarItemList: SideBarItemType[] = [
  {
    id: 1,
    name: "Dashboard",
    icon: icons.dashboard,
  },
  {
    id: 2,
    name: "Employees",
    icon: icons.diversity,
  },
  // {
  //   id: 3,
  //   name: "Setting",
  //   icon: icons.setting,
  // },
];

const logOutIcon = (): JSX.Element => {
  return (
    <div className="flex h-12 w-12 items-center justify-center rounded-full bg-violet-900 text-white">
      {React.createElement(icons.logout)}
    </div>
  );
};
const LeftSideBar = () => {
  const [logoutTrigger, { isLoading }] = useLogoutMutation();
  const dispatch = useAppDispatch();

  const logoutHandler = async () => {
    await logoutTrigger(null);
    dispatch(logout());
  };

  const sidebarItems = sidebarItemList.map((item) => (
    <SideBarItem key={item.id} {...item} />
  ));
  return (
    <>
      <SideBarContainer className="flex h-full w-60 flex-col gap-12 rounded-xl px-4 py-6">
        <SideBarItemsGroup className="flex justify-center">
          <SideBarLogo Logo={FeelMeLogo} />
        </SideBarItemsGroup>
        <SideBarItemsGroup className="flex flex-1 flex-col gap-2">
          {sidebarItems}
        </SideBarItemsGroup>
        <SideBarItemsGroup className="flex flex-col align-middle">
          <SideBarButton
            className={`${isLoading && "animate-pulse"}`}
            disabled={isLoading ? true : false}
            onClick={logoutHandler}
            icon={logOutIcon}
            name="Logout"
          />
        </SideBarItemsGroup>
      </SideBarContainer>
    </>
  );
};

export default LeftSideBar;
