import React, { useState } from "react";
import { FormText } from "../../atoms/FormText/FormText";

import styles from "./UserCreateForm.module.css"

export const UserCreateForm:React.FC = () => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    
    return (
        <>
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
        </>
    )
}