// FeelThem

const default_nav = "#";

export const SideBarContainer = (
  props: React.HTMLAttributes<HTMLDivElement>
) => {
  return <aside {...props}>{props.children}</aside>;
};

export const SideBarItemsGroup = (
  props: React.HTMLAttributes<HTMLDivElement>
) => {
  return <div {...props}>{props.children}</div>;
};

export const SideBarLogo = ({ Logo, to = default_nav }: any) => {
  return (
    <a
      href={to}
      className="text-center transition duration-100 ease-in-out hover:scale-105 active:scale-100"
    >
      <Logo />
    </a>
  );
};

interface SideBarButton {
  name?: string;
  icon?: any;
  disabled?: boolean;
}
export const SideBarButton = ({
  icon: Icon,
  name,
  disabled = false,
  ...props
}: React.HTMLAttributes<HTMLDivElement> & SideBarButton) => {
  return (
    <div {...props}>
      <button
        className="flex items-center gap-4 px-4 transition ease-in-out hover:scale-105"
        disabled
      >
        {Icon && <Icon />}
        {name && <div className="font-medium">{name}</div>}
      </button>
    </div>
  );
};
