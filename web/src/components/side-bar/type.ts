import { SvgIconComponent } from "@mui/icons-material";

export interface SideBarItemType {
    id: number;
    name: string;
    icon?: SvgIconComponent;
    to?: string;
}