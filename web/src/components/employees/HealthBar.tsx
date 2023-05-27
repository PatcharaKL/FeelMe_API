export const HealthBar = ({ hp }: { hp: number }) => {
    const calHpColor = () => {
      if (hp <= 20) {
        return "#E11D48";
      }
      if (hp <= 50) {
        return "#FBBF24";
      }
      if (hp <= 100) {
        return "#22C55E";
      }
      return "#303030";
    };
  
    const hp_style = {
      width: hp + "%",
      backgroundColor: calHpColor(),
    };
  
    return (
      <div className="w-full">
        <div className="text-md font-semibold">
          HP: <span className="text-violet-700">{hp}/100</span>
        </div>
        <div className="flex h-4 flex-col items-start justify-center rounded-md bg-violet-200">
          <div className={`h-2 w-10/12 rounded-md`} style={hp_style}></div>
        </div>
      </div>
    );
  };