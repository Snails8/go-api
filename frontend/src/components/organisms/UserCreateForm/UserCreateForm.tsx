import React, { useState } from "react";
import { FormText } from "../../atoms/FormText/FormText";

import styles from "./UserCreateForm.module.css"

export const UserCreateForm:React.FC = () => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    
    async function handleSubmit(e: any) {
        e.preventDefault();
        const fromData = {
            name: name,
            email: email,
        }

        const res = await fetch("http://localhost:7001/api/v1/users/create",{
            method: "POST",
            body: JSON.stringify(fromData)
        })
    }

    return (
        <form onSubmit={handleSubmit}>
            <div className={`${styles.form_container}`}>
                <div className={`${styles.form_item}`}>
                    <div className={`${styles.container_inner}`}>
                        <span className={`${styles.form_label}`}>名前</span>
                            <FormText
                                type="text"
                                id="name"
                                defaultValue={name}
                                required={true}
                                width={400}
                                height={30}
                                padding={10}
                                onChangeHandler={(value: string) => {
                                    setName(value)
                                }}
                            />
                    </div>
                </div>
                <div className={`${styles.form_item}`}>
                    <div className={`${styles.container_inner}`}>
                        <span className={`${styles.form_label}`}>メール</span>
                            <FormText
                                type="email"
                                id="email"
                                defaultValue={email}
                                required={true}
                                width={400}
                                height={30}
                                padding={10}
                                onChangeHandler={(value: string) => {
                                    setEmail(value)
                                }}
                            />
                    </div>
                </div>
            </div>

            <div>
                <button type="submit" className="btn btn-success">検索</button>
            </div>
        </form>
    )
}