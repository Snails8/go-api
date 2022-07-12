import React from "react";
import { UsersTable } from "../../components/organisms/UsersTable/UsersTable";
import { BasePage } from "../../components/templates/BasePage/BasePage";

export const Users:React.FC = () => {
   
    return (
        <>
          <BasePage>
            <UsersTable />  
          </BasePage>
        </>
    )
}