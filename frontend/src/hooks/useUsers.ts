import { useEffect, useState } from 'react';

export interface User {
    /**
     * 
     * @type {number}
     * @memberof User
     */
    id: number;

    /**
     * 
     * @type {string}
     * @memberof User
     */
    name: string;
}

export const useUsers = () => {
    const [users, setUsers] = useState<User[] | undefined>(undefined);

    useEffect(() => {
        fetch('http://localhost:7001/api/v1/users')
            .then(res => res.json())
            .then((users) => {
                setUsers(users)
            },
            (error) => {
                console.log(error);
                const errData = {

                }
            }
            )
    }, [])

    return {
        users
    }
}