import smileLogo from "../../src/assets/smile.png";

export const FeelMeLogo = () => {
  return (
    <>
      <div className="flex select-none justify-center">
        <img className="h-8 object-scale-down" src={smileLogo}></img>
      </div>
      <div className="select-none text-4xl font-semibold text-gray-700">
        Feel<span className="font-extrabold text-violet-700">them</span>
      </div>
      <div className="select-none text-start text-sm font-light text-gray-600">
        ---- feel your people ----
      </div>
    </>
  );
};
