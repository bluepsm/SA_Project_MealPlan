import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
import CreateComponentIcon from '@material-ui/icons/AddCircleOutline';
import SignOut from '@material-ui/icons/Settings';

import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarPinButton,
} from '@backstage/core';

export const AppSidebar = () => (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    <SidebarItem icon={HomeIcon} to="welcome" text="Home" />
    {/* <SidebarItem icon={CreateComponentIcon} to="create" text="Create..." />
    <SidebarItem icon={CreateComponentIcon} to="welcome" text="Welcome" /> */}
    <SidebarItem
      icon={CreateComponentIcon}
      to="meal_plan"
      text="Create mealplan"
    />

    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    {/* <SidebarUserSettings  /> */}
    <SidebarPinButton />
  </Sidebar>
);
