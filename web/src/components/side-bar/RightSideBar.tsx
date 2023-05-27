// FeelThem

import React from "react";
import { images } from "../../assets/images";
import {
  SideBarContainer,
  SideBarItemsGroup,
  SideBarLogo,
  SideBarButton,
} from "./SideBar";
import { icons } from "../../assets/icons";
import { SvgIconComponent } from "@mui/icons-material";
import { useAppSelector } from "../../app/hooks";
import { useGetEmployeeQuery } from "../../services/feelme_api";

const avatar = () => {
  const id = useAppSelector((state) => state.user.id);
  const { data: user, isLoading: isUserLoading } = useGetEmployeeQuery(id);
  // const name = useAppSelector((state) => state.auth.name);
  // const position = useAppSelector((state) => state.auth.position);
  const AvatarImage = () => {
    return (
      <img
        className="rounded-full"
        src={user.avatar_url}
        alt="user avatar"
      />
    );
  };
  return (
    <div className="flex flex-col items-center gap-3 text-center overflow-hidden">
      {user && !isUserLoading && (
        <>
          <div className="h-16 w-16">
            <AvatarImage />
          </div>
          <div>
            <div className="text-lg font-bold text-violet-900 truncate">{user.name} {user.surname.slice(0, 1)+"."}</div>
            <div className="text-xs text-gray-700 truncate">{user.position_name}</div>
          </div>
        </>
      )}
    </div>
  );
};

const IconWrapper = (icon: SvgIconComponent): JSX.Element => {
  return (
    <div className="flex h-12 w-12 items-center justify-center">
      {React.createElement(icon, {
        className:
          "w-fit h-fit text-violet-800 hover:text-violet-600 hover:scale-125",
      })}
    </div>
  );
};

const RightSideBar = () => {
  return (
    <>
      <SideBarContainer className="flex h-full w-24 flex-col gap-12 rounded-xl py-6 px-2 pt-6">
        <SideBarItemsGroup className="flex justify-center">
          <SideBarLogo Logo={avatar} />
        </SideBarItemsGroup>
        <SideBarItemsGroup className="flex flex-1 flex-col items-center justify-end gap-4">
          {/* <SideBarButton icon={() => IconWrapper(icons.message)} />
          <SideBarButton icon={() => IconWrapper(icons.notification)} /> */}
        </SideBarItemsGroup>
      </SideBarContainer>
    </>
  );
};

export default RightSideBar;
