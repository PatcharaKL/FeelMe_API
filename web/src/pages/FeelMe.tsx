import { useAppSelector } from "../app/hooks";
import { Board } from "../components/board/Board";
import LeftSideBar from "../components/side-bar/LeftSideBar";
import RightSideBar from "../components/side-bar/RightSideBar";

const FeelMe = () => {
  return (
    <>
      <div className="h-screen w-screen bg-gradient-to-br from-white to-violet-50 p-4">
        <div className="flex h-full gap-4">
          <LeftSideBar />
          <Board />
          <RightSideBar />
        </div>
      </div>
    </>
  );
};
export default FeelMe;
