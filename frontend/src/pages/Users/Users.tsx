import React from "react";
import { UsersTable } from "../../components/organisms/UsersTable/UsersTable";
import { BasePage } from "../../components/templates/BasePage/BasePage";

import {
  User,
  useUsers
} from '../../hooks';

export const Users:React.FC = () => {
  // TODO::コンパイルを通すため一時的にany
  // 解消方法はfetch の型がany になってしまうのでその修正
  // const users = useUsers();
  const users:any = useUsers();

  console.log(users)

    return (
        <>
          <BasePage>
            <UsersTable 
              users={users}
            />  
          </BasePage>
        </>
    )
}