import React, { ReactNode } from 'react';
import { Header } from '../../../components/atoms/Header/Header';
import { Sidebar } from '../../../components/atoms/Sidebar/Sidebar';

import styles from './BasePage.module.css'

export interface BasePageProps {
    children: ReactNode
}

export const BasePage: React.FC<BasePageProps> = ({
    children
}) => {

    return (
        <>
            <Header 
                pageTitle='title'
            />
            <Sidebar />
            <div className={`${styles.base}`}>
                {children}
            </div>
        </>  
    )
}