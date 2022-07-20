import React from "react";
import { UsersTable } from "../../components/organisms/UsersTable/UsersTable";
import { BasePage } from "../../components/templates/BasePage/BasePage";

import {
  useUsers
} from '../../hooks/useUsers';

export const Users:React.FC = () => {
  const users = useUsers();
  
    return (
        <>
          <BasePage>
            <UsersTable />  
          </BasePage>
        </>
    )
}